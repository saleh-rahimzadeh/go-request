package request

import "time"

//──────────────────────────────────────────────────────────────────────────────────────────────────

const (
	MAX_TIMEOUT               = 300 * time.Second
	MAX_RETRY_SLEEP_DURATIION = 10
)

// Content Types
//──────────────────────────────────────────────────────────────────────────────────────────────────

type ContentType string

const (
	HTTP_JSON ContentType = "application/json"
	HTTP_FORM ContentType = "application/x-www-form-urlencoded"
)
