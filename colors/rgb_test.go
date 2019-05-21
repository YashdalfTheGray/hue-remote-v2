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
	}{
		{
			desc: "builds a color out of rgb values",
			r:    10,
			g:    45,
			b:    234,
			out:  colors.RGB{10, 45, 234},
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
			in:   colors.RGB{200, 90, 50},
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
		out  string
	}{
		{
			desc: "converts an RGB color to hex string",
			in:   colors.RGB{0, 0, 0},
			out:  "#000000",
		},
		{
			desc: "converts another RGB color to hex string",
			in:   colors.RGB{68, 138, 255},
			out:  "#448aff",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result := tC.in.ToHexCode(); result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result)
			}
		})
	}
}

func TestRGBToHSL(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.RGB
		out  string
	}{
		{
			desc: "converts a non-gray RGB color to HSL",
			in:   colors.NewRGB(13, 166, 242),
			out:  colors.HSL{200, 90, 50}.String(),
		},
		{
			desc: "converts a gray RGB color to HSL",
			in:   colors.NewRGB(128, 128, 128),
			out:  colors.HSL{0, 0, 50}.String(),
		},
		{
			desc: "converts a dark RGB color to HSL",
			in:   colors.NewRGB(40, 34, 90),
			out:  colors.HSL{246, 45, 24}.String(),
		},
		{
			desc: "converts a mostly green RGB color to HSL",
			in:   colors.NewRGB(40, 150, 90),
			out:  colors.HSL{147, 58, 37}.String(),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result := tC.in.ToHSL(); result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result.String())
			}
		})
	}
}
