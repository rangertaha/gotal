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

type IndicatorPlugin func(opts ...internal.ConfigOption) (internal.Plugin, error)

type IndicatorFunc func(opts ...internal.ConfigOption) (internal.Series, internal.Stream, error)

var (
	INDICATORS = map[string]IndicatorPlugin{}
	GROUPS     = map[GroupType][]IndicatorPlugin{}
)

func Add(id string, plugin IndicatorPlugin, groups ...GroupType) error {
	id = strings.ToLower(id)

	if _, ok := INDICATORS[id]; ok {
		return fmt.Errorf("indicator %s already exists", id)
	}
	INDICATORS[id] = plugin

	for _, group := range groups {
		if _, ok := GROUPS[group]; !ok {
			GROUPS[group] = []IndicatorPlugin{}
		}
		GROUPS[group] = append(GROUPS[group], plugin)
	}
	return nil
}

func Get(id string) (IndicatorPlugin, error) {
	id = strings.ToLower(id)

	if plugin, ok := INDICATORS[id]; ok {
		return plugin, nil
	}
	return nil, fmt.Errorf("indicator %s not found", id)
}

// Group returns all indicators in a group
func Group(id GroupType) ([]IndicatorPlugin, error) {
	if group, ok := GROUPS[id]; ok {
		return group, nil
	}
	return nil, fmt.Errorf("group %s not found", id)
}

func All() (indicatorPlugins []IndicatorPlugin) {
	for _, plugin := range INDICATORS {
		indicatorPlugins = append(indicatorPlugins, plugin)
	}
	return indicatorPlugins
}

func Func(name string) IndicatorFunc {
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
