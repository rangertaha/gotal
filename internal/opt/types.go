package opt

import (
	"time"

	"github.com/rangertaha/gotal/internal/series"
	"github.com/rangertaha/gotal/internal/stream"
	"github.com/rangertaha/gotal/internal/tick"
)

func (o *Option) Int(key string, defaults ...any) int {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(int)
}

func (o *Option) String(key string, defaults ...any) string {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(string)
}

func (o *Option) Strings(key string, defaults ...any) []string {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.([]string)
}

func (o *Option) Duration(key string, defaults ...any) time.Duration {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(time.Duration)
}

func (o *Option) Bool(key string, defaults ...any) bool {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(bool)
}

func (o *Option) Float(key string, defaults ...any) float64 {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(float64)
}

func (o *Option) Time(key string, defaults ...any) time.Time {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(time.Time)
}

func (o *Option) Tick(key string, defaults ...any) *tick.Tick {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(*tick.Tick)
}

func (o *Option) Series(key string, defaults ...any) *series.Series {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(*series.Series)
}

func (o *Option) Stream(key string, defaults ...any) *stream.Stream {
	defaultValue := o.getDefault(defaults...)

	v := o.Get(key, defaultValue)
	return v.(*stream.Stream)
}
