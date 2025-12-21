package gotal

import (
	"github.com/rangertaha/gotal/internal/plugins/brokers"
	_ "github.com/rangertaha/gotal/internal/plugins/brokers/all"
)

var (
	Coinbase brokers.BrokerFunc
)

func init() {
	Coinbase = brokers.Func("coinbase")
}
