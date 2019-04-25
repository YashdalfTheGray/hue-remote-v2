package models

// StatusResponse is the server status response payload,
// gets sent back on GET /
type StatusResponse struct {
	Status string `json:"status"`
}

// NewStatusResponse creates a new response to send out with the
// given server ID
func NewStatusResponse() StatusResponse {
	return StatusResponse{Status: "ok"}
}
