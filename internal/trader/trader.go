package trader

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal/internal"
)

func New(opts ...func(t *trader)) internal.Trader {
	t := &trader{
		// asset
		asset: "",

		// plugins
		providers:  []internal.Provider{},
		brokers:    []internal.Broker{},
		strategies: []internal.Strategy{},
		storages:   []internal.Storage{},
		indicators: []internal.Indicator{},
	}

	for _, opt := range opts {
		opt(t)
	}

	return t
}

type trader struct {
	asset string

	// plugins
	providers  []internal.Provider
	brokers    []internal.Broker
	strategies []internal.Strategy
	storages   []internal.Storage
	indicators []internal.Indicator
}

func (t *trader) Init(paths ...string) error {
	fmt.Println("Initializing trader with paths", paths)
	return nil
}

func (t *trader) Fill(start, end time.Time, duration time.Duration, provider string) error {
	fmt.Println("Fill data from", start, "to", end, "duration", duration, "provider", provider)

	return nil
}

func (t *trader) Train(start, end time.Time) error {
	fmt.Println("Train model from", start, "to", end)
	return nil
}

func (t *trader) Test(start, end time.Time) error {
	fmt.Println("Test model from", start, "to", end)
	return nil
}

func (t *trader) Live(start, end time.Time) error {
	fmt.Println("Live trading from", start, "to", end)
	return nil
}

func (t *trader) Exec(start, end time.Time) error {
	fmt.Println("Executing trader from", start, "to", end)
	return nil
}
