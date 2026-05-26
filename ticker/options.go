package ticker

import "github.com/rangertaha/gotal"

type Option func(*Ticker) error

func WithName(name string) Option {
	return func(s *Ticker) error { s.name = name; return nil }
}

func WithTags(t map[string]string) Option {
	return func(s *Ticker) error {
		for k, v := range t {
			s.tags[k] = v
		}
		return nil
	}
}

func WithTicks(ticks ...gotal.Tick) Option {
	return func(s *Ticker) error {
		return s.Add(ticks...)
	}
}
