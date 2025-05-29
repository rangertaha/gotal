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

	// Create a new CCI calculator with standard period (20)
	cci := indicators.NewCCI(20)

	// Calculate CCI
	_, err := cci.Calculate(high, low, close)
	if err != nil {
		fmt.Printf("Error calculating CCI: %v\n", err)
		return
	}

	// Get results
	result := cci.GetResult()

	// Print results
	fmt.Println("Commodity Channel Index (CCI):")
	fmt.Println("Parameters: Period=20")
	fmt.Println("\nDetailed Analysis:")

	for i := 0; i < len(result.CCI); i++ {
		fmt.Printf("\nDay %d:\n", i+20) // Adjust index to match CCI output
		fmt.Printf("  CCI: %.2f\n", result.CCI[i])
		fmt.Printf("  Typical Price: %.2f\n", result.TP[i])
		fmt.Printf("  Mean TP: %.2f\n", result.MeanTP[i])
		fmt.Printf("  Mean Deviation: %.2f\n", result.MeanDev[i])

		// Print signals
		if cci.IsExtremelyOverbought(i) {
			fmt.Println("  Signal: Extremely Overbought")
		} else if cci.IsOverbought(i) {
			fmt.Println("  Signal: Overbought")
		} else if cci.IsExtremelyOversold(i) {
			fmt.Println("  Signal: Extremely Oversold")
		} else if cci.IsOversold(i) {
			fmt.Println("  Signal: Oversold")
		}

		if cci.HasZeroLineCross(i) {
			fmt.Println("  Signal: Zero Line Cross")
		}

		if cci.HasBullishDivergence(close, i) {
			fmt.Println("  Signal: Bullish Divergence")
		} else if cci.HasBearishDivergence(close, i) {
			fmt.Println("  Signal: Bearish Divergence")
		}
	}

	// Print CCI analysis
	fmt.Println("\nCCI Analysis:")
	fmt.Println("1. CCI is a momentum oscillator")
	fmt.Println("2. CCI > 100 indicates overbought conditions")
	fmt.Println("3. CCI < -100 indicates oversold conditions")
	fmt.Println("4. CCI > 200 indicates extremely overbought conditions")
	fmt.Println("5. CCI < -200 indicates extremely oversold conditions")
	fmt.Println("6. Zero line crosses can signal trend changes")
	fmt.Println("7. Divergences can signal potential reversals")
	fmt.Println("8. Best used in ranging markets")
	fmt.Println("9. Can be used to identify overbought/oversold conditions")
	fmt.Println("10. Can be used to spot divergences and reversals")
}
