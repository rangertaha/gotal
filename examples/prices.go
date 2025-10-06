package examples

import (
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
			tick.WithFields(map[string]float64{"price": price}),
			tick.WithTags(Tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
		price++
	}

	return ticks
}
