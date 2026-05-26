package gotal_test

import (
	"testing"
	"time"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/tick"
	"github.com/rangertaha/gotal/ticker"
)

func ohlcvFixture(t *testing.T) *ticker.Ticker {
	t.Helper()
	highs := []float64{102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131}
	lows := []float64{99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128}
	opens := []float64{100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129}
	closes := []float64{101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130}
	vols := []float64{1000, 1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400, 2500, 2600, 2700, 2800, 2900, 3000, 3100, 3200, 3300, 3400, 3500, 3600, 3700, 3800, 3900}

	base := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	ticks := make([]gotal.Tick, len(closes))
	for i := range closes {
		ticks[i] = tick.NewTick(base.AddDate(0, 0, i), 24*time.Hour, map[string]float64{
			"open": opens[i], "high": highs[i], "low": lows[i], "close": closes[i], "volume": vols[i],
		}, nil, nil)
	}
	ts, err := ticker.New("ohlcv", ticker.WithTicks(ticks...))
	if err != nil {
		t.Fatal(err)
	}
	return ts
}

func TestOverlap(t *testing.T) {
	out := gotal.Overlap(ohlcvFixture(t))
	// SMA, EMA, WMA, DEMA, TEMA, TRIMA, MIDPOINT, MIDPRICE, BBANDS_*, HMA — at least these output fields should appear.
	for _, name := range []string{"sma", "ema", "wma", "dema", "tema", "trima", "midpoint", "midprice", "bbands_middle", "hma"} {
		if out.Fields().Get(name) == nil {
			t.Errorf("Overlap: missing field %q", name)
		}
	}
}

func TestMomentum(t *testing.T) {
	out := gotal.Momentum(ohlcvFixture(t))
	for _, name := range []string{"mom", "roc", "rsi", "willr", "cci", "macd", "apo", "ppo", "cmo", "trix", "aroon_up", "aroon_down", "aroonosc", "stoch_k", "mfi"} {
		if out.Fields().Get(name) == nil {
			t.Errorf("Momentum: missing field %q", name)
		}
	}
}

func TestVolume(t *testing.T) {
	out := gotal.Volume(ohlcvFixture(t))
	for _, name := range []string{"obv", "ad", "adosc", "vwma", "vwap"} {
		if out.Fields().Get(name) == nil {
			t.Errorf("Volume: missing field %q", name)
		}
	}
}

func TestVolatility(t *testing.T) {
	out := gotal.Volatility(ohlcvFixture(t))
	for _, name := range []string{"trange", "atr", "natr"} {
		if out.Fields().Get(name) == nil {
			t.Errorf("Volatility: missing field %q", name)
		}
	}
}

func TestPrice(t *testing.T) {
	out := gotal.Price(ohlcvFixture(t))
	for _, name := range []string{"avgprice", "medprice", "typprice", "wclprice", "hlc3", "ohlc4"} {
		if out.Fields().Get(name) == nil {
			t.Errorf("Price: missing field %q", name)
		}
	}
}

func TestStatistic(t *testing.T) {
	out := gotal.Statistic(ohlcvFixture(t))
	for _, name := range []string{"stddev", "variance", "linearreg", "linearreg_slope", "tsf"} {
		if out.Fields().Get(name) == nil {
			t.Errorf("Statistic: missing field %q", name)
		}
	}
}

func TestMathOps(t *testing.T) {
	out := gotal.MathOps(ohlcvFixture(t))
	for _, name := range []string{"max", "min", "sumwindow", "add", "sub", "mult", "div", "maxindex", "minindex"} {
		if out.Fields().Get(name) == nil {
			t.Errorf("MathOps: missing field %q", name)
		}
	}
}

func TestMathTransforms(t *testing.T) {
	out := gotal.MathTransforms(ohlcvFixture(t))
	for _, name := range []string{"cos", "sin", "tan", "sqrt", "exp", "ln", "log10", "ceil", "floor"} {
		if out.Fields().Get(name) == nil {
			t.Errorf("MathTransforms: missing field %q", name)
		}
	}
}

func TestRunAll(t *testing.T) {
	out := gotal.RunAll(ohlcvFixture(t))
	// RunAll should produce >= 50 output fields from all the real implementations.
	got := len(out.Fields().Names())
	if got < 50 {
		t.Errorf("RunAll produced only %d fields, expected at least 50", got)
	}
}

func TestGroups(t *testing.T) {
	groups := gotal.Groups()
	if len(groups) < 10 {
		t.Errorf("Groups() returned %d entries, expected at least 10", len(groups))
	}
	// Each group should have at least one registered indicator (real or stub).
	for _, g := range groups {
		if g == "trend" || g == "other" {
			continue // these are catch-all / unused
		}
		ids := gotal.Group(g)
		if len(ids) == 0 {
			t.Errorf("group %q has no registered indicators", g)
		}
	}
}
