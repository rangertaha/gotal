package trader

import (
	"time"

	"github.com/rangertaha/gotal/pkg/trading"
)

func main() {
	start := time.Now().AddDate(-10, 0, 0)
	end := time.Now()

	// Create a new trader
	t := trading.New(
		trading.WithProvider("polygon"),
		trading.WithBroker("coinbase"),
		trading.WithStrategy("macd"),
	)
	t.Backfill(start, end)

	// Train the strategy model
	t.Train(start, end.AddDate(-1, 0, 0))

	// Test the strategy model
	t.Test(start.AddDate(-1, 0, 0), end.AddDate(-1, 0, 0))

	// Live trading with mock broker
	t.Live(end.AddDate(0, 0, 1))

	// Execute the trader
	t.Exec(end.AddDate(0, 0, 1))
}
