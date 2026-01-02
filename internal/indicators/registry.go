package indicators

import (
	"fmt"
	"strings"

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

// type IndicatorPlugin func(opts ...internal.ConfigOption) (internal.Plugin, error)

type indicatorPlugin func(opts ...internal.ConfigOption) (internal.Indicator, error)

type indicatorFunc func(series internal.Series, opts ...internal.ConfigOption) (internal.Series, error)

var (
	INDICATORS = map[string]indicatorPlugin{}
	GROUPS     = map[GroupType][]indicatorPlugin{}
)

func Add(id string, plugin indicatorPlugin, groups ...GroupType) error {
	id = strings.ToLower(id)

	if _, ok := INDICATORS[id]; ok {
		return fmt.Errorf("indicator %s already exists", id)
	}
	INDICATORS[id] = plugin

	// add indicator to groups
	for _, group := range groups {
		if _, ok := GROUPS[group]; !ok {
			GROUPS[group] = []indicatorPlugin{}
		}
		GROUPS[group] = append(GROUPS[group], plugin)
	}
	return nil
}

func Get(id string) indicatorPlugin {
	id = strings.ToLower(id)

	if plugin, ok := INDICATORS[id]; ok {
		return plugin
	}
	panic(fmt.Errorf("indicator %s not found", id))
}

// Group returns all indicators in a group
func Group(id GroupType) []indicatorPlugin {
	if group, ok := GROUPS[id]; ok {
		return group
	}
	panic(fmt.Errorf("group %s not found", id))
}

func All() (indicatorPlugins []indicatorPlugin) {
	for _, plugin := range INDICATORS {
		indicatorPlugins = append(indicatorPlugins, plugin)
	}
	return indicatorPlugins
}

func Func(id string) internal.IndicatorFunc {
	pluginFunc := Get(id)



	return func(input internal.Series, opts ...internal.ConfigOption) (internal.Series, error) {
		plugin, err := pluginFunc(opts...)
		if err != nil {
			return nil, err
		}


		return plugin.Compute(input), nil
	}
}

