package volatility_test

import (
	"testing"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/indicators/testutil"
)

func TestTRANGE(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1, 1, 1},
		[]float64{10, 12, 11, 15},
		[]float64{8, 9, 8, 12},
		[]float64{9, 11, 10, 13},
		nil,
	)
	out, err := gotal.Trange(ts)
	if err != nil {
		t.Fatal(err)
	}
	want := []float64{2, 3, 3, 5}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "trange"), want, 1e-9)
}

func TestATR(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1, 1, 1, 1, 1},
		[]float64{10, 12, 11, 15, 14, 16},
		[]float64{8, 9, 8, 12, 11, 13},
		[]float64{9, 11, 10, 13, 12, 15},
		nil,
	)
	out, err := gotal.Atr(ts, 3)
	if err != nil {
		t.Fatal(err)
	}
	if testutil.Outputs(t, out, "atr")[2] <= 0 {
		t.Errorf("ATR[2] should be positive")
	}
}

func TestNATR(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1, 1, 1, 1, 1},
		[]float64{10, 12, 11, 15, 14, 16},
		[]float64{8, 9, 8, 12, 11, 13},
		[]float64{9, 11, 10, 13, 12, 15},
		nil,
	)
	out, err := gotal.Natr(ts, 3)
	if err != nil {
		t.Fatal(err)
	}
	if testutil.Outputs(t, out, "natr")[2] <= 0 {
		t.Errorf("NATR[2] should be positive")
	}
}
