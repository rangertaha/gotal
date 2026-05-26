// Package testutil provides fixture builders shared by every indicator's
// _test.go file. It depends on the public gotal/ticker packages, so it can
// only be imported from test files (the indicator packages themselves must
// not import it, to avoid cycles).
package testutil

import (
	"math"
	"testing"
	"time"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/tick"
	"github.com/rangertaha/gotal/ticker"
)

// MakeSeries builds a *ticker.Ticker from a slice of close prices, one per day.
func MakeSeries(t testing.TB, name string, closes []float64) *ticker.Ticker {
	t.Helper()
	base := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	ticks := make([]gotal.Tick, len(closes))
	for i, c := range closes {
		ticks[i] = tick.NewTick(
			base.AddDate(0, 0, i),
			24*time.Hour,
			map[string]float64{"close": c},
			nil, nil,
		)
	}
	ts, err := ticker.New(name, ticker.WithTicks(ticks...))
	if err != nil {
		t.Fatalf("MakeSeries: %v", err)
	}
	return ts
}

// MakeOHLCV builds a Ticker from parallel OHLCV slices. Pass nil volume to
// skip the volume field.
func MakeOHLCV(t testing.TB, name string, open, high, low, closes, volume []float64) *ticker.Ticker {
	t.Helper()
	base := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	n := len(closes)
	ticks := make([]gotal.Tick, n)
	for i := 0; i < n; i++ {
		fields := map[string]float64{
			"open":  open[i],
			"high":  high[i],
			"low":   low[i],
			"close": closes[i],
		}
		if volume != nil {
			fields["volume"] = volume[i]
		}
		ticks[i] = tick.NewTick(base.AddDate(0, 0, i), 24*time.Hour, fields, nil, nil)
	}
	ts, err := ticker.New(name, ticker.WithTicks(ticks...))
	if err != nil {
		t.Fatalf("MakeOHLCV: %v", err)
	}
	return ts
}

// Outputs returns the values vector for the named computed field, or fails the test.
func Outputs(t testing.TB, ts gotal.TimeSeries, field string) []float64 {
	t.Helper()
	v := ts.Fields().Get(field)
	if v == nil {
		t.Fatalf("field %q not found on series %q", field, ts.Name())
	}
	return v.Values()
}

// NearlyEqual reports whether |a-b| <= eps. NaN inputs are never equal.
func NearlyEqual(a, b, eps float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) {
		return false
	}
	return math.Abs(a-b) <= eps
}

// AssertSlicesClose fails the test if got and want differ by more than eps at any index.
func AssertSlicesClose(t testing.TB, got, want []float64, eps float64) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %d want %d", len(got), len(want))
	}
	for i := range got {
		if !NearlyEqual(got[i], want[i], eps) {
			t.Errorf("index %d: got %.6f want %.6f", i, got[i], want[i])
		}
	}
}
