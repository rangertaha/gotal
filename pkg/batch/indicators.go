package batch

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/plugins/indicators"
	_ "github.com/rangertaha/gotal/internal/plugins/indicators/all"
)

var (
	SMA, EMA, RSI, MACD internal.BatchFunc
)

func init() {
	SMA = indicators.Batch("sma")
	EMA = indicators.Batch("ema")
	RSI = indicators.Batch("rsi")
	MACD = indicators.Batch("macd")
}
