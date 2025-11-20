// This example shows how to use the SMA indicator.
// It creates a time series of prices and then calculates the SMA.
// It then prints the result.

// Example:
// go run examples/indicators/ma/main.go

// Output:
//
// $\sqrt{3x-1}+(1+x)^2$

package main

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal/examples"
	. "github.com/rangertaha/gotal/pkg/indicators"
)

func main() {
	// Create a new series of price ticks
	fmt.Printf("\nCreate a new series of price ticks\n\n")
	prices := examples.PricesSeries(1000, time.Second)
	prices.Print()

	// // Convert price series to OHLC series
	fmt.Printf("\nConvert price series to OHLCV series\n\n")
	ohlcv := OHLCV(prices, WithPeriod(25), OnField("price"))
	ohlcv.Print()


	// Create a new series of sma 3
	fmt.Printf("\nCreate a new series of sma with period 3\n\n")
	sma := SMA(ohlcv, WithPeriod(3), OnField("close"))
	sma.Print()

	// Create a new series of ema 3
	fmt.Printf("\nCreate a new series of ema with period 3\n\n")
	ema := EMA(ohlcv, WithPeriod(3), OnField("close"))
	ema.Print()

	// Create a new series of wma 3
	fmt.Printf("\nCreate a new series of wma with period 3\n\n")
	wma := WMA(ohlcv, WithPeriod(3), OnField("close"))
	wma.Print()
}
