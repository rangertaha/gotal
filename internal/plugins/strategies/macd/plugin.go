package macd

import (
	_ "embed"
)

const (
	PluginID          = "MACD"
	PluginName        = "Moving Average Convergence Divergence."
	PluginDescription = "MAC is used to identify trends in the price of a security."
)

//go:embed macd.hcl
var PluginHCLTemplate string

func (p *macd) ID() string {
	return PluginID
}

func (p *macd) Title() string {
	return PluginName
}

func (p *macd) Description() string {
	return PluginDescription
}

func (p *macd) Template() string {
	return PluginHCLTemplate
}
