package ma

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/config"
	"github.com/rangertaha/gotal/internal/indicators"
	"github.com/rangertaha/gotal/internal/tick"
)

const smaPluginID = "SMA"

type sma struct {
	Name   string
	Source string
	Period int

	window []float64
	sum    float64
}

func NewSMA(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := config.New(opts...)

	return &sma{
		Name:   c.GetStr("name", "sma"),
		Source: c.GetStr("source", "close"),
		Period: c.GetInt("period", 14),
	}, nil
}

func (i *sma) Reset() error {
	i.window = nil
	i.sum = 0
	return nil
}
func (i *sma) Ready() bool { return len(i.window) >= i.Period }

func (i *sma) Compute(input internal.TimeSeries) internal.TimeSeries {
	if input == nil {
		return input
	}
	ticks := input.Ticks()
	n := len(ticks)
	if n == 0 || i.Period <= 0 {
		return input
	}

	src := make([]float64, n)
	for j, t := range ticks {
		fields, ok := t.Fields(i.Source)
		if !ok {
			continue
		}
		src[j] = fields[i.Source]
	}

	out := make([]float64, n)
	var sum float64
	for j := 0; j < n; j++ {
		sum += src[j]
		if j >= i.Period {
			sum -= src[j-i.Period]
			out[j] = sum / float64(i.Period)
		} else if j == i.Period-1 {
			out[j] = sum / float64(i.Period)
		}
	}

	input.Fields().Set(i.Name, out)
	return input
}

func (i *sma) Process(input internal.Tick) internal.Tick {
	if input == nil || i.Period <= 0 {
		return input
	}
	fields, ok := input.Fields(i.Source)
	if !ok {
		return input
	}
	x := fields[i.Source]

	i.window = append(i.window, x)
	i.sum += x
	if len(i.window) > i.Period {
		i.sum -= i.window[0]
		i.window = i.window[1:]
	}

	if len(i.window) < i.Period {
		return input
	}

	val := i.sum / float64(i.Period)
	return tick.WithSignals(input, map[string]float64{i.Name: val})
}

func init() {
	if err := indicators.Add(smaPluginID, NewSMA, indicators.OVERLAP); err != nil {
		panic(err)
	}
}
