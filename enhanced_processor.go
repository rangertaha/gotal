package gotal

import (
	"context"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/tick"
)

// ===== ENHANCED UNIFIED PROCESSOR =====

// ProcessorMode determines how the processor operates
type ProcessorMode string

const (
	ModeAuto   ProcessorMode = "auto"   // Automatically detect and adapt
	ModeBatch  ProcessorMode = "batch"  // Process entire series at once
	ModeStream ProcessorMode = "stream" // Process real-time data streams
)

// UnifiedProcessor elegantly handles both stream and batch processing
type UnifiedProcessor interface {
	// Core processing (mode-agnostic)
	Process(input *tick.Tick) *tick.Tick

	// Batch processing (optimized)
	Batch(series *series.Series) *series.Series

	// Stream processing (real-time)
	Stream(ctx context.Context, input <-chan *tick.Tick) <-chan *tick.Tick

	// Configuration
	WithMode(mode ProcessorMode) UnifiedProcessor
	Reset() UnifiedProcessor
	Clone() UnifiedProcessor
}

// ===== ENHANCED VERSION OF YOUR SMA =====

type EnhancedSMA struct {
	// Your existing fields
	Name   string
	Input  string
	Output string
	Period int

	// Enhanced fields for unified processing
	mode   ProcessorMode
	window []float64
	series *series.Series

	// Function composition (your current approach enhanced)
	processFn func(*tick.Tick) *tick.Tick
	batchFn   func(*series.Series) *series.Series
	streamFn  func(context.Context, <-chan *tick.Tick) <-chan *tick.Tick
}

func NewEnhancedSMA(opts ...internal.OptFunc) *EnhancedSMA {
	sma := &EnhancedSMA{
		Name:   "sma",
		Input:  "close",
		Output: "sma",
		Period: 20,
		mode:   ModeAuto,
		window: make([]float64, 0),
		series: series.New("sma"),
	}

	// Apply options...

	// Set up function composition
	sma.processFn = sma.processCore
	sma.batchFn = sma.batchCore
	sma.streamFn = sma.streamCore

	return sma
}

// ===== CORE PROCESSING LOGIC (SHARED) =====

func (sma *EnhancedSMA) processCore(input *tick.Tick) *tick.Tick {
	if !input.HasField(sma.Input) {
		return tick.NewEmpty()
	}

	value := input.GetField(sma.Input)
	sma.addToWindow(value)

	if len(sma.window) >= sma.Period {
		avg := sma.calculateAverage()
		return input.Clone().SetField(sma.Output, avg)
	}

	return tick.NewEmpty()
}

// ===== MODE IMPLEMENTATIONS =====

// Process - Single tick processing (works for both modes)
func (sma *EnhancedSMA) Process(input *tick.Tick) *tick.Tick {
	return sma.processFn(input)
}

// Batch - Optimized batch processing
func (sma *EnhancedSMA) Batch(input *series.Series) *series.Series {
	if sma.mode == ModeStream {
		// Fallback to stream processing for consistency
		return sma.batchViaStream(input)
	}
	return sma.batchFn(input)
}

func (sma *EnhancedSMA) batchCore(input *series.Series) *series.Series {
	output := series.New(sma.Name)
	sma.Reset() // Reset state for batch processing

	for _, t := range input.Ticks() {
		if result := sma.processCore(t); !result.IsEmpty() {
			output.Add(result)
		}
	}

	return output
}

// Stream - Real-time stream processing
func (sma *EnhancedSMA) Stream(ctx context.Context, input <-chan *tick.Tick) <-chan *tick.Tick {
	return sma.streamFn(ctx, input)
}

func (sma *EnhancedSMA) streamCore(ctx context.Context, input <-chan *tick.Tick) <-chan *tick.Tick {
	output := make(chan *tick.Tick, 100) // Buffered channel

	go func() {
		defer close(output)
		sma.Reset() // Reset state for stream processing

		for {
			select {
			case <-ctx.Done():
				return
			case t, ok := <-input:
				if !ok {
					return
				}
				if result := sma.processCore(t); !result.IsEmpty() {
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

// ===== UTILITY METHODS =====

func (sma *EnhancedSMA) WithMode(mode ProcessorMode) UnifiedProcessor {
	clone := sma.Clone().(*EnhancedSMA)
	clone.mode = mode
	return clone
}

func (sma *EnhancedSMA) Reset() UnifiedProcessor {
	sma.window = sma.window[:0] // Clear window efficiently
	sma.series.Reset()
	return sma
}

func (sma *EnhancedSMA) Clone() UnifiedProcessor {
	return &EnhancedSMA{
		Name:   sma.Name,
		Input:  sma.Input,
		Output: sma.Output,
		Period: sma.Period,
		mode:   sma.mode,
		window: make([]float64, 0, sma.Period),
		series: series.New(sma.Name),
	}
}

// ===== HELPER METHODS =====

func (sma *EnhancedSMA) addToWindow(value float64) {
	sma.window = append(sma.window, value)
	if len(sma.window) > sma.Period {
		sma.window = sma.window[1:] // Remove oldest value
	}
}

func (sma *EnhancedSMA) calculateAverage() float64 {
	if len(sma.window) == 0 {
		return 0
	}

	sum := 0.0
	for _, v := range sma.window {
		sum += v
	}
	return sum / float64(len(sma.window))
}

func (sma *EnhancedSMA) batchViaStream(input *series.Series) *series.Series {
	// Convert series to stream and back (for consistency when in stream mode)
	inputChan := make(chan *tick.Tick, len(input.Ticks()))
	for _, t := range input.Ticks() {
		inputChan <- t
	}
	close(inputChan)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	outputChan := sma.Stream(ctx, inputChan)

	output := series.New(sma.Name)
	for result := range outputChan {
		output.Add(result)
	}

	return output
}

// ===== ELEGANT USAGE EXAMPLES =====

func ExampleUnifiedUsage() {
	// Same processor, multiple modes
	sma := NewEnhancedSMA()

	// 1. Auto mode (adapts automatically)
	sma = sma.WithMode(ModeAuto).(*EnhancedSMA)

	// 2. Batch processing
	historicalData := loadHistoricalData()
	results := sma.Batch(historicalData)

	// 3. Stream processing
	ctx := context.Background()
	liveData := connectToLiveFeed()
	liveResults := sma.Reset().Stream(ctx, liveData)

	// 4. Single tick processing
	tick := &tick.Tick{}
	result := sma.Process(tick)
}

// ===== REGISTRY PATTERN ENHANCED =====

type ProcessorRegistry struct {
	processors map[string]func(...internal.OptFunc) UnifiedProcessor
}

func (r *ProcessorRegistry) Register(name string, constructor func(...internal.OptFunc) UnifiedProcessor) {
	r.processors[name] = constructor
}

func (r *ProcessorRegistry) Create(name string, opts ...internal.OptFunc) UnifiedProcessor {
	if constructor, exists := r.processors[name]; exists {
		return constructor(opts...)
	}
	return nil
}

// Usage becomes extremely clean
func CreateProcessor(name string, mode ProcessorMode, opts ...internal.OptFunc) UnifiedProcessor {
	registry := &ProcessorRegistry{processors: make(map[string]func(...internal.OptFunc) UnifiedProcessor)}
	registry.Register("sma", func(opts ...internal.OptFunc) UnifiedProcessor {
		return NewEnhancedSMA(opts...)
	})

	return registry.Create(name, opts...).WithMode(mode)
}
