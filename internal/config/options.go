package config

import (
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
	if len(defaults) > 0 {
		return defaults[0]
	}
	return nil
}

// Retrieval types

func (c *Config) GetStr(key string, defaults ...string) string {
	if v, ok := c.kv[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	return ""
}

func (c *Config) GetInt(key string, defaults ...int) int {
	if v, ok := c.kv[key]; ok {
		if i, ok := v.(int); ok {
			return i
		}
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	return 0
}

func (c *Config) GetFloat(key string, defaults ...float32) float32 {
	if v, ok := c.kv[key]; ok {
		switch n := v.(type) {
		case float32:
			return n
		case float64:
			return float32(n)
		}
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	return 0
}

func (c *Config) GetBool(key string, defaults ...bool) bool {
	if v, ok := c.kv[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	return false
}

func (c *Config) GetTime(key string, defaults ...time.Time) time.Time {
	if v, ok := c.kv[key]; ok {
		if t, ok := v.(time.Time); ok {
			return t
		}
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	return time.Time{}
}

func (c *Config) GetDuration(key string, defaults ...time.Duration) time.Duration {
	if v, ok := c.kv[key]; ok {
		if d, ok := v.(time.Duration); ok {
			return d
		}
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	return 0
}

func (c *Config) GetSeries(key string, defaults ...internal.TimeSeries) internal.TimeSeries {
	if v, ok := c.kv[key]; ok {
		if s, ok := v.(internal.TimeSeries); ok {
			return s
		}
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	return nil
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
