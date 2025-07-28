package example

import (
	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/indicators"
)

type example struct {
	metrics []*gotal.Metric
	opts    *indicators.Option
}

func Series(m []*gotal.Metric, opts ...indicators.Opts) []*gotal.Metric {
	opt, err := indicators.NewOption(opts...)
	if err != nil {
		panic(err)
	}
	example := &example{
		metrics: m,
		opts:    opt,
	}

	return example.Calculate()
}

func Stream(m chan *gotal.Metric, opts ...indicators.Opts) chan *gotal.Metric {
	opt, err := indicators.NewOption(opts...)
	if err != nil {
		panic(err)
	}
	example := &example{opts: opt}
	out := make(chan *gotal.Metric)

	for metric := range m {
		example.Add(metric)
		for _, metric := range example.Calculate() {
			out <- metric
		}
	}

	return out
}

func (e *example) Calculate() []*gotal.Metric {

	return e.metrics
}

func (e *example) Add(m *gotal.Metric) {
	e.metrics = append(e.metrics, m)
}

func init() {
	series := func(m []*gotal.Metric, opts ...indicators.Opts) []*gotal.Metric {
		// Default options
		opts = append(opts, indicators.Period(10))

		return Series(m, opts...)
	}
	stream := func(m chan *gotal.Metric, opts ...indicators.Opts) chan *gotal.Metric {
		// Default options
		opts = append(opts, indicators.Period(10))

		return Stream(m, opts...)
	}
	indicators.Add("example", series, stream)
}
