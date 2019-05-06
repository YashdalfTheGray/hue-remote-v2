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
