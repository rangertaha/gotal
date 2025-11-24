// This example shows how to use the MACD (Moving Average Convergence Divergence) indicator.
// MACD is a trend-following momentum indicator that shows the relationship between two moving averages of a security's price.
// The MACD is calculated by subtracting the 26-period EMA from the 12-period EMA.
// A nine-day EMA of the MACD called the "signal line" is then plotted on top of the MACD line.
// The histogram represents the difference between the MACD line and the signal line.

// It creates a time series of prices and then calculates the MACD indicator.
// It then prints the MACD line, signal line, and histogram values.

// Example:
// go run examples/indicators/macd/main.go

// Output:
// MACD Line, Signal Line, Histogram values for each time period

package main

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal/examples"
	. "github.com/rangertaha/gotal/pkg/indicators"
)

func main() {
	// Create a new series of price ticks
	fmt.Printf("Create a new series of price ticks\n\n")
	prices := examples.SineSeries(time.Second, 100, 0.1, 0)

	// Display the original price data
	fmt.Println("Original Price Data:")
	prices.Print()

	// Create MACD indicator with default parameters (12, 26, 9)
	fmt.Printf("\n\nCalculating MACD indicator...\n")
	macd := MACD(prices,
		WithFastPeriod(12),
		WithSlowPeriod(26),
		WithSignalPeriod(9),
	)

	// Display MACD results
	fmt.Println("\nMACD Results:")
	fmt.Println("Time\t\t\tMACD Line\tSignal Line\tHistogram")
	fmt.Println("----\t\t\t---------\t----------\t--------")

	// Print MACD values
	for i := 0; i < macd.Len(); i++ {
		tick := macd.At(i)
		if tick != nil {
			macdValue := tick.GetField("macd")
			signalValue := tick.GetField("signal")
			histogramValue := tick.GetField("histogram")

			fmt.Printf("%s\t%.4f\t\t%.4f\t\t%.4f\n",
				tick.Timestamp().Format("15:04:05"),
				macdValue,
				signalValue,
				histogramValue)
		}
	}
}
