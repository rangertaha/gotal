package main

import (
	"fmt"
	"log"

	"github.com/rangertaha/gotal/internal/config"
	_ "github.com/rangertaha/gotal/internal/plugins/brokers/all"
	_ "github.com/rangertaha/gotal/internal/plugins/indicators/all"
	_ "github.com/rangertaha/gotal/internal/plugins/providers/all"
	_ "github.com/rangertaha/gotal/internal/plugins/strategies/all"
)

func main() {
	// Load config from file
	cfg, err := config.Load("sine.hcl")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	for _, provider := range cfg.Providers {
		fmt.Printf("Provider: %+v %+v\n\n", provider.Type, provider.Name)
	}
	for _, indicator := range cfg.Indicators {
		fmt.Printf("Indicator: %+v %+v\n\n", indicator.Type, indicator.Name)
	}
	for _, strategy := range cfg.Strategies {
		fmt.Printf("Strategy: %+v %+v\n\n", strategy.Type, strategy.Name)
	}
	for _, broker := range cfg.Brokers {
		fmt.Printf("Broker: %+v %+v\n\n", broker.Type, broker.Name)
	}
}
