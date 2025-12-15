package ma

import (
	"math"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/opt"
	"github.com/rangertaha/gotal/internal/plugins"
	"github.com/rangertaha/gotal/internal/plugins/indicators"
	"github.com/rangertaha/gotal/internal/series"
	"github.com/rangertaha/gotal/internal/tick"
)

// EMAt​=α⋅Pricet​+(1−α)⋅EMAt−1
// EMA = (Close - EMA(previous day)) * multiplier + EMA(previous day)
// multiplier = 2 / (N + 1)
// N is the number of days in the EMA
// EMA(previous day) is the EMA of the previous day
// Close is the closing price of the current day
const emaPluginID = "EMA"
const emaPluginName = "Exponential Moving Average"
const emaPluginDescription = "Exponential Moving Average is a technical indicator that smooths out price data by giving more weight to recent prices."
const emaPluginHCL = `
indicator "ema" {
  period = 14
  alpha = 0.1
}
`

type ema struct {
	plugins.Plugin

	// EMA parameters
	Period int     `hcl:"period"`         // period to compute the EMA
	Alpha  float64 `hcl:"alpha,optional"` // alpha to compute the EMA
}

func emaNew(opts ...internal.PluginOptions) internal.Plugin {
	params := opt.New(opts...)

	e := &ema{
		Plugin: plugins.Plugin{
			PID:         emaPluginID,
			Title:       emaPluginName,
			Summary:     emaPluginDescription,
			Template:    emaPluginHCL,
			Params:      params,
			Series:      params.Series("series", nil),
			Fields:      []string{params.String("input", "value")}, // input field names to compute the EMA
			Initialized: false,
		},
		Period: params.Int("period", math.NaN()),
		Alpha:  params.Float("alpha", 2/(float64(params.Int("period", math.NaN()))+1)),
	}

	// e.Init(params...)


	return e
}

func (i *ema) Init(opts ...internal.PluginOptions) error {
	// i.Params = opt.New(opts...)

	// i.Initialized = false
	// i.Period = i.Params.Int("period", math.NaN())
	// i.Alpha = i.Params.Float("alpha", 2/(float64(i.Period)+1))
	// i.Series = i.Params.Series("series", nil)
	// i.Fields = []string{i.Params.String("input", "value")} // input field names to compute the EMA

	return nil
}

func (i *ema) Compute(input *series.Series) (output *series.Series) {
	output = series.New(i.ID())
	return
}

func (i *ema) Process(input *tick.Tick) (output *tick.Tick) {
	i.Series.Push(input)
	return i.Series.Pop()
}

func init() {
	indicators.Add(emaPluginID, emaNew, indicators.TREND)
}
