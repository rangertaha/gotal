package ma

import (
	_ "embed"
)

const (
	emaPluginID          = "EMA"
	emaPluginName        = "Exponential Moving Average"
	emaPluginDescription = "Exponential Moving Average is a technical indicator that smooths out price data by giving more weight to recent prices."
)

//go:embed ema.hcl
var emaPluginHCLTemplate string

func (p *ema) ID() string {
	return emaPluginID
}

func (p *ema) Title() string {
	return emaPluginName
}

func (p *ema) Description() string {
	return emaPluginDescription
}

func (p *ema) Template() string {
	return emaPluginHCLTemplate
}
