package gen

import (
	"fmt"
	"math"
	"time"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/tick"
)

type Sine struct {
	Name      string        `hcl:"name,label"`         // Name (default: sine)
	Periods   int           `hcl:"periods,optional"`   // Number of periods in dataset (default: 100)
	Amplitude float64       `hcl:"amplitude,optional"` // Amplitude (default: 1.0)
	Frequency float64       `hcl:"frequency,optional"` // Frequency (default: 1.0)
	Phase     float64       `hcl:"phase,optional"`     // Phase (default: 0.0)
	Offset    float64       `hcl:"offset,optional"`    // Offset (default: 0.0)
	Duration  time.Duration `hcl:"duration,optional"`  // Duration (default: 1s)
}

func (s *Sine) Init() error {
	fmt.Printf("Sine init: %+v\n", s)
	return nil
}

func (s *Sine) Compute(input internal.Series) internal.Series {
	fmt.Printf("Sine compute: %+v\n", s)
	t := time.Now()

	// Apply the offset to the starting time
	t = t.Add(time.Duration(s.Offset) * s.Duration)

	for i := 0; i < s.Periods; i++ {
		// Generate one complete sine wave cycle (2π radians)
		for i := 0.0; i <= 2*math.Pi; i += s.Frequency {
			value := (s.Amplitude * math.Sin(i)) + s.Amplitude
			tick := tick.New(
				tick.WithTime(t),
				tick.WithDuration(s.Duration),
				tick.WithFields(map[string]float64{s.Name: value}),
				// tick.WithTags(map[string]string{"sine": "true"}),
			)
			input.Add(tick)
			t = t.Add(s.Duration)
		}
	}

	return input
}

func (s *Sine) Process(input internal.Stream) (output internal.Stream) {
	return nil
}
