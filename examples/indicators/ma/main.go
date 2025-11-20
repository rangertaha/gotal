// This example shows how to use the SMA indicator.
// It creates a time series of prices and then calculates the SMA.
// It then prints the result.

// Example:
// go run examples/indicators/ma/main.go

// Output:
//
//

package main

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal/examples"
	ind "github.com/rangertaha/gotal/pkg/indicators"
)

func main() {
	// Create a new series of price ticks
	fmt.Printf("\nCreate a new series of price ticks\n\n")
	prices := examples.PricesSeries(1000, time.Second)
	prices.Print()

	// // Convert price series to OHLC series
	fmt.Printf("\nConvert price series to OHLCV series\n\n")
	ohlcv := ind.OHLCV(prices, ind.WithPeriod(25), ind.OnField("price"))
	ohlcv.Print()

	// Create a new series of sma 3
	fmt.Printf("\nCreate a new series of sma with period 3\n\n")
	sma := ind.SMA(ohlcv, ind.WithPeriod(3), ind.OnField("close"))
	sma.Print()

	// Create a new series of ema 3
	fmt.Printf("\nCreate a new series of ema with period 3\n\n")
	ema := ind.EMA(ohlcv, ind.WithPeriod(3), ind.OnField("close"))
	ema.Print()

	// Create a new series of wma 3
	fmt.Printf("\nCreate a new series of wma with period 3\n\n")
	wma := ind.WMA(ohlcv, ind.WithPeriod(3), ind.OnField("close"))
	wma.Print()
}
