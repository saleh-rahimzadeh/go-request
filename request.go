package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	net_url "net/url"
	"slices"
	"strings"
	"time"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

type Request interface {
	SendJson(c Demand, data any) (Result, Response, bool)
	SendForm(c Demand, data map[string]string) (Result, Response, bool)
	SendQuery(c Demand, data any) (Result, Response, bool)
}

type request struct {
	Timeout time.Duration
	Retries []time.Duration
}

//┌ Instance
//└─────────────────────────────────────────────────────────────────────────────────────────────────

// New create a new Request instance
// timeout define connection time limit (refer to (http.Client).Timeout), maximum is MAX_TIMEOUT
// retries define time pause between retries, empty array means only one try
func New(timeout time.Duration, retries []time.Duration) Request {
	var timeoutValue time.Duration = max(timeout, MAX_TIMEOUT)

	var retriesValue []time.Duration = make([]time.Duration, 0)
	for retry := range slices.Values(retries) {
		retriesValue = append(retriesValue, retry)
	}

	return request{
		Timeout: timeoutValue,
		Retries: retriesValue,
	}
}

//┌ Methods
//└─────────────────────────────────────────────────────────────────────────────────────────────────

// SendJson send http request with JSON payload
// if data is nil then request will be sent without body
// if can't encode json then empty body will be sent
func (r request) SendJson(c Demand, data any) (Result, Response, bool) {
	dataByte, err := json.Marshal(data)
	if err != nil {
		dataByte = []byte{}
	}
	return r.perform(
		c.ContentType(HTTP_JSON),
		bytes.NewBuffer(dataByte),
	)
}

// SendForm send http request with www form payload
func (r request) SendForm(c Demand, data map[string]string) (Result, Response, bool) {
	formData := net_url.Values{}
	for k, v := range data {
		formData[k] = []string{v}
	}
	encodedData := formData.Encode()
	return r.perform(
		c.ContentType(HTTP_FORM),
		strings.NewReader(encodedData),
	)
}

// SendQuery send http request with query string payload
// data is payload of data in `map[string]string` or `url.Values` format
// if data is invalid not query param will be set
func (r request) SendQuery(c Demand, data any) (Result, Response, bool) {
	if data != nil {
		params := c.URI.Query()

		switch payload := data.(type) {
		case map[string]string:
			for k, v := range payload {
				params.Add(k, v)
			}
		case net_url.Values:
			for k, v := range payload {
				for _, vv := range v {
					params.Add(k, vv)
				}
			}
		}

		c.URI.RawQuery = params.Encode()
	}
	return r.perform(c, nil)
}

//┌ Internal Methods
//└─────────────────────────────────────────────────────────────────────────────────────────────────

// perform http request
// it silently discards and return unsuccess if c.Error contains error
// return Result on success
// return Response with properties and errors
// return True on success
func (r request) perform(c Demand, body io.Reader) (result Result, response Response, isSuccess bool) {
	if c.Error != nil {
		isSuccess = false
		return //↩️ ∅
	}

	var (
		err   error
		start = time.Now()
	)

	defer func(start time.Time) {
		response.TotalElapsed = time.Since(start)
	}(start)

	if len(r.Retries) == 0 {
		r.Retries = []time.Duration{0}
	}

	for retryIndex, duration := range r.Retries {
		response.Retries = retryIndex + 1

		begin := time.Now()
		result, err = r.send(c, body)
		response.Elapsed = time.Since(begin)

		if err == nil {
			isSuccess = true
			return //↩️ ∅
		}

		response.Errors = append(response.Errors, err)

		time.Sleep(duration)
	}

	// All retries failed
	isSuccess = false
	return //↩️ ∅
}

func (r request) send(c Demand, body io.Reader) (Result, error) {
	httpRequest, err := http.NewRequest(c.Method, c.GetUrl(), body)
	if err != nil {
		return Result{}, err
	}
	if c.Type != "" {
		httpRequest.Header.Set("Content-Type", c.Type)
	}
	if c.Token != "" {
		httpRequest.Header.Add("Authorization", c.Token)
	}

	for k, v := range c.Headers {
		httpRequest.Header.Add(k, v)
	}

	client := &http.Client{
		Timeout: r.Timeout,
	}
	response, err := client.Do(httpRequest)
	if err != nil {
		return Result{}, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return Result{}, err
	}

	bodyObject := map[string]any{}
	err = json.Unmarshal(responseBody, &bodyObject)
	if err != nil {
		bodyObject = nil
	}

	var result = Result{
		StatusCode: response.StatusCode,
		Body:       responseBody,
		BodyObject: bodyObject,
		IsOK:       false,
	}

	if response.StatusCode == http.StatusOK {
		result.IsOK = true
	}

	return result, nil
}
