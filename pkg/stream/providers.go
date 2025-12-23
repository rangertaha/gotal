package batch

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/plugins/providers"
	_ "github.com/rangertaha/gotal/internal/plugins/providers/all"
)

var (
	Generator, Polygon, Yahoo, Binance internal.StreamFunc
)

func init() {
	Generator = providers.Stream("generator")
	Polygon = providers.Stream("massive")
	Yahoo = providers.Stream("yahoo")
	Binance = providers.Stream("binance")
}
