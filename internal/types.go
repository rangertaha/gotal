package internal

import (
	"time"

	"github.com/hashicorp/hcl/v2"
)

type Configurator interface {
	Set(key string, value any)
	Get(key string, defaults ...any) any
	Merge(body hcl.Body)
	Decode(ctx *hcl.EvalContext, cfgs ...string) error
}
type ConfigOption func(Configurator) error

type PluginFunc func(...ConfigOption) (Series, Stream, error)

type BatchFunc func(...ConfigOption) (Series, error)

type StreamFunc func(...ConfigOption) (Stream, error)

type Plugin interface {
	ID() string
	Title() string
	Description() string
}

type Initializer interface {
	Init(Configurator) error
}

type Provider interface {
}

type Processor interface {
	Process(Tick) Tick
	Stream() Stream
	Compute() Series
	Reset() error
	Ready() bool
}

// type Streamer interface {
// 	Stream(input Stream) Stream
// }

// type Serieser interface {
// 	Series(input Series) Series
// }

//	type Ticker interface {
//		Tick(input Tick) Tick
//	}
//
// Series is a collection of ticks
type Series interface {
	Name(names ...string) string

	// Crud methods
	Get(index int) Tick
	Add(ticks ...Tick) error
	Delete(index ...int) error
	Update(ticks ...Tick) error

	// Retrieval methods
	Ticks(index ...int) []Tick // returns all or a subset of tickskl
	Head(n int) Series
	Tail(n int) Series
	Slice(start, end int) Series
	Copy() Series

	// // Time methods
	// Duration() time.Duration
	// SetDuration(duration time.Duration)
	// Timestamp() time.Time
	// TimeRange() (time.Time, time.Time)
	// Timestamps() []time.Time

	// Collection operations
	Len() int
	IsEmpty() bool

	// // Access methods
	// At(index int) Tick
	// AtTime(timestamp time.Time) Tick

	// // Collection manipulation
	// Head(n int) Series
	// Tail(n int) Series
	// Slice(start, end int) Series
	// Copy() Series

	// // Data operations
	// Pop() Tick
	// Push(ticks ...Tick)

	// // Utility
	// Spawn() Series

	// Output methods
	Print()
}

// Stream is a channel of ticks
type Stream interface {
	// Update(input Stream) Stream
	// AddError(err error)
	// Start()               // Starts the stream processing
	// Stop()                // Stops the stream processing
	// Push(input Tick)      // Pushes a tick into the stream
	// Pop() Tick            // Pops the next tick from the stream (if applicable)
	// Channel() <-chan Tick // Returns the underlying channel of ticks (read-only)
	// Error() <-chan error  // Returns the underlying error channel (read-only)
	// Close()               // Closes the stream and all resources
	// IsClosed() bool       // Checks if the stream is closed
	// Len() int             // Returns the current length of the stream buffer (if buffered)
	Ready() bool // Returns true if the stream is ready
	Print()      // Prints the stream to the console
}

type Tick interface {
	// Interface compliance methods
	Update(input Tick) Tick
	// AddError(err error)

	// Core identification
	ID() string
	SetID(id string)

	// Time methods
	Time() time.Time
	SetTime(timestamp time.Time)
	Epock() int64
	SetEpock(epock int64)
	Duration() time.Duration
	SetDuration(duration time.Duration)

	// Field methods
	Fields() map[string]float64
	GetField(key string) float64
	SetField(key string, value float64)
	SetFields(fields map[string]float64)
	HasField(key string) bool
	HasFields(keys ...string) bool
	RemoveField(key string)
	FieldNames() []string

	// Utility methods
	Len() int
	IsEmpty() bool
	Reset()
	ForEach(fn func(key string, value float64) float64)
}
