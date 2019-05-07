package colors_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/colors"
	"github.com/yashdalfthegray/hue-remote-v2/utils"
)

func TestNewHSL(t *testing.T) {
	testCases := []struct {
		desc    string
		h, s, l float64
		out     colors.HSL
		err     bool
	}{
		{
			desc: "builds a new HSL out of given numbers",
			h:    250.4,
			s:    65,
			l:    35,
			out:  utils.Must(colors.NewHSL(250.4, 65, 35)).(colors.HSL),
			err:  false,
		},
		{
			desc: "errors when numbers are out of bounds",
			h:    378,
			s:    65,
			l:    35,
			out:  utils.Must(colors.NewHSL(0, 0, 0)).(colors.HSL),
			err:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			color, err := colors.NewHSL(tC.h, tC.s, tC.l)
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

func TestHSLString(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HSL
		out  string
	}{
		{
			desc: "builds a string out of an HSL color",
			in:   utils.Must(colors.NewHSL(200, 90, 50)).(colors.HSL),
			out:  "hsl(200, 90%, 50%)",
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

func TestHSLToRGB(t *testing.T) {
	testCases := []struct {
		desc string
		in   colors.HSL
		out  colors.RGB
	}{
		{
			desc: "converts a non-gray HSL color to RGB",
			in:   utils.Must(colors.NewHSL(200, 90, 50)).(colors.HSL),
			out:  colors.NewRGB(12, 165, 242),
		},
		{
			desc: "converts a gray HSL color to RGB",
			in:   utils.Must(colors.NewHSL(0, 0, 50)).(colors.HSL),
			out:  colors.NewRGB(127, 127, 127),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := tC.in.ToRGB()
			if result.String() != tC.out.String() {
				t.Errorf("Expected %s but got %s", tC.out.String(), result.String())
			}
		})
	}
}
