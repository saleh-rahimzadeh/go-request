package request

import (
	"errors"
	"fmt"
	net_url "net/url"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

type Claim struct {
	URI     *net_url.URL
	Token   string
	Type    string
	Method  string
	Headers map[string]string
	Error   error
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func NewClaim(method string, rawUrl string, path string, params map[string]string) *Claim {
	claim := &Claim{
		URI:     nil,
		Token:   "",
		Type:    "",
		Method:  method,
		Headers: make(map[string]string),
	}
	uri, err := net_url.ParseRequestURI(rawUrl)
	if err != nil {
		claim.Error = errors.Join(claim.Error, err)
		return claim
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

	claim.URI = uri
	return claim
}

// ──────────────────────────────────────────────────────────────────────────────────────────────────

func (c *Claim) GetUrl() string {
	return c.URI.String()
}

func (c *Claim) ContentType(ctype ContentType) *Claim {
	if ctype == "" {
		c.Error = errors.Join(c.Error, errors.New("content type is empty"))
		return c
	}
	c.Type = string(ctype)
	return c
}

func (c *Claim) AuthorizationBearer(token string) *Claim {
	if token == "" {
		c.Error = errors.Join(c.Error, errors.New("token is empty"))
		return c
	}
	c.Token = fmt.Sprintf("Bearer %s", token)
	return c
}

func (c *Claim) Authorization(token string) *Claim {
	if token == "" {
		c.Error = errors.Join(c.Error, errors.New("token is empty"))
		return c
	}
	c.Token = token
	return c
}

func (c *Claim) AddHeader(name, value string) *Claim {
	c.Headers[name] = value
	return c
}
