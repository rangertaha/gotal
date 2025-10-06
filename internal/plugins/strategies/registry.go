package indicators

import (
	"fmt"
	"strings"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/stream"
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

type NewIndicatorFunc func(opts ...internal.OptFunc) internal.Indicator

var INDICATORS = map[string]NewIndicatorFunc{}

var GROUPS = map[GroupType][]NewIndicatorFunc{}

// var IndicatorFn IndicatorFunc

func Add(name string, fn NewIndicatorFunc, groups ...GroupType) error {
	name = strings.ToLower(name)

	if _, ok := INDICATORS[name]; ok {
		return fmt.Errorf("indicator %s already exists", name)
	}
	// TODO: Need to convert internal.Indicator to IndicatorsFunc
	// For now, creating a placeholder
	// var indicatorsFunc IndicatorsFunc
	INDICATORS[name] = fn

	for _, group := range groups {
		if _, ok := GROUPS[group]; !ok {
			GROUPS[group] = []NewIndicatorFunc{}
		}
		GROUPS[group] = append(GROUPS[group], fn)
	}
	return nil
}

func Get(name string) (NewIndicatorFunc, error) {
	name = strings.ToLower(name)

	if indicator, ok := INDICATORS[name]; ok {
		return indicator, nil
	}
	return nil, fmt.Errorf("indicator %s not found", name)
}

// Group returns all indicators in a group
func Group(group GroupType) ([]NewIndicatorFunc, error) {
	if _, ok := GROUPS[group]; !ok {
		return nil, fmt.Errorf("group %s not found", group)
	}
	return GROUPS[group], nil
}

func Series(name string) (internal.SeriesFunc, error) {
	name = strings.ToLower(name)

	if indicator, ok := INDICATORS[name]; ok {
		indicatorFn := internal.SeriesFunc(func(input *series.Series, opts ...internal.OptFunc) (output *series.Series) {
			return indicator(opts...).Compute(input)
		})
		return indicatorFn, nil
	}
	return nil, fmt.Errorf("indicator %s not found", name)
}

func Stream(name string) (internal.StreamFunc, error) {
	name = strings.ToLower(name)

	if indicator, ok := INDICATORS[name]; ok {
		indicatorFn := internal.StreamFunc(func(input *stream.Stream, opts ...internal.OptFunc) (output *stream.Stream) {
			output = stream.New(name)
			go func() {
				ind := indicator(opts...)
				for t := range input.Ticks() {
					if tick := ind.Process(t); tick != nil {
						output.Add(tick)
					}
				}
			}()
			return output
		})
		return indicatorFn, nil
	}
	return nil, fmt.Errorf("indicator %s not found", name)
}
