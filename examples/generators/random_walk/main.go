package main

import (
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

	// Generate a random walk series
	prices := gen.RandomWalk(time.Second, 2.0, 0.1, 50, Tags)

	// Display table
	prices.Print()
}
