package config

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal/internal"
)

type Config struct {
	kv map[string]any
}

func New(opts ...internal.ConfigOption) internal.Configurator {
	o := &Config{
		kv: make(map[string]any),
	}

	for _, opt := range opts {
		if err := opt(o); err != nil {
			return nil
		}
	}

	return o
}

func (c *Config) Set(key string, value any) {
	c.kv[key] = value
}

func (c *Config) Get(key string, defaults ...any) any {
	if value, ok := c.kv[key]; ok {
		return value
	}
	if len(defaults) == 0 {
		panic(fmt.Sprintf("The option '%s' is required", key))
	}
	return defaults[0]
}

// Retrieval types

func (c *Config) GetStr(key string, defaults ...string) string {
	return c.Get(key).(string)
}

func (c *Config) GetInt(key string, defaults ...int) int {
	return c.Get(key).(int)
}

func (c *Config) GetFloat(key string, defaults ...float32) float32 {
	return c.Get(key).(float32)
}

func (c *Config) GetBool(key string, defaults ...bool) bool {
	return c.Get(key).(bool)
}

func (c *Config) GetTime(key string, defaults ...time.Time) time.Time {
	return c.Get(key).(time.Time)
}

func (c *Config) GetDuration(key string, defaults ...time.Duration) time.Duration {
	return c.Get(key).(time.Duration)
}

func (c *Config) GetSeries(key string, defaults ...internal.Series) internal.Series {
	return c.Get(key).(internal.Series)
}

// Options functions used as function arguments to the New function
func WithParams(params ...any) internal.ConfigOption {
	return func(conf internal.Configurator) error {
		return nil
	}
}

func With(key string, value any) internal.ConfigOption {
	return func(conf internal.Configurator) error {
		conf.Set(key, value)
		return nil
	}
}
