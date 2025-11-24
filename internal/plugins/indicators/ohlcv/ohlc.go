package ohlcv

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/tick"
	"github.com/rangertaha/gotal/internal/plugins/indicators"
)

type ohlc struct {
	Name     string        `hcl:"name,optional,default=ohlc"`   // name of the data series
	Input    string        `hcl:"input,optional,default=value"` // input field name to compute the OHLCV on
	Duration time.Duration `hcl:"duration"`                     // period to compute the OHLCV on

	// series is the series of ticks to compute the OHLC on
	series *series.Series
}

func NewOHLC(opts ...internal.OptFunc) *ohlc {
	cfg := opt.New(opts...)
	name := cfg.Name("ohlc")
	duration := cfg.GetDuration("duration", 14*time.Minute)

	return &ohlc{
		Name:     name,
		Duration: duration,
		Input:    cfg.Field("value", "price"),

		series: series.New(name),
	}
}

func (i *ohlc) Compute(input *series.Series) (output *series.Series) {
	output = series.New(i.Name)
	i.series.Reset()

	for _, t := range input.Ticks() {
		if t := i.Process(t); !t.IsEmpty() {
			output.Add(t)
		}
	}

	return
}

func (i *ohlc) Process(input *tick.Tick) (output *tick.Tick) {

	// check if the series has the required field
	if !i.series.HasField(i.Input) {
		panic(fmt.Sprintf("series is missing field %v", i.Input))
	}

	// add the input tick to the series
	i.series.Push(input)

	if i.series.Len() == int(i.Duration.Seconds()) {
		output = tick.New(
			tick.WithTimestamp(input.Timestamp()),
			tick.WithDuration(input.Duration()),
			tick.WithFields(map[string]float64{
				"open":   i.series.First(i.Input),
				"high":   i.series.Max(i.Input),
				"low":    i.series.Min(i.Input),
				"close":  i.series.Last(i.Input),
				"volume": i.series.Sum(i.Input),
			}),
			tick.WithTags(input.Tags()),
		)

		i.series.Reset()
		return output
	}

	return nil
}

func init() {
	indicators.Add("ohlc", func(opts ...internal.OptFunc) internal.Indicator {
		return NewOHLC(opts...)
	})
}
