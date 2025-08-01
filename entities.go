package request

import "time"

//──────────────────────────────────────────────────────────────────────────────────────────────────

// Result of performing request
type Result struct {
	// Body represents the response body.
	Body []byte

	// BodyObject represents the response body marshaled as `map[string]any`,
	// nil if the body is not a valid JSON
	BodyObject map[string]any

	// StatusCode http status code, e.g. http.StatusOK
	StatusCode int

	// IsOK is status ok
	// Indeed does response got http.StatusOK
	IsOK bool
}

// Properties of perfoming request
type Properties struct {
	// Elapsed time spend to getting last response
	Elapsed time.Duration

	// TotalElapsed total time spend to getting responses
	TotalElapsed time.Duration

	// Retries number of retries performed
	Retries int

	// Errors contains all errors that occurred during the request
	Errors []error
}
