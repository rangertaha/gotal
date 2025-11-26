package main

import (
	"fmt"

	"github.com/rangertaha/gotal/internal/pkg/series"
	_ "github.com/rangertaha/gotal/internal/plugins/providers/all"
	"github.com/rangertaha/gotal/pkg/providers/stream"
)

func main() {
	var Tags = map[string]string{
		"symbol":   "SOL-USD",
		"exchange": "COINBASE",
		"currency": "USD",
		"asset":    "SOL",
	}

	// Generate a new series of mock waves
	prices := stream.Mock()
	fmt.Println("mock wave generated")
	prices.Print()

	// Save to CSV
	fmt.Println("CSV: Saving to file")
	if err := prices.Save("mock.csv"); err != nil {
		fmt.Println(err)
	}

	// Load from CSV
	fmt.Println("CSV: Loading from file")
	ts, err := series.Load("mock.csv")
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

	if err := plot.Save("mock.png", 8, 4); err != nil {
		fmt.Printf("Error saving plot: %v\n", err)
	}

	if err := plot.Show(8, 4); err != nil {
		fmt.Printf("Error showing plot: %v\n", err)
	}

	// if err := plot.Show(); err != nil {
	// 	fmt.Print(err)
	// }

	// if err := plot.Save("mock.png"); err != nil {
	// 	fmt.Print(err)
	// }
}
