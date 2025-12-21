package gen

import (
	"errors"
	"fmt"
	"time"

	_ "embed"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/opt"
	"github.com/rangertaha/gotal/internal/plugins/providers"
	"github.com/rangertaha/gotal/internal/series"
)

const (
	PluginID          = "gen"
	PluginName        = "Generator"
	PluginDescription = "Generates mock data for testing and development"
)

//go:embed gen.hcl
var pluginHCLTemplate string

type Generator struct {
	SeriesName string `hcl:"name,optional"` // Series name of the generator

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

func NewGenerator(opts ...internal.ConfigOption) (internal.Plugin, error) {
	// Initial plugin with default values
	p := &Generator{
		series: series.New(PluginID),
	}
	config := opt.New(p)
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	
	return p, nil
}

func (p *Generator) Init() error {
	fmt.Printf("Generator init: %+v\n", p)
	var errs error

	for _, sine := range p.Sine {
		if err := sine.Init(); err != nil {
			errs = errors.Join(errs, err)
		}
	}
	for _, square := range p.Square {
		if err := square.Init(); err != nil {
			errs = errors.Join(errs, err)
		}
	}
	for _, triangle := range p.Triangle {
		if err := triangle.Init(); err != nil {
			errs = errors.Join(errs, err)
		}
	}
	for _, sawtooth := range p.Sawtooth {
		if err := sawtooth.Init(); err != nil {
			errs = errors.Join(errs, err)
		}
	}
	return errs
}

func (p *Generator) ID() string {
	return PluginID
}

func (p *Generator) Name() string {
	return PluginName
}

func (p *Generator) Description() string {
	return PluginDescription
}

func (p *Generator) Reset() error {
	return nil
}

func (p *Generator) Ready() bool {
	return true
}

func (p *Generator) Compute(input internal.Series) internal.Series {
	if input.IsEmpty() {
		input = p.series
	}
	for _, sine := range p.Sine {
		input = sine.Compute(input)
	}
	for _, square := range p.Square {
		input = square.Compute(input)
	}
	for _, triangle := range p.Triangle {
		input = triangle.Compute(input)
	}
	for _, sawtooth := range p.Sawtooth {
		input = sawtooth.Compute(input)
	}
	return input
}
func (p *Generator) Stream(input internal.Stream) internal.Stream {

	return input
}

func (p *Generator) Process(input internal.Stream) internal.Stream {
	// for _, sine := range p.Sine {
	// 	input.Update(sine.Process(input))
	// }
	// for _, square := range p.Square {
	// 	input.Update(square.Process(input))
	// }
	// for _, triangle := range p.Triangle {
	// 	input.Update(triangle.Process(input))
	// }
	// for _, sawtooth := range p.Sawtooth {
	// 	input.Update(sawtooth.Process(input))
	// }
	return input
}

func init() {
	providers.Add(PluginID, NewGenerator)
}
