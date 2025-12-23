package providers

import (
	"fmt"
	"strings"

	"github.com/rangertaha/gotal/internal"
)

type ProviderFunc func(opts ...internal.ConfigOption) (internal.Plugin, error)

var PROVIDERS = map[string]ProviderFunc{}

func Add(id string, plugin ProviderFunc) error {
	id = strings.ToLower(id)

	if _, ok := PROVIDERS[id]; ok {
		return fmt.Errorf("indicator %s already exists", id)
	}
	PROVIDERS[id] = plugin

	return nil
}

func Get(id string) (ProviderFunc, error) {
	id = strings.ToLower(id)

	if plugin, ok := PROVIDERS[id]; ok {
		return plugin, nil
	}
	return nil, fmt.Errorf("provider %s not found", id)
}

func All() (providerPlugins []ProviderFunc) {
	for _, plugin := range PROVIDERS {
		providerPlugins = append(providerPlugins, plugin)
	}
	return providerPlugins
}

func Batch(name string) internal.BatchFunc {
	return func(opts ...internal.ConfigOption) (internal.Series, error) {
		fmt.Printf("Batch: %s\n", name)

		// Lookup the plugin
		plg, err := Get(name)
		if err != nil {
			return nil, err
		}

		// Create and initialize a new plugin instance
		plugin, err := plg(opts...)
		if err != nil {
			return nil, err
		}

		// If the plugin is a processor, compute the series
		if processor, ok := plugin.(internal.Processor); ok {
			return processor.Compute(), nil
		}
		fmt.Printf("Plugin: %+v\n", plugin)

		return nil, nil
	}
}

func Stream(name string) internal.StreamFunc {
	return func(opts ...internal.ConfigOption) (internal.Stream, error) {
		return nil, nil
	}
}
