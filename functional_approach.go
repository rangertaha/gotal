package gotal

import (
	"context"

	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/tick"
)

// ===== FUNCTIONAL APPROACH (MOST ELEGANT) =====

// ProcessorFn is the core processing function type
type ProcessorFn func(*tick.Tick, *State) (*tick.Tick, *State)

// State holds processor state (for stateful operations like moving averages)
type State struct {
	Data   map[string]any
	Window []float64
	Count  int64
}

// NewState creates a new processing state
func NewState() *State {
	return &State{
		Data:   make(map[string]any),
		Window: make([]float64, 0),
		Count:  0,
	}
}

// Processor wraps a ProcessorFn with metadata and state management
type Processor struct {
	Name      string
	Fn        ProcessorFn
	InitState func() *State
	ResetFn   func(*State) *State

	// Current state (for stream processing)
	state *State
}

// NewProcessor creates a functional processor
func NewProcessor(name string, fn ProcessorFn, initState func() *State) *Processor {
	p := &Processor{
		Name:      name,
		Fn:        fn,
		InitState: initState,
		ResetFn:   func(s *State) *State { return initState() },
	}
	p.state = p.InitState()
	return p
}

// ===== CORE OPERATIONS (ELEGANT & UNIFIED) =====

// Process single tick (works for both stream and batch)
func (p *Processor) Process(input *tick.Tick) *tick.Tick {
	result, newState := p.Fn(input, p.state)
	p.state = newState
	return result
}

// Batch processes entire series
func (p *Processor) Batch(series *series.Series) *series.Series {
	output := series.New(p.Name)
	p.state = p.InitState() // Reset state for batch

	for _, t := range series.Ticks() {
		if result := p.Process(t); !result.IsEmpty() {
			output.Add(result)
		}
	}

	return output
}

// Stream processes real-time data
func (p *Processor) Stream(ctx context.Context, input <-chan *tick.Tick) <-chan *tick.Tick {
	output := make(chan *tick.Tick, 100)

	go func() {
		defer close(output)
		p.state = p.InitState() // Reset state for stream

		for {
			select {
			case <-ctx.Done():
				return
			case t, ok := <-input:
				if !ok {
					return
				}
				if result := p.Process(t); !result.IsEmpty() {
					select {
					case output <- result:
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}()

	return output
}

// Reset processor state
func (p *Processor) Reset() *Processor {
	p.state = p.ResetFn(p.state)
	return p
}

// ===== PROCESSOR BUILDERS (ELEGANT FACTORY FUNCTIONS) =====

// SMA creates a Simple Moving Average processor
func SMA(period int, inputField, outputField string) *Processor {
	return NewProcessor("sma", func(input *tick.Tick, state *State) (*tick.Tick, *State) {
		if !input.HasField(inputField) {
			return tick.NewEmpty(), state
		}

		value := input.GetField(inputField)

		// Add to window
		newWindow := append(state.Window, value)
		if len(newWindow) > period {
			newWindow = newWindow[1:] // Remove oldest
		}

		newState := &State{
			Data:   state.Data,
			Window: newWindow,
			Count:  state.Count + 1,
		}

		if len(newWindow) >= period {
			sum := 0.0
			for _, v := range newWindow {
				sum += v
			}
			avg := sum / float64(len(newWindow))

			result := input.Clone().SetField(outputField, avg)
			return result, newState
		}

		return tick.NewEmpty(), newState
	}, func() *State {
		return NewState()
	})
}

// EMA creates an Exponential Moving Average processor
func EMA(period int, inputField, outputField string) *Processor {
	alpha := 2.0 / (float64(period) + 1.0)

	return NewProcessor("ema", func(input *tick.Tick, state *State) (*tick.Tick, *State) {
		if !input.HasField(inputField) {
			return tick.NewEmpty(), state
		}

		value := input.GetField(inputField)
		newState := &State{
			Data:  make(map[string]any),
			Count: state.Count + 1,
		}

		var ema float64
		if state.Count == 0 {
			ema = value // First value
		} else {
			prevEMA := state.Data["ema"].(float64)
			ema = alpha*value + (1-alpha)*prevEMA
		}

		newState.Data["ema"] = ema

		result := input.Clone().SetField(outputField, ema)
		return result, newState
	}, func() *State {
		return NewState()
	})
}

// ===== PROCESSOR COMPOSITION (PIPELINE PATTERN) =====

// Pipeline composes multiple processors
type Pipeline struct {
	processors []*Processor
	name       string
}

// NewPipeline creates a new processing pipeline
func NewPipeline(name string) *Pipeline {
	return &Pipeline{
		processors: make([]*Processor, 0),
		name:       name,
	}
}

// Add appends a processor to the pipeline
func (p *Pipeline) Add(processor *Processor) *Pipeline {
	p.processors = append(p.processors, processor)
	return p
}

// Process runs the entire pipeline on a single tick
func (p *Pipeline) Process(input *tick.Tick) *tick.Tick {
	current := input
	for _, processor := range p.processors {
		if current.IsEmpty() {
			break
		}
		current = processor.Process(current)
	}
	return current
}

// Batch runs the pipeline on an entire series
func (p *Pipeline) Batch(input *series.Series) *series.Series {
	// Reset all processors
	for _, processor := range p.processors {
		processor.Reset()
	}

	output := series.New(p.name)
	for _, t := range input.Ticks() {
		if result := p.Process(t); !result.IsEmpty() {
			output.Add(result)
		}
	}

	return output
}

// Stream runs the pipeline on a real-time stream
func (p *Pipeline) Stream(ctx context.Context, input <-chan *tick.Tick) <-chan *tick.Tick {
	output := make(chan *tick.Tick, 100)

	go func() {
		defer close(output)

		// Reset all processors
		for _, processor := range p.processors {
			processor.Reset()
		}

		for {
			select {
			case <-ctx.Done():
				return
			case t, ok := <-input:
				if !ok {
					return
				}
				if result := p.Process(t); !result.IsEmpty() {
					select {
					case output <- result:
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}()

	return output
}

// ===== ELEGANT USAGE EXAMPLES =====

func ExampleFunctionalApproach() {
	// 1. Single processors
	sma20 := SMA(20, "close", "sma20")
	ema12 := EMA(12, "close", "ema12")

	// 2. Pipeline composition
	pipeline := NewPipeline("technical_analysis").
		Add(sma20).
		Add(ema12).
		Add(SMA(5, "sma20", "sma_of_sma")) // SMA of SMA!

	// 3. Works seamlessly for both batch and stream

	// Batch processing
	historicalData := loadHistoricalData()
	batchResults := pipeline.Batch(historicalData)

	// Stream processing
	ctx := context.Background()
	liveData := connectToLiveFeed()
	streamResults := pipeline.Stream(ctx, liveData)

	// Process results...
}

// ===== ADVANCED COMPOSITION PATTERNS =====

// Parallel processes multiple pipelines in parallel
func Parallel(pipelines ...*Pipeline) *Pipeline {
	return NewPipeline("parallel").Add(NewProcessor("parallel",
		func(input *tick.Tick, state *State) (*tick.Tick, *State) {
			result := input.Clone()

			// Run all pipelines and merge results
			for _, pipeline := range pipelines {
				pipeResult := pipeline.Process(input)
				if !pipeResult.IsEmpty() {
					// Merge fields from pipeline result
					for _, field := range pipeResult.FieldNames() {
						result.SetField(field, pipeResult.GetField(field))
					}
				}
			}

			return result, state
		}, func() *State { return NewState() }))
}

// Conditional creates conditional processing
func Conditional(condition func(*tick.Tick) bool, processor *Processor) *Processor {
	return NewProcessor("conditional", func(input *tick.Tick, state *State) (*tick.Tick, *State) {
		if condition(input) {
			return processor.Process(input), state
		}
		return input, state // Pass through unchanged
	}, func() *State { return NewState() })
}

// Example: Complex strategy
func ExampleComplexStrategy() {
	// Create conditional processors
	macdStrategy := NewPipeline("macd_strategy").
		Add(EMA(12, "close", "ema12")).
		Add(EMA(26, "close", "ema26")).
		Add(NewProcessor("macd", func(input *tick.Tick, state *State) (*tick.Tick, *State) {
			if input.HasField("ema12") && input.HasField("ema26") {
				macd := input.GetField("ema12") - input.GetField("ema26")
				return input.Clone().SetField("macd", macd), state
			}
			return tick.NewEmpty(), state
		}, func() *State { return NewState() }))

	rsiStrategy := NewPipeline("rsi_strategy").
		Add(SMA(14, "close", "rsi")) // Simplified RSI

	// Combine strategies
	multiStrategy := Parallel(macdStrategy, rsiStrategy)

	// Works for both stream and batch!
	results := multiStrategy.Batch(loadHistoricalData())
}
