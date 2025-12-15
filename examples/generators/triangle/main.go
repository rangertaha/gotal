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

	// Generate a new series of triangle waves
	prices := gen.Triangle(time.Second, 100, 0.1, 0, Tags)

	// Display table
	fmt.Println("Triangle wave generator")
	prices.Print()

	// Save to file
	// plot := prices.Plot()
	// plot.SavePlot("triangle.png")
}
