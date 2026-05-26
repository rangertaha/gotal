package price_test

import (
	"testing"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/indicators/testutil"
)

func TestAVGPRICE(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{10, 20, 30}, []float64{12, 22, 32},
		[]float64{8, 18, 28}, []float64{11, 21, 31}, nil)
	out, err := gotal.Avgprice(ts)
	if err != nil {
		t.Fatal(err)
	}
	want := []float64{(10 + 12 + 8 + 11) / 4.0, (20 + 22 + 18 + 21) / 4.0, (30 + 32 + 28 + 31) / 4.0}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "avgprice"), want, 1e-9)
}

func TestMEDPRICE(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1}, []float64{12, 22}, []float64{8, 18}, []float64{10, 20}, nil)
	out, err := gotal.Medprice(ts)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "medprice"), []float64{10, 20}, 1e-9)
}

func TestTYPPRICE(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1}, []float64{12, 22}, []float64{8, 18}, []float64{10, 20}, nil)
	out, err := gotal.Typprice(ts)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "typprice"), []float64{10, 20}, 1e-9)
}

func TestWCLPRICE(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1}, []float64{12, 22}, []float64{8, 18}, []float64{10, 20}, nil)
	out, err := gotal.Wclprice(ts)
	if err != nil {
		t.Fatal(err)
	}
	want := []float64{(12 + 8 + 2*10) / 4.0, (22 + 18 + 2*20) / 4.0}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "wclprice"), want, 1e-9)
}

func TestHLC3(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1}, []float64{12, 22}, []float64{8, 18}, []float64{10, 20}, nil)
	out, err := gotal.Hlc3(ts)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "hlc3"), []float64{10, 20}, 1e-9)
}

func TestOHLC4(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{2, 4}, []float64{12, 22}, []float64{8, 18}, []float64{10, 20}, nil)
	out, err := gotal.Ohlc4(ts)
	if err != nil {
		t.Fatal(err)
	}
	want := []float64{(2 + 12 + 8 + 10) / 4.0, (4 + 22 + 18 + 20) / 4.0}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "ohlc4"), want, 1e-9)
}

func TestHEIKINASHI(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{10, 11, 12}, []float64{12, 13, 14},
		[]float64{9, 10, 11}, []float64{11, 12, 13}, nil)
	out, err := gotal.Heikinashi(ts)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range []string{"ha_open", "ha_high", "ha_low", "ha_close"} {
		if len(testutil.Outputs(t, out, f)) != 3 {
			t.Fatalf("HA %s length wrong", f)
		}
	}
}
