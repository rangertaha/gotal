package mock

import (
	"fmt"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/sig"
	"github.com/rangertaha/gotal/internal/pkg/tick"
	"github.com/rangertaha/gotal/internal/plugins/indicators"
)

// EMAt​=α⋅Pricet​+(1−α)⋅EMAt−1
// EMA = (Close - EMA(previous day)) * multiplier + EMA(previous day)
// multiplier = 2 / (N + 1)
// N is the number of days in the EMA
// EMA(previous day) is the EMA of the previous day
// Close is the closing price of the current day

type InputSeries struct {
	Name   string   `hcl:"name,optional"`
	Fields []string `hcl:"fields,optional"`
}

type output struct {
	Fields []string `hcl:"fields,optional"`
}

type mock struct {
	Name   string  `hcl:"name,optional,default=ema"` // name of the data series
	Input  input   `hcl:"input,optional"`            // input field name to compute the EMA
	Output output  `hcl:"output,optional"`           // output field name for the EMA
	Period int     `hcl:"period"`                    // period to compute the EMA
	Alpha  float64 `hcl:"alpha,optional"`            // alpha to compute the EMA

	// series is the series of ticks to compute the EMA
	series *series.Series
	// previousEMA stores the previous EMA value for calculation
	previousEMA float64
	// initialized tracks if we have enough data to start EMA calculation
	initialized bool
}

func NewMock(opts ...internal.OptFunc) *mock {
	cfg := opt.New(opts...)
	period := cfg.Period(14)

	return &mock{
		Name:   cfg.Name("ema"),
		Period: period,
		Input:  cfg.Field("value"),
		Output: cfg.Output(fmt.Sprintf("%s%d", "ema", period)),
		Alpha:  cfg.GetFloat("alpha", 2.0/(float64(period)+1.0)),

		series:      series.New(cfg.Name("ema")),
		initialized: false,
		previousEMA: 0,
	}
}

func (i *mock) Compute(input *series.Series) (output *series.Series) {
	output = series.New(i.Name)
	i.series.Reset()

	for _, t := range input.Ticks() {
		if t := i.Process(t); !t.IsEmpty() {
			output.Add(t)
		}
	}

	return
}

func (i *mock) Process(input *tick.Tick) (output *tick.Tick) {

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
		output = i.calculate(input)

		// remove the first tick from the series
		i.series.Pop()
	}

	return
}

func (i *mock) calculate(input *tick.Tick) (output *tick.Tick) {
	// get the current price value
	currentPrice := input.GetField(i.Input)

	var emaValue float64

	if !i.initialized {
		// for the first calculation, use the current price as the initial EMA
		emaValue = currentPrice
		i.initialized = true
	} else {
		// calculate EMA using the formula: EMA = α * Price + (1 - α) * PreviousEMA
		emaValue = i.Alpha*currentPrice + (1-i.Alpha)*i.previousEMA
	}

	// store the current EMA value for the next calculation
	i.previousEMA = emaValue

	// create a new tick with the EMA value
	output = tick.New(
		tick.WithTimestamp(input.Timestamp()),
		tick.WithDuration(input.Duration()),
		tick.WithFields(map[string]float64{i.Output: emaValue}),
		tick.WithTags(input.Tags()),
		tick.WithSignals(map[sig.Signal]sig.Strength{}))

	return
}

func init() {
	indicators.Add("ema", func(opts ...internal.OptFunc) internal.Indicator {
		return NewMock(opts...)
	}, indicators.TREND)
}
