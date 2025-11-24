package series

import (
	"math/rand"
	"sort"
	"time"

	"github.com/rangertaha/gotal/internal/pkg/tick"
)

func Sort(ticks []*tick.Tick) []*tick.Tick {
	sort.Slice(ticks, func(i, j int) bool {
		return ticks[i].Timestamp().Before(ticks[j].Timestamp())
	})
	return ticks
}

func Random(name string, start, end time.Time, duration time.Duration, fields []string, opts ...SeriesOptions) (s *Series) {

	s = &Series{
		name:  name,
		ticks: make([]*tick.Tick, 0),
		tags:  make(map[string]string),
	}

	// Generate random ticks between start and end time at given duration intervals
	for t := start; t.Before(end); t = t.Add(duration) {
		// Create new tick with timestamp and duration
		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
		)

		// Generate random values for each field
		values := make(map[string]float64)
		for _, field := range fields {

			// Generate random value between 1-100
			values[field] = rand.Float64()*99 + 1
		}
		tick.SetFields(values)

		// Add tick to series
		s.ticks = append(s.ticks, tick)
	}

	// Sort ticks by timestamp
	sort.Slice(s.ticks, func(i, j int) bool {
		return s.ticks[i].Timestamp().Before(s.ticks[j].Timestamp())
	})

	for _, opt := range opts {
		opt(s)
	}

	return

}
