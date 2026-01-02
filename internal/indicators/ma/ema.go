package ma

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/config"
	"github.com/rangertaha/gotal/internal/indicators"
	"github.com/rangertaha/gotal/internal/series"
)

const emaPluginID = "EMA"

type ema struct {
	Name   string `hcl:"name,optional"`   // name of the series
	Source string `hcl:"source,optional"` // source of the series

	// algorithm parameters
	Period int     `hcl:"period"`         // period to compute the EMA
	Alpha  float32 `hcl:"alpha,optional"` // alpha to compute the EMA

	// series internal state
	series internal.Series
}

func NewEMA(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := config.New(opts...)


	p := &ema{
		Name:   c.GetStr("name", "ema"),
		Source: c.GetStr("source", "value"),
		Period: c.GetInt("period", 14),
		Alpha:  c.GetFloat("alpha", 0.1),
	}
	p.series = series.New(p.Name)

	return p, nil
}

// func (i *ema) Init(config internal.Configurator) error {
// 	// i.Name = config.Get("name").(string)
// 	// i.Source = config.Get("source").(string)
// 	// i.Period = config.Get("period").(int)
// 	// i.Alpha = config.Get("alpha").(float64)

// 	return nil
// }

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
	if err := indicators.Add(emaPluginID, NewEMA, indicators.TREND); err != nil {
		panic(err)
	}
}
