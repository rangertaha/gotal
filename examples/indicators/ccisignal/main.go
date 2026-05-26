// Example: ccisignal signal — emits +100 (buy), -100 (sell), or 0 (hold).
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/tick"
	"github.com/rangertaha/gotal/ticker"
)

func main() {
	// A rise-then-crash series produces oversold/overbought crossings.
	highs := []float64{102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 131, 122, 112, 102, 97, 94, 92, 90, 88, 86, 84, 82, 80, 78, 76, 74}
	lows := []float64{99, 101, 103, 105, 107, 109, 111, 113, 115, 117, 119, 121, 123, 125, 127, 129, 128, 119, 109, 99, 94, 91, 89, 87, 85, 83, 81, 79, 77, 75, 73, 71}
	opens := []float64{100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 130, 121, 111, 101, 96, 93, 91, 89, 87, 85, 83, 81, 79, 77, 75, 73}
	closes := []float64{101, 103, 105, 107, 109, 111, 113, 115, 117, 119, 121, 123, 125, 127, 129, 131, 130, 121, 111, 101, 96, 93, 91, 89, 87, 85, 83, 81, 79, 77, 75, 73}
	vols := []float64{1000, 1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400, 2500, 2400, 2300, 2200, 2100, 2000, 1900, 1800, 1700, 1600, 1500, 1400, 1300, 1200, 1100, 1000, 900}

	base := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	ticks := make([]gotal.Tick, len(closes))
	for i := range closes {
		ticks[i] = tick.NewTick(base.AddDate(0, 0, i), 24*time.Hour, map[string]float64{
			"open": opens[i], "high": highs[i], "low": lows[i], "close": closes[i], "volume": vols[i],
		}, nil, nil)
	}
	ts, err := ticker.New("ohlcv", ticker.WithTicks(ticks...))
	if err != nil {
		log.Fatal(err)
	}

	out, err := gotal.Get("CCISIGNAL")(ts)
	if err != nil {
		log.Fatal(err)
	}

	sig := out.Fields().Get("ccisignal").Values()
	var buys, sells int
	for i, v := range sig {
		switch v {
		case 100:
			buys++
			fmt.Printf("  BUY  at tick %d (close=%.2f)\n", i, closes[i])
		case -100:
			sells++
			fmt.Printf("  SELL at tick %d (close=%.2f)\n", i, closes[i])
		}
	}
	fmt.Printf("# ccisignal: %d buys, %d sells across %d ticks\n", buys, sells, len(sig))
}
