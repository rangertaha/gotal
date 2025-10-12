package macd

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/tick"
	"github.com/rangertaha/gotal/internal/plugins/providers"
)

type polygon struct {
	Name  string `hcl:"name,optional"`  // name of the data series
	Input string `hcl:"input,optional"` // field to compute the MACD on

}

func New(opts ...internal.OptFunc) *polygon {
	cfg := opt.New(opts...)

	return &polygon{
		Name:  cfg.Name("macd"),
		Input: cfg.Field("value"),
	}
}

func (i *polygon) Compute(input *series.Series) (output *series.Series) {
	output = series.New(i.Name)

	return
}

func (i *polygon) Process(input *tick.Tick) (output *tick.Tick) {

	return
}

func init() {
	providers.Add("polygon", func(opts ...internal.OptFunc) internal.Provider {
		return New(opts...)
	})
}
