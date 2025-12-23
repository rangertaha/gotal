package macd

import (
	"errors"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/opt"
	"github.com/rangertaha/gotal/internal/plugins/strategies"
	"github.com/rangertaha/gotal/internal/series"
)

type macd struct {
	Name string `hcl:"name,optional"` // name of the strategy

	// connection parameters
	FastPeriod   int `hcl:"fast"`    // Fast period
	SlowPeriod   int `hcl:"slow"`    // Slow period
	SignalPeriod int `hcl:"signal"`  // Signal period

	Source  string `hcl:"source,optional"`  // Source field name (e.g. "close", "open", "high", "low", "volume")
	OMAType string `hcl:"omatype,optional"` // Moving average type for the oscillator (e.g. "ema", "sma", "wma")
	SMAType string `hcl:"smatype,optional"` // Moving average type for the signal (e.g. "ema", "sma", "wma")

	// series internal state
	series internal.Series
}

func New(opts ...internal.ConfigOption) (internal.Plugin, error) {

	p := &macd{
		FastPeriod:   12,
		SlowPeriod:   26,
		SignalPeriod: 9,

		series: series.New(PluginID),
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

func (p *macd) Init() error {
	if p.FastPeriod <= 0 {
		return errors.New("fast period is required")
	}
	if p.SlowPeriod <= 0 {
		return errors.New("slow period is required")
	}
	if p.SignalPeriod <= 0 {
		return errors.New("signal period is required")
	}
	return nil
}

func (i *macd) Compute(input internal.Series) (output internal.Series) {

	return
}

func (i *macd) Process(input internal.Tick) (output internal.Tick) {

	return
}

func init() {
	strategies.Add(PluginID, New)
}
