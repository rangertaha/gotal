package main

import (
	"time"

	"github.com/rangertaha/gotal/pkg/gen"
)

func main() {
	var Tags = map[string]string{
		"symbol":   "ETH-USD",
		"exchange": "COINBASE", 
		"currency": "USD",
		"asset":    "ETH",
	}

	// Generate a sawtooth wave
	prices := gen.Sawtooth(time.Second, 100, 0.2, 0, Tags)

	// Display table
	prices.Print()
}
