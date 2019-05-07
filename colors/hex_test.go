package colors_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/colors"
	"github.com/yashdalfthegray/hue-remote-v2/utils"
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
		out  colors.RGB
		err  bool
	}{
		{
			desc: "converts a hex string to RGB",
			in:   colors.NewHexCode("#deadaf"),
			out:  colors.NewRGB(222, 173, 175),
			err:  false,
		},
		{
			desc: "errors on string without the # sign",
			in:   colors.NewHexCode("deadaf"),
			out:  colors.NewRGB(0, 0, 0),
			err:  true,
		},
		{
			desc: "errors on invalid characters",
			in:   colors.NewHexCode("#jklmno"),
			out:  colors.NewRGB(0, 0, 0),
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
			if result.String() != tC.out.String() {
				t.Errorf("expected %s but got %s", tC.out, result.String())
			}
		})
	}
}
func TestHexToHSL(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HexCode
		out  colors.HSL
		err  bool
	}{
		{
			desc: "converts a hex string to HSL",
			in:   colors.NewHexCode("#deadaf"),
			out:  utils.Must(colors.NewHSL(357.55, 42.61, 77.45)).(colors.HSL),
			err:  false,
		},
		{
			desc: "errors on string without the # sign",
			in:   colors.NewHexCode("deadaf"),
			out:  utils.Must(colors.NewHSL(0, 0, 0)).(colors.HSL),
			err:  true,
		},
		{
			desc: "errors on invalid characters",
			in:   colors.NewHexCode("#jklmno"),
			out:  utils.Must(colors.NewHSL(0, 0, 0)).(colors.HSL),
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
			if result.String() != tC.out.String() {
				t.Errorf("expected %s but got %s", tC.out, result.String())
			}
		})
	}
}
