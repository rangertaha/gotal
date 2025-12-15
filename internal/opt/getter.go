package opt

import (
	"time"

	"github.com/rangertaha/gotal/internal"
)



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
