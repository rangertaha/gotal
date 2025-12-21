package ma

import (
	"math"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/opt"
	"github.com/rangertaha/gotal/internal/plugins/indicators"
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

	// EMA parameters
	Period int     `hcl:"period"`         // period to compute the EMA
	Alpha  float64 `hcl:"alpha,optional"` // alpha to compute the EMA
}

func NewEMA(opts ...internal.ConfigOption) (internal.Plugin, error) {
	params := opt.New(opts...)

	e := &ema{
		Period: params.Get("period", math.NaN()).(int),
		Alpha:  params.Get("alpha", 2/(float64(params.Get("period", math.NaN()).(int))+1)).(float64),
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

func (i *ema) Init() error {
	// i.Params = opt.New(opts...)

	// i.Initialized = false
	// i.Period = i.Params.Int("period", math.NaN())
	// i.Alpha = i.Params.Float("alpha", 2/(float64(i.Period)+1))
	// i.Series = i.Params.Series("series", nil)
	// i.Fields = []string{i.Params.String("input", "value")} // input field names to compute the EMA

	return nil
}

func (i *ema) Reset() error {
	return nil
}

func (i *ema) Ready() bool {
	return true
}

func (i *ema) Compute(input internal.Series) (output internal.Series) {
	return input
}

func (i *ema) Process(input internal.Tick) (output internal.Tick) {
	return input
}

func init() {
	indicators.Add(emaPluginID, NewEMA, indicators.TREND)
}
