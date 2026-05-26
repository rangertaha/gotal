// Copyright (C) 2026 Rangertaha <rangertaha@gmail.com>
// SPDX-License-Identifier: GPL-3.0-or-later
package ticker

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/plot"
)

// Ticker is a named, time-ordered batch of ticks. It implements gotal.TimeSeries.
type Ticker struct {
	name   string
	tags   tags
	fields fields
	ticks  []gotal.Tick
}

func New(name string, options ...Option) (*Ticker, error) {
	t := &Ticker{
		name:   name,
		tags:   make(tags),
		fields: make(fields),
		ticks:  make([]gotal.Tick, 0),
	}

	for _, option := range options {
		if err := option(t); err != nil {
			return nil, err
		}
	}

	return t, nil
}

func sortTicks(ticks []gotal.Tick) {
	sort.Slice(ticks, func(i, j int) bool {
		return ticks[i].Time().Before(ticks[j].Time())
	})
}

func (s *Ticker) Name() string { return s.name }

func (s *Ticker) Tags(names ...string) gotal.Tags {
	if len(names) == 0 {
		return s.tags
	}
	return s.tags.Filter(names...)
}

func (s *Ticker) Fields(names ...string) gotal.Fields {
	if len(names) == 0 {
		return s.fields
	}
	return s.fields.Filter(names...)
}

func (s *Ticker) Ticks() []gotal.Tick { return s.ticks }

func (s *Ticker) IsEmpty() bool { return len(s.ticks) == 0 }
func (s *Ticker) Len() int      { return len(s.ticks) }

func (s *Ticker) With(options ...Option) error {
	for _, option := range options {
		if err := option(s); err != nil {
			return err
		}
	}
	return nil
}

func (s *Ticker) Add(ticks ...gotal.Tick) error {
	s.ticks = append(s.ticks, ticks...)
	sortTicks(s.ticks)
	return nil
}

func (s *Ticker) First() gotal.Tick {
	if len(s.ticks) == 0 {
		return nil
	}
	return s.ticks[0]
}

func (s *Ticker) Last() gotal.Tick {
	if len(s.ticks) == 0 {
		return nil
	}
	return s.ticks[len(s.ticks)-1]
}

func (s *Ticker) At(index int) gotal.Tick {
	if index < 0 || index >= len(s.ticks) {
		return nil
	}
	return s.ticks[index]
}

func (s *Ticker) Range(start, end int) (*Ticker, error) {
	if start < 0 || end > len(s.ticks) || start > end {
		return nil, errors.New("index out of range")
	}
	return New(s.name, WithTicks(s.ticks[start:end]...))
}

// Print writes a human-readable table of ticks and field signals to stdout.
func (s *Ticker) Print() {
	tickFieldNames := s.tickFieldNames()
	computedFields := s.fields.Names()
	sort.Strings(computedFields)

	fmt.Printf("# %s (%d ticks)\n", s.name, len(s.ticks))
	header := "time"
	for _, name := range tickFieldNames {
		header += "," + name
	}
	for _, name := range computedFields {
		header += "," + name
	}
	fmt.Println(header)

	for i, tick := range s.ticks {
		row := tick.Time().UTC().Format("2006-01-02T15:04:05Z")
		tickFields, _ := tick.Fields()
		for _, name := range tickFieldNames {
			if v, ok := tickFields[name]; ok {
				row += "," + strconv.FormatFloat(v, 'f', -1, 64)
			} else {
				row += ","
			}
		}
		for _, name := range computedFields {
			v := s.fields.Get(name)
			if v == nil {
				row += ","
				continue
			}
			values := v.Values()
			if i < len(values) {
				row += "," + strconv.FormatFloat(values[i], 'f', -1, 64)
			} else {
				row += ","
			}
		}
		fmt.Println(row)
	}
}

func (s *Ticker) tickFieldNames() []string {
	seen := map[string]struct{}{}
	for _, tick := range s.ticks {
		fs, ok := tick.Fields()
		if !ok {
			continue
		}
		for name := range fs {
			seen[name] = struct{}{}
		}
	}
	names := make([]string, 0, len(seen))
	for name := range seen {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// Save writes the ticker as a CSV file. Header is time + tick-field names + computed field names.
func (s *Ticker) Save(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	tickFieldNames := s.tickFieldNames()

	computedFields := s.fields.Names()
	sort.Strings(computedFields)

	header := append([]string{"time"}, tickFieldNames...)
	header = append(header, computedFields...)
	if err := w.Write(header); err != nil {
		return err
	}

	for i, tick := range s.ticks {
		row := []string{tick.Time().UTC().Format("2006-01-02T15:04:05Z")}
		tickFields, _ := tick.Fields()
		for _, name := range tickFieldNames {
			if v, ok := tickFields[name]; ok {
				row = append(row, strconv.FormatFloat(v, 'f', -1, 64))
			} else {
				row = append(row, "")
			}
		}
		for _, name := range computedFields {
			v := s.fields.Get(name)
			if v == nil {
				row = append(row, "")
				continue
			}
			values := v.Values()
			if i < len(values) {
				row = append(row, strconv.FormatFloat(values[i], 'f', -1, 64))
			} else {
				row = append(row, "")
			}
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}

// Plot returns a gotal.Plot for the given fields (defaults to all fields if none specified).
func (s *Ticker) Plot(fields ...string) gotal.Plot {
	return plot.New(s, fields...)
}

var _ gotal.TimeSeries = (*Ticker)(nil)
