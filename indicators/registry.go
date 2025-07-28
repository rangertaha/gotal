package indicators

import (
	"fmt"

	"github.com/rangertaha/gotal"
)

type Series func([]*gotal.Metric, ...Opts) []*gotal.Metric

type Stream func(chan *gotal.Metric, ...Opts) chan *gotal.Metric

type indicator struct {
	Series Series
	Stream Stream
}

var Indicators = map[string]indicator{}

var Groups = map[string]map[string]indicator{}

func Add(name string, series Series, stream Stream, groups ...string) error {
	if _, ok := Indicators[name]; ok {
		return fmt.Errorf("indicator %s already exists", name)
	}
	Indicators[name] = indicator{
		Series: series,
		Stream: stream,
	}
	for _, group := range groups {
		if _, ok := Groups[group]; !ok {
			Groups[group] = map[string]indicator{}
		}
		Groups[group][name] = indicator{
			Series: series,
			Stream: stream,
		}
	}
	return nil
}

func Get(name string) (Series, Stream, error) {
	if indicator, ok := Indicators[name]; ok {
		return indicator.Series, indicator.Stream, nil
	}
	return nil, nil, fmt.Errorf("indicator %s not found", name)
}

func GetGroup(group string) ([]Series, []Stream, error) {
	if indicators, ok := Groups[group]; ok {
		series := make([]Series, 0, len(indicators))
		streams := make([]Stream, 0, len(indicators))
		for _, indicator := range indicators {
			series = append(series, indicator.Series)
			streams = append(streams, indicator.Stream)
		}
		return series, streams, nil
	}
	return nil, nil, fmt.Errorf("group %s not found", group)
}


