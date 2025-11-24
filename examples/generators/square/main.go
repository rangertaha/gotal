package main

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal/pkg/gen"
)

func main() {
	var Tags = map[string]string{
		"symbol":   "BTC-USD",
		"exchange": "BINANCE",
		"currency": "USD",
		"asset":    "BTC",
	}

	// Generate a new series of square waves
	prices := gen.Square(time.Second, 100, 0.1, 0, Tags)

	// Display table
	fmt.Println("Square wave generator")
	prices.Print()

	// Save to file
	// plot := prices.Plot()
	// plot.SavePlot("square.png")
}
