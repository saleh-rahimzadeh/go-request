package request

//──────────────────────────────────────────────────────────────────────────────────────────────────

// LastError get latest error if exists
func LastError(e []error) error {
	if len(e) == 0 {
		return nil
	}
	return e[len(e)-1]
}
