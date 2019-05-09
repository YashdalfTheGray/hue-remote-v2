package models_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/models"
)

func TestNewErrorResponse(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		out  models.ErrorResponse
	}{
		{
			desc: "creates a new error response with message",
			in:   "test message",
			out:  models.NewErrorResponse("test message"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			message := models.NewErrorResponse(tC.in)
			if message.Message != tC.out.Message {
				t.Errorf("Expected message to be %s but was %s", tC.out.Message, message.Message)
			}
		})
	}
}
