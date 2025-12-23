package coinbase

import (
	_ "embed"
)

const (
	PluginID          = "CB"
	PluginName        = "Coinbase"
	PluginDescription = "Coinbase broker"
)

//go:embed coinbase.hcl
var pluginHCLTemplate string

func (p *coinbase) ID() string {
	return PluginID
}

func (p *coinbase) Title() string {
	return PluginName
}

func (p *coinbase) Description() string {
	return PluginDescription
}

func (p *coinbase) Template() string {
	return pluginHCLTemplate
}
