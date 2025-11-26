package main

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/tick"
	_ "github.com/rangertaha/gotal/internal/plugins/providers/all"
	"github.com/rangertaha/gotal/pkg/providers"
)

var (
	Tags = map[string]string{
		"symbol":   "SOL-USD",
		"exchange": "COINBASE",
		"currency": "USD",
		"asset":    "SOL",
	}
	Field = map[string]float64{
		"start":    float64(time.Now().Unix()),
		"stop":     float64(time.Now().Add(1 * time.Hour).Unix()),
		"duration": float64(time.Duration(1 * time.Minute)),
	}
)

func main() {
	ticks := tick.New(tick.WithFields(Field))
	series := series.New("mock", series.WithFields(Field), series.WithTicks(ticks))

	// Generate a new series of mock data
	prices := providers.Mock(series, providers.With("type", "sine"))
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
