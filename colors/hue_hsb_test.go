package colors_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/colors"
)

func TestNewHueHSB(t *testing.T) {
	testCases := []struct {
		desc    string
		h, s, b int
		out     colors.HueHSB
		err     bool
	}{
		{
			desc: "Creates a HueHSB color out of valid values",
			h:    56234,
			s:    190,
			b:    250,
			out:  colors.HueHSB{56234, 190, 250},
			err:  false,
		},
		{
			desc: "throws an error when the channels are out of range",
			h:    56234,
			s:    400,
			b:    250,
			out:  colors.HueHSB{0, 0, 0},
			err:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			color, err := colors.NewHueHSB(tC.h, tC.s, tC.b)
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

func TestHueHSBString(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HueHSB
		out  string
	}{
		{
			desc: "converts a HueHSB color into a string for display",
			in:   colors.HueHSB{54789, 162, 200},
			out:  "hue(54789, 162, 200)",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if result := tC.in.String(); result != tC.out {
				t.Errorf("Expected %s but got %s", tC.out, result)
			}
		})
	}
}

func TestHueHSBToHSV(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HueHSB
		out  string
	}{
		{
			desc: "converts a HueHSB color into its HSV equivalent",
			in:   colors.HueHSB{54789, 162, 200},
			out:  "hsv(301, 64%, 79%)",
		},
		{
			desc: "converts another HueHSB color into its HSV equivalent",
			in:   colors.HueHSB{34565, 200, 100},
			out:  "hsv(190, 79%, 40%)",
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
