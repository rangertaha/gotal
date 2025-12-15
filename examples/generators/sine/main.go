package main

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/pkg/gen"
)

func main() {
	var Tags = map[string]string{
		"symbol":   "BTC-USD",
		"exchange": "BINANCE",
		"currency": "USD",
		"asset":    "BTC",
	}

	// Generate a new series of sine waves
	prices := gen.Sine("price", time.Second, 100, 0.1, 0, 4, Tags)
	fmt.Println("Sine wave generated")
	prices.Print()

	// Save to CSV
	fmt.Println("CSV: Saving to file")
	if err := prices.Save("sine.csv"); err != nil {
		fmt.Println(err)
	}

	// Load from CSV
	fmt.Println("CSV: Loading from file")
	ts, err := series.Load("sine.csv")
	if err != nil {
		fmt.Println("Errors:")
		fmt.Println(err)
	}

	// Display table
	if ts != nil {
		ts.Print()
	}

	// Create and save plot
	plot, err := prices.Plot("price")
	if err != nil {
		fmt.Printf("Error creating plot: %v\n", err)
	}

	if err := plot.Save("sine.png", 8, 4); err != nil {
		fmt.Printf("Error saving plot: %v\n", err)
	}
	
	if err := plot.Show(8, 4); err != nil {
		fmt.Printf("Error showing plot: %v\n", err)
	}

	// if err := plot.Show(); err != nil {
	// 	fmt.Print(err)
	// }

	// if err := plot.Save("sine.png"); err != nil {
	// 	fmt.Print(err)
	// }
}
