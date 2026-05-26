// Example: ATAN
package main

import (
	"log"
	"time"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/tick"
	"github.com/rangertaha/gotal/ticker"
)

func main() {
	closes := []float64{100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119}
	base := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	ticks := make([]gotal.Tick, len(closes))
	for i, c := range closes {
		ticks[i] = tick.NewTick(base.AddDate(0, 0, i), 24*time.Hour, map[string]float64{"close": c}, nil, nil)
	}
	ts, err := ticker.New("prices", ticker.WithTicks(ticks...))
	if err != nil {
		log.Fatal(err)
	}

	out, err := gotal.ATAN(ts)
	if err != nil {
		log.Fatal(err)
	}
	out.Print()
}
