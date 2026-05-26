// Copyright (C) 2026 Rangertaha <rangertaha@gmail.com>
// SPDX-License-Identifier: GPL-3.0-or-later
package stream

import (
	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/config"
	"github.com/rangertaha/gotal/internal/indicators"
	_ "github.com/rangertaha/gotal/internal/indicators/all"
)

// indicator resolves a registered indicator constructor by id and applies opts.
func indicator(id string, opts ...gotal.ConfigOption) gotal.Indicator {
	plugin := indicators.Get(id)
	ind, err := plugin(opts...)
	if err != nil {
		panic(err)
	}
	return ind
}

// EMA pipes in through the registered EMA indicator and emits per-tick signals.
func EMA(in *Stream, opts ...gotal.ConfigOption) *Stream {
	return Apply(in, indicator("ema", opts...))
}

// SMA pipes in through the registered SMA indicator.
func SMA(in *Stream, opts ...gotal.ConfigOption) *Stream {
	return Apply(in, indicator("sma", opts...))
}

// Ema is a typed convenience wrapper for EMA.
func Ema(in *Stream, period int, alpha float64) *Stream {
	return EMA(in,
		config.With("period", period),
		config.With("alpha", alpha),
	)
}

// Sma is a typed convenience wrapper for SMA.
func Sma(in *Stream, period int) *Stream {
	return SMA(in, config.With("period", period))
}
