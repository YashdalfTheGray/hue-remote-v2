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
			in:   colors.RGB{13, 166, 242},
			out:  "hsl(200, 90%, 50%)",
		},
		{
			desc: "converts a gray RGB color to HSL",
			in:   colors.RGB{128, 128, 128},
			out:  "hsl(0, 0%, 50%)",
		},
		{
			desc: "converts a dark RGB color to HSL",
			in:   colors.RGB{40, 34, 90},
			out:  "hsl(246, 45%, 24%)",
		},
		{
			desc: "converts a mostly green RGB color to HSL",
			in:   colors.RGB{40, 150, 90},
			out:  "hsl(147, 58%, 37%)",
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

func TestRGBToHSV(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.RGB
		out  string
	}{
		{
			desc: "converts a non-gray RGB color to HSL",
			in:   colors.RGB{13, 166, 242},
			out:  "hsv(200, 95%, 95%)",
		},
		{
			desc: "converts a gray RGB color to HSL",
			in:   colors.RGB{128, 128, 128},
			out:  "hsv(0, 0%, 50%)",
		},
		{
			desc: "converts black color to HSL",
			in:   colors.RGB{0, 0, 0},
			out:  "hsv(0, 0%, 0%)",
		},
		{
			desc: "converts a dark RGB color to HSL",
			in:   colors.RGB{40, 34, 90},
			out:  "hsv(246, 62%, 35%)",
		},
		{
			desc: "converts a mostly green RGB color to HSL",
			in:   colors.RGB{40, 150, 90},
			out:  "hsv(147, 73%, 59%)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result := tC.in.ToHSV(); result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result.String())
			}
		})
	}
}

func TestRGBToHueHSB(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.RGB
		out  string
	}{
		{
			desc: "converts an RGB color into HueHSB",
			in:   colors.RGB{42, 190, 235},
			out:  "hue(35314, 208, 234)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result := tC.in.ToHueHSB(); result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result.String())
			}
		})
	}
}
