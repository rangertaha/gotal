// Package mathop implements TA-Lib's Math Operator functions.
package mathop

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/indicators"
	"github.com/rangertaha/gotal/internal/indicators/util"
)

type windowed struct {
	Name   string
	Source string
	Period int
	fn     func([]float64, int) []float64
}

func (i *windowed) Reset() error                          { return nil }
func (i *windowed) Ready() bool                           { return true }
func (i *windowed) Process(t internal.Tick) internal.Tick { return t }
func (i *windowed) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		return i.fn(src, i.Period)
	})
}

func ctor(name string, defaultPeriod int, fn func([]float64, int) []float64) func(opts ...internal.ConfigOption) (internal.Indicator, error) {
	return func(opts ...internal.ConfigOption) (internal.Indicator, error) {
		c := util.Cfg(opts...)
		return &windowed{
			Name:   c.GetStr("name", name),
			Source: c.GetStr("source", "close"),
			Period: c.GetInt("period", defaultPeriod),
			fn:     fn,
		}, nil
	}
}

// ---------- Binary operators (two source fields) ----------

type binary struct {
	Name, Source1, Source2 string
	op                     func(a, b float64) float64
}

func (i *binary) Reset() error                          { return nil }
func (i *binary) Ready() bool                           { return true }
func (i *binary) Process(t internal.Tick) internal.Tick { return t }
func (i *binary) Compute(ts internal.TimeSeries) internal.TimeSeries {
	a := util.FieldOf(ts, i.Source1)
	b := util.FieldOf(ts, i.Source2)
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	out := make([]float64, n)
	for j := 0; j < n; j++ {
		out[j] = i.op(a[j], b[j])
	}
	return util.AttachField(ts, i.Name, out)
}

func binaryCtor(name string, op func(a, b float64) float64) func(opts ...internal.ConfigOption) (internal.Indicator, error) {
	return func(opts ...internal.ConfigOption) (internal.Indicator, error) {
		c := util.Cfg(opts...)
		return &binary{
			Name:    c.GetStr("name", name),
			Source1: c.GetStr("source1", "high"),
			Source2: c.GetStr("source2", "low"),
			op:      op,
		}, nil
	}
}

// ---------- Index variants ----------

type windowedIdx struct {
	Name   string
	Source string
	Period int
	fn     func([]float64, int) []float64
}

func (i *windowedIdx) Reset() error                          { return nil }
func (i *windowedIdx) Ready() bool                           { return true }
func (i *windowedIdx) Process(t internal.Tick) internal.Tick { return t }
func (i *windowedIdx) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		return i.fn(src, i.Period)
	})
}

func rollingMaxIndex(src []float64, period int) []float64 {
	n := len(src)
	out := make([]float64, n)
	if period <= 0 || n < period {
		return out
	}
	for i := period - 1; i < n; i++ {
		idx := i - period + 1
		for j := i - period + 2; j <= i; j++ {
			if src[j] > src[idx] {
				idx = j
			}
		}
		out[i] = float64(idx)
	}
	return out
}
func rollingMinIndex(src []float64, period int) []float64 {
	n := len(src)
	out := make([]float64, n)
	if period <= 0 || n < period {
		return out
	}
	for i := period - 1; i < n; i++ {
		idx := i - period + 1
		for j := i - period + 2; j <= i; j++ {
			if src[j] < src[idx] {
				idx = j
			}
		}
		out[i] = float64(idx)
	}
	return out
}

// All runs every implemented Math Operator indicator against ts.
func All(ts internal.TimeSeries, opts ...internal.ConfigOption) internal.TimeSeries {
	return util.RunGroup(ts, indicators.MATHOP, opts...)
}

func init() {
	util.Must(indicators.Add("MAX", ctor("max", 30, util.RollingMax), indicators.MATHOP))
	util.Must(indicators.Add("MIN", ctor("min", 30, util.RollingMin), indicators.MATHOP))
	util.Must(indicators.Add("SUMWINDOW", ctor("sumwindow", 30, util.RollingSum), indicators.MATHOP))

	util.Must(indicators.Add("ADD", binaryCtor("add", func(a, b float64) float64 { return a + b }), indicators.MATHOP))
	util.Must(indicators.Add("SUB", binaryCtor("sub", func(a, b float64) float64 { return a - b }), indicators.MATHOP))
	util.Must(indicators.Add("MULT", binaryCtor("mult", func(a, b float64) float64 { return a * b }), indicators.MATHOP))
	util.Must(indicators.Add("DIV", binaryCtor("div", func(a, b float64) float64 {
		if b == 0 {
			return 0
		}
		return a / b
	}), indicators.MATHOP))

	idx := func(name string, fn func([]float64, int) []float64) func(opts ...internal.ConfigOption) (internal.Indicator, error) {
		return func(opts ...internal.ConfigOption) (internal.Indicator, error) {
			c := util.Cfg(opts...)
			return &windowedIdx{
				Name:   c.GetStr("name", name),
				Source: c.GetStr("source", "close"),
				Period: c.GetInt("period", 30),
				fn:     fn,
			}, nil
		}
	}
	util.Must(indicators.Add("MAXINDEX", idx("maxindex", rollingMaxIndex), indicators.MATHOP))
	util.Must(indicators.Add("MININDEX", idx("minindex", rollingMinIndex), indicators.MATHOP))
}
