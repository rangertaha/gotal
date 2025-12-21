package coinbase

import (
	"errors"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/plugins/brokers"
)

const (
	PluginID          = "COINBASE"
	PluginName        = "Coinbase"
	PluginDescription = "Coinbase broker"
	PluginHCL         = `
broker "coinbase" {
	api_key = "YOUR_API_KEY"
	base_url = "https://api.coinbase.com"
	version = "v2"
	account {
		id = "YOUR_ACCOUNT_ID"
		name = "YOUR_ACCOUNT_NAME"
		currency = "YOUR_ACCOUNT_CURRENCY"
		balance = "YOUR_ACCOUNT_BALANCE"
		allowances {
			buy = "0" // max amount of money that can be bought	
			sell = "0" // max amount of money that can be sold
			transfer = "0" // max amount of money that can be transferred
			deposit = "0" // max amount of money that can be deposited
			withdraw = "0" // max amount of money that can be withdrawn
		}
		limits {
			buy = "0" // max amount of money that can be bought	
			sell = "0" // max amount of money that can be sold
			transfer = "0" // max amount of money that can be transferred
			deposit = "0" // max amount of money that can be deposited
			withdraw = "0" // max amount of money that can be withdrawn
		}
		fees {
			buy = "0" // fee for buying
			sell = "0" // fee for selling	
			transfer = "0" // fee for transferring
			deposit = "0" // fee for depositing
			withdraw = "0" // fee for withdrawing
		}
	}
}
`
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
}

func New(opts ...internal.ConfigOption) (internal.Plugin, error) {

	c := &coinbase{
		BaseURL: CoinbaseBaseURL,
		Version: CoinbaseVersion,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
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
