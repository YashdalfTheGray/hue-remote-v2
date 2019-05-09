package models_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/models"
)

func TestNewMissingEnvError(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		out  models.MissingEnvError
	}{
		{
			desc: "creates a new missing env error",
			in:   "TEST_ENV",
			out:  models.NewMissingEnvError("TEST_ENV"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			testErr := models.NewMissingEnvError(tC.in)
			if testErr.Error() != tC.out.Error() {
				t.Errorf("Expected error strings to be the same\n\tExpected - %s\n\tGot - %s", tC.out.Error(), testErr.Error())
			}
		})
	}
}

// func TestNilMissingEnvError(t *testing.T) {
// 	var nilEnvError *models.MissingEnvError = nil

// 	if nilEnvError.Error() != "" {
// 		t.Errorf("Expected to be able to call method MissingEnvError.Error() with nil receiver")
// 	}
// }
