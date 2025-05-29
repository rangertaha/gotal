package main

import (
	"fmt"

	"github.com/rangertaha/gotal/indicators"
)

func main() {
	// Sample price data (daily closing prices)
	prices := []float64{
		44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84, 46.08,
		45.89, 46.03, 45.61, 46.28, 46.28, 46.00, 46.03, 46.41, 46.22, 45.64,
		46.21, 46.25, 45.71, 46.45, 45.78, 45.35, 44.03, 44.18, 44.22, 44.57,
	}

	// Create a new RSI calculator with period 14 (standard period)
	rsi := indicators.NewRSI(14)

	// Calculate RSI
	result, err := rsi.Calculate(prices)
	if err != nil {
		fmt.Printf("Error calculating RSI: %v\n", err)
		return
	}

	// Print results
	fmt.Println("Relative Strength Index (RSI) with period 14:")
	for i, value := range result.Values {
		status := "Neutral"
		if rsi.IsOverbought(value) {
			status = "Overbought"
		} else if rsi.IsOversold(value) {
			status = "Oversold"
		}
		fmt.Printf("RSI[%d] = %.2f (%s)\n", i, value, status)
	}

	// Print some analysis
	fmt.Println("\nRSI Analysis:")
	fmt.Println("1. RSI values range from 0 to 100")
	fmt.Println("2. Values above 70 typically indicate overbought conditions")
	fmt.Println("3. Values below 30 typically indicate oversold conditions")
	fmt.Println("4. RSI can be used to identify potential trend reversals")
	fmt.Println("5. Divergence between price and RSI can signal potential trend changes")
}
