package colors_test

import (
	"github.com/yashdalfthegray/hue-remote-v2/colors"
	"testing"
)

func TestRGBToHexCode(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.RGB
		out  colors.HexCode
	}{
		{
			desc: "converts an RGB color to hex string",
			in:   colors.NewRGB(0, 0, 0),
			out:  "#000000",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.in.ToHexCode() != tC.out {
				t.Errorf("Expected %s but got %s", tC.in.ToHexCode(), tC.out)
			}
		})
	}
}
