package coinbase

import (
	"errors"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/tick"
)

type Account struct {
	ID        string     `hcl:"id"`               // Account ID
	Name      string     `hcl:"name"`             // Account Name
	Currency  string     `hcl:"currency"`         // Account Currency
	Balance   float64    `hcl:"balance"`          // Account Balance
	Allowance *allowance `hcl:"allowances,block"` // Account Allowances
	Limit     *limit     `hcl:"limits,block"`     // Account Limits
	Fee       *fee       `hcl:"fees,block"`       // Account Fees

	// orders []*order
}

type allowance struct {
	Buy      float64 `hcl:"buy"`      // Max amount of money that can be bought
	Sell     float64 `hcl:"sell"`     // Max amount of money that can be sold
	Transfer float64 `hcl:"transfer"` // Max amount of money that can be transferred
	Deposit  float64 `hcl:"deposit"`  // Max amount of money that can be deposited
	Withdraw float64 `hcl:"withdraw"` // Max amount of money that can be withdrawn
}

type limit struct {
	Buy      float64 `hcl:"buy"`      // Max amount of money that can be bought
	Sell     float64 `hcl:"sell"`     // Max amount of money that can be sold
	Transfer float64 `hcl:"transfer"` // Max amount of money that can be transferred
	Deposit  float64 `hcl:"deposit"`  // Max amount of money that can be deposited
	Withdraw float64 `hcl:"withdraw"` // Max amount of money that can be withdrawn
}

type fee struct {
	Buy      float64 `hcl:"buy"`      // Fee for buying
	Sell     float64 `hcl:"sell"`     // Fee for selling
	Transfer float64 `hcl:"transfer"` // Fee for transferring
	Deposit  float64 `hcl:"deposit"`  // Fee for depositing
	Withdraw float64 `hcl:"withdraw"` // Fee for withdrawing
}

func (a *Account) Init(opts ...internal.PluginOptions) error {
	return nil
}

func (a *Account) Process(input *tick.Tick) *tick.Tick {
	return nil
}

func (p *Account) Buy(symbol string, quantity float64) (err error) {
	return errors.New("not implemented")
}

func (a *Account) Sell(symbol string, quantity float64) (err error) {
	return errors.New("not implemented")
}

func (a *Account) Transfer(symbol string, quantity float64) (err error) {
	return errors.New("not implemented")
}

func (a *Account) Deposit(symbol string, quantity float64) (err error) {
	return errors.New("not implemented")
}

func (a *Account) Withdraw(symbol string, quantity float64) (err error) {
	return errors.New("not implemented")
}

func (a *Account) GetAllowance() (allowance *allowance, err error) {
	return a.Allowance, nil
}

func (a *Account) GetLimit() (limit *limit, err error) {
	return a.Limit, nil
}

func (a *Account) GetFee() (fee *fee, err error) {
	return a.Fee, nil
}

func (p *Account) GetBalance() (balance float64, err error) {
	return p.Balance, nil
}

func (a *Account) GetCurrency() (currency string, err error) {
	return a.Currency, nil
}

func (a *Account) GetName() (name string, err error) {
	return a.Name, nil
}
