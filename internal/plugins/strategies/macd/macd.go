package macd

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/tick"
	"github.com/rangertaha/gotal/internal/plugins/strategies"
)

type macd struct {
	Name  string `hcl:"name,optional"`  // name of the data series
	Input string `hcl:"input,optional"` // field to compute the MACD on
}

func New(opts ...internal.OptFunc) *macd {
	cfg := opt.New(opts...)

	return &macd{
		Name:  cfg.Name("macd"),
		Input: cfg.Field("value"),
	}
}

func (i *macd) Compute(input *series.Series) (output *series.Series) {

	return
}

func (i *macd) Process(input *tick.Tick) (output *tick.Tick) {

	return
}

func init() {
	strategies.Add("macd", func(opts ...internal.OptFunc) internal.Strategy {
		return New(opts...)
	})

}
