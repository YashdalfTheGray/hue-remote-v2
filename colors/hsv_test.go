package colors_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/colors"
)

func TestNewHSV(t *testing.T) {
	testCases := []struct {
		desc    string
		h, s, v float64
		out     colors.HSV
		err     bool
	}{
		{
			desc: "creates a new HSV color out of valid channels",
			h:    45,
			s:    90,
			v:    100,
			out:  colors.HSV{45, 90, 100},
			err:  false,
		},
		{
			desc: "returns an error for invalid channel values",
			h:    450,
			s:    90,
			v:    100,
			out:  colors.HSV{},
			err:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			color, err := colors.NewHSV(tC.h, tC.s, tC.v)
			if !tC.err && err != nil {
				t.Errorf("Expected to not have error but got error %s", err.Error())
			} else if tC.err && err == nil {
				t.Errorf("Expected error but got nil")
			} else if color.String() != tC.out.String() {
				t.Errorf("Expecting %s but got %s", tC.out.String(), color.String())
			}
		})
	}
}

func TestHSVString(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HSV
		out  string
	}{
		{
			desc: "builds a string out of an HSV color",
			in:   colors.HSV{200, 90, 50},
			out:  "hsv(200, 90%, 50%)",
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

func TestHSVToRGB(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HSV
		out  string
	}{
		{
			desc: "converts an HSV color to RGB",
			in:   colors.HSV{216.23, 82.81, 50.2},
			out:  "rgb(22, 64, 128)",
		},
		{
			desc: "converts an orange HSV color to RGB",
			in:   colors.HSV{23, 100, 100},
			out:  "rgb(255, 97, 0)",
		},
		{
			desc: "converts a green HSV color to RGB",
			in:   colors.HSV{82, 100, 100},
			out:  "rgb(161, 255, 0)",
		},
		{
			desc: "converts a teal HSV color to RGB",
			in:   colors.HSV{170, 100, 100},
			out:  "rgb(0, 255, 212)",
		},
		{
			desc: "converts a blue HSV color to RGB",
			in:   colors.HSV{226, 100, 100},
			out:  "rgb(0, 59, 255)",
		},
		{
			desc: "converts a purple HSV color to RGB",
			in:   colors.HSV{280, 100, 100},
			out:  "rgb(170, 0, 255)",
		},
		{
			desc: "converts another red HSV color to RGB",
			in:   colors.HSV{344, 100, 100},
			out:  "rgb(255, 0, 67)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rgbColor := tC.in.ToRGB()
			if rgbColor.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, rgbColor.String())
			}
		})
	}
}

func TestHSVToHexCode(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HSV
		out  string
	}{
		{
			desc: "converts an HSV color to RGB",
			in:   colors.HSV{216.23, 82.81, 50.2},
			out:  "#164080",
		},
		{
			desc: "converts an orange HSV color to RGB",
			in:   colors.HSV{23, 100, 100},
			out:  "#ff6100",
		},
		{
			desc: "converts a green HSV color to RGB",
			in:   colors.HSV{82, 100, 100},
			out:  "#a1ff00",
		},
		{
			desc: "converts a teal HSV color to RGB",
			in:   colors.HSV{170, 100, 100},
			out:  "#00ffd4",
		},
		{
			desc: "converts a blue HSV color to RGB",
			in:   colors.HSV{226, 100, 100},
			out:  "#003bff",
		},
		{
			desc: "converts a purple HSV color to RGB",
			in:   colors.HSV{280, 100, 100},
			out:  "#aa00ff",
		},
		{
			desc: "converts another red HSV color to RGB",
			in:   colors.HSV{344, 100, 100},
			out:  "#ff0043",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rgbColor := tC.in.ToHexCode()
			if rgbColor.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, rgbColor.String())
			}
		})
	}
}

func TestHSVToHSL(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HSV
		out  string
	}{
		{
			desc: "converts a bright HSV color to RGB",
			in:   colors.HSV{344, 100, 100},
			out:  "hsl(344, 100%, 50%)",
		},
		{
			desc: "converts a dark HSV color to RGB",
			in:   colors.HSV{198, 60, 1},
			out:  "hsl(198, 43%, 1%)",
		},
		{
			desc: "converts a black HSV color to RGB",
			in:   colors.HSV{198, 60, 0.5},
			out:  "hsl(198, 0%, 0%)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rgbColor := tC.in.ToHSL()
			if rgbColor.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, rgbColor.String())
			}
		})
	}
}
