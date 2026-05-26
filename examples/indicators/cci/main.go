// Example: CCI
package main

import (
	"log"
	"time"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/tick"
	"github.com/rangertaha/gotal/ticker"
)

func main() {
	opens := []float64{100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119}
	highs := []float64{102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121}
	lows := []float64{99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118}
	closes := []float64{101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120}

	base := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	ticks := make([]gotal.Tick, len(closes))
	for i := range closes {
		ticks[i] = tick.NewTick(base.AddDate(0, 0, i), 24*time.Hour, map[string]float64{
			"open": opens[i], "high": highs[i], "low": lows[i], "close": closes[i],
		}, nil, nil)
	}
	ts, err := ticker.New("ohlc", ticker.WithTicks(ticks...))
	if err != nil {
		log.Fatal(err)
	}

	out, err := gotal.Cci(ts, 14)
	if err != nil {
		log.Fatal(err)
	}
	out.Print()
}
