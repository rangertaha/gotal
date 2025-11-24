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
	ind "github.com/rangertaha/gotal/pkg/indicators"
)

func main() {
	// Create a new series of price ticks
	fmt.Printf("\nCreate a new series of price ticks\n\n")
	prices := examples.PricesSeries(1000, time.Second)
	prices.Print()

}
