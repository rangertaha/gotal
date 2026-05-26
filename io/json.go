// Copyright (C) 2026 Rangertaha <rangertaha@gmail.com>
// SPDX-License-Identifier: GPL-3.0-or-later
package io

import (
	"errors"

	"github.com/rangertaha/gotal"
)

type jsonSource struct {
	path string
}

func JSON(path string) Source  { return &jsonSource{path: path} }
func JSONL(path string) Source { return &jsonSource{path: path} }

func (s *jsonSource) Ticks() ([]gotal.Tick, error) {
	return nil, errors.New("io.JSON: not implemented")
}

func (s *jsonSource) Stream() gotal.Stream {
	return nil
}
