package examples

import (
	"math"
	"time"

	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/tick"
)

var Tags = map[string]string{
	"symbol":   "BTCUSDT",
	"exchange": "binance",
	"currency": "USD",
	"asset":    "BTC",
}

// Prices creates a historical price dataset
// It's used for testing and examples.
func PricesSeries(count int, duration time.Duration) *series.Series {
	ticks := series.New("prices")

	t := time.Now()
	price := 0.0
	for i := 0; i < count; i++ {
		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"value": price}),
			tick.WithTags(Tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
		price++
	}

	return ticks
}

func SineSeries(duration time.Duration, amplitude float64, frequency float64, offset float64) *series.Series {
	ticks := series.New("sine")
	t := time.Now()

	// Apply the offset to the starting time
	t = t.Add(time.Duration(offset) * duration)

	// Generate one complete sine wave cycle (2π radians)
	for i := 0.0; i <= 2*math.Pi; i += frequency {
		value := (amplitude * math.Sin(i)) + amplitude
		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(Tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}
