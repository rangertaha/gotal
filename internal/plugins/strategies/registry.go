package strategies

import (
	"fmt"
	"strings"

	"github.com/rangertaha/gotal/internal"
)

type StrategyPlugin func(opts ...internal.ConfigOption) (internal.Plugin, error)

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

func Batch(name string) internal.BatchFunc {
	return func(opts ...internal.ConfigOption) (internal.Series, error) {
		return nil, nil
	}
}

func Stream(name string) internal.StreamFunc {
	return func(opts ...internal.ConfigOption) (internal.Stream, error) {
		return nil, nil
	}
}
