package colors_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/colors"
	"github.com/yashdalfthegray/hue-remote-v2/utils"
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
			out:  utils.Must(colors.NewHSV(45, 90, 100)).(colors.HSV),
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
