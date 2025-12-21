package main

import (
	"fmt"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/opt"
	_ "github.com/rangertaha/gotal/internal/plugins/providers/all"
)

func main() {
	// get a new generator with the sine.hcl config file
	genSeries, genStream, err := gotal.Generator(opt.WithFile("sine.hcl"))
	if err != nil {
		fmt.Println("Failed to get generator function")
		return
	}

	fmt.Println("Generator series: ", genSeries)
	fmt.Println("Generator stream: ", genStream)
	fmt.Println("Generator series: ", err)
	fmt.Println("-----------------------------------------------------------------------------")

	// get a new generator with the config content
	genSeries2, genStream2, err2 := gotal.Generator(opt.WithHCL(`
name = "prices"
start = date("2021/01/01")
end = date("2021/01/02")
interval = minutes(1)

sine "price" {
  periods   = 100            // Number of periods (points)
  amplitude = 1.0            // Amplitude of the sine wave
  frequency = 1.0            // Frequency of the sine wave
  phase     = 0.0            // Phase offset (radians)
  offset    = 0.0            // Value offset
}
	`))
	if err2 != nil {
		fmt.Println("Failed to get generator function", err2)
		return
	}

	fmt.Println("Generator series: ", genSeries2)
	fmt.Println("Generator stream: ", genStream2)
	fmt.Println("Generator series: ", err2)
	fmt.Println("-----------------------------------------------------------------------------")

	// get a new generator with the config content
	genSeries3, genStream3, err3 := gotal.Generator(opt.WithJSON(`
		{
			"name": "prices",
			"start": 1609459200,
			"end": 1609459200,
			"interval": 60,
			"sine": {
				"price": {
					"periods": 100,
					"amplitude": 1.0,
					"frequency": 1.0,
					"phase": 0,
					"offset": 0
				}
			}
		}
		`))
	if err3 != nil {
		fmt.Println("Failed to get generator function", err3)
		return
	}

	// genSeries.Print()
	fmt.Println("Generator series: ", genSeries3)
	fmt.Println("Generator stream: ", genStream3)
	fmt.Println("Generator series: ", err3)
	fmt.Println("-----------------------------------------------------------------------------")

	// get a new generator with the config content
	genSeries4, genStream4, err4 := gotal.Generator(
		opt.With("start", "1609459200"),
		opt.With("end", "1609459200"),
		opt.With("interval", "1"),
		opt.With("sine", "1m"),
		opt.With("name", "new_series"),
		opt.WithHCL(`
		  sine "volume" {
			periods = 100
			amplitude = 1.0
			frequency = 1.0
			phase = 0
			offset = 0
		}`),
	)
	if err4 != nil {
		fmt.Println("Failed to get generator function", err4)
		return
	}

	fmt.Println("Generator series: ", genSeries4)
	fmt.Println("Generator stream: ", genStream4)
	fmt.Println("Generator series: ", err4)
	fmt.Println("-----------------------------------------------------------------------------")
	//-----------------------------------------------------------------------------

	//-----------------------------------------------------------------------------

}
