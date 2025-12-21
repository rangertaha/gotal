package gotal

import (
	"github.com/rangertaha/gotal/internal/plugins/indicators"
	_ "github.com/rangertaha/gotal/internal/plugins/indicators/all"
)

var (
	SMA, EMA, RSI, MACD indicators.IndicatorFunc
)

func init() {
	SMA = indicators.Func("ema")
	EMA = indicators.Func("ema")
	RSI = indicators.Func("rsi")
	MACD = indicators.Func("macd")
}
