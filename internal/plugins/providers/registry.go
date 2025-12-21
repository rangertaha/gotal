package providers

import (
	"fmt"
	"strings"

	"github.com/rangertaha/gotal/internal"
)

type ProviderPlugin func(opts ...internal.ConfigOption) (internal.Plugin, error)

type ProviderFunc func(opts ...internal.ConfigOption) (internal.Series, internal.Stream, error)

var PROVIDERS = map[string]ProviderPlugin{}

func Add(id string, plugin ProviderPlugin) error {
	id = strings.ToLower(id)

	if _, ok := PROVIDERS[id]; ok {
		return fmt.Errorf("indicator %s already exists", id)
	}
	PROVIDERS[id] = plugin

	return nil
}

func Get(id string) (ProviderPlugin, error) {
	id = strings.ToLower(id)

	if plugin, ok := PROVIDERS[id]; ok {
		return plugin, nil
	}
	return nil, fmt.Errorf("provider %s not found", id)
}

func All() (providerPlugins []ProviderPlugin) {
	for _, plugin := range PROVIDERS {
		providerPlugins = append(providerPlugins, plugin)
	}
	return providerPlugins
}

func Func(name string) ProviderFunc {
	return func(opts ...internal.ConfigOption) (internal.Series, internal.Stream, error) {

		plg, err := Get(name)
		if err != nil {
			return nil, nil, err
		}
		plugin, err := plg(opts...)
		if err != nil {
			return nil, nil, err
		}

		if initializer, ok := plugin.(internal.Initializer); ok {
			if err := initializer.Init(); err != nil {
				return nil, nil, err
			}
		}

		if processor, ok := plugin.(internal.Processor); ok {
			return processor.Compute(), processor.Stream(), nil
		}

		return nil, nil, nil
	}
}
