package main

import (
	"fmt"

	"github.com/rangertaha/gotal/indicators"
	"github.com/rangertaha/gotal/types"
)

func main() {
	// Sample price data
	prices := []float64{
		10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0,
	}

	// Create a new SMA calculator with period 3
	sma := indicators.NewMovingAverage(types.SMA, 3)

	// Calculate SMA
	result, err := sma.Calculate(prices)
	if err != nil {
		fmt.Printf("Error calculating SMA: %v\n", err)
		return
	}

	// Print results
	fmt.Println("Simple Moving Average (SMA) with period 3:")
	for i, value := range result.Values {
		fmt.Printf("SMA[%d] = %.2f\n", i, value)
	}

	// Create a new EMA calculator with period 3
	ema := indicators.NewMovingAverage(types.EMA, 3)

	// Calculate EMA
	result, err = ema.Calculate(prices)
	if err != nil {
		fmt.Printf("Error calculating EMA: %v\n", err)
		return
	}

	// Print results
	fmt.Println("\nExponential Moving Average (EMA) with period 3:")
	for i, value := range result.Values {
		fmt.Printf("EMA[%d] = %.2f\n", i, value)
	}

	// Create a new WMA calculator with period 3
	wma := indicators.NewMovingAverage(types.WMA, 3)

	// Calculate WMA
	result, err = wma.Calculate(prices)
	if err != nil {
		fmt.Printf("Error calculating WMA: %v\n", err)
		return
	}

	// Print results
	fmt.Println("\nWeighted Moving Average (WMA) with period 3:")
	for i, value := range result.Values {
		fmt.Printf("WMA[%d] = %.2f\n", i, value)
	}
}
