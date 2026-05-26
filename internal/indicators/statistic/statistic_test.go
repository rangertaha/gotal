package statistic_test

import (
	"testing"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/indicators/testutil"
)

func TestSTDDEV(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{2, 4, 4, 4, 5, 5, 7, 9})
	out, err := gotal.Stddev(ts, 8)
	if err != nil {
		t.Fatal(err)
	}
	if !testutil.NearlyEqual(testutil.Outputs(t, out, "stddev")[7], 2, 1e-9) {
		t.Errorf("STDDEV[7] should be 2")
	}
}

func TestVARIANCE(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{2, 4, 4, 4, 5, 5, 7, 9})
	out, err := gotal.Var_(ts, 8)
	if err != nil {
		t.Fatal(err)
	}
	if !testutil.NearlyEqual(testutil.Outputs(t, out, "variance")[7], 4, 1e-9) {
		t.Errorf("VARIANCE[7] should be 4")
	}
}

func TestLINEARREG(t *testing.T) {
	closes := []float64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	ts := testutil.MakeSeries(t, "p", closes)
	out, err := gotal.LINEARREG(ts, gotal.With("period", 5))
	if err != nil {
		t.Fatal(err)
	}
	if !testutil.NearlyEqual(testutil.Outputs(t, out, "linearreg")[9], 19, 1e-9) {
		t.Errorf("LINEARREG[9] should be 19")
	}
}

func TestLINEARREG_SLOPE(t *testing.T) {
	closes := []float64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	ts := testutil.MakeSeries(t, "p", closes)
	out, err := gotal.LINEARREG_SLOPE(ts, gotal.With("period", 5))
	if err != nil {
		t.Fatal(err)
	}
	if !testutil.NearlyEqual(testutil.Outputs(t, out, "linearreg_slope")[9], 1, 1e-9) {
		t.Errorf("slope = 1")
	}
}

func TestTSF(t *testing.T) {
	closes := []float64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	ts := testutil.MakeSeries(t, "p", closes)
	out, err := gotal.TSF(ts, gotal.With("period", 5))
	if err != nil {
		t.Fatal(err)
	}
	if !testutil.NearlyEqual(testutil.Outputs(t, out, "tsf")[9], 20, 1e-9) {
		t.Errorf("TSF[9] should forecast 20")
	}
}

func TestBETA(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 2, 3, 4, 5, 6, 7, 8},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8},
		[]float64{2, 4, 6, 8, 10, 12, 14, 16},
		nil,
	)
	out, err := gotal.BETA(ts, gotal.With("source1", "close"), gotal.With("source2", "open"), gotal.With("period", 5))
	if err != nil {
		t.Fatal(err)
	}
	if !testutil.NearlyEqual(testutil.Outputs(t, out, "beta")[7], 2, 1e-9) {
		t.Errorf("BETA should be 2")
	}
}

func TestCORREL(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 2, 3, 4, 5, 6, 7, 8},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8},
		[]float64{2, 4, 6, 8, 10, 12, 14, 16},
		nil,
	)
	out, err := gotal.CORREL(ts, gotal.With("source1", "close"), gotal.With("source2", "open"), gotal.With("period", 5))
	if err != nil {
		t.Fatal(err)
	}
	if !testutil.NearlyEqual(testutil.Outputs(t, out, "correl")[7], 1, 1e-9) {
		t.Errorf("CORREL should be 1")
	}
}
