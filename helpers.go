package request

//──────────────────────────────────────────────────────────────────────────────────────────────────

func SendJson(c *Claim, data any) (Result, []error, error) {
	var req request
	return req.SendJson(c, data)
}

func SendForm(c *Claim, data map[string]string) (Result, []error, error) {
	var req request
	return req.SendForm(c, data)
}

func SendQuery(c *Claim, data map[string]string) (Result, []error, error) {
	var req request
	return req.SendQuery(c, data)
}
