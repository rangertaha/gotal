package main

import (
	"fmt"

	"github.com/rangertaha/gotal/indicators"
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

	// Create a new ADX calculator with standard period (14)
	adx := indicators.NewADX(14)

	// Calculate ADX
	_, err := adx.Calculate(high, low, close)
	if err != nil {
		fmt.Printf("Error calculating ADX: %v\n", err)
		return
	}

	// Get results
	result := adx.GetResult()

	// Print results
	fmt.Println("Average Directional Index (ADX):")
	fmt.Println("Parameters: Period=14")
	fmt.Println("\nDetailed Analysis:")

	for i := 0; i < len(result.ADX); i++ {
		fmt.Printf("\nDay %d:\n", i+27) // Adjust index to match ADX output
		fmt.Printf("  ADX: %.2f\n", result.ADX[i])
		fmt.Printf("  +DI: %.2f\n", result.PlusDI[i])
		fmt.Printf("  -DI: %.2f\n", result.MinusDI[i])

		// Print trend strength
		if adx.IsVeryStrongTrend(i) {
			fmt.Println("  Trend Strength: Very Strong")
		} else if adx.IsStrongTrend(i) {
			fmt.Println("  Trend Strength: Strong")
		} else {
			fmt.Println("  Trend Strength: Weak")
		}

		// Print trend direction
		if adx.IsBullishTrend(i) {
			fmt.Println("  Trend Direction: Bullish")
		} else if adx.IsBearishTrend(i) {
			fmt.Println("  Trend Direction: Bearish")
		} else {
			fmt.Println("  Trend Direction: Neutral")
		}

		// Print trend reversal
		if adx.HasTrendReversal(i) {
			fmt.Println("  Signal: Trend Reversal")
		}
	}

	// Print ADX analysis
	fmt.Println("\nADX Analysis:")
	fmt.Println("1. ADX is a trend strength indicator")
	fmt.Println("2. ADX > 25 indicates a strong trend")
	fmt.Println("3. ADX > 50 indicates a very strong trend")
	fmt.Println("4. +DI > -DI indicates a bullish trend")
	fmt.Println("5. -DI > +DI indicates a bearish trend")
	fmt.Println("6. Trend reversals occur when DI lines cross")
	fmt.Println("7. Best used in trending markets")
	fmt.Println("8. Can be used to confirm trend strength")
	fmt.Println("9. Higher ADX values indicate stronger trends")
	fmt.Println("10. Lower ADX values indicate ranging markets")
}
