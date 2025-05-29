package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/rangertaha/gota/indicators"
)

func main() {
	// Set random seed for reproducibility
	rand.Seed(42)

	// Generate sample price data
	prices := generateSampleData(100)

	// Create HMM with default configuration
	config := indicators.DefaultHMMConfig()
	hmm := indicators.NewHMM(config)

	// Fit the model to the data
	hmm.Fit(prices)

	// Make predictions
	result := hmm.Predict(prices)

	// Print results
	fmt.Println("Hidden Markov Model Analysis Results:")
	fmt.Println("====================================")

	// Print state sequence
	fmt.Println("\nState Sequence:")
	for i, state := range result.States {
		fmt.Printf("Time %d: %s (Probability: %.2f%%)\n",
			i+1,
			hmm.GetStateDescription(state),
			result.Probabilities[i]*100)
	}

	// Print transition matrix
	fmt.Println("\nTransition Matrix:")
	for i := range result.Transitions {
		fmt.Printf("From %s:\n", hmm.GetStateDescription(i))
		for j := range result.Transitions[i] {
			fmt.Printf("  To %s: %.2f%%\n",
				hmm.GetStateDescription(j),
				result.Transitions[i][j]*100)
		}
	}
}

// generateSampleData creates sample price data with different market regimes
func generateSampleData(n int) []float64 {
	prices := make([]float64, n)
	prices[0] = 100.0 // Initial price

	// Parameters for different regimes
	regimes := []struct {
		mean       float64
		volatility float64
		duration   int
	}{
		{0.001, 0.01, 30},   // Bullish
		{-0.001, 0.015, 20}, // Bearish
		{0.0001, 0.005, 50}, // Sideways
	}

	currentRegime := 0
	regimeCount := 0

	for i := 1; i < n; i++ {
		// Check if we need to switch regimes
		if regimeCount >= regimes[currentRegime].duration {
			currentRegime = (currentRegime + 1) % len(regimes)
			regimeCount = 0
		}

		// Generate return based on current regime
		return_ := regimes[currentRegime].mean + regimes[currentRegime].volatility*rand.NormFloat64()
		prices[i] = prices[i-1] * math.Exp(return_)
		regimeCount++
	}

	return prices
}
