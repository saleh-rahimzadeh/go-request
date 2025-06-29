package request

//──────────────────────────────────────────────────────────────────────────────────────────────────

type Result struct {
	Body       []byte
	BodyObject map[string]any
	StatusCode int
	Elapsed    int64
	IsOK       bool
}
