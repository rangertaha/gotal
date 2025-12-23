package batch

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/plugins/brokers"
	_ "github.com/rangertaha/gotal/internal/plugins/brokers/all"
)

var (
	Coinbase internal.StreamFunc
)

func init() {
	Coinbase = brokers.Stream("coinbase")
}
