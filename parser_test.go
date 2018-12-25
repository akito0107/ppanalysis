package ppanalysis

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  AnalysisBody
	}{
		{
			name: "simple case",
			in: `{
	"github.com/akito0107/blogplayground/errwrapsample": {
		"mustwrap": [
			{
				"posn": "/Users/akito/go/src/github.com/akito0107/blogplayground/errwrapsample/openfile.go:11:3",
				"message": "should be use errors.Wrap() or errors.Wrapf()"
			}
		]
	}
}`,
			out: map[string]map[string][]AnalysisMessage{
				"github.com/akito0107/blogplayground/errwrapsample": {
					"mustwrap": []AnalysisMessage{
						{
							Posn:    "/Users/akito/go/src/github.com/akito0107/blogplayground/errwrapsample/openfile.go:11:3",
							Message: "should be use errors.Wrap() or errors.Wrapf()",
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			act, err := Parse(bytes.NewBufferString(c.in))
			if err != nil {
				t.Errorf("%s: %+v", c.in, err)
			}

			if !reflect.DeepEqual(act, c.out) {
				t.Errorf("must be same %+v, %+v", act, c.out)
			}

		})
	}
}
