package momentum_test

import (
	"testing"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/indicators/testutil"
)

func TestMOM(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{1, 2, 4, 7, 11})
	out, err := gotal.Mom(ts, 2)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "mom"), []float64{0, 0, 3, 5, 7}, 1e-9)
}

func TestROC(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{100, 110, 121, 133.1})
	out, err := gotal.Roc(ts, 1)
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "roc"), []float64{0, 10, 10, 10}, 1e-9)
}

func TestROCP(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{100, 110, 121})
	out, err := gotal.ROCP(ts, gotal.With("period", 1))
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "rocp"), []float64{0, 0.1, 0.1}, 1e-9)
}

func TestROCR(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{100, 110, 121})
	out, err := gotal.ROCR(ts, gotal.With("period", 1))
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "rocr"), []float64{0, 1.1, 1.1}, 1e-9)
}

func TestROCR100(t *testing.T) {
	ts := testutil.MakeSeries(t, "p", []float64{100, 110, 121})
	out, err := gotal.ROCR100(ts, gotal.With("period", 1))
	if err != nil {
		t.Fatal(err)
	}
	testutil.AssertSlicesClose(t, testutil.Outputs(t, out, "rocr100"), []float64{0, 110, 110}, 1e-9)
}

func TestWILLR(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1, 1, 1, 1},
		[]float64{10, 12, 11, 13, 12},
		[]float64{8, 9, 10, 11, 11},
		[]float64{9, 11, 10, 13, 11},
		nil,
	)
	out, err := gotal.Willr(ts, 3)
	if err != nil {
		t.Fatal(err)
	}
	got := testutil.Outputs(t, out, "willr")
	if !testutil.NearlyEqual(got[2], -50, 1e-9) {
		t.Errorf("WILLR[2] = %v, want -50", got[2])
	}
}

func TestRSI(t *testing.T) {
	closes := make([]float64, 30)
	for i := range closes {
		closes[i] = 100 + float64(i)
	}
	ts := testutil.MakeSeries(t, "p", closes)
	out, err := gotal.Rsi(ts, 14)
	if err != nil {
		t.Fatal(err)
	}
	got := testutil.Outputs(t, out, "rsi")
	if !testutil.NearlyEqual(got[29], 100, 1e-6) {
		t.Errorf("RSI on monotonic rise should be 100, got %v", got[29])
	}
}

func TestCCI(t *testing.T) {
	ts := testutil.MakeOHLCV(t, "p",
		[]float64{1, 1, 1, 1, 1, 1, 1},
		[]float64{10, 11, 12, 13, 14, 15, 16},
		[]float64{8, 9, 10, 11, 12, 13, 14},
		[]float64{9, 10, 11, 12, 13, 14, 15},
		nil,
	)
	out, err := gotal.Cci(ts, 3)
	if err != nil {
		t.Fatal(err)
	}
	if len(testutil.Outputs(t, out, "cci")) != 7 {
		t.Fatal("CCI length wrong")
	}
}

func TestMACD(t *testing.T) {
	closes := make([]float64, 60)
	for i := range closes {
		closes[i] = 100 + float64(i)
	}
	ts := testutil.MakeSeries(t, "p", closes)
	out, err := gotal.MACD(ts)
	if err != nil {
		t.Fatal(err)
	}
	line := testutil.Outputs(t, out, "macd")
	if line[50] <= 0 {
		t.Errorf("MACD line should be positive on uptrend, got %v", line[50])
	}
}

func TestAPO(t *testing.T) {
	closes := make([]float64, 40)
	for i := range closes {
		closes[i] = 100 + float64(i)
	}
	ts := testutil.MakeSeries(t, "p", closes)
	out, err := gotal.APO(ts)
	if err != nil {
		t.Fatal(err)
	}
	if testutil.Outputs(t, out, "apo")[39] <= 0 {
		t.Errorf("APO should be positive on uptrend")
	}
}

func TestPPO(t *testing.T) {
	closes := make([]float64, 40)
	for i := range closes {
		closes[i] = 100 + float64(i)
	}
	ts := testutil.MakeSeries(t, "p", closes)
	out, err := gotal.PPO(ts)
	if err != nil {
		t.Fatal(err)
	}
	if testutil.Outputs(t, out, "ppo")[39] <= 0 {
		t.Errorf("PPO should be positive on uptrend")
	}
}

func TestCMO(t *testing.T) {
	closes := []float64{100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115}
	ts := testutil.MakeSeries(t, "p", closes)
	out, err := gotal.CMO(ts, gotal.With("period", 14))
	if err != nil {
		t.Fatal(err)
	}
	got := testutil.Outputs(t, out, "cmo")
	if !testutil.NearlyEqual(got[15], 100, 1e-9) {
		t.Errorf("CMO on monotonic rise = %v, want 100", got[15])
	}
}

func TestTRIX(t *testing.T) {
	closes := make([]float64, 50)
	for i := range closes {
		closes[i] = 100 + float64(i)
	}
	ts := testutil.MakeSeries(t, "p", closes)
	out, err := gotal.TRIX(ts, gotal.With("period", 14))
	if err != nil {
		t.Fatal(err)
	}
	if testutil.Outputs(t, out, "trix")[49] <= 0 {
		t.Errorf("TRIX should be positive on uptrend")
	}
}

func TestAROON(t *testing.T) {
	highs := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	lows := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	ts := testutil.MakeOHLCV(t, "p", lows, highs, lows, highs, nil)
	out, err := gotal.AROON(ts, gotal.With("period", 14))
	if err != nil {
		t.Fatal(err)
	}
	up := testutil.Outputs(t, out, "aroon_up")
	dn := testutil.Outputs(t, out, "aroon_down")
	if !testutil.NearlyEqual(up[15], 100, 1e-9) {
		t.Errorf("Aroon up on monotonic rise = %v, want 100", up[15])
	}
	if up[15] <= dn[15] {
		t.Errorf("AROON_UP should exceed AROON_DOWN on uptrend")
	}
}

func TestAROONOSC(t *testing.T) {
	highs := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	lows := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	ts := testutil.MakeOHLCV(t, "p", lows, highs, lows, highs, nil)
	out, err := gotal.AROONOSC(ts, gotal.With("period", 14))
	if err != nil {
		t.Fatal(err)
	}
	if testutil.Outputs(t, out, "aroonosc")[15] <= 0 {
		t.Errorf("AROONOSC on uptrend should be positive")
	}
}

func TestSTOCH(t *testing.T) {
	highs := []float64{10, 12, 11, 13, 14, 15, 14, 16, 18, 17}
	lows := []float64{8, 9, 9, 10, 11, 12, 11, 13, 15, 14}
	closes := []float64{9, 11, 10, 12, 13, 14, 13, 15, 17, 16}
	ts := testutil.MakeOHLCV(t, "p", lows, highs, lows, closes, nil)
	out, err := gotal.STOCH(ts)
	if err != nil {
		t.Fatal(err)
	}
	if len(testutil.Outputs(t, out, "stoch_k")) != 10 {
		t.Fatal("STOCH length wrong")
	}
}
