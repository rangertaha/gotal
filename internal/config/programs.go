/*
 * HXE - Host-based Process Execution Engine
 * Copyright (C) 2025 Rangertaha <rangertaha@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package config

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/rangertaha/hxe/internal/services/program/models"

	_ "embed"
)

const (
	PROGRAM_CONFIG_DIR = "programs"
)

var (
	//go:embed program.hcl
	DefaultProgramConfig []byte
)

type (
	ProgramConfig struct {
		Programs []*models.Program `hcl:"program,block"`
	}
	// Service struct {
	// 	ID        string `hcl:"id,label"`
	// 	Directory string `hcl:"directory,optional"`
	// 	Conn      *nats.Conn
	// 	Config    hcl.Body `hcl:"config,remain"`
	// }
)

// New creates a new configuration
func LoadProgramConfig(filename string) ([]*models.Program, error) {
	p := &ProgramConfig{}
	// Load config
	if err := hclsimple.DecodeFile(filename, CtxFunctions, p); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	if p.Programs == nil {
		return nil, fmt.Errorf("no programs found in config file")
	}

	return p.Programs, nil
}
