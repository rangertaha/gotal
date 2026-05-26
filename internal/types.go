// Copyright (C) 2026 Rangertaha <rangertaha@gmail.com>
// SPDX-License-Identifier: GPL-3.0-or-later
package internal

import (
	"time"

	"gonum.org/v1/gonum/stat"
)

type Configurator interface {
	Set(key string, value any)
	Get(key string, defaults ...any) any

	GetStr(key string, defaults ...string) string
	GetInt(key string, defaults ...int) int
	GetFloat(key string, defaults ...float32) float32
	GetBool(key string, defaults ...bool) bool
	GetTime(key string, defaults ...time.Time) time.Time
	GetDuration(key string, defaults ...time.Duration) time.Duration
	GetSeries(key string, defaults ...TimeSeries) TimeSeries
}

type ConfigOption func(Configurator) error

type IndicatorFunc func(TimeSeries, ...ConfigOption) (TimeSeries, error)
type StreamIndicatorFunc func(Stream, ...ConfigOption) Stream
type BatchFunc func(...ConfigOption) (TimeSeries, error)
type StreamFunc func(...ConfigOption) (Stream, error)
type PluginFunc func(...ConfigOption) (TimeSeries, Stream, error)

type Initializer interface {
	Init(Configurator) error
}

type Indicator interface {
	Reset() error
	Ready() bool
	Process(Tick) Tick
	Compute(TimeSeries) TimeSeries
}

type Tick interface {
	Time() time.Time
	Duration() time.Duration
	Fields(names ...string) (map[string]float64, bool)
	Signals(names ...string) (map[string]float64, bool)
	Tags(names ...string) (map[string]string, bool)
}

type TimeSeries interface {
	Name() string
	Len() int
	IsEmpty() bool
	Ticks() []Tick
	At(i int) Tick
	First() Tick
	Last() Tick
	Add(ticks ...Tick) error
	Fields(names ...string) Fields
	Tags(names ...string) Tags

	Print()
	Save(path string) error
	Plot(fields ...string) Plot
}

type Stream interface {
	Ticks() <-chan Tick
	Errors() <-chan error
	Close()
	Ready() bool
	Print()
}

type Plot interface {
	Save(filename string, width, height int) error
	Show(width, height int) error
}

type Fields interface {
	Get(name string) Vector
	Set(name string, values []float64)
	Delete(name string)
	Names() []string
	Filter(names ...string) Fields
	Len() int
}

type Tags interface {
	Get(name string) string
	Set(name string, value string)
	Delete(name string)
	Names() []string
	Values() []string
	Filter(names ...string) Tags
	Len() int
}

type Vector interface {
	Values() []float64
	Range() float64
	Sum() float64
	Min() float64
	Max() float64
	First() float64
	Last() float64
	Mean() float64
	StdDev(weights ...float64) float64
	Variance(weights ...float64) float64
	Norm(norm float64) float64
	Quantile(p float64, c stat.CumulantKind, weights []float64) float64
}
