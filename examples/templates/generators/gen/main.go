package main

import (
	"fmt"
	"log"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/config"
	"github.com/rangertaha/gotal/internal/opts"
	t "github.com/rangertaha/gotal/internal/plugins/templates"
	_ "github.com/rangertaha/gotal/internal/plugins/templates/all"
)

func main() {

	// Load config from file
	cfg, err := config.Load("sine.hcl")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("%+v\n", cfg)
	for _, provider := range cfg.Providers {
		fmt.Printf("%+v\n%+v\n", provider, provider.Config)
	}
	for _, indicator := range cfg.Indicators {
		fmt.Printf("%+v\n", indicator)
	}
	for _, strategy := range cfg.Strategies {
		fmt.Printf("%+v\n", strategy)
	}
	for _, broker := range cfg.Brokers {
		fmt.Printf("%+v\n", broker)
	}
	//-----------------------------------------------------------------------------
	// Generate a new generator

	// Use default options
	{
		pfunc, err := t.Get("GEN")
		if err != nil {
			fmt.Println("Failed to get generator - plugin is nil")
			return
		}
		plg, err := pfunc()
		if err != nil {
			fmt.Println("Failed to get generator - plugin is nil")
			return
		}

		fmt.Printf("%s %s %s\n", plg.ID(), plg.Name(), plg.Description())

		if proc, ok := plg.(internal.Processor); ok {
			output := proc.Process(nil)
			fmt.Printf("Processor executed, output: %+v\n", output)
		} else {
			fmt.Println("Plugin does not implement Processor interface")
		}
	}
	//-----------------------------------------------------------------------------
	{
		pfunc, err := t.Get("GEN")
		if err != nil {
			fmt.Println("Failed to get generator - plugin is nil")
			return
		}
		cfg := opts.WithConfig("sine.hcl")

		plg, err := pfunc(cfg)
		if err != nil {
			fmt.Println("Failed to get generator - plugin is nil")
			return
		}

		fmt.Printf("%s %s %s\n", plg.ID(), plg.Name(), plg.Description())

		if proc, ok := plg.(internal.Processor); ok {
			output := proc.Process(nil)
			fmt.Printf("Processor executed, output: %+v\n", output)
		} else {
			fmt.Println("Plugin does not implement Processor interface")
		}
	}
	//-----------------------------------------------------------------------------

	//-----------------------------------------------------------------------------
	// // Use custom options
	// {
	// 	// function parameters
	// 	freq := opt.With("freq", 100)

	// 	// function call
	// 	series, _ := p.Mock(freq)

	// 	// check response is not empty
	// 	if series.IsEmpty() {

	// 		fmt.Println("Failed to generate sine wave - series is empty")
	// 		return
	// 	}

	// 	// print response
	// 	series.Print()
	// }

	// // Load options from config file
	// {
	// 	// function parameters
	// 	configFile := opt.WithConfig("sine.hcl")

	// 	// function call
	// 	series, _ := p.Mock(configFile)

	// 	// check response is not empty
	// 	if series.IsEmpty() {
	// 		fmt.Println("Failed to generate sine wave - series is empty")
	// 		return
	// 	}

	// 	// print response
	// 	series.Print()
	// }

	// // Load options from config directory
	// {
	// 	// get all config directory
	// 	conf := opt.WithConfig("files")

	// 	// get all providers
	// 	for _, provider := range conf.Providers() {
	// 		series, _ := provider()
	// 		if series.IsEmpty() {
	// 			fmt.Println("Failed to get series for provider", provider.ID())
	// 		}

	// 		// print series
	// 		series.Print()
	// 	}
	// }

}
