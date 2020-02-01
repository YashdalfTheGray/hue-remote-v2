package utils_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/models"

	"github.com/yashdalfthegray/hue-remote-v2/utils"
)

func TestCheckEnv(t *testing.T) {
	testCases := []struct {
		desc string
		in   func(string) (string, bool)
		out  error
	}{
		{
			desc: "returns nil when all variables are found",
			in:   func(key string) (string, bool) { return "", true },
			out:  nil,
		},
		{
			desc: "returns error when all variables are not found",
			in:   func(key string) (string, bool) { return "", false },
			out:  models.NewMissingEnvError("HUE_BRIDGE_ADDRESS"),
		},
		{
			desc: "returns error when one of the variables is not found",
			in: func(key string) (string, bool) {
				if key == "HUE_BRIDGE_ADDRESS" {
					return "", true
				}
				return "", false
			},
			out: models.NewMissingEnvError("HUE_BRIDGE_USERNAME"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result := utils.CheckEnv(tC.in); result != tC.out {
				t.Errorf("Expected %s but got %s", tC.out.Error(), result.Error())
			}
		})
	}
}
