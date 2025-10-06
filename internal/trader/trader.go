package trader

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal/internal"
)

func New(opts ...internal.OptFunc) internal.Trader {
	return &trader{
		providers:  []internal.Provider{},
		brokers:    []internal.Broker{},
		strategies: []internal.Strategy{},
		storages:   []internal.Storage{},
	}
}

type trader struct {
	providers []internal.Provider
	brokers   []internal.Broker
	strategies []internal.Strategy
	storages   []internal.Storage
}

func (t *trader) Backfill(start, end time.Time) {
	fmt.Println("Backfill data from", start, "to", end)
}

func (t *trader) Train(start, end time.Time) {
	fmt.Println("Train model from", start, "to", end)
}

func (t *trader) Test(start, end time.Time) {
	fmt.Println("Test model from", start, "to", end)
}

func (t *trader) Live(end time.Time) {
	fmt.Println("Live trading to", end)
}

func (t *trader) Exec(end time.Time) {
	fmt.Println("Executing trader to", end)
}
