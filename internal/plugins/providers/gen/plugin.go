package gen

import (
	_ "embed"
)

const (
	PluginID          = "generator"
	PluginName        = "Price Generator"
	PluginDescription = "Generates mock data for testing and development"
)

//go:embed gen.hcl
var pluginHCLTemplate string

func (p *Generator) ID() string {
	return PluginID
}

func (p *Generator) Title() string {
	return PluginName
}

func (p *Generator) Description() string {
	return PluginDescription
}

func (p *Generator) Template() string {
	return pluginHCLTemplate
}
