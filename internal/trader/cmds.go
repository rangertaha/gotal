package trader

import (
	"time"

	"github.com/rangertaha/gotal/internal"
)

func Load(opts ...func(t *trader)) (trader internal.Trader, err error) {
	trader, err = New(opts...)
	return
}

func Init(paths ...string) error {
	return nil
}

func Fill(start, end time.Time, duration time.Duration, provider string, opts ...func(t *trader)) error {
	trader, err := Load(opts...)
	if err != nil {
		return err
	}
	return trader.Fill(start, end, duration, provider)
}

func Train(start, end time.Time, opts ...func(t *trader)) error {
	trader, err := Load(opts...)
	if err != nil {
		return err
	}
	return trader.Train(start, end)
}

func Test(start, end time.Time, opts ...func(t *trader)) error {
	trader, err := Load(opts...)
	if err != nil {
		return err
	}
	return trader.Test(start, end)
}

func Live(start, end time.Time, opts ...func(t *trader)) error {
	trader, err := Load(opts...)
	if err != nil {
		return err
	}
	return trader.Live(start, end)
}

func Exec(start, end time.Time, opts ...func(t *trader)) error {
	trader, err := Load(opts...)
	if err != nil {
		return err
	}
	return trader.Exec(start, end)
}
