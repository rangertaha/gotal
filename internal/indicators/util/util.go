// Package util contains shared helpers used across indicator implementations.
package util

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/config"
	"github.com/rangertaha/gotal/internal/indicators"
)

// FieldOf extracts the named field's values across all ticks in a TimeSeries.
func FieldOf(ts internal.TimeSeries, source string) []float64 {
	ticks := ts.Ticks()
	out := make([]float64, len(ticks))
	for i, t := range ticks {
		if fields, ok := t.Fields(source); ok {
			out[i] = fields[source]
		}
	}
	return out
}

// FieldsOf extracts multiple named fields across all ticks. Returned slices
// align by index.
func FieldsOf(ts internal.TimeSeries, sources ...string) [][]float64 {
	out := make([][]float64, len(sources))
	for i, s := range sources {
		out[i] = FieldOf(ts, s)
	}
	return out
}

// AttachField stores a computed series under name in the TimeSeries' fields
// map and returns the (mutated) TimeSeries.
func AttachField(ts internal.TimeSeries, name string, values []float64) internal.TimeSeries {
	ts.Fields().Set(name, values)
	return ts
}

// Unary runs fn over the named source field and attaches the result.
func Unary(ts internal.TimeSeries, source, name string, fn func([]float64) []float64) internal.TimeSeries {
	if ts == nil || ts.Len() == 0 {
		return ts
	}
	return AttachField(ts, name, fn(FieldOf(ts, source)))
}

// Cfg is shorthand for config.New(opts...).
func Cfg(opts ...internal.ConfigOption) internal.Configurator {
	return config.New(opts...)
}

// SMA computes a simple moving average over src with the given period.
// The first period-1 outputs are zero.
func SMA(src []float64, period int) []float64 {
	n := len(src)
	out := make([]float64, n)
	if period <= 0 || n == 0 {
		return out
	}
	var sum float64
	for i := 0; i < n; i++ {
		sum += src[i]
		if i >= period {
			sum -= src[i-period]
		}
		if i >= period-1 {
			out[i] = sum / float64(period)
		}
	}
	return out
}

// EMA computes an exponential moving average over src. alpha <= 0 uses 2/(period+1).
func EMA(src []float64, period int, alpha float64) []float64 {
	n := len(src)
	out := make([]float64, n)
	if n == 0 {
		return out
	}
	if alpha <= 0 {
		alpha = 2.0 / float64(period+1)
	}
	out[0] = src[0]
	for i := 1; i < n; i++ {
		out[i] = alpha*src[i] + (1-alpha)*out[i-1]
	}
	return out
}

// RollingMax returns the max of src in a trailing window of size period.
func RollingMax(src []float64, period int) []float64 {
	n := len(src)
	out := make([]float64, n)
	if period <= 0 || n < period {
		return out
	}
	for i := period - 1; i < n; i++ {
		m := src[i-period+1]
		for j := i - period + 2; j <= i; j++ {
			if src[j] > m {
				m = src[j]
			}
		}
		out[i] = m
	}
	return out
}

// RollingMin returns the min of src in a trailing window of size period.
func RollingMin(src []float64, period int) []float64 {
	n := len(src)
	out := make([]float64, n)
	if period <= 0 || n < period {
		return out
	}
	for i := period - 1; i < n; i++ {
		m := src[i-period+1]
		for j := i - period + 2; j <= i; j++ {
			if src[j] < m {
				m = src[j]
			}
		}
		out[i] = m
	}
	return out
}

// RollingSum returns the sum of src in a trailing window of size period.
func RollingSum(src []float64, period int) []float64 {
	n := len(src)
	out := make([]float64, n)
	if period <= 0 {
		return out
	}
	var sum float64
	for i := 0; i < n; i++ {
		sum += src[i]
		if i >= period {
			sum -= src[i-period]
		}
		if i >= period-1 {
			out[i] = sum
		}
	}
	return out
}

// Must panics on err. Used in init() registrations.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Signal constants. We use TA-Lib's CDL convention so that signal indicators
// and pattern indicators share an output space.
const (
	SignalBuy  = 100
	SignalSell = -100
	SignalHold = 0
)

// ThresholdCross emits SignalBuy on the bar values cross *below* lower (after
// being at-or-above) and SignalSell on the bar values cross *above* upper. All
// other bars are SignalHold. Use for oversold/overbought style oscillators.
func ThresholdCross(values []float64, upper, lower float64) []float64 {
	n := len(values)
	out := make([]float64, n)
	for i := 1; i < n; i++ {
		if values[i-1] >= lower && values[i] < lower {
			out[i] = SignalBuy
		} else if values[i-1] <= upper && values[i] > upper {
			out[i] = SignalSell
		}
	}
	return out
}

// ZeroCross emits SignalBuy when values cross from <=0 to >0 and SignalSell
// when crossing from >=0 to <0. Hold otherwise.
func ZeroCross(values []float64) []float64 {
	n := len(values)
	out := make([]float64, n)
	for i := 1; i < n; i++ {
		switch {
		case values[i-1] <= 0 && values[i] > 0:
			out[i] = SignalBuy
		case values[i-1] >= 0 && values[i] < 0:
			out[i] = SignalSell
		}
	}
	return out
}

// LineCrossover emits SignalBuy when a crosses above b and SignalSell when
// a crosses below b. Hold otherwise. Both inputs must be the same length.
func LineCrossover(a, b []float64) []float64 {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	out := make([]float64, n)
	for i := 1; i < n; i++ {
		switch {
		case a[i-1] <= b[i-1] && a[i] > b[i]:
			out[i] = SignalBuy
		case a[i-1] >= b[i-1] && a[i] < b[i]:
			out[i] = SignalSell
		}
	}
	return out
}

// RunIDs constructs and runs every indicator id in the list against ts. Stubs
// and indicators whose constructor errors out are skipped silently — this is
// best-effort "run everything that can run" semantics.
func RunIDs(ts internal.TimeSeries, ids []string, opts ...internal.ConfigOption) internal.TimeSeries {
	if ts == nil {
		return ts
	}
	for _, id := range ids {
		if indicators.IsStub(id) {
			continue
		}
		ind, err := indicators.Get(id)(opts...)
		if err != nil {
			continue
		}
		ts = ind.Compute(ts)
	}
	return ts
}

// RunGroup runs every implemented indicator registered under the given group.
func RunGroup(ts internal.TimeSeries, group indicators.GroupType, opts ...internal.ConfigOption) internal.TimeSeries {
	return RunIDs(ts, indicators.Group(group), opts...)
}
