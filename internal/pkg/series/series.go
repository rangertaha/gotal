package series

import (
	"time"

	"github.com/rangertaha/gotal/internal/pkg/tick"
)

// Series represents a collection of market events, capturing the most granular form of market data.
type Series struct {
	name  string
	ticks []*tick.Tick

	// metadata
	tags     map[string]string
}

// NewSeries creates a new Series of ticks
func New(name string, opts ...SeriesOptions) (s *Series) {
	s = &Series{
		name:  name,
		ticks: make([]*tick.Tick, 0),
		tags:  make(map[string]string),
	}

	for _, opt := range opts {
		opt(s)
	}

	return

}

// Name returns the name of the Series collection.
func (s *Series) Name() string {
	return s.name
}

// SetName sets the name of the Series collection.
func (s *Series) SetName(name string) {
	s.name = name
}

// Duration returns the duration of the Series collection.
func (s *Series) Duration() time.Duration {
	if len(s.ticks) == 0 {
		return 0
	}
	return s.ticks[len(s.ticks)-1].Duration()
}

// SetDuration sets the duration of the Series collection.
func (s *Series) SetDuration(duration time.Duration) {
	// s.duration = duration
	for _, tick := range s.ticks {
		tick.SetDuration(duration)
	}
}

// Timestamps returns the timestamps of the Series collection.
func (s *Series) Timestamp() time.Time {
	if len(s.ticks) == 0 {
		return time.Time{}
	}
	return s.ticks[len(s.ticks)-1].Timestamp()
}

func (s *Series) TimeRange() (time.Time, time.Time) {
	if len(s.ticks) == 0 {
		return time.Time{}, time.Time{}
	}
	return s.ticks[0].Timestamp(), s.ticks[len(s.ticks)-1].Timestamp()
}

func (s *Series) Timestamps() (out []time.Time) {
	for _, tick := range s.ticks {
		out = append(out, tick.Timestamp())
	}
	return out
}

// Reset the ticks only to nil
func (s *Series) Reset() *Series {
	s.ticks = make([]*tick.Tick, 0)
	return s
}

// Set the ticks
func (s *Series) Set(ticks ...*tick.Tick) {
	s.ticks = Sort(ticks)
	s.update()
}

// Append adds one or more ticks to the end of the collection.
func (s *Series) Add(ticks ...*tick.Tick) *Series {
	s.ticks = append(s.ticks, ticks...)
	s.ticks = Sort(s.ticks)
	s.update()
	return s
}

// Update updates the series with the given ticks.
func (s *Series) Update(seriesInputs ...*Series) *Series {
	for _, seriesInput := range seriesInputs {
		for _, tickInput := range seriesInput.Ticks() {
			for i, tick := range s.ticks {
				if tick.Timestamp().Equal(tickInput.Timestamp()) {
					s.ticks[i].Update(tickInput)
				}
			}
		}
	}

	s.ticks = Sort(s.ticks)
	s.update()
	return s
}

// Len returns the number of ticks in the collection.
func (s *Series) Len() int {
	return len(s.ticks)
}

// IsEmpty returns true if the series is empty.
func (s *Series) IsEmpty() bool {
	return len(s.ticks) == 0
}

// Head returns the first n ticks.
func (s *Series) Head(n int) *Series {
	return &Series{ticks: s.ticks[n:], name: s.name, tags: s.tags}
}

// Tail returns the last n ticks.
func (s *Series) Tail(n int) *Series {
	return &Series{ticks: s.ticks[:n], name: s.name, tags: s.tags}
}

// Range returns a new Series collection containing ticks between start and end indices.
func (s *Series) Slice(start, end int) *Series {
	return &Series{name: s.name, ticks: s.ticks[start:end], tags: s.tags}
}

// Series returns the ticks of the Series collection.
func (s *Series) Ticks() []*tick.Tick {
	if len(s.ticks) == 0 {
		return []*tick.Tick{}
	}
	return s.ticks
}

// Copy returns a new Series collection with the same ticks.
func (s *Series) Copy(opts ...SeriesOptions) *Series {
	series := &Series{
		name:  s.name,
		ticks: s.ticks,
		// meta:         s.meta,
		tags: s.tags,
		// fields:       s.fields,
		// duration:     s.duration,
		// timestamp:    s.timestamp,
		// defaultField: s.defaultField,
	}

	for _, opt := range opts {
		opt(series)
	}

	return series
}

// At returns the tick at the specified index.
func (s *Series) At(index int) *tick.Tick {
	return s.ticks[index]
}

func (s *Series) AtTime(timestamp time.Time) *tick.Tick {
	for _, tick := range s.ticks {
		if tick.Timestamp().Equal(timestamp) {
			return tick
		}
	}
	return nil
}

// Pop returns the last tick and removes it from the Series collection.
func (s *Series) Pop() *tick.Tick {
	last := s.At(s.Len() - 1)
	s.ticks = s.ticks[:s.Len()-1]
	return last
}

// Push adds a tick to the beginning of the Series collection.
func (s *Series) Push(ticks ...*tick.Tick) {
	s.ticks = append(ticks, s.ticks...)
}

// Shift removes the first tick and returns it.
func (s *Series) Shift(tick *tick.Tick) *tick.Tick {
	s.Push(tick)
	return s.Pop()
}

// Apply applies the given options to each tick in the series.
func (s *Series) Apply(opts ...tick.TickOptions) {
	for _, tick := range s.ticks {
		for _, opt := range opts {
			opt(tick)
		}
	}
	s.ticks = Sort(s.ticks)
}

func (s *Series) Moving(start, stop, window int) (series []*Series) {
	if stop < 0 {
		stop = s.Len() - 1 + stop
	}

	for i := start; i < stop-window; i++ {
		series = append(series, s.Slice(i, i+window))
	}
	return series
}

func (s *Series) update() {
	// s.timestamp = tick.Timestamp()
	// s.duration = tick.Duration()

	// update the tags with common tags from the ticks
	if len(s.ticks) > 0 {
		s.tags = s.ticks[len(s.ticks)-1].Tags()
	}
}

func (s *Series) Spawn(opts ...SeriesOptions) *Series {
	series := &Series{
		name:   s.name,
		ticks:  make([]*tick.Tick, 0),
		tags:   s.tags,
	}

	for _, opt := range opts {
		opt(series)
	}

	return series
}
