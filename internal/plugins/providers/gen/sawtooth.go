package gen

import (
	"fmt"

	"github.com/rangertaha/gotal/internal"
)

type Sawtooth struct {
	Name      string  `hcl:"name,label"`         // Name (default: sawtooth)
	Periods   int     `hcl:"periods,optional"`   // Number of periods in dataset (default: 100)
	Amplitude float64 `hcl:"amplitude,optional"` // Amplitude (default: 1.0)
	Frequency float64 `hcl:"frequency,optional"` // Frequency (default: 1.0)
	Phase     float64 `hcl:"phase,optional"`     // Phase (default: 0.0)
	Offset    float64 `hcl:"offset,optional"`    // Offset (default: 0.0)
}

func (s *Sawtooth) Init(config internal.Configurator) (err error) {
	fmt.Printf("Sawtooth init: %+v\n", s)
	return nil
}

func (s *Sawtooth) Compute(input internal.Series) (output internal.Series) {
	fmt.Printf("Sawtooth compute: %+v\n", s)
	return input
}

func (s *Sawtooth) Process(input internal.Stream) (output internal.Stream) {
	return input
}
