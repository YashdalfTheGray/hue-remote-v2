package colors_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/colors"
)

func TestHexToString(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HexCode
		out  string
	}{
		{
			desc: "prints out the string that's the hex code",
			in:   colors.HexCode("#164080"),
			out:  "#164080",
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

func TestHexToRGB(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HexCode
		out  string
		err  bool
	}{
		{
			desc: "converts a hex string to RGB",
			in:   colors.HexCode("#deadaf"),
			out:  "rgb(222, 173, 175)",
			err:  false,
		},
		{
			desc: "errors on string without the # sign",
			in:   colors.HexCode("deadaf"),
			out:  "rgb(0, 0, 0)",
			err:  true,
		},
		{
			desc: "errors on invalid characters",
			in:   colors.HexCode("#jklmno"),
			out:  "rgb(0, 0, 0)",
			err:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := tC.in.ToRGB()
			if !tC.err && err != nil {
				t.Errorf(err.Error())
			}
			if tC.err && err == nil {
				t.Errorf("Expected error to not be nil")
			}
			if result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result.String())
			}
		})
	}
}
func TestHexToHSL(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HexCode
		out  string
		err  bool
	}{
		{
			desc: "converts a hex string to HSL",
			in:   colors.HexCode("#deadaf"),
			out:  "hsl(358, 43%, 77%)",
			err:  false,
		},
		{
			desc: "errors on string without the # sign",
			in:   colors.HexCode("deadaf"),
			out:  "hsl(0, 0%, 0%)",
			err:  true,
		},
		{
			desc: "errors on invalid characters",
			in:   colors.HexCode("#jklmno"),
			out:  "hsl(0, 0%, 0%)",
			err:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := tC.in.ToHSL()
			if !tC.err && err != nil {
				t.Errorf(err.Error())
			}
			if tC.err && err == nil {
				t.Errorf("Expected error to not be nil")
			}
			if result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result.String())
			}
		})
	}
}

func TestHexToHSV(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HexCode
		out  string
		err  bool
	}{
		{
			desc: "converts a HexCode color into HSV",
			in:   colors.HexCode("#deadaf"),
			out:  "hsv(358, 22%, 87%)",
			err:  false,
		},
		{
			desc: "errors on invalid characters",
			in:   colors.NewHexCode("#jklmno"),
			out:  "hsv(0, 0%, 0%)",
			err:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := tC.in.ToHSV()
			if !tC.err && err != nil {
				t.Errorf(err.Error())
			}
			if tC.err && err == nil {
				t.Errorf("Expected error to not be nil")
			}
			if result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result.String())
			}
		})
	}
}

func TestHexToHueHSB(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HexCode
		out  string
		err  bool
	}{
		{
			desc: "converts a HexCode color into HSV",
			in:   colors.HexCode("#deadaf"),
			out:  "hue(65089, 56, 221)",
			err:  false,
		},
		{
			desc: "errors on invalid characters",
			in:   colors.HexCode("#jklmno"),
			out:  "hue(0, 0, 0)",
			err:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := tC.in.ToHueHSB()
			if !tC.err && err != nil {
				t.Errorf(err.Error())
			}
			if tC.err && err == nil {
				t.Errorf("Expected error to not be nil")
			}
			if result.String() != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result.String())
			}
		})
	}
}
