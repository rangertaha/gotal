package trader

import (
	"time"

	"github.com/rangertaha/gotal/internal"
)

func Load(opts ...func(t *trader)) internal.Trader {
	trader := New(opts...)
	return trader
}

func Backfill(start, end time.Time, duration time.Duration, provider string, opts ...func(t *trader)) error {
	trader := Load(opts...)
	return trader.Backfill(start, end, duration, provider)
}

func Train(start, end time.Time, opts ...func(t *trader)) error {
	trader := Load(opts...)
	return trader.Train(start, end)
}

func Test(start, end time.Time, opts ...func(t *trader)) error {
	trader := Load(opts...)
	return trader.Test(start, end)
}

func Live(start, end time.Time, opts ...func(t *trader)) error {
	trader := Load(opts...)
	return trader.Live(start, end)
}

func Exec(start, end time.Time, opts ...func(t *trader)) error {
	trader := Load(opts...)
	return trader.Exec(start, end)
}
