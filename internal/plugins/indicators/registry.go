package indicators

import (
	"fmt"

	"github.com/rangertaha/gotal/internal"
)

type GroupType string

const (
	TREND      GroupType = "trend"
	MOMENTUM   GroupType = "momentum"
	VOLATILITY GroupType = "volatility"
	VOLUME     GroupType = "volume"
	CYCLE      GroupType = "cycle"
	OTHER      GroupType = "other"
)

type PluginFunc func(opts ...internal.PluginOptions) internal.Plugin

var (
	INDICATORS = map[string]PluginFunc{}
	GROUPS     = map[GroupType][]PluginFunc{}
)

func Add(id string, plugin PluginFunc, groups ...GroupType) error {

	if _, ok := INDICATORS[id]; ok {
		return fmt.Errorf("indicator %s already exists", id)
	}
	INDICATORS[id] = plugin

	for _, group := range groups {
		if _, ok := GROUPS[group]; !ok {
			GROUPS[group] = []PluginFunc{}
		}
		GROUPS[group] = append(GROUPS[group], plugin)
	}
	return nil
}

func Get(id string) (PluginFunc, error) {
	if plugin, ok := INDICATORS[id]; ok {
		return plugin, nil
	}
	return nil, fmt.Errorf("indicator %s not found", id)
}

// Group returns all indicators in a group
func Group(id GroupType) ([]PluginFunc, error) {
	if group, ok := GROUPS[id]; ok {
		return group, nil
	}
	return nil, fmt.Errorf("group %s not found", id)
}


func All() (plugins []internal.Plugin) {
	for _, plugin := range INDICATORS {
		plugins = append(plugins, plugin())
	}
	return plugins
}
