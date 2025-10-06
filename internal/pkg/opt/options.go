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

func (o *Option) Get(key string, defaults ...any) any {
	defaultValue := o.getDefault(defaults...)
	if v, ok := o.Params[key]; ok {
		return v
	}

	if defaultValue != nil {
		return defaultValue
	}
	return nil
}

func (o *Option) GetInt(key string, defaults ...any) int {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(int)
}

func (o *Option) GetString(key string, defaults ...any) string {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(string)
}

func (o *Option) GetStrings(key string, defaults ...any) []string {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.([]string)
}

func (o *Option) GetDuration(key string, defaults ...any) time.Duration {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(time.Duration)
}

func (o *Option) GetBool(key string, defaults ...any) bool {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(bool)
}

func (o *Option) GetFloat(key string, defaults ...any) float64 {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(float64)
}

func (o *Option) GetTime(key string, defaults ...any) time.Time {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(time.Time)
}

func (o *Option) Ticker(key string, defaults ...any) internal.Ticker {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(internal.Ticker)
}

func (o *Option) Name(s ...any) string {
	return o.GetString("name", s...)
}

func (o *Option) Suffix(s ...any) string {
	return o.GetString("suffix", s...)
}

func (o *Option) Period(n ...any) int {
	return o.GetInt("period", n...)
}

// FastPeriod for MACD
func (o *Option) FastPeriod(n ...any) int {
	return o.GetInt("fastPeriod", n...)
}

// SlowPeriod for MACD
func (o *Option) SlowPeriod(n ...any) int {
	return o.GetInt("slowPeriod", n...)
}

// SignalPeriod for MACD
func (o *Option) SignalPeriod(n ...any) int {
	return o.GetInt("signalPeriod", n...)
}

func (o *Option) Field(s ...any) string {
	return o.GetString("field", s...)
}

func (o *Option) Fields(s ...any) []string {
	return o.GetStrings("field", s...)
}

func (o *Option) Input(s ...any) string {
	return o.GetString("input", s...)
}

func (o *Option) Output(s ...any) string {
	return o.GetString("output", s...)
}

func (o *Option) Inputs(s ...any) []string {
	return o.GetStrings("inputs", s...)
}

func (o *Option) MAType(s ...any) string {
	return o.GetString("maType", s...)
}

func (o *Option) getDefault(defaults ...any) any {
	if len(defaults) > 0 {
		return defaults[0]
	}
	return nil
}

