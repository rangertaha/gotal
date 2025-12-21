package series

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rangertaha/gotal/internal"
)

// Series represents a collection of market events, capturing the most granular form of market data.
type Series struct {
	name  string
	ticks []internal.Tick
}

// NewSeries creates a new Series of ticks
func New(name string, opts ...Option) internal.Series {
	s := &Series{
		name:  name,
		ticks: make([]internal.Tick, 0),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func Sort(ticks []internal.Tick) []internal.Tick {
	sort.Slice(ticks, func(i, j int) bool {
		return ticks[i].Time().Before(ticks[j].Time())
	})
	return ticks
}

// Name returns the name of the Series collection.
func (s *Series) Name(names ...string) string {
	if len(names) > 0 {
		s.name = strings.TrimSpace(strings.Join(names, ""))
	}
	return s.name
}

// Get returns the tick at the specified index.
func (s *Series) Get(index int) internal.Tick {
	return s.ticks[index]
}

// Add adds one or more ticks to the end of the collection.
func (s *Series) Add(ticks ...internal.Tick) error {
	s.ticks = append(s.ticks, ticks...)
	return nil
}

// Delete deletes the tick at the specified index.
func (s *Series) Delete(index ...int) error {
	for _, idx := range index {
		if idx < 0 || idx >= len(s.ticks) {
			return fmt.Errorf("index out of range")
		}
		s.ticks = append(s.ticks[:idx], s.ticks[idx+1:]...)
	}
	return nil
}

// Update updates the tick at the specified index.
func (s *Series) Update(ticks ...internal.Tick) error {
	s.ticks = append(s.ticks, ticks...)
	return nil
}

// Series returns the ticks of the Series collection.
func (s *Series) Ticks(index ...int) (ticks []internal.Tick) {
	if len(index) > 0 {
		ticks = make([]internal.Tick, 0)
		for _, idx := range index {
			if idx <= len(s.ticks) {
				ticks = append(ticks, s.ticks[idx])
			}
		}
		return ticks
	}
	return s.ticks
}

// Head returns the first n ticks.
func (s *Series) Head(n int) internal.Series {
	return &Series{ticks: s.ticks[:n], name: s.name}
}

// Tail returns the last n ticks.
func (s *Series) Tail(n int) internal.Series {
	return &Series{ticks: s.ticks[len(s.ticks)-n:], name: s.name}
}

// Slice returns a new Series collection containing ticks between start and end indices.
func (s *Series) Slice(start, end int) internal.Series {
	return &Series{ticks: s.ticks[start:end], name: s.name}
}

// Copy returns a new Series collection with the same ticks.
func (s *Series) Copy() internal.Series {
	series := &Series{
		name:  s.name,
		ticks: s.ticks,
	}
	return series
}

// IsEmpty returns true if the Series collection is empty.
func (s *Series) IsEmpty() bool {
	return len(s.ticks) == 0
}

// Len returns the number of ticks in the Series collection.
func (s *Series) Len() int {
	return len(s.ticks)
}
