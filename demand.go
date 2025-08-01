package request

import (
	"errors"
	"fmt"
	net_url "net/url"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

type Demand struct {
	URI     net_url.URL
	Token   string
	Type    string
	Method  string
	Headers map[string]string
	Error   error
}

//┌ Instance
//└─────────────────────────────────────────────────────────────────────────────────────────────────

// BuildDemand create a new demand instance
// method one of http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, ...
// url raw url interpreted only as an absolute URI
// path relative path
func BuildDemand(method string, url string, path string) Demand {
	demand := Demand{
		URI:     net_url.URL{},
		Token:   "",
		Type:    "",
		Method:  method,
		Headers: make(map[string]string),
	}

	uri, err := net_url.ParseRequestURI(url)
	if err != nil {
		demand.Error = errors.Join(demand.Error, err)
		return demand
	}
	if path != "" {
		uri.Path += path
	}

	demand.URI = *uri
	return demand
}

//┌ Methods
//└─────────────────────────────────────────────────────────────────────────────────────────────────

// GetUrl get a valid URL string
func (c Demand) GetUrl() string {
	return c.URI.String()
}

// ContentType set content type header
func (c Demand) ContentType(ctype ContentType) Demand {
	if ctype == "" {
		c.Error = errors.Join(c.Error, ErrDemandContentTypeEmpty)
		return c
	}
	c.Type = string(ctype)
	return c
}

// AuthorizationBearer set bearer authorization header
func (c Demand) AuthorizationBearer(token string) Demand {
	if token == "" {
		c.Error = errors.Join(c.Error, ErrDemandTokenEmpty)
		return c
	}
	c.Token = fmt.Sprintf("Bearer %s", token)
	return c
}

// Authorization set authorization header
func (c Demand) Authorization(token string) Demand {
	if token == "" {
		c.Error = errors.Join(c.Error, ErrDemandTokenEmpty)
		return c
	}
	c.Token = token
	return c
}

// AddHeader add additional header
func (c Demand) Header(name, value string) Demand {
	c.Headers[name] = value
	return c
}

// Parameter add query parameters to the URL
// params is payload of data in types:
// `map[string]string`,
// `map[string]any`,
// `url.Values`
func (c Demand) Parameter(params any) Demand {
	if params == nil {
		c.Error = errors.Join(c.Error, ErrDemandParamEmpty)
		return c
	}

	query := c.URI.Query()

	switch payload := params.(type) {
	case map[string]string:
		for k, v := range payload {
			query.Add(k, v)
		}
	case map[string]any:
		for k, v := range payload {
			query.Add(k, fmt.Sprintf("%v", v))
		}
	case net_url.Values:
		for k, v := range payload {
			for _, vv := range v {
				query.Add(k, vv)
			}
		}
	}

	c.URI.RawQuery = query.Encode()

	return c
}
