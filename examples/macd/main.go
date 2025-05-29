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
		43.42, 42.66, 43.13, 43.82, 44.28, 44.00, 43.46, 43.19, 43.22, 43.57,
		44.15, 44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84,
	}

	// Create a new MACD calculator with standard periods
	// Fast Period: 12, Slow Period: 26, Signal Period: 9
	macd := indicators.NewMACD(12, 26, 9)

	// Calculate MACD
	_, err := macd.Calculate(prices)
	if err != nil {
		fmt.Printf("Error calculating MACD: %v\n", err)
		return
	}

	// Get detailed results
	macdResult := macd.GetResult()

	// Print results
	fmt.Println("Moving Average Convergence Divergence (MACD):")
	fmt.Println("Periods: Fast=12, Slow=26, Signal=9")
	fmt.Println("\nDetailed Analysis:")

	for i := 0; i < len(macdResult.MACDLine); i++ {
		fmt.Printf("\nDay %d:\n", i+34) // Adjust index to match MACD output
		fmt.Printf("  Price: %.2f\n", prices[i+34])
		fmt.Printf("  MACD Line: %.4f\n", macdResult.MACDLine[i])
		fmt.Printf("  Signal Line: %.4f\n", macdResult.SignalLine[i])
		fmt.Printf("  Histogram: %.4f\n", macdResult.Histogram[i])

		// Print signals
		if macd.IsBullish(i) {
			fmt.Println("  Signal: Bullish")
		} else if macd.IsBearish(i) {
			fmt.Println("  Signal: Bearish")
		} else {
			fmt.Println("  Signal: Neutral")
		}

		// Print divergences
		if macd.HasBullishDivergence(prices, i) {
			fmt.Println("  Divergence: Bullish")
		} else if macd.HasBearishDivergence(prices, i) {
			fmt.Println("  Divergence: Bearish")
		}
	}

	// Print MACD analysis
	fmt.Println("\nMACD Analysis:")
	fmt.Println("1. MACD Line = Fast EMA (12) - Slow EMA (26)")
	fmt.Println("2. Signal Line = EMA of MACD Line (9)")
	fmt.Println("3. Histogram = MACD Line - Signal Line")
	fmt.Println("4. Bullish Signal: Histogram crosses above zero")
	fmt.Println("5. Bearish Signal: Histogram crosses below zero")
	fmt.Println("6. Divergence: Price and MACD move in opposite directions")
}
