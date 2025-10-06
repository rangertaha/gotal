package internal

import (
	"time"

	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/stream"
	"github.com/rangertaha/gotal/internal/pkg/tick"
)

type Options interface {
	// Primitive objects
	Set(key string, value any)
	Get(key string, defaults ...any) any
	GetInt(key string, defaults ...any) int
	GetString(key string, defaults ...any) string
	GetFloat(key string, defaults ...any) float64
	GetBool(key string, defaults ...any) bool
	GetDuration(key string, defaults ...any) time.Duration
	GetTime(key string, defaults ...any) time.Time

	// Objects
	Period(n ...any) int
	FastPeriod(n ...any) int
	SlowPeriod(n ...any) int
	MAType(s ...any) string

	//
	Name(s ...any) string
	Suffix(s ...any) string
	Field(s ...any) string
	Fields(s ...any) []string
	Output(s ...any) string
}

type OptFunc func(Options)

// type IndicatorFunc IndicatorFn

// type IndicatorFn interface {
// 	Stream(stream *stream.Stream, opts ...OptFunc) *stream.Stream
// 	Series(series *series.Series, opts ...OptFunc) *series.Series
// }

type Indicator interface {
	Process(*tick.Tick) *tick.Tick
	Compute(*series.Series) *series.Series
}

type Exporter interface {
	Read(input ...*series.Series)
	Write() error
}

type Ticker interface{}

type Streamer interface{}

type IndicatorFunc func(*series.Series, ...OptFunc) *series.Series

func (i IndicatorFunc) Stream(input *stream.Stream, opts ...OptFunc) *stream.Stream {
	return input
}

type SeriesFunc func(*series.Series, ...OptFunc) *series.Series

type StreamFunc func(*stream.Stream, ...OptFunc) *stream.Stream

type Strategy interface {
}

type Provider interface {
}

type Broker interface {
}

type Storage interface {
}

type Trader interface {
	Backfill(start, end time.Time) // backfill historical prices from providers
	Train(start, end time.Time)    // train the strategy model
	Test(start, end time.Time)     // test the trained model
	Live(end time.Time)            // live trading with mock broker
	Exec(end time.Time)            // execute the trader
}
