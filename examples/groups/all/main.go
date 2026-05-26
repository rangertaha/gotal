// Example: gotal.RunAll runs every implemented all indicator against the
// input TimeSeries and attaches each indicator's output as a named field.
package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/tick"
	"github.com/rangertaha/gotal/ticker"
)

func main() {
	highs := []float64{102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141}
	lows := []float64{99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138}
	opens := []float64{100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139}
	closes := []float64{101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140}
	vols := []float64{1000, 1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400, 2500, 2600, 2700, 2800, 2900, 3000, 3100, 3200, 3300, 3400, 3500, 3600, 3700, 3800, 3900, 4000, 4100, 4200, 4300, 4400, 4500, 4600, 4700, 4800, 4900}

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

	out := gotal.RunAll(ts)

	names := out.Fields().Names()
	sort.Strings(names)
	fmt.Printf("# %s ran on %d ticks — %d output fields:\n", "all", out.Len(), len(names))
	for _, name := range names {
		fmt.Println(" ", name)
	}
}
