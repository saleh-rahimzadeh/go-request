package request

import (
	"errors"
	"time"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

const (
	MAX_TIMEOUT time.Duration = 30 * time.Minute
)

//┌ Content Types
//└─────────────────────────────────────────────────────────────────────────────────────────────────

type ContentType string

const (
	HTTP_JSON ContentType = "application/json"
	HTTP_FORM ContentType = "application/x-www-form-urlencoded"
)

//┌ Errors
//└─────────────────────────────────────────────────────────────────────────────────────────────────

var (
	ErrDemandContentTypeEmpty error = errors.New("content type is empty")
	ErrDemandTokenEmpty       error = errors.New("token is empty")
	ErrDemandParamEmpty       error = errors.New("params is empty")
)
