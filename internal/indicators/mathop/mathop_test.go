package mathop_test

import (
	"testing"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/indicators/testutil"
)

func TestMAX(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{1, 5, 3, 8, 2})
	out, err := gotal.MAX(ts, gotal.With("period", 3))
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "max"), []float64{0, 0, 5, 8, 8}, 1e-9)
}

func TestMIN(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{1, 5, 3, 8, 2})
	out, err := gotal.MIN(ts, gotal.With("period", 3))
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "min"), []float64{0, 0, 1, 3, 2}, 1e-9)
}

func TestSUMWINDOW(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{1, 2, 3, 4, 5})
	out, err := gotal.SUMWINDOW(ts, gotal.With("period", 3))
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "sumwindow"), []float64{0, 0, 6, 9, 12}, 1e-9)
}

func TestADD(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1, 1}, []float64{10, 12, 14},
		[]float64{8, 10, 12}, []float64{9, 11, 13}, nil)
	out, err := gotal.ADD(ts, gotal.With("source1", "high"), gotal.With("source2", "low"))
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "add"), []float64{18, 22, 26}, 1e-9)
}

func TestSUB(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1, 1}, []float64{10, 12, 14},
		[]float64{8, 10, 12}, []float64{9, 11, 13}, nil)
	out, err := gotal.SUB(ts, gotal.With("source1", "high"), gotal.With("source2", "low"))
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "sub"), []float64{2, 2, 2}, 1e-9)
}

func TestMULT(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1}, []float64{10, 20}, []float64{2, 5}, []float64{5, 10}, nil)
	out, err := gotal.MULT(ts, gotal.With("source1", "high"), gotal.With("source2", "low"))
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "mult"), []float64{20, 100}, 1e-9)
}

func TestDIV(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1}, []float64{10, 20}, []float64{2, 5}, []float64{5, 10}, nil)
	out, err := gotal.DIV(ts, gotal.With("source1", "high"), gotal.With("source2", "low"))
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "div"), []float64{5, 4}, 1e-9)
}

func TestMAXINDEX(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{1, 5, 3, 8, 2})
	out, err := gotal.MAXINDEX(ts, gotal.With("period", 3))
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "maxindex"), []float64{0, 0, 1, 3, 3}, 1e-9)
}

func TestMININDEX(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{1, 5, 3, 8, 2})
	out, err := gotal.MININDEX(ts, gotal.With("period", 3))
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "minindex"), []float64{0, 0, 0, 2, 4}, 1e-9)
}
