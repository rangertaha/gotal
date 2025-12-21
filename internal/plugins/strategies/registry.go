package strategies

import (
	"fmt"

	"github.com/rangertaha/gotal/internal"
)


type PluginFunc func(opts ...internal.PluginOptions) internal.Plugin

var (
	STRATEGIES = map[string]PluginFunc{}
)

func Add(id string, plugin PluginFunc) error {

	if _, ok := STRATEGIES[id]; ok {
		return fmt.Errorf("strategy %s already exists", id)
	}
	STRATEGIES[id] = plugin

	return nil
}

func Get(id string) (PluginFunc, error) {
	if plugin, ok := STRATEGIES[id]; ok {
		return plugin, nil
	}
	return nil, fmt.Errorf("strategy %s not found", id)
}

func All() (plugins []internal.Plugin) {
	for _, plugin := range STRATEGIES {
		plugins = append(plugins, plugin())
	}
	return plugins
}
	