package main

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal/pkg/gen"
)

func main() {
	var Tags = map[string]string{
		"symbol":   "NOISE",
		"exchange": "SIMULATION", 
		"currency": "USD",
		"asset":    "TEST",
	}

	fmt.Println("=== WHITE NOISE ===")
	whiteNoise := gen.WhiteNoise(time.Second, 50, 30, Tags)
	whiteNoise.Print()

	fmt.Println("\n=== PINK NOISE ===") 
	pinkNoise := gen.PinkNoise(time.Second, 50, 30, Tags)
	pinkNoise.Print()
}
