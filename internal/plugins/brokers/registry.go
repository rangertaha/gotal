package brokers

import (
	"fmt"

	"github.com/rangertaha/gotal/internal"
)

type PluginFunc func(opts ...internal.PluginOptions) internal.Plugin

var (
	BROKERS	 = map[string]PluginFunc{}
)

func Add(id string, plugin PluginFunc) error {

	if _, ok := BROKERS[id]; ok {
		return fmt.Errorf("broker %s already exists", id)
	}
	return nil
}

func Get(id string) (PluginFunc, error) {
	if plugin, ok := BROKERS[id]; ok {
		return plugin, nil
	}
	return nil, fmt.Errorf("broker %s not found", id)
}

func All() (plugins []internal.Plugin) {
	for _, plugin := range BROKERS {
		plugins = append(plugins, plugin())
	}
	return plugins
}
