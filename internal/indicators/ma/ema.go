package ma

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/config"
	"github.com/rangertaha/gotal/internal/indicators"
	"github.com/rangertaha/gotal/internal/tick"
)

const emaPluginID = "EMA"

type ema struct {
	Name   string
	Source string
	Period int
	Alpha  float64

	prev    float64
	hasPrev bool
}

func NewEMA(opts ...internal.ConfigOption) (internal.Indicator, error) {
	c := config.New(opts...)

	period := c.GetInt("period", 14)
	alpha := float64(c.GetFloat("alpha", 0))
	if alpha <= 0 {
		alpha = 2.0 / float64(period+1)
	}

	return &ema{
		Name:   c.GetStr("name", "ema"),
		Source: c.GetStr("source", "close"),
		Period: period,
		Alpha:  alpha,
	}, nil
}

func (i *ema) Reset() error {
	i.prev = 0
	i.hasPrev = false
	return nil
}
func (i *ema) Ready() bool { return i.hasPrev }

func (i *ema) Compute(input internal.TimeSeries) internal.TimeSeries {
	if input == nil {
		return input
	}
	ticks := input.Ticks()
	n := len(ticks)
	if n == 0 {
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
	out[0] = src[0]
	for j := 1; j < n; j++ {
		out[j] = i.Alpha*src[j] + (1-i.Alpha)*out[j-1]
	}

	input.Fields().Set(i.Name, out)
	return input
}

func (i *ema) Process(input internal.Tick) internal.Tick {
	if input == nil {
		return input
	}
	fields, ok := input.Fields(i.Source)
	if !ok {
		return input
	}
	x := fields[i.Source]

	var val float64
	if !i.hasPrev {
		val = x
		i.hasPrev = true
	} else {
		val = i.Alpha*x + (1-i.Alpha)*i.prev
	}
	i.prev = val

	return tick.WithSignals(input, map[string]float64{i.Name: val})
}

func init() {
	if err := indicators.Add(emaPluginID, NewEMA, indicators.OVERLAP); err != nil {
		panic(err)
	}
}
