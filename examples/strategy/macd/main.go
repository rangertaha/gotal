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
	"fmt"
	"time"

	"github.com/rangertaha/gotal/internal/pkg/internal/tick"
	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/sig"
	"github.com/rangertaha/gotal/pkg/indicators"
	"github.com/rangertaha/gotal/pkg/providers"
	"github.com/rangertaha/gotal/pkg/strategies"
	"github.com/rangertaha/gotal/pkg/traders"
)

func main() {

	trader := traders.Trader("btc-001")

	// Download new price ohlcv bars
	client := providers.Polygon(providers.WithAPIKey("XXXX"))

	// Download ohlcv bars
	ohlcv := client.Download("ohlcv.csv",
		providers.WithDataset("aggregates"),
		providers.WithSymbol("BTC-USD"),
		providers.WithDuration(1*time.Minute),
		providers.WithStartDate(time.Now().Add(-time.Minute*1)),
		providers.WithEndDate(time.Now()))

	// Create a new macd indicator
	macd := indicators.MACD(ohlcv,
		indicators.OnField("close"),
		indicators.WithShortPeriod(12),
		indicators.WithLongPeriod(26),
		indicators.WithSignalPeriod(9),
	)

	// Get the account client
	account := brokers.Coinbase("XXXX", brokers.WithAPIKey("XXXX"))

	// Create a new strategy
	stratagy := strategies.New(
		strategies.WithName("macd01"),
		strategies.WithSeries(macd),
		strategies.WithAccount(account),
	)

	stratagy.OnTrain(func(train *strategies.Training, start time.Time, end time.Time) {
		fmt.Println(train)
	})

	stratagy.OnTest(func(test *strategies.Testing) {
		fmt.Println(test)
	})

	stratagy.OnLive(func(live *strategies.Living) {
		fmt.Println(live)
	})

	trader.Add(stratagy)

	// Start the trader
	trader.Start()

}
