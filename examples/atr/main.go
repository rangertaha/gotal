package main

import (
	"fmt"

	"github.com/rangertaha/gota/indicators"
)

func main() {
	// Sample price data (daily OHLC)
	high := []float64{
		44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84, 46.08,
		45.89, 46.03, 45.61, 46.28, 46.28, 46.00, 46.03, 46.41, 46.22, 45.64,
		46.21, 46.25, 45.71, 46.45, 45.78, 45.35, 44.03, 44.18, 44.22, 44.57,
		43.42, 42.66, 43.13, 43.82, 44.28, 44.00, 43.46, 43.19, 43.22, 43.57,
		44.15, 44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84,
	}

	low := []float64{
		44.09, 43.61, 43.61, 43.33, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84,
		45.89, 46.03, 45.61, 46.28, 46.28, 46.00, 46.03, 46.41, 46.22, 45.64,
		46.21, 46.25, 45.71, 46.45, 45.78, 45.35, 44.03, 44.18, 44.22, 44.57,
		43.42, 42.66, 43.13, 43.82, 44.28, 44.00, 43.46, 43.19, 43.22, 43.57,
		44.15, 44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84,
	}

	close := []float64{
		44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84, 46.08, 45.89,
		46.03, 45.61, 46.28, 46.28, 46.00, 46.03, 46.41, 46.22, 45.64, 46.21,
		46.25, 45.71, 46.45, 45.78, 45.35, 44.03, 44.18, 44.22, 44.57, 43.42,
		42.66, 43.13, 43.82, 44.28, 44.00, 43.46, 43.19, 43.22, 43.57, 44.15,
		44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84, 46.08,
	}

	// Create a new ATR calculator with standard period (14)
	atr := indicators.NewATR(14)

	// Calculate ATR
	_, err := atr.Calculate(high, low, close)
	if err != nil {
		fmt.Printf("Error calculating ATR: %v\n", err)
		return
	}

	// Get results
	atrValues := atr.GetResult()

	// Print results
	fmt.Println("Average True Range (ATR):")
	fmt.Println("Period: 14")
	fmt.Println("\nDetailed Analysis:")

	for i := 0; i < len(atrValues); i++ {
		fmt.Printf("\nDay %d:\n", i+15) // Adjust index to match ATR output
		fmt.Printf("  ATR: %.4f\n", atrValues[i])
		fmt.Printf("  Volatility Ratio: %.2f\n", atr.GetVolatilityRatio(i))

		// Print volatility conditions
		if atr.IsVolatilityHigh(i, 1.5) {
			fmt.Println("  Volatility: High (Above 1.5x average)")
		} else if atr.IsVolatilityLow(i, 1.5) {
			fmt.Println("  Volatility: Low (Below 0.67x average)")
		} else {
			fmt.Println("  Volatility: Normal")
		}
	}

	// Print ATR analysis
	fmt.Println("\nATR Analysis:")
	fmt.Println("1. ATR measures market volatility")
	fmt.Println("2. Higher ATR indicates higher volatility")
	fmt.Println("3. Lower ATR indicates lower volatility")
	fmt.Println("4. ATR can be used to set stop-loss levels")
	fmt.Println("5. ATR can help identify potential breakouts")
	fmt.Println("6. ATR is non-directional (doesn't indicate trend direction)")
	fmt.Println("7. ATR can be used to normalize other indicators")
	fmt.Println("8. ATR is useful for position sizing")
}
