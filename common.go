package request

import (
	"errors"
	"time"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

const (
	MAX_TIMEOUT               = 30 * time.Minute
	MAX_RETRY_SLEEP_DURATIION = 10
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
	ErrContentEmpty error = errors.New("content type is empty")
	ErrTokenEmpty   error = errors.New("token is empty")
)
