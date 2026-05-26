package ma_test

import (
	"testing"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/indicators/testutil"
)

func TestSMA(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{1, 2, 3, 4, 5})
	out, err := gotal.Sma(ts, 3)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "sma"), []float64{0, 0, 2, 3, 4}, 1e-9)
}

func TestEMA(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{1, 2, 3, 4, 5})
	out, err := gotal.Ema(ts, 3, 0)
	if err != nil {
		t.Fatal(err)
	}
	// alpha = 2/(3+1) = 0.5
	want := []float64{1, 1.5, 2.25, 3.125, 4.0625}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "ema"), want, 1e-9)
}
