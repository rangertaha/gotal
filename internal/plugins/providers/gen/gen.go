package gen

import (
	"errors"
	"time"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/opt"
	"github.com/rangertaha/gotal/internal/plugins/providers"
	"github.com/rangertaha/gotal/internal/series"
	"github.com/rangertaha/gotal/internal/stream"
)

type Generator struct {
	Name string `hcl:"name,optional"` // Series name of the generator

	// Timing parameters (optional)
	StartDate int64         `hcl:"start,optional"`    // Start time
	EndDate   int64         `hcl:"end,optional"`      // End time
	Interval  time.Duration `hcl:"interval,optional"` // Interval (e.g. 1m, 5m, 15m, 30m, 1h, 4h, 1d)

	Sine     []*Sine     `hcl:"sine,block"`
	Square   []*Square   `hcl:"square,block"`
	Triangle []*Triangle `hcl:"triangle,block"`
	Sawtooth []*Sawtooth `hcl:"sawtooth,block"`

	// Series internal state
	series internal.Series
}

func NewGenerator(opts ...internal.ConfigOption) (p internal.Plugin, err error) {
	// Initial plugin with default values
	p = &Generator{
		Name:      PluginID,
		StartDate: time.Now().AddDate(-10, 0, 0).Unix(),
		EndDate:   time.Now().Unix(),
		Interval:  time.Minute * 1,
		Sine:      []*Sine{},
		Square:    []*Square{},
		Triangle:  []*Triangle{},
		Sawtooth:  []*Sawtooth{},
		series:    series.New(PluginID),
	}

	// Create a new configurator with the plugin
	config := opt.New(p)
	for _, opt := range opts {
		if err = opt(config); err != nil {
			return nil, err
		}
	}

	if initializer, ok := p.(internal.Initializer); ok {
		if err = initializer.Init(config); err != nil {
			return nil, err
		}
	}

	return p, nil
}

// Init initializes plugin inputs with the provided configurator
func (p *Generator) Init(config internal.Configurator) (err error) {

	// Initialize sub-plugins
	for _, sine := range p.Sine {
		if err = sine.Init(config); err != nil {
			err = errors.Join(err, err)
		}
	}
	for _, square := range p.Square {
		if err = square.Init(config); err != nil {
			err = errors.Join(err, err)
		}
	}
	for _, triangle := range p.Triangle {
		if err = triangle.Init(config); err != nil {
			err = errors.Join(err, err)
		}
	}
	for _, sawtooth := range p.Sawtooth {
		if err = sawtooth.Init(config); err != nil {
			err = errors.Join(err, err)
		}
	}
	return
}

func (p *Generator) Reset() error {
	return nil
}

func (p *Generator) Ready() bool {
	return true
}

func (p *Generator) Compute() internal.Series {
	// Compute the series for each sine wave
	for _, sine := range p.Sine {
		p.series = sine.Compute(p.series)
	}
	for _, square := range p.Square {
		p.series = square.Compute(p.series)
	}
	for _, triangle := range p.Triangle {
		p.series = triangle.Compute(p.series)
	}
	for _, sawtooth := range p.Sawtooth {
		p.series = sawtooth.Compute(p.series)
	}
	return p.series
}
func (p *Generator) Stream() internal.Stream {
	return stream.New(p.Name)
}

func (p *Generator) Process(input internal.Tick) internal.Tick {

	return input
}

func init() {
	providers.Add(PluginID, NewGenerator)
}
