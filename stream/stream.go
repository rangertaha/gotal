// Copyright (C) 2026 Rangertaha <rangertaha@gmail.com>
// SPDX-License-Identifier: GPL-3.0-or-later
package stream

import (
	"fmt"
	"sync"

	"github.com/rangertaha/gotal"
)

// Stream is a live channel of ticks. It satisfies gotal.Stream.
type Stream struct {
	name   string
	ticks  chan gotal.Tick
	errs   chan error
	closed bool
	mu     sync.Mutex
}

// New returns an inbound stream backed by the given channels. If errs is nil,
// an internal error channel is created. The stream takes ownership of the
// channels — call Close to release.
func New(name string, ticks <-chan gotal.Tick, errs ...<-chan error) *Stream {
	out := make(chan gotal.Tick)
	errOut := make(chan error, 1)

	s := &Stream{
		name:  name,
		ticks: out,
		errs:  errOut,
	}

	go func() {
		defer close(out)
		for t := range ticks {
			out <- t
		}
	}()

	if len(errs) > 0 && errs[0] != nil {
		go func() {
			for e := range errs[0] {
				errOut <- e
			}
		}()
	}

	return s
}

// FromChan wraps existing channels without spawning a goroutine.
// Useful for internal Apply chains.
func FromChan(name string, ticks chan gotal.Tick, errs chan error) *Stream {
	if errs == nil {
		errs = make(chan error, 1)
	}
	return &Stream{
		name:  name,
		ticks: ticks,
		errs:  errs,
	}
}

func (s *Stream) Name() string             { return s.name }
func (s *Stream) Ticks() <-chan gotal.Tick { return s.ticks }
func (s *Stream) Errors() <-chan error     { return s.errs }
func (s *Stream) Ready() bool              { return s.ticks != nil }

func (s *Stream) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return
	}
	s.closed = true
	close(s.errs)
}

// Print drains the stream to stdout. Returns once the upstream closes.
func (s *Stream) Print() {
	fmt.Printf("# stream %s\n", s.name)
	for tick := range s.ticks {
		fields, _ := tick.Fields()
		signals, _ := tick.Signals()
		fmt.Printf("%s fields=%v signals=%v\n",
			tick.Time().UTC().Format("2006-01-02T15:04:05Z"),
			fields, signals,
		)
	}
}

var _ gotal.Stream = (*Stream)(nil)
