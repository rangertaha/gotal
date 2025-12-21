package config

import (
	_ "embed"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/rangertaha/gotal/internal"
)

var (
	//go:embed default.hcl
	DefaultConfig []byte
)

type Config struct {
	Debug      bool      `hcl:"debug,optional"`
	Version    string    `hcl:"version,optional"`
	Providers  []*Plugin `hcl:"provider,block"`
	Indicators []*Plugin `hcl:"indicator,block"`
	Strategies []*Plugin `hcl:"strategy,block"`
	Brokers    []*Plugin `hcl:"broker,block"`
}

type Plugin struct {
	Type   string   `hcl:"type,label"`
	Name   string   `hcl:"name,label"`
	Config hcl.Body `hcl:",remain"`
}

func New(options ...func(*Config) error) (*Config, error) {
	cfg := &Config{
		Debug:      false,
		Version:    internal.VERSION,
		Providers:  []*Plugin{},
		Indicators: []*Plugin{},
		Strategies: []*Plugin{},
		Brokers:    []*Plugin{},
	}
	for _, option := range options {
		if err := option(cfg); err != nil {
			return nil, err
		}
	}
	return cfg, nil
}

func Load(path string) (cfg *Config, err error) {
	if cfg, err = New(); err != nil {
		return nil, err
	}
	if err = hclsimple.DecodeFile(path, CtxFunctions, cfg); err != nil {
		return nil, err
	}
	return
}

// func (c *Config) GetPlugin(name string) (*Plugin, error) {
// 	for _, plugin := range c.Providers {
// 		if plugin. == name {
// 			return plugin, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("plugin %s not found", name)
// }
