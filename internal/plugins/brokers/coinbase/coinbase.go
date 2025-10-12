package coinbase

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/tick"
	"github.com/rangertaha/gotal/internal/plugins/brokers"
)

type coinbase struct {
	Name  string `hcl:"name,optional"`  // name of the data series
	Input string `hcl:"input,optional"` // field to compute the MACD on

}

func New(opts ...internal.OptFunc) *coinbase {
	cfg := opt.New(opts...)

	return &coinbase{
		Name:  cfg.Name("macd"),
		Input: cfg.Field("value"),
	}
}

func (i *coinbase) Compute(input *series.Series) (output *series.Series) {
	output = series.New(i.Name)

	return
}

func (i *coinbase) Process(input *tick.Tick) (output *tick.Tick) {

	return
}

func init() {
	brokers.Add("coinbase", func(opts ...internal.OptFunc) internal.Broker {
		return New(opts...)
	})
}
