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

func TestRGBString(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.RGB
		out  string
	}{
		{
			desc: "builds a string out of an RGB color",
			in:   colors.NewRGB(200, 90, 50),
			out:  "rgb(200, 90, 50)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.in.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, tC.in.String())
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
		{
			desc: "converts another RGB color to hex string",
			in:   colors.NewRGB(68, 138, 255),
			out:  "#448aff",
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

func TestRGBToHSL(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.RGB
		out  colors.HSL
	}{
		{
			desc: "converts a non-gray RGB color to HSL",
			in:   colors.NewRGB(13, 166, 242),
			out:  utils.Must(colors.NewHSL(200, 90, 50)).(colors.HSL),
		},
		{
			desc: "converts a gray RGB color to HSL",
			in:   colors.NewRGB(128, 128, 128),
			out:  utils.Must(colors.NewHSL(0, 0, 50)).(colors.HSL),
		},
		{
			desc: "converts a dark RGB color to HSL",
			in:   colors.NewRGB(40, 34, 90),
			out:  utils.Must(colors.NewHSL(246, 45, 24)).(colors.HSL),
		},
		{
			desc: "converts a mostly green RGB color to HSL",
			in:   colors.NewRGB(40, 150, 90),
			out:  utils.Must(colors.NewHSL(147, 58, 37)).(colors.HSL),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := tC.in.ToHSL()
			if result.String() != tC.out.String() {
				t.Errorf("Expected %s but got %s", tC.out.String(), result.String())
			}
		})
	}
}
