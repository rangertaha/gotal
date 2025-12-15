package main

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal/pkg/gen"
)

func main() {
	var Tags = map[string]string{
		"symbol":   "TREND-TEST",
		"exchange": "SIMULATION", 
		"currency": "USD",
		"asset":    "TEST",
	}

	fmt.Println("=== LINEAR TREND ===")
	linearTrend := gen.LinearTrend(time.Hour, 100, 0.5, 2.0, 20, Tags)
	linearTrend.Print()

	fmt.Println("\n=== EXPONENTIAL TREND ===")
	expTrend := gen.ExponentialTrend(time.Hour, 100, 0.02, 1.0, 20, Tags)
	expTrend.Print()
}
