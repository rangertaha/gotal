package main

import (
	"fmt"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/opt"
	_ "github.com/rangertaha/gotal/internal/plugins/providers/all"
	"github.com/rangertaha/gotal/pkg/batch"
)

func main() {
	var err error
	var dataset internal.Series

	fmt.Println("1. -----------------------------------------------------------------------------")
	{
		if dataset, err = batch.Generator(opt.WithFile("sine.hcl")); dataset != nil {
			dataset.Print()
		}
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	fmt.Println("2. -----------------------------------------------------------------------------")
	{
		if dataset, err = batch.Generator(opt.WithHCL(`
		sine "price" {
			periods = 100
			amplitude = 1.0
			frequency = 1.0
			phase = 0
			offset = 0
		}`)); dataset != nil {
			dataset.Print()
		}

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	fmt.Println("3. -----------------------------------------------------------------------------")
	{
		if dataset, err = batch.Generator(opt.WithJSON(`
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
		}`)); dataset != nil {
			dataset.Print()
		}

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	fmt.Println("4. -----------------------------------------------------------------------------")
	{
		if dataset, err = batch.Generator(
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
		}`)); dataset != nil {
			dataset.Print()
		}

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}
