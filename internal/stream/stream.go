package stream

import (
	"fmt"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/tick"
)

// Ticks represents a collection of market events, capturing the most granular form of market data.
type Stream struct {
	name  string
	ticks <-chan *tick.Tick

	// metadata
	tags map[string]string
}

func New(name string, opts ...StreamOptions) internal.Stream {
	s := &Stream{
		name:  name,
		ticks: make(chan *tick.Tick, 0),
		tags:  make(map[string]string),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s

}

// Name returns the name of the Ticks collection.
func (t *Stream) Name() string {
	return t.name
}

// SetName sets the name of the Ticks collection.
func (t *Stream) SetName(name string) {
	t.name = name
}

func (t *Stream) Ticks() <-chan *tick.Tick {
	return t.ticks
}

func (t *Stream) Add(ticks ...*tick.Tick) {
	ch := make(chan *tick.Tick)
	for _, tick := range ticks {
		ch <- tick
	}
	close(ch)
	t.ticks = ch
}

// Save saves the Series collection to a file.
func (t *Stream) Save(filename string, outputs ...string) error {
	return nil
}

// Print prints the Series collection to the console.
func (t *Stream) Print() {
	for tick := range t.ticks {
		fmt.Printf("%+v\n", tick)
	}
}

func (t *Stream) Ready() bool {
	return len(t.ticks) == 0
}