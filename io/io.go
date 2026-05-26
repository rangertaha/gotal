// Copyright (C) 2026 Rangertaha <rangertaha@gmail.com>
// SPDX-License-Identifier: GPL-3.0-or-later
package io

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/ticker"
)

// Source produces ticks either as a slice (batch) or as a Stream (live).
type Source interface {
	Ticks() ([]gotal.Tick, error)
	Stream() gotal.Stream
}

// Read dispatches by file extension and returns a batch *ticker.Ticker.
func Read(path string) (*ticker.Ticker, error) {
	var src Source
	switch ext := strings.ToLower(filepath.Ext(path)); ext {
	case ".csv":
		src = CSV(path)
	case ".json":
		src = JSON(path)
	case ".jsonl", ".ndjson":
		src = JSONL(path)
	default:
		return nil, fmt.Errorf("io: unsupported extension %q", ext)
	}
	ticks, err := src.Ticks()
	if err != nil {
		return nil, err
	}
	name := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	return ticker.New(name, ticker.WithTicks(ticks...))
}
