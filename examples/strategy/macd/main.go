// This example shows how to use the MACD strategy.
// It creates a time series of prices and then calculates the MACD.
// It then prints the result.

// Example:
// go run examples/strategy/macd/main.go

// Output:
//
//

package main

import (
	"time"

	"github.com/rangertaha/gotal/pkg/brokers"
	"github.com/rangertaha/gotal/pkg/indicators"
	"github.com/rangertaha/gotal/pkg/providers"
	"github.com/rangertaha/gotal/pkg/strategies"
)

func main() {

	// Download historical ohlcv prices
	ohlcv := providers.Mock(
		providers.WithDataset("sine"),
		providers.WithName("ohlcv"),
		providers.WithCache("ohlcv.csv"),
		providers.WithSymbol("BTC"),
		providers.WithDuration(1*time.Minute),
		providers.WithStartDate(time.Now().AddDate(-10, 0, 0)),
		providers.WithEndDate(time.Now()))

	// Create a new macd indicator
	macd := indicators.MACD(ohlcv,
		indicators.OnField("close"),
		indicators.WithFastPeriod(12),
		indicators.WithSlowPeriod(26),
		indicators.WithSignalPeriod(9),
	)

	// Get the broker account client
	account := brokers.Mock()

	// Create a new strategy
	stratagy := strategies.MACD(
		strategies.WithIndicator(macd),
		strategies.WithBroker(account),
	)

	// Print the strategy results
	stratagy.Print()

}
