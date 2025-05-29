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

	// Create a new Stochastic calculator with standard parameters
	// KPeriod=14, DPeriod=3, Slowing=3
	stoch := indicators.NewStochastic(14, 3, 3)

	// Calculate Stochastic
	_, err := stoch.Calculate(high, low, close)
	if err != nil {
		fmt.Printf("Error calculating Stochastic: %v\n", err)
		return
	}

	// Get results
	result := stoch.GetResult()

	// Print results
	fmt.Println("Stochastic Oscillator:")
	fmt.Println("Parameters: K=14, D=3, Slowing=3")
	fmt.Println("\nDetailed Analysis:")

	for i := 0; i < len(result.K); i++ {
		fmt.Printf("\nDay %d:\n", i+20) // Adjust index to match Stochastic output
		fmt.Printf("  %K: %.2f\n", result.K[i])
		fmt.Printf("  %D: %.2f\n", result.D[i])

		// Print signals
		if stoch.IsOverbought(i) {
			fmt.Println("  Signal: Overbought")
		} else if stoch.IsOversold(i) {
			fmt.Println("  Signal: Oversold")
		}

		if stoch.HasBullishCross(i) {
			fmt.Println("  Signal: Bullish Cross (%K crossed above %D)")
		} else if stoch.HasBearishCross(i) {
			fmt.Println("  Signal: Bearish Cross (%K crossed below %D)")
		}

		if stoch.HasBullishDivergence(close, i) {
			fmt.Println("  Signal: Bullish Divergence")
		} else if stoch.HasBearishDivergence(close, i) {
			fmt.Println("  Signal: Bearish Divergence")
		}
	}

	// Print Stochastic analysis
	fmt.Println("\nStochastic Oscillator Analysis:")
	fmt.Println("1. Stochastic Oscillator is a momentum indicator")
	fmt.Println("2. %K is the main line (Fast Stochastic)")
	fmt.Println("3. %D is the signal line (SMA of %K)")
	fmt.Println("4. Overbought condition: %K > 80")
	fmt.Println("5. Oversold condition: %K < 20")
	fmt.Println("6. Bullish cross: %K crosses above %D")
	fmt.Println("7. Bearish cross: %K crosses below %D")
	fmt.Println("8. Divergences can signal potential reversals")
	fmt.Println("9. Best used in ranging markets")
	fmt.Println("10. Can be used to identify overbought/oversold conditions")
}
