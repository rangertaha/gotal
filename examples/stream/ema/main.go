// Example: pipe a CSV of prices through a streaming EMA.
package main

import (
	"fmt"

	"github.com/rangertaha/gotal/io"
	"github.com/rangertaha/gotal/stream"
)

func main() {
	in := io.CSV("../../indicators/ema/prices.csv").Stream().(*stream.Stream)
	out := stream.Ema(in, 5, 0)

	for tick := range out.Ticks() {
		fields, _ := tick.Fields()
		signals, _ := tick.Signals()
		fmt.Printf("%s close=%.2f signals=%v\n",
			tick.Time().UTC().Format("2006-01-02"),
			fields["close"],
			signals,
		)
	}
}
