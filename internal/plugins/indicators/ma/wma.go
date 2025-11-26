package ma

import (
	"fmt"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/sig"
	"github.com/rangertaha/gotal/internal/pkg/tick"
	"github.com/rangertaha/gotal/internal/plugins/indicators"
)

type wma struct {
	Name   string `hcl:"name,optional"`                // name of the data series
	Input  string `hcl:"input,optional,default=value"` // input field name to compute the WMA
	Output string `hcl:"output,optional"`              // output field name for the WMA
	Period int    `hcl:"period"`                       // period to compute the WMA on

	// series is the series of ticks to compute the SMA on
	series *series.Series
}

func NewWMA(opts ...internal.OptFunc) *wma {
	cfg := opt.New(opts...)
	name := cfg.Name("wma")
	period := cfg.Period(24)

	return &wma{
		Name:   name,
		Period: period,
		Input:  cfg.Field("value"),
		Output: cfg.Output(fmt.Sprintf("%s%d", "wma", period)),

		series: series.New(name),
	}
}

func (i *wma) Compute(input *series.Series) (output *series.Series) {
	output = series.New(i.Name)
	i.series.Reset()

	for _, t := range input.Ticks() {
		if t := i.Process(t); !t.IsEmpty() {
			output.Add(t)
		}
	}

	return
}

func (i *wma) Process(input *tick.Tick) (output *tick.Tick) {
	// check if the series has the required field
	if !i.series.HasField(i.Input) {
		panic(fmt.Sprintf("series is missing field %v", i.Input))
	}

	// add the input tick to the series
	i.series.Push(input)

	// create a new empty tick
	output = tick.New()

	// if the series is not long enough, return false
	if i.series.Len() > i.Period {
		// calculate the average
		output = i.calculate(i.series)

		// remove the first tick from the series
		i.series.Pop()
	}

	return
}

func (i *wma) calculate(input *series.Series) (output *tick.Tick) {
	// calculate the average
	value := input.Mean(i.Input)

	// create a new tick with the SMA value
	output = tick.New(
		tick.WithTime(i.series.Timestamp()),
		tick.WithDuration(i.series.Duration()),
		tick.WithFields(map[string]float64{i.Output: value}),
		tick.WithTags(i.series.Tags()),
		tick.WithSignals(map[sig.Signal]sig.Strength{}))

	return
}

func init() {
	indicators.Add("wma", func(opts ...internal.OptFunc) internal.Indicator {
		return NewWMA(opts...)
	}, indicators.TREND)
}
