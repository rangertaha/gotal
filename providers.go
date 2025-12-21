package gotal

import (
	"github.com/rangertaha/gotal/internal/plugins/providers"
	_ "github.com/rangertaha/gotal/internal/plugins/providers/all"
)

var (
	Generator, Polygon, Yahoo, Binance providers.ProviderFunc
)

func init() {
	Generator = providers.Func("gen")
	// Polygon = providers.Func("polygon")
	// Yahoo = providers.Func("yahoo")
	// Binance = providers.Func("binance")
}
