package internal

import (
	"time"
)

type Configurator interface {
	Set(key string, value any)
	Get(key string, defaults ...any) any

	// decoding methods
	GetStr(key string, defaults ...string) string
	GetInt(key string, defaults ...int) int
	GetFloat(key string, defaults ...float32) float32
	GetBool(key string, defaults ...bool) bool
	GetTime(key string, defaults ...time.Time) time.Time
	GetDuration(key string, defaults ...time.Duration) time.Duration

	// Dataset
	GetSeries(key string, defaults ...Series) Series
}
type ConfigOption func(Configurator) error

type PluginFunc func(...ConfigOption) (Series, Stream, error)

type BatchFunc func(...ConfigOption) (Series, error)

type IndicatorFunc func(Series, ...ConfigOption) (Series, error)

type StreamFunc func(...ConfigOption) (Stream, error)

type Initializer interface {
	Init(Configurator) error
}

type Indicator interface {
	Reset() error
	Ready() bool
	Process(Tick) Tick
	Compute(Series) Series
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

	FieldNames() []string

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
	Save(filename ...string) error
	Plot(fields ...string) Plot
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

type Plot interface {
	Show(width, height int) error
	Save(filename string, width, height int) error
}
