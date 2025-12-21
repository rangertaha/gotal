package series

import (
	"time"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/tick"
)

type Option func(*Series)

// type SeriesOptions func(*Series) error

func WithName(name string) Option {
	return func(s *Series) { s.name = name }
}

// func WithMeta(meta map[string]any) SeriesOptions {
// 	return func(s *Series) { s.meta = meta }
// }

// func WithDuration(duration time.Duration) SeriesOptions {
// 	return func(s *Series) { s.duration = duration }
// }

// func WithPeriod(period int) SeriesOptions {
// 	return func(s *Series) { s.meta["period"] = period }
// }

// func WithFields(fields map[string][]float64) SeriesOptions {
// 	return func(s *Series) {
// 		s.fields = fields
// 	}
// }

func WithTicks(ticks ...internal.Tick) Option {
	return func(s *Series) { s.ticks = ticks }
}

func WithFields(fields ...map[string]float64) Option {
	return func(s *Series) {
		for _, field := range fields {
			tick := tick.New()
			for k, v := range field {
				if k == "time" {
					tick.SetTime(time.Unix(int64(v), 0))
				} else {
					tick.SetField(k, v)
				}
			}
			s.Add(tick)
		}
	}
}
