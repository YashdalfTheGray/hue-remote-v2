package models

// StatusResponse is the server status response payload,
// gets sent back on GET /
type StatusResponse struct {
	Status           string `json:"status"`
	RedisServerFound bool   `json:"redisServerFound"`
	BridgeFound      bool   `json:"bridgeFound"`
	BridgeUserFound  bool   `json:"bridgeUserFound"`
	APITokenFound    bool   `json:"apiTokenFound"`
}

// NewStatusResponse creates a new response to send out with the
// given server ID
func NewStatusResponse() StatusResponse {
	return StatusResponse{Status: "ok"}
}
