package main

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal/pkg/gen"
)

func main() {
	var Tags = map[string]string{
		"symbol":   "STEP-TEST",
		"exchange": "SIMULATION", 
		"currency": "USD",
		"asset":    "TEST",
	}

	fmt.Println("=== STEP FUNCTION ===")
	stepFunc := gen.Step(time.Minute, 50, 150, 15, 30, Tags)
	stepFunc.Print()

	fmt.Println("\n=== RAMP FUNCTION ===")
	rampFunc := gen.Ramp(time.Minute, 50, 150, 30, Tags)
	rampFunc.Print()
}
