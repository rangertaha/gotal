package macd

import (
	"errors"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/plugins/strategies"
)

const (
	PluginID          = "MACD"
	PluginName        = "Moving Average Convergence Divergence."
	PluginDescription = "MAC is used to identify trends in the price of a security."
	PluginHCL         = `
strategy "macd" {
	fast = 12  # Fast period
	slow = 26  # Slow period
	signal = 9  # Signal period
}
`
)

type macd struct {

	// Connection parameters
	FastPeriod   int `hcl:"fast"`   // Fast period
	SlowPeriod   int `hcl:"slow"`   // Slow period
	SignalPeriod int `hcl:"signal"` // Signal period
}

func New(opts ...internal.ConfigOption) (internal.Plugin, error) {

	p := &macd{
		FastPeriod:   12,
		SlowPeriod:   26,
		SignalPeriod: 9,
	}
	return p
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
