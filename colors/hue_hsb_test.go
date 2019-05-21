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
