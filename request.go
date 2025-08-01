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
	SendJson(c Demand, data any) (Result, Properties, bool)
	SendForm(c Demand, data any) (Result, Properties, bool)
	Send(c Demand) (Result, Properties, bool)
}

type request struct {
	Timeout time.Duration
	Retries []time.Duration
}

//┌ Instance
//└─────────────────────────────────────────────────────────────────────────────────────────────────

// New create a new Request instance
// timeout define connection time limit (refer to (http.Client).Timeout), maximum is MAX_TIMEOUT
// retries define time pause between retries, length retries represents number of retries to perform, empty array means only one try
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
func (r request) SendJson(c Demand, data any) (Result, Properties, bool) {
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
func (r request) SendForm(c Demand, data any) (Result, Properties, bool) {
	var body io.Reader

	switch payload := data.(type) {
	case map[string]string:
		formData := net_url.Values{}
		for k, v := range payload {
			formData[k] = []string{v}
		}
		encodedData := formData.Encode()
		body = strings.NewReader(encodedData)
	case net_url.Values:
		encodedData := payload.Encode()
		body = strings.NewReader(encodedData)
	case string:
		body = strings.NewReader(payload)
	default:
		dataByte, err := json.Marshal(payload)
		if err != nil {
			dataByte = []byte{}
		}
		body = bytes.NewBuffer(dataByte)
	}

	return r.perform(
		c.ContentType(HTTP_FORM),
		body,
	)
}

// Send http request
// It send request without any payload
func (r request) Send(c Demand) (Result, Properties, bool) {
	return r.perform(c, nil)
}

//┌ Internal Methods
//└─────────────────────────────────────────────────────────────────────────────────────────────────

// perform http request
// it silently discards and return unsuccess if c.Error contains error
// return Result on success
// return Response with properties and errors
// return True on success
func (r request) perform(c Demand, body io.Reader) (result Result, response Properties, isSuccess bool) {
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
		result, err = r.do(c, body)
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

func (r request) do(c Demand, body io.Reader) (Result, error) {
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
	if c.Type == string(HTTP_JSON) && len(responseBody) != 0 {
		if json.Unmarshal(responseBody, &bodyObject) != nil {
			bodyObject = nil
		}
	} else {
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
