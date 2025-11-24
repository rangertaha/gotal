package main

import (
	"time"

	"github.com/rangertaha/gotal/pkg/gen"
)

func main() {
	var Tags = map[string]string{
		"symbol":   "AAPL",
		"exchange": "NASDAQ", 
		"currency": "USD",
		"asset":    "AAPL",
	}

	// Generate Geometric Brownian Motion (Black-Scholes model)
	// Parameters: duration, initialPrice, drift, volatility, samples, tags
	prices := gen.GBM(time.Minute, 150.0, 0.05, 0.2, 100, Tags)

	// Display table
	prices.Print()
}
