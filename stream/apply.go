// Copyright (C) 2026 Rangertaha <rangertaha@gmail.com>
// SPDX-License-Identifier: GPL-3.0-or-later
package stream

import (
	"github.com/rangertaha/gotal"
)

// Apply runs ind.Process on every tick flowing through in and emits the
// resulting tick on a downstream Stream. Errors are forwarded.
func Apply(in *Stream, ind gotal.Indicator) *Stream {
	outTicks := make(chan gotal.Tick)
	outErrs := make(chan error, 1)

	go func() {
		defer close(outTicks)
		for tick := range in.Ticks() {
			outTicks <- ind.Process(tick)
		}
	}()

	go func() {
		for err := range in.Errors() {
			outErrs <- err
		}
	}()

	return FromChan(in.Name(), outTicks, outErrs)
}
