package strategies

import (
	"fmt"
	"strings"

	"github.com/rangertaha/gotal/internal"
)

type StrategyPlugin func(opts ...internal.ConfigOption) (internal.Plugin, error)

type StrategyFunc func(opts ...internal.ConfigOption) (internal.Series, internal.Stream, error)

var (
	STRATEGIES = map[string]StrategyPlugin{}
)

func Add(id string, plugin StrategyPlugin) error {
	id = strings.ToLower(id)

	if _, ok := STRATEGIES[id]; ok {
		return fmt.Errorf("strategy %s already exists", id)
	}
	STRATEGIES[id] = plugin

	return nil
}

func Get(id string) (StrategyPlugin, error) {
	id = strings.ToLower(id)

	if plugin, ok := STRATEGIES[id]; ok {
		return plugin, nil
	}
	return nil, fmt.Errorf("strategy %s not found", id)
}

func All() (strategyPlugins []StrategyPlugin) {
	for _, plugin := range STRATEGIES {
		strategyPlugins = append(strategyPlugins, plugin)
	}
	return strategyPlugins
}

func Func(name string) StrategyFunc {
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
