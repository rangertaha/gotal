package signal_test

import (
	"testing"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/indicators/testutil"
)

func TestRSISIGNAL(t *testing.T) {
	// Series: builds up gradually then crashes — should yield at least one
	// overbought (-100) and one oversold (+100) crossing event.
	closes := []float64{
		// climb
		100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130,
		// crash
		129, 120, 110, 100, 95, 92, 90, 88, 86, 84, 82, 80, 78, 76, 74, 72,
	}
	ts := testutil.MakeSeries(t, "p", closes)
	out, err := gotal.Get("RSISIGNAL")(ts)
	if err != nil {
		t.Fatal(err)
	}
	sig := testutil.Outputs(t, out, "rsisignal")
	var sells, buys int
	for _, v := range sig {
		switch v {
		case 100:
			buys++
		case -100:
			sells++
		}
	}
	if sells == 0 && buys == 0 {
		t.Errorf("RSISIGNAL produced no crossings; got %v", sig)
	}
}

func TestMOMSIGNAL(t *testing.T) {
	// Series that crosses zero in momentum (up then down).
	closes := []float64{100, 105, 110, 115, 120, 119, 110, 100, 90, 80, 85, 95, 110}
	ts := testutil.MakeSeries(t, "p", closes)
	out, err := gotal.Get("MOMSIGNAL")(ts, gotal.With("period", 3))
	if err != nil {
		t.Fatal(err)
	}
	sig := testutil.Outputs(t, out, "momsignal")
	var crossings int
	for _, v := range sig {
		if v != 0 {
			crossings++
		}
	}
	if crossings == 0 {
		t.Errorf("MOMSIGNAL expected at least one zero-cross, got %v", sig)
	}
}

func TestMACDCross(t *testing.T) {
	// Long-enough series for MACD line + signal to cross at least once.
	closes := make([]float64, 80)
	for i := 0; i < 40; i++ {
		closes[i] = 100 + float64(i) // uptrend
	}
	for i := 40; i < 80; i++ {
		closes[i] = 140 - float64(i-40) // downtrend
	}
	ts := testutil.MakeSeries(t, "p", closes)
	out, err := gotal.Get("MACDSIGNAL_CROSS")(ts)
	if err != nil {
		t.Fatal(err)
	}
	sig := testutil.Outputs(t, out, "macdcross")
	var crossings int
	for _, v := range sig {
		if v != 0 {
			crossings++
		}
	}
	if crossings == 0 {
		t.Errorf("MACD crossover expected on trend reversal, got none")
	}
}

func TestSignalGroup(t *testing.T) {
	// Same series as MACD test.
	closes := make([]float64, 80)
	for i := 0; i < 40; i++ {
		closes[i] = 100 + float64(i)
	}
	for i := 40; i < 80; i++ {
		closes[i] = 140 - float64(i-40)
	}
	highs := make([]float64, 80)
	lows := make([]float64, 80)
	for i := range closes {
		highs[i] = closes[i] + 1
		lows[i] = closes[i] - 1
	}
	ts := testutil.MakeOHLCV(t, "p", closes, highs, lows, closes, nil)
	out := gotal.Signal(ts)
	// At least these field names should appear.
	for _, name := range []string{"rsisignal", "momsignal", "rocsignal", "trixsignal", "pposignal", "cmosignal", "macdcross"} {
		if out.Fields().Get(name) == nil {
			t.Errorf("gotal.Signal: missing %q", name)
		}
	}
}
