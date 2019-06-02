package utils_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/utils"
)

func TestTempToMired(t *testing.T) {
	testCases := []struct {
		desc    string
		in, out uint16
	}{
		{
			desc: "converts a color temperature to mired value",
			in:   3000,
			out:  333,
		},
		{
			desc: "properly handles minimum color temperature",
			in:   2000,
			out:  500,
		},
		{
			desc: "properly handles maximum color temperature",
			in:   6500,
			out:  154,
		},
		{
			desc: "properly handles lower than minimum color temperature",
			in:   1500,
			out:  500,
		},
		{
			desc: "properly handles higher than maximum color temperature",
			in:   7000,
			out:  153,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result := utils.TempToMired(tC.in); result != tC.out {
				t.Errorf("Expected %d but got %d", tC.out, result)
			}
		})
	}
}
