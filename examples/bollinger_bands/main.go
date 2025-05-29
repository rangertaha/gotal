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

	// Create a new Bollinger Bands calculator with standard parameters
	// Period: 20, Standard Deviation: 2.0
	bb := indicators.NewBollingerBands(20, 2.0)

	// Calculate Bollinger Bands
	_, err := bb.Calculate(prices)
	if err != nil {
		fmt.Printf("Error calculating Bollinger Bands: %v\n", err)
		return
	}

	// Get detailed results
	bbResult := bb.GetResult()

	// Print results
	fmt.Println("Bollinger Bands:")
	fmt.Println("Period: 20, Standard Deviation: 2.0")
	fmt.Println("\nDetailed Analysis:")

	for i := 0; i < len(bbResult.MiddleBand); i++ {
		fmt.Printf("\nDay %d:\n", i+20) // Adjust index to match BB output
		fmt.Printf("  Price: %.2f\n", prices[i+19])
		fmt.Printf("  Middle Band: %.4f\n", bbResult.MiddleBand[i])
		fmt.Printf("  Upper Band: %.4f\n", bbResult.UpperBand[i])
		fmt.Printf("  Lower Band: %.4f\n", bbResult.LowerBand[i])
		fmt.Printf("  Band Width: %.4f\n", bbResult.Width[i])
		fmt.Printf("  %%B: %.4f\n", bbResult.PercentB[i])

		// Print signals
		if bb.IsOverbought(i) {
			fmt.Println("  Signal: Overbought")
		} else if bb.IsOversold(i) {
			fmt.Println("  Signal: Oversold")
		} else {
			fmt.Println("  Signal: Neutral")
		}

		// Print band conditions
		if bb.IsSqueeze(i) {
			fmt.Println("  Band Condition: Squeeze (Narrowing)")
		} else if bb.IsExpansion(i) {
			fmt.Println("  Band Condition: Expansion (Widening)")
		} else {
			fmt.Println("  Band Condition: Stable")
		}
	}

	// Print Bollinger Bands analysis
	fmt.Println("\nBollinger Bands Analysis:")
	fmt.Println("1. Middle Band = 20-period Simple Moving Average")
	fmt.Println("2. Upper Band = Middle Band + (2.0 * Standard Deviation)")
	fmt.Println("3. Lower Band = Middle Band - (2.0 * Standard Deviation)")
	fmt.Println("4. Band Width = (Upper - Lower) / Middle")
	fmt.Println("5. %B = (Price - Lower) / (Upper - Lower)")
	fmt.Println("6. Overbought: Price above Upper Band or %B > 1.0")
	fmt.Println("7. Oversold: Price below Lower Band or %B < 0.0")
	fmt.Println("8. Squeeze: Bands narrowing (decreasing width)")
	fmt.Println("9. Expansion: Bands widening (increasing width)")
}
