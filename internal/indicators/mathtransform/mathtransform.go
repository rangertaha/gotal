// Package mathtransform implements TA-Lib's Math Transform functions (per-element math).
package mathtransform

import (
	"math"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/indicators"
	"github.com/rangertaha/gotal/internal/indicators/util"
)

type unary struct {
	Name   string
	Source string
	fn     func(float64) float64
}

func (i *unary) Reset() error                          { return nil }
func (i *unary) Ready() bool                           { return true }
func (i *unary) Process(t internal.Tick) internal.Tick { return t }
func (i *unary) Compute(ts internal.TimeSeries) internal.TimeSeries {
	return util.Unary(ts, i.Source, i.Name, func(src []float64) []float64 {
		out := make([]float64, len(src))
		for j, v := range src {
			out[j] = i.fn(v)
		}
		return out
	})
}

func ctor(name string, fn func(float64) float64) func(opts ...internal.ConfigOption) (internal.Indicator, error) {
	return func(opts ...internal.ConfigOption) (internal.Indicator, error) {
		c := util.Cfg(opts...)
		return &unary{
			Name:   c.GetStr("name", name),
			Source: c.GetStr("source", "close"),
			fn:     fn,
		}, nil
	}
}

// All runs every implemented Math Transform indicator against ts.
func All(ts internal.TimeSeries, opts ...internal.ConfigOption) internal.TimeSeries {
	return util.RunGroup(ts, indicators.MATHTRANSFORM, opts...)
}

func init() {
	util.Must(indicators.Add("ACOS", ctor("acos", math.Acos), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("ASIN", ctor("asin", math.Asin), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("ATAN", ctor("atan", math.Atan), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("COS", ctor("cos", math.Cos), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("SIN", ctor("sin", math.Sin), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("TAN", ctor("tan", math.Tan), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("COSH", ctor("cosh", math.Cosh), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("SINH", ctor("sinh", math.Sinh), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("TANH", ctor("tanh", math.Tanh), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("CEIL", ctor("ceil", math.Ceil), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("FLOOR", ctor("floor", math.Floor), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("EXP", ctor("exp", math.Exp), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("LN", ctor("ln", math.Log), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("LOG10", ctor("log10", math.Log10), indicators.MATHTRANSFORM))
	util.Must(indicators.Add("SQRT", ctor("sqrt", math.Sqrt), indicators.MATHTRANSFORM))
}
