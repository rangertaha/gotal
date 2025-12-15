package plugins

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/series"
	"github.com/rangertaha/gotal/internal/stream"
)

type Plugin struct {
	PID      string   `hcl:"id"`
	Title    string   `hcl:"name"`
	Summary  string   `hcl:"description"`
	Fields   []string `hcl:"inputs,optional,default=[value]"` // input field names to compute
	Template string   `hcl:"-"`                               // template to compute the plugin

	// input data
	Params      internal.Options // input parameters
	Series      *series.Series    // input data series
	Initialized bool              // ready to compute
}

type PluginFunc func(plugin *Plugin) (series *series.Series, stream *stream.Stream)

func (p *Plugin) ID() string {
	return p.PID
}

func (p *Plugin) Name() string {
	return p.Title
}

func (p *Plugin) Description() string {
	return p.Summary
}

func (p *Plugin) Ready() bool {
	return p.Initialized
}

func (p *Plugin) Options() internal.Options {
	return p.Params
}

func (p *Plugin) Batch() *series.Series {
	return series.New(p.Name())
}

func (p *Plugin) Stream() *stream.Stream {
	return stream.New(p.Name())
}
