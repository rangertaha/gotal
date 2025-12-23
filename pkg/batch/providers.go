package batch

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/plugins/providers"
	_ "github.com/rangertaha/gotal/internal/plugins/providers/all"
)

var (
	Generator, Polygon, Yahoo, Binance internal.BatchFunc
)

func init() {
	Generator = providers.Batch("generator")
	// Polygon = providers.Func("polygon")
	// Yahoo = providers.Func("yahoo")
	// Binance = providers.Func("binance")
}
