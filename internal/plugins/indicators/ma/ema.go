package ma

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/opt"
	"github.com/rangertaha/gotal/internal/plugins/indicators"
	"github.com/rangertaha/gotal/internal/series"
)

type ema struct {

	// function parameters
	SeriesName string  `hcl:"name,optional"`  // name of the series
	Period     int     `hcl:"period"`         // period to compute the EMA
	Alpha      float64 `hcl:"alpha,optional"` // alpha to compute the EMA

	// series internal state
	series internal.Series
}

func NewEMA(opts ...internal.ConfigOption) (internal.Plugin, error) {
	p := &ema{
		Period: 14,
		Alpha:  0.1,
		series: series.New(emaPluginID),
	}

	config := opt.New(p)
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	if err := p.Init(); err != nil {
		return nil, err
	}

	return p, nil
}

func (i *ema) Init() error {

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
