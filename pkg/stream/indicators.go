package batch

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/plugins/indicators"
	_ "github.com/rangertaha/gotal/internal/plugins/indicators/all"
)

var (
	SMA, EMA, RSI, MACD internal.StreamFunc
)

func init() {
	SMA = indicators.Stream("sma")
	EMA = indicators.Stream("ema")
	RSI = indicators.Stream("rsi")
	MACD = indicators.Stream("macd")
}
