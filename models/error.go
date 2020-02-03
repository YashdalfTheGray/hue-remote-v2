package models

import "errors"

// ErrorResponse is the server status response payload,
// gets sent back on any request
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// NewErrorResponse creates a new response to send out with the
// given server ID
func NewErrorResponse(err string) ErrorResponse {
	return ErrorResponse{Status: "error", Message: err}
}

// NewArgumentError created a new error instance with the right
// message and returns it, optionally accounting for if an
// argument was passed in.
func NewArgumentError(arg string) error {
	if arg == "" {
		return errors.New("The arguments given are invalid")
	}

	return errors.New("Argument " + arg + " is invalid")
}
