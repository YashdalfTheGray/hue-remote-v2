package utils_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/utils"
)

func TestValidateColorString(t *testing.T) {
	testTable := []struct {
		in  string
		out bool
	}{
		{"#ffffff", true},
		{"#123456", true},
		{"#b16a12", true},
		{"3ca12b", false},
		{"", false},
		{"#fffffff", false},
		{"gibberish", false},
		{"(*&#(#(#*&", false},
	}

	for _, testData := range testTable {
		if result := utils.ValidateColorString(testData.in); result != testData.out {
			t.Errorf("ValidateColorString with input %s - expected %t but got %t", testData.in, testData.out, result)
		}
	}
}
