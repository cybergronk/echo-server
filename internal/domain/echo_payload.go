package domain

// Echo defines an echo contract.
// If any of the fields below is missing from a request, expect errors.
type Echo struct {
	What       string `json:"what"`
	StatusCode int    `json:"status_code"`
}
