package trading

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	"github.com/rangertaha/gotal/internal/trader"
)

var (
	// Trading options
	WithBroker    = opt.WithBroker
	WithProvider  = opt.WithProvider
	WithStrategy  = opt.WithStrategy
	WithIndicator = opt.WithIndicator
)

func New(opts ...internal.OptFunc) internal.Trader {
	return trader.New(opts...)
}
