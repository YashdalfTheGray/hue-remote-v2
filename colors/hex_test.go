package colors_test

import (
	"testing"

	"github.com/yashdalfthegray/hue-remote-v2/colors"
)

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
