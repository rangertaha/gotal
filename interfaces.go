package gotal

import "time"

// import "time"

// type Serieser interface {
// 	Name() string
// 	Metrics() []Metricer
// 	Push(metric Metricer)
// 	Pop() Metricer
// 	Len() int
// 	Cap() int
// 	IsEmpty() bool
// 	IsFull() bool
// }

// type Metricer interface {
// 	Measurement() string
// 	Timestamp() time.Time
// 	Tags() map[string]string
// 	Fields() map[string]float64
// }

// // Chovler represents an assent candlestick with Close, High, Open, Volume, Low
// type Candler interface {
// 	Timestamp() time.Time
// 	Open() float64
// 	High() float64
// 	Low() float64
// 	Close() float64
// 	Volume() float64
// }

// type Patterner interface {
// 	Timestamp() time.Time
// 	Name() string
// 	Bullish() float64
// 	Bearish() float64
// 	Reversal() float64
// }

// type Indicater interface {
// 	Timestamp() time.Time
// 	Value() float64
// }

type IndicatorFunc func([]*Metric, ...Option) []*Metric

type IndicatorStreamFunc func(chan *Metric, ...Option) chan *Metric

type Indicator interface {
	Add(metric *Metric)
	Calculate() []*Metric
	Stream(metrics chan *Metric) chan *Metric
}

type Metric struct {
	Timestamp time.Time
	Value     float64
}

type Option func(*Metric)
	