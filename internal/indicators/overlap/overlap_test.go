package overlap_test

import (
	"testing"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/indicators/testutil"
)

func TestWMA(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{1, 2, 3, 4, 5})
	out, err := gotal.Wma(ts, 3)
	if err != nil {
		t.Fatal(err)
	}
	// denom = 3*4/2 = 6
	want := []float64{0, 0, 14.0 / 6, 20.0 / 6, 26.0 / 6}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "wma"), want, 1e-9)
}

func TestDEMA(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{1, 2, 3, 4, 5, 6, 7, 8})
	out, err := gotal.Dema(ts, 3)
	if err != nil {
		t.Fatal(err)
	}
	got := testutil.Outputs(t, out, "dema")
	if got[7] < 6 || got[7] > 9 {
		t.Errorf("DEMA[7] = %v, expected near 8", got[7])
	}
}

func TestTEMA(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	out, err := gotal.Tema(ts, 3)
	if err != nil {
		t.Fatal(err)
	}
	got := testutil.Outputs(t, out, "tema")
	if got[9] < 8 || got[9] > 11 {
		t.Errorf("TEMA[9] = %v, expected near 10", got[9])
	}
}

func TestTRIMA(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{1, 2, 3, 4, 5, 6, 7, 8})
	out, err := gotal.Trima(ts, 3)
	if err != nil {
		t.Fatal(err)
	}
	if len(testutil.Outputs(t, out, "trima")) != 8 {
		t.Fatal("trima length wrong")
	}
}

func TestMIDPOINT(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{1, 5, 3, 8, 2})
	out, err := gotal.Midpoint(ts, 3)
	if err != nil {
		t.Fatal(err)
	}
	want := []float64{0, 0, 3, 5.5, 5}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "midpoint"), want, 1e-9)
}

func TestMIDPRICE(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1, 1, 1, 1},
		[]float64{2, 4, 3, 7, 5},
		[]float64{1, 2, 1, 3, 2},
		[]float64{1.5, 3, 2, 5, 4},
		nil,
	)
	out, err := gotal.Midprice(ts, 3)
	if err != nil {
		t.Fatal(err)
	}
	want := []float64{0, 0, 2.5, 4, 4}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "midprice"), want, 1e-9)
}

func TestBBANDS(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{2, 4, 4, 4, 5, 5, 7, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21})
	out, err := gotal.BBANDS(ts, gotal.With("period", 5), gotal.With("nbdevup", 2.0), gotal.With("nbdevdn", 2.0))
	if err != nil {
		t.Fatal(err)
	}
	upper := testutil.Outputs(t, out, "bbands_upper")
	mid := testutil.Outputs(t, out, "bbands_middle")
	lower := testutil.Outputs(t, out, "bbands_lower")
	for j := 10; j < 20; j++ {
		if upper[j] < mid[j] || mid[j] < lower[j] {
			t.Errorf("BBANDS band ordering broken at %d: U=%v M=%v L=%v", j, upper[j], mid[j], lower[j])
		}
	}
}

func TestHMA(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25})
	out, err := gotal.Hma(ts, 9)
	if err != nil {
		t.Fatal(err)
	}
	if len(testutil.Outputs(t, out, "hma")) != 16 {
		t.Fatal("HMA length wrong")
	}
}
