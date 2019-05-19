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
			out:  colors.RGB{222, 173, 175}.String(),
			err:  false,
		},
		{
			desc: "errors on string without the # sign",
			in:   colors.HexCode("deadaf"),
			out:  colors.RGB{0, 0, 0}.String(),
			err:  true,
		},
		{
			desc: "errors on invalid characters",
			in:   colors.HexCode("#jklmno"),
			out:  colors.RGB{0, 0, 0}.String(),
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
				t.Errorf("expecting error to not be nil")
			}
			if result.String() != tC.out {
				t.Errorf("expected %s but got %s", tC.out, result.String())
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
			out:  colors.HSL{357.55, 42.61, 77.45}.String(),
			err:  false,
		},
		{
			desc: "errors on string without the # sign",
			in:   colors.HexCode("deadaf"),
			out:  colors.HSL{0, 0, 0}.String(),
			err:  true,
		},
		{
			desc: "errors on invalid characters",
			in:   colors.HexCode("#jklmno"),
			out:  colors.HSL{0, 0, 0}.String(),
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
				t.Errorf("expecting error to not be nil")
			}
			if result.String() != tC.out {
				t.Errorf("expected %s but got %s", tC.out, result.String())
			}
		})
	}
}

func TestToHSV(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HexCode
		out  string
		err  bool
	}{
		{
			desc: "converts a HexCode color into HSV",
			in:   colors.HexCode("#deadaf"),
			out:  colors.HSV{357.55, 22.07, 87.06}.String(),
			err:  false,
		},
		{
			desc: "errors on invalid characters",
			in:   colors.NewHexCode("#jklmno"),
			out:  colors.HSV{0, 0, 0}.String(),
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
				t.Errorf("expecting error to not be nil")
			}
			if result.String() != tC.out {
				t.Errorf("expected %s but got %s", tC.out, result.String())
			}
		})
	}
}

func TestToHueHSB(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HexCode
		out  string
		err  bool
	}{
		{
			desc: "converts a HexCode color into HSV",
			in:   colors.HexCode("#deadaf"),
			out:  colors.HueHSB{65089, 56, 221}.String(),
			err:  false,
		},
		{
			desc: "errors on invalid characters",
			in:   colors.HexCode("#jklmno"),
			out:  colors.HueHSB{0, 0, 0}.String(),
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
				t.Errorf("expecting error to not be nil")
			}
			if result.String() != tC.out {
				t.Errorf("expected %s but got %s", tC.out, result.String())
			}
		})
	}
}
