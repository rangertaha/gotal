package mathtransform_test

import (
	"math"
	"testing"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/indicators/testutil"
)

type unaryCase struct {
	name string
	fn   gotal.IndicatorFunc
	out  string
	in   []float64
	want []float64
}

func TestMathTransforms(t *testing.T) {
	cases := []unaryCase{
		{"COS", gotal.COS, "cos", []float64{0, math.Pi}, []float64{1, -1}},
		{"SIN", gotal.SIN, "sin", []float64{0, math.Pi / 2}, []float64{0, 1}},
		{"TAN", gotal.TAN, "tan", []float64{0, math.Pi / 4}, []float64{0, 1}},
		{"COSH", gotal.COSH, "cosh", []float64{0}, []float64{1}},
		{"SINH", gotal.SINH, "sinh", []float64{0}, []float64{0}},
		{"TANH", gotal.TANH, "tanh", []float64{0}, []float64{0}},
		{"ACOS", gotal.ACOS, "acos", []float64{1, -1}, []float64{0, math.Pi}},
		{"ASIN", gotal.ASIN, "asin", []float64{0, 1}, []float64{0, math.Pi / 2}},
		{"ATAN", gotal.ATAN, "atan", []float64{0, 1}, []float64{0, math.Pi / 4}},
		{"CEIL", gotal.CEIL, "ceil", []float64{1.2, 2.7, -1.1}, []float64{2, 3, -1}},
		{"FLOOR", gotal.FLOOR, "floor", []float64{1.2, 2.7, -1.1}, []float64{1, 2, -2}},
		{"EXP", gotal.EXP, "exp", []float64{0, 1}, []float64{1, math.E}},
		{"LN", gotal.LN, "ln", []float64{1, math.E}, []float64{0, 1}},
		{"LOG10", gotal.LOG10, "log10", []float64{1, 10, 100}, []float64{0, 1, 2}},
		{"SQRT", gotal.SQRT, "sqrt", []float64{0, 4, 9}, []float64{0, 2, 3}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ts := testutil.MakeSeries(t, "p", tc.in)
			out, err := tc.fn(ts)
			if err != nil {
				t.Fatal(err)
			}
			testutil.AssertSlicesClose(t, testutil.Outputs(t, out, tc.out), tc.want, 1e-9)
		})
	}
}
