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
// params query parameters to append to the URL
func BuildDemand(method string, url string, path string, params map[string]string) Demand {
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

	if len(params) > 0 {
		q := uri.Query()
		for key, value := range params {
			q.Set(key, value)
		}
		uri.RawQuery = q.Encode()
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
		c.Error = errors.Join(c.Error, ErrContentEmpty)
		return c
	}
	c.Type = string(ctype)
	return c
}

// AuthorizationBearer set bearer authorization header
func (c Demand) AuthorizationBearer(token string) Demand {
	if token == "" {
		c.Error = errors.Join(c.Error, ErrTokenEmpty)
		return c
	}
	c.Token = fmt.Sprintf("Bearer %s", token)
	return c
}

// Authorization set authorization header
func (c Demand) Authorization(token string) Demand {
	if token == "" {
		c.Error = errors.Join(c.Error, ErrTokenEmpty)
		return c
	}
	c.Token = token
	return c
}

// AddHeader add additional header
func (c Demand) AddHeader(name, value string) Demand {
	c.Headers[name] = value
	return c
}
