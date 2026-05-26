// Example: AD
package main

import (
	"log"
	"time"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/tick"
	"github.com/rangertaha/gotal/ticker"
)

func main() {
	opens := []float64{100, 101, 102, 103, 104, 105, 106, 107, 108, 109}
	highs := []float64{102, 103, 104, 105, 106, 107, 108, 109, 110, 111}
	lows := []float64{99, 100, 101, 102, 103, 104, 105, 106, 107, 108}
	closes := []float64{101, 102, 103, 102, 105, 106, 105, 108, 109, 108}
	volume := []float64{1000, 1200, 900, 1500, 1100, 1300, 1400, 1600, 1700, 1800}

	base := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	ticks := make([]gotal.Tick, len(closes))
	for i := range closes {
		ticks[i] = tick.NewTick(base.AddDate(0, 0, i), 24*time.Hour, map[string]float64{
			"open": opens[i], "high": highs[i], "low": lows[i], "close": closes[i], "volume": volume[i],
		}, nil, nil)
	}
	ts, err := ticker.New("ohlcv", ticker.WithTicks(ticks...))
	if err != nil {
		log.Fatal(err)
	}

	out, err := gotal.Ad(ts)
	if err != nil {
		log.Fatal(err)
	}
	out.Print()
}
