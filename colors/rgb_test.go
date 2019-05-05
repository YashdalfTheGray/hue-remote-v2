package colors_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/colors"
	"github.com/yashdalfthegray/hue-remote-v2/utils"
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
			out:  utils.Must(colors.NewRGB(10, 45, 234)).(colors.RGB),
			err:  false,
		},
		{
			desc: "throws an error for values being out of bounds",
			r:    10,
			g:    45,
			b:    234,
			out:  utils.Must(colors.NewRGB(10, 45, 234)).(colors.RGB),
			err:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rgbBlue, err := colors.NewRGB(tC.r, tC.g, tC.b)
			if !tC.err && err != nil {
				t.Errorf("Expected error to be nil")
			} else if tC.err && err == nil {
				t.Errorf("Expected error but got nil")
			} else if rgbBlue.String() != tC.out.String() {
				t.Errorf("Expected %s but got %s", tC.out.String(), rgbBlue.String())
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
			in:   utils.Must(colors.NewRGB(0, 0, 0)).(colors.RGB),
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
