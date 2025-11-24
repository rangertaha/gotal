package opt

import (
	"time"

	"github.com/rangertaha/gotal/internal"
)

var PERIOD int = 10
var FIELDS []string = []string{"value"}

type Option struct {
	Params map[string]any
}

func New(opts ...internal.OptFunc) internal.Options {
	cfg := &Option{
		Params: make(map[string]any),
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func SetPeriod(period int) {
	PERIOD = period
}

func WithName(n string) internal.OptFunc {
	return func(c internal.Options) { c.Set("name", n) }
}

func WithPeriod(p int) internal.OptFunc {
	return func(c internal.Options) { c.Set("period", p) }
}

func WithDuration(d time.Duration) internal.OptFunc {
	return func(c internal.Options) { c.Set("duration", d) }
}

func WithInput(i string) internal.OptFunc {
	return func(c internal.Options) { c.Set("input", i) }
}

func WithOutput(o string) internal.OptFunc {
	return func(c internal.Options) { c.Set("output", o) }
}

func WithInputs(i ...string) internal.OptFunc {
	return func(c internal.Options) { c.Set("inputs", i) }
}

func WithField(f string) internal.OptFunc {
	return func(c internal.Options) { c.Set("field", f) }
}

func WithFields(f ...string) internal.OptFunc {
	return func(c internal.Options) { c.Set("fields", f) }
}

func WithFieldMap(fmap map[string]string) internal.OptFunc {
	return func(c internal.Options) {
		for k, v := range fmap {
			c.Set(k, v)
		}
	}
}

// WithFastPeriod for MACD
func WithFastPeriod(p int) internal.OptFunc {
	return func(c internal.Options) { c.Set("fastPeriod", p) }
}

// WithSlowPeriod for MACD
func WithSlowPeriod(p int) internal.OptFunc {
	return func(c internal.Options) { c.Set("slowPeriod", p) }
}

// WithSignalPeriod for MACD
func WithSignalPeriod(p int) internal.OptFunc {
	return func(c internal.Options) { c.Set("signalPeriod", p) }
}

// WithTimePeriod for MACD
func WithTimePeriod(p time.Duration) internal.OptFunc {
	return func(c internal.Options) { c.Set("timePeriod", p) }
}

func WithMAType(p string) internal.OptFunc {
	return func(c internal.Options) { c.Set("maType", p) }
}

func With(name string, value any) internal.OptFunc {
	return func(c internal.Options) {
		c.Set(name, value)
	}
}


// WithIndicator for trading
func WithIndicator(i string, opts ...internal.OptFunc) internal.OptFunc {
	return func(c internal.Options) { c.Set("indicator", i) }
}

func WithStrategy(s string, opts ...internal.OptFunc) internal.OptFunc {
	return func(c internal.Options) { c.Set("strategy", s) }
}

func WithBroker(b string) internal.OptFunc {
	return func(c internal.Options) { c.Set("broker", b) }
}

func WithProvider(p string, opts ...internal.OptFunc) internal.OptFunc {
	return func(c internal.Options) { c.Set("provider", p) }
}



func (o *Option) Set(key string, value any) {
	o.Params[key] = value
}
