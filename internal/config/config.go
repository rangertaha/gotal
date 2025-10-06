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
	"os"
	"path/filepath"
)

const (
	CONFIG_DIR      = "hxe"
	CONFIG_FILE     = "config.hcl"
	PROGRAMS_DIR    = "programs"
	DEFAULT_SUBJECT = "hxe"
)

// var (
// 	//go:embed config.hcl
// 	DefaultConfig []byte

// 	//go:embed agent.db
// 	DefaultSqlite []byte
// )

// type (
// 	Config struct {
// 		ID      string `hcl:"id,optional"`
// 		Debug   bool   `hcl:"debug,optional"`
// 		Version string `hcl:"version,optional"`
// 		Banner  bool   `hcl:"banner,optional"`

// 		// Config
// 		configFile string
// 		Directory  string `hcl:"directory,optional"`

// 		Server   Server    `hcl:"server,block"`
// 		Services []Service `hcl:"service,block"`
// 	}
// 	Service struct {
// 		ID        string   `hcl:"id,label"`
// 		Directory string   `hcl:"directory,optional"`
// 		Config    hcl.Body `hcl:"config,remain"`
// 	}
// )

// // func init() {
// // 	// log.SetGlobalLevel(zerolog.TraceLevel)
// // 	// log.SetGlobalLevel(zerolog.ErrorLevel)
// // }

// // New creates a new configuration
// func New(options ...func(*Config) error) (*Config, error) {
// 	s := &Config{
// 		Banner:  true,
// 		Debug:   false,
// 		Version: internal.VERSION,
// 	}

// 	// Apply config options
// 	for _, opt := range options {
// 		err := opt(s)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	if s.Banner {
// 		internal.PrintBanner()
// 	}

// 	if s.Debug {
// 		log.SetGlobalLevel(zerolog.TraceLevel)
// 	}

// 	return s, nil
// }

// func CliOptions(ctx context.Context, cmd *cli.Command) func(c *Config) error {
// 	return func(c *Config) error {
// 		if cmd.String("config") != "" {
// 			// if config file is provided, use it
// 			c.configFile = cmd.String("config")
// 			if err := FileOption(c.configFile)(c); err != nil {
// 				return err
// 			}
// 		} else {
// 			// if config file is not provided, use default options
// 			if err := DefaultOptions()(c); err != nil {
// 				return err
// 			}
// 		}
// 		return nil
// 	}
// }

// func FileOption(path string) func(*Config) error {
// 	return func(c *Config) error {
// 		if path == "" {
// 			return fmt.Errorf("config file path is required")
// 		}

// 		// // Ask user if they want to create a config file
// 		// if c.ConfigFile == "" {
// 		// 	fmt.Println("No config file provided, would you like to create one?")
// 		// 	var answer string
// 		// 	fmt.Scanln(&answer)
// 		// 	if answer == "y" {
// 		// 		c.ConfigFile = path
// 		// 	}
// 		// }
// 		return nil
// 	}
// }

// func DefaultOptions() func(*Config) error {
// 	return func(c *Config) error {
// 		if c.Directory == "" {
// 			userConfigDir, err := os.UserConfigDir()
// 			if err != nil {
// 				return fmt.Errorf("error getting user config directory: %w", err)
// 			}
// 			c.Directory = filepath.Join(userConfigDir, CONFIG_DIR)
// 		}

// 		if err := c.Load(); err != nil {
// 			return fmt.Errorf("error loading config file: %w", err)
// 		}

// 		if err := c.LoadDatabase(); err != nil {
// 			return fmt.Errorf("error loading database file: %w", err)
// 		}

// 		return nil
// 	}
// }

// func (c *Config) Load() (err error) {
// 	// Create config file if it doesn't exist
// 	if c.configFile == "" {
// 		c.configFile = filepath.Join(c.Directory, CONFIG_FILE)
// 		if err := createFileIfNotExists(c.configFile, DefaultConfig); err != nil {
// 			return fmt.Errorf("error creating config file: %w", err)
// 		}
// 	}

// 	if err = hclsimple.DecodeFile(c.configFile, CtxFunctions, c); err != nil {
// 		return fmt.Errorf("error parsing config file: %w", err)
// 	}

// 	return
// }

// func (c *Config) LoadDatabase() (err error) {
// 	log.Info().Msg("setting up database")

// 	dbFile := filepath.Join(c.Directory, "agent.db")
// 	if err := createFileIfNotExists(dbFile, DefaultSqlite); err != nil {
// 		return fmt.Errorf("error creating database file: %w", err)
// 	}

// 	log.Info().Str("file", dbFile).Msg("using existing SQLite database")
// 	db.DB, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
// 	if err != nil {
// 		return err
// 	}
// 	db.SetDB(db.DB)

// 	return
// }

func createFileIfNotExists(filename string, contents []byte) error {
	// Check if file exists
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// Create directory if it doesn't exist
		dir := filepath.Dir(filename)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		// Write default config to file
		if err := os.WriteFile(filename, contents, 0644); err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}

		fmt.Printf("Created file: %s\n", filename)

	} else if err != nil {
		return fmt.Errorf("error checking file: %w", err)
	}

	return nil
}
