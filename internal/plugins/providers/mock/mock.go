package macd

import (
	"time"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/tick"
	"github.com/rangertaha/gotal/internal/plugins/providers"
)

type mock struct {
	Name  string `hcl:"name,optional"`  // name of the data series
	Field string `hcl:"input,optional"` // field name

	Cache     string        `hcl:"cache,optional"`      // cache file name
	Dataset   string        `hcl:"dataset,optional"`    // dataset name
	Symbol    string        `hcl:"symbol,optional"`     // symbol name
	StartDate time.Time     `hcl:"start_date,optional"` // start date
	EndDate   time.Time     `hcl:"end_date,optional"`   // end date
	Interval  time.Duration `hcl:"interval,optional"`   // interval

}

func New(opts ...internal.OptFunc) *mock {
	cfg := opt.New(opts...)

	return &mock{
		Name:      cfg.Name("mock"),
		Field:     cfg.Field("value"),
		Cache:     cfg.GetString("cache", "mock"),
		Dataset:   cfg.GetString("dataset", "mock"),
		Symbol:    cfg.GetString("symbol", "ASSET"),
		StartDate: cfg.GetTime("start_date", time.Now().AddDate(-1, 0, 0)),
		EndDate:   cfg.GetTime("end_date", time.Now()),
		Interval:  cfg.Duration("interval", 1*time.Minute),
	}
}

func (p *mock) Compute(input *series.Series) (output *series.Series) {
	output = series.New(p.Name)

	return
}

func (p *mock) Process(input *tick.Tick) (output *tick.Tick) {

	return
}

func init() {
	providers.Add("mock", func(opts ...internal.OptFunc) internal.Provider {
		return New(opts...)
	})
}
