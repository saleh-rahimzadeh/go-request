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
	SendJson(c *Claim, data any) (Result, []error, error)
	SendForm(c *Claim, data map[string]string) (Result, []error, error)
	SendQuery(c *Claim, data any) (Result, []error, error)
}

type request struct {
	Timeout time.Duration
	Retries []time.Duration
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

// NewRequest create a new Request instance
// timeout define connection time limit
// retries define time pause between retries
func NewRequest(timeout int64, retries []int64) Request {
	var timeoutValue time.Duration = max(time.Duration(timeout)*time.Second, MAX_TIMEOUT)

	var retriesValue []time.Duration = make([]time.Duration, 0)
	for retry := range slices.Values(retries) {
		retriesValue = append(retriesValue, time.Second*time.Duration(retry))
	}

	return request{
		Timeout: timeoutValue,
		Retries: retriesValue,
	}
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func (r request) SendJson(c *Claim, data any) (Result, []error, error) {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return Result{}, []error{err}, err
	}
	c.ContentType(HTTP_JSON)
	return r.perform(*c, bytes.NewBuffer(dataByte))
}

func (r request) SendForm(c *Claim, data map[string]string) (Result, []error, error) {
	formData := net_url.Values{}
	for k, v := range data {
		formData[k] = []string{v}
	}
	encodedData := formData.Encode()
	c.ContentType(HTTP_FORM)
	return r.perform(*c, strings.NewReader(encodedData))
}

func (r request) SendQuery(c *Claim, data any) (Result, []error, error) {
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
	default:
		panic("SendQuery data is not valid")
	}

	c.URI.RawQuery = params.Encode()
	return r.perform(*c, nil)
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

// perform perform http request
// return Result on success
// return an array of retry errors
// return last error
func (r request) perform(c Claim, body io.Reader) (Result, []error, error) {
	var (
		result Result
		err    error
	)

	if len(r.Retries) == 0 {
		result, err = r.send(c, body)
		return result, []error{err}, err
	}

	var aerr []error

	for _, schedule := range r.Retries {
		result, err = r.send(c, body)
		if err == nil {
			return result, nil, nil
		}

		aerr = append(aerr, err)

		time.Sleep(schedule)
	}

	// All retries failed
	return Result{}, aerr, err
}

func (r request) send(c Claim, body io.Reader) (Result, error) {
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

	startTime := time.Now()

	client := &http.Client{
		Timeout: r.Timeout,
	}
	response, err := client.Do(httpRequest)
	if err != nil {
		return Result{}, err
	}
	defer response.Body.Close()

	elapsed := time.Since(startTime)

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
		Body:       responseBody,
		Elapsed:    elapsed.Milliseconds(),
		StatusCode: response.StatusCode,
		BodyObject: bodyObject,
		IsOK:       false,
	}

	if response.StatusCode == http.StatusOK {
		result.IsOK = true
	}

	return result, nil
}
