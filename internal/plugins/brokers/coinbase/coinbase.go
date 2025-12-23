package coinbase

import (
	_ "embed"
	"errors"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/opt"
	"github.com/rangertaha/gotal/internal/plugins/brokers"
	"github.com/rangertaha/gotal/internal/series"
)

const (
	CoinbaseBaseURL = "https://api.coinbase.com"
	CoinbaseVersion = "v2"
)

type coinbase struct {

	// connection parameters
	APIKey  string `hcl:"api_key"`  // API key
	BaseURL string `hcl:"base_url"` // Base URL
	Version string `hcl:"version"`  // Version

	// account parameters
	Accounts []Account `hcl:"accounts,block"` // Accounts

	// series internal state
	series internal.Series
}

func New(opts ...internal.ConfigOption) (internal.Plugin, error) {

	p := &coinbase{
		BaseURL: CoinbaseBaseURL,
		Version: CoinbaseVersion,
		series:  series.New(PluginID),
	}

	config := opt.New(p)
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	if err := p.Init(); err != nil {
		return nil, err
	}

	return p, nil
}

func (c *coinbase) Init() error {
	errs := errors.New("coinbase plugin initialization errors")

	for _, account := range c.Accounts {
		errs = errors.Join(errs, account.Init())
	}
	return errs
}

// Compute is not implemented for coinbase plugin
func (p *coinbase) Compute(input internal.Series) internal.Series {
	return input
}

func (p *coinbase) Process(input internal.Tick) internal.Tick {
	return input
}

// GetAccounts returns the accounts for the coinbase plugin
func (p *coinbase) GetAccounts() []Account {
	return p.Accounts
}

func init() {
	brokers.Add(PluginID, New)
}
