package main

import (
	"log"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/series"
)

func main() {
	input := series.New("price")
	ema, err := gotal.EMA(input)
	if err != nil {
		log.Fatal(err)
	} 

	// print the ema series to the console
	ema.Print()

	// save the ema series to a csv file
	ema.Save("ema.csv")

	// plot the ema series
	plot := ema.Plot()

	// save the plot to a png file
	plot.Save("ema.png", 800, 600)
}
