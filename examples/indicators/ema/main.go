// Example: compute EMA over a batch CSV of prices.
package main

import (
	"log"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/io"
)

func main() {
	ts, err := io.Read("prices.csv")
	if err != nil {
		log.Fatal(err)
	}

	ema, err := gotal.EMA(ts,
		gotal.With("name", "ema"),
		gotal.With("source", "close"),
		gotal.With("period", 5),
	)
	if err != nil {
		log.Fatal(err)
	}

	ema.Print()

	if err := ema.Save("ema.csv"); err != nil {
		log.Fatal(err)
	}
}
