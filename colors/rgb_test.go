package colors_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/colors"
)

func TestNewRGB(t *testing.T) {
	testCases := []struct {
		desc    string
		r, g, b uint8
		out     colors.RGB
		err     bool
	}{
		{
			desc: "builds a color out of rgb values",
			r:    10,
			g:    45,
			b:    234,
			out:  colors.NewRGB(10, 45, 234),
			err:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if c := colors.NewRGB(tC.r, tC.g, tC.b); c.String() != tC.out.String() {
				t.Errorf("Expected %s but got %s", tC.out.String(), c.String())
			}
		})
	}
}

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
