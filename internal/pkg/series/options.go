package series

import (
	"github.com/rangertaha/gotal/internal/pkg/tick"
)

type SeriesOptions func(*Series)

func WithName(name string) SeriesOptions {
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

func WithTicks(ticks ...*tick.Tick) SeriesOptions {
	return func(s *Series) { s.ticks = ticks }
}

// func From(series *Series) SeriesOptions {
// 	return func(s *Series) {
// 		s.name = series.name
// 		s.duration = series.duration
// 		s.ticks = series.ticks
// 		s.meta = series.meta
// 		s.defaultField = series.defaultField
// 		s.tags = series.tags
// 	}
// }
