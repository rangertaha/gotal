package volume_test

import (
	"testing"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/indicators/testutil"
)

func TestOBV(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0},
		[]float64{10, 11, 10, 12, 12, 11},
		[]float64{100, 150, 200, 250, 50, 300},
	)
	out, err := gotal.Obv(ts)
	if err != nil {
		t.Fatal(err)
	}
	want := []float64{100, 250, 50, 300, 300, 0}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "obv"), want, 1e-9)
}

func TestAD(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1, 1}, []float64{10, 12, 11},
		[]float64{8, 9, 8}, []float64{9, 11, 10},
		[]float64{100, 200, 150})
	out, err := gotal.Ad(ts)
	if err != nil {
		t.Fatal(err)
	}
	if len(testutil.Outputs(t, out, "ad")) != 3 {
		t.Fatal("AD length wrong")
	}
}

func TestADOSC(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		[]float64{10, 12, 11, 13, 14, 15, 14, 16, 18, 17, 19, 20},
		[]float64{8, 9, 9, 10, 11, 12, 11, 13, 15, 14, 16, 17},
		[]float64{9, 11, 10, 12, 13, 14, 13, 15, 17, 16, 18, 19},
		[]float64{100, 110, 90, 120, 130, 140, 130, 150, 170, 160, 180, 190},
	)
	out, err := gotal.ADOSC(ts)
	if err != nil {
		t.Fatal(err)
	}
	if len(testutil.Outputs(t, out, "adosc")) != 12 {
		t.Fatal("ADOSC length wrong")
	}
}

func TestMFI(t *testing.T) {
	highs := []float64{10, 11, 12, 13, 14, 15, 14, 16, 18, 17, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28}
	lows := []float64{8, 9, 9, 10, 11, 12, 11, 13, 15, 14, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	closes := []float64{9, 11, 10, 12, 13, 14, 13, 15, 17, 16, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}
	vols := []float64{100, 110, 90, 120, 130, 140, 130, 150, 170, 160, 180, 190, 200, 210, 220, 230, 240, 250, 260, 270}
	ts := testutil.MakeOHLCV(t, "p", lows, highs, lows, closes, vols)
	out, err := gotal.MFI(ts, gotal.With("period", 14))
	if err != nil {
		t.Fatal(err)
	}
	if testutil.Outputs(t, out, "mfi")[19] < 50 {
		t.Errorf("MFI on uptrend should be > 50")
	}
}

func TestVWMA(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1, 1, 1, 1},
		[]float64{1, 1, 1, 1, 1},
		[]float64{1, 1, 1, 1, 1},
		[]float64{10, 20, 30, 40, 50},
		[]float64{1, 1, 1, 1, 1},
	)
	out, err := gotal.Get("VWMA")(ts, gotal.With("period", 3))
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "vwma"), []float64{0, 0, 20, 30, 40}, 1e-9)
}

func TestVWAP(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1, 1}, []float64{10, 12, 14},
		[]float64{8, 10, 12}, []float64{9, 11, 13},
		[]float64{100, 100, 100})
	out, err := gotal.Vwap(ts)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "vwap"), []float64{9, 10, 11}, 1e-9)
}
