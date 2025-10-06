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
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/nats-io/nats.go"

	_ "embed"

	"github.com/urfave/cli/v3"
)

const (
	DefaultPort        = 3143
	DefaultHost        = "0.0.0.0"
	DefaultBaseURL     = "nats://0.0.0.0:3143"
	DefaultUsername    = ""
	DefaultPassword    = ""
	DefaultTimeout     = 30 * time.Second
	CLIENT_CONFIG_FILE = "client.hcl"
)

//go:embed client.hcl
var DefaultClientConfig []byte

type (
	ClientConfig struct {
		Clients    []*Client `hcl:"client,block"`
		profile    string
		configFile string
		configDir  string
	}
	Client struct {
		Name     string        `hcl:"name,label"`
		UseIPC   bool          `hcl:"ipc,optional"`
		Host     string        `hcl:"host,optional"`
		Url      string        `hcl:"url,optional"`
		Port     int           `hcl:"port,optional"`
		Debug    bool          `hcl:"debug,optional"`
		Token    string        `hcl:"token,optional"`
		Password string        `hcl:"password,optional"`
		Username string        `hcl:"username,optional"`
		Timeout  time.Duration `hcl:"timeout,optional"`
	}
)

// New creates a new configuration
func NewClientConfig(options ...func(*ClientConfig) error) (client *Client, err error) {
	s := &ClientConfig{}

	// Apply config options
	for _, opt := range options {
		err := opt(s)
		if err != nil {
			return nil, err
		}
	}

	// Get the client configuration
	if client, err = s.Profile(); err != nil {
		return nil, fmt.Errorf("failed to load client: %w", err)
	}

	return client, nil
}

func ClientCliOpts(ctx context.Context, cmd *cli.Command) func(c *ClientConfig) error {
	return func(c *ClientConfig) error {
		if cmd.String("config") != "" {
			// if config file is provided, use it
			c.configFile = cmd.String("config")
			if err := ClientFileOption(c.configFile)(c); err != nil {
				return err
			}
		}

		if cmd.String("profile") != "" {
			c.profile = cmd.String("profile")
		}

		for i := range c.Clients {

			c.Clients[i].UseIPC = false
			if cmd.Bool("ipc") {
				c.Clients[i].UseIPC = true
			}

			c.Clients[i].Host = DefaultHost
			if cmd.String("host") != "" {
				c.Clients[i].Host = cmd.String("host")
			}

			c.Clients[i].Url = DefaultBaseURL
			if cmd.Int("port") != 0 {
				c.Clients[i].Port = cmd.Int("port")
			}

			c.Clients[i].Username = DefaultUsername
			if cmd.String("username") != "" {
				c.Clients[i].Username = cmd.String("username")
			}

			c.Clients[i].Password = DefaultPassword
			if cmd.String("password") != "" {
				c.Clients[i].Password = cmd.String("password")
			}

			c.Clients[i].Token = ""
			if cmd.String("token") != "" {
				c.Clients[i].Token = cmd.String("token")
			}

			if cmd.Bool("timeout") {
				c.Clients[i].Timeout = cmd.Duration("timeout")
			}

		}

		return nil
	}
}

func ClientFileOption(path string) func(*ClientConfig) (err error) {
	return func(c *ClientConfig) (err error) {
		if path == "" {
			return fmt.Errorf("config file path is required")
		}

		if err = hclsimple.DecodeFile(path, CtxFunctions, c); err != nil {
			return fmt.Errorf("error parsing config file: %w", err)
		}
		return nil
	}
}

func ClientDefaultOptions() func(*ClientConfig) (err error) {
	return func(c *ClientConfig) (err error) {
		if c.configDir == "" {
			userConfigDir, err := os.UserConfigDir()
			if err != nil {
				return fmt.Errorf("error getting user config directory: %w", err)
			}
			c.configDir = filepath.Join(userConfigDir, CONFIG_DIR)
			if c.configFile == "" {
				c.configFile = filepath.Join(c.configDir, CLIENT_CONFIG_FILE)
				if err := createFileIfNotExists(c.configFile, DefaultClientConfig); err != nil {
					return fmt.Errorf("error creating config file: %w", err)
				}
			}
		}

		if err = hclsimple.DecodeFile(c.configFile, CtxFunctions, c); err != nil {
			return fmt.Errorf("error parsing config file: %w", err)
		}

		return nil
	}
}

func ClientProfileOpts(profile string) func(*ClientConfig) (err error) {
	return func(c *ClientConfig) (err error) {
		if profile != "" {
			c.profile = profile
		}
		return nil
	}
}

func ClientConfigOpts(config *ClientConfig) func(*ClientConfig) (err error) {
	return func(c *ClientConfig) (err error) {
		if config != nil {
			c.Clients = config.Clients
		}
		return nil
	}
}

// func (c *ClientConfig) Load() (err error) {
// 	// Create config file if it doesn't exist
// 	if c.configFile == "" {
// 		c.configFile = filepath.Join(c.configDir, CLIENT_CONFIG_FILE)
// 		if err := createFileIfNotExists(c.configFile, DefaultClientConfig); err != nil {
// 			return fmt.Errorf("error creating config file: %w", err)
// 		}
// 	}

// 	if err = hclsimple.DecodeFile(c.configFile, CtxFunctions, c); err != nil {
// 		return fmt.Errorf("error parsing config file: %w", err)
// 	}

// 	return
// }

// Profile returns the client configuration for the specified profile name
func (c *ClientConfig) Profile(names ...string) (*Client, error) {
	if len(names) == 0 {
		names = append(names, c.profile)
	}

	// Return the first client that matches the profile name
	for _, name := range names {
		for _, client := range c.Clients {
			if client.Name == name {
				return client, nil
			}
		}
	}
	return nil, fmt.Errorf("client not found")
}

func (c *Client) Options() (opts nats.Option) {
	return func(o *nats.Options) error {

		nats.Name("hxe-client")(o)
		// nats.InProcessServer(ns)(o),
		nats.FlusherTimeout(c.Timeout)(o)

		o.Name = c.Name
		o.Url = c.Url
		o.User = c.Username
		o.Password = c.Password
		o.Token = c.Token
		o.Timeout = c.Timeout
		o.Servers = []string{c.Url}
		if c.UseIPC {
			o.Servers = []string{"nats://localhost:4222"}
		}
		return nil
	}
}
