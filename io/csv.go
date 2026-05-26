// Copyright (C) 2026 Rangertaha <rangertaha@gmail.com>
// SPDX-License-Identifier: GPL-3.0-or-later
package io

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/internal/tick"
	"github.com/rangertaha/gotal/stream"
)

// CSVOption configures CSV parsing.
type CSVOption func(*csvSource)

type csvSource struct {
	path        string
	timeCol     string
	timeLayouts []string
	duration    time.Duration
}

// CSV returns a Source that reads ticks from a CSV file. The header row
// determines field names; the column named "time" (configurable) becomes
// the tick timestamp; all other numeric columns become fields.
func CSV(path string, opts ...CSVOption) Source {
	s := &csvSource{
		path:    path,
		timeCol: "time",
		timeLayouts: []string{
			time.RFC3339Nano,
			time.RFC3339,
			"2006-01-02 15:04:05",
			"2006-01-02T15:04:05",
			"2006-01-02",
		},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func CSVTimeColumn(name string) CSVOption {
	return func(s *csvSource) { s.timeCol = name }
}

func CSVTimeLayouts(layouts ...string) CSVOption {
	return func(s *csvSource) { s.timeLayouts = layouts }
}

func CSVDuration(d time.Duration) CSVOption {
	return func(s *csvSource) { s.duration = d }
}

func (s *csvSource) Ticks() ([]gotal.Tick, error) {
	f, err := os.Open(s.path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.TrimLeadingSpace = true

	header, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("io.CSV: reading header: %w", err)
	}

	timeIdx := -1
	for i, col := range header {
		if strings.EqualFold(col, s.timeCol) {
			timeIdx = i
			break
		}
	}
	if timeIdx < 0 {
		return nil, fmt.Errorf("io.CSV: time column %q not found in header %v", s.timeCol, header)
	}

	ticks := make([]gotal.Tick, 0, 128)
	row := 1
	for {
		record, err := r.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			if strings.Contains(err.Error(), "EOF") {
				break
			}
			return nil, fmt.Errorf("io.CSV: row %d: %w", row, err)
		}
		row++

		ts, err := s.parseTime(record[timeIdx])
		if err != nil {
			return nil, fmt.Errorf("io.CSV: row %d time %q: %w", row, record[timeIdx], err)
		}

		fields := map[string]float64{}
		for i, col := range header {
			if i == timeIdx {
				continue
			}
			if i >= len(record) {
				continue
			}
			v, err := strconv.ParseFloat(strings.TrimSpace(record[i]), 64)
			if err != nil {
				continue
			}
			fields[col] = v
		}

		ticks = append(ticks, tick.NewTick(ts, s.duration, fields, nil, nil))
	}

	return ticks, nil
}

func (s *csvSource) Stream() gotal.Stream {
	out := make(chan gotal.Tick)
	errs := make(chan error, 1)

	go func() {
		defer close(out)
		ticks, err := s.Ticks()
		if err != nil {
			errs <- err
			return
		}
		for _, t := range ticks {
			out <- t
		}
	}()

	return stream.FromChan(s.path, out, errs)
}

func (s *csvSource) parseTime(raw string) (time.Time, error) {
	raw = strings.TrimSpace(raw)
	if epoch, err := strconv.ParseInt(raw, 10, 64); err == nil {
		switch {
		case epoch > 1e18:
			return time.Unix(0, epoch), nil
		case epoch > 1e15:
			return time.Unix(0, epoch*int64(time.Microsecond)), nil
		case epoch > 1e12:
			return time.Unix(0, epoch*int64(time.Millisecond)), nil
		default:
			return time.Unix(epoch, 0), nil
		}
	}
	for _, layout := range s.timeLayouts {
		if ts, err := time.Parse(layout, raw); err == nil {
			return ts, nil
		}
	}
	return time.Time{}, fmt.Errorf("no matching time layout")
}
