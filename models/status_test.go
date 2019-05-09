package models_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/models"
)

func TestNewStatusResponse(t *testing.T) {
	testCases := []struct {
		desc string
		out  models.StatusResponse
	}{
		{
			desc: "creates a new error response with message",
			out:  models.NewStatusResponse(),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			message := models.NewStatusResponse()
			if message.Status != "ok" {
				t.Errorf("Expected status to be \"ok\"")
			}
		})
	}
}
