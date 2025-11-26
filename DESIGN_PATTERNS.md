# Elegant Stream/Batch Design Patterns

## 1. Unified Processor Pattern (Most Elegant)

```go
// Core abstraction that works for both modes
type Processor[T any] interface {
    Process(input T) T
    Configure(opts ...OptFunc) Processor[T]
}

// Generic processor that adapts to both stream and batch
type UnifiedProcessor[T any] struct {
    name     string
    stateful bool
    state    any
    processFn func(T, any) (T, any)
    resetFn   func() any
}

func (p *UnifiedProcessor[T]) Process(input T) T {
    if p.stateful {
        result, newState := p.processFn(input, p.state)
        p.state = newState
        return result
    }
    result, _ := p.processFn(input, nil)
    return result
}

// Automatic adaptation for both modes
func (p *UnifiedProcessor[T]) Batch(series *Series) *Series {
    p.state = p.resetFn()
    output := NewSeries(p.name)
    for _, tick := range series.Ticks() {
        if result := p.Process(tick); !result.IsEmpty() {
            output.Add(result)
        }
    }
    return output
}

func (p *UnifiedProcessor[T]) Stream() StreamProcessor {
    return func(input <-chan *Tick) <-chan *Tick {
        output := make(chan *Tick)
        go func() {
            defer close(output)
            p.state = p.resetFn()
            for tick := range input {
                if result := p.Process(tick); !result.IsEmpty() {
                    output <- result
                }
            }
        }()
        return output
    }
}
```

## 2. Functional Pipeline Pattern

```go
// Composable functions that work in both modes
type ProcessorFunc func(*Tick) *Tick

type Pipeline struct {
    processors []ProcessorFunc
    stateful   map[int]any
}

func (p *Pipeline) Add(processor ProcessorFunc) *Pipeline {
    p.processors = append(p.processors, processor)
    return p
}

func (p *Pipeline) Process(tick *Tick) *Tick {
    current := tick
    for _, processor := range p.processors {
        current = processor(current)
        if current == nil {
            break
        }
    }
    return current
}

// Auto-adapts to both modes
func (p *Pipeline) Batch(series *Series) *Series {
    output := NewSeries("pipeline")
    for _, tick := range series.Ticks() {
        if result := p.Process(tick); result != nil {
            output.Add(result)
        }
    }
    return output
}

func (p *Pipeline) Stream(input <-chan *Tick) <-chan *Tick {
    output := make(chan *Tick)
    go func() {
        defer close(output)
        for tick := range input {
            if result := p.Process(tick); result != nil {
                output <- result
            }
        }
    }()
    return output
}
```

## 3. Mode-Aware Interface Pattern

```go
type ProcessingMode int

const (
    BatchMode ProcessingMode = iota
    StreamMode
    AutoMode
)

type ModeAwareProcessor interface {
    SetMode(mode ProcessingMode)
    Process(input any) any
    
    // Optional optimizations
    BatchProcess(series *Series) *Series
    StreamProcess(stream <-chan *Tick) <-chan *Tick
}

type BaseProcessor struct {
    mode ProcessingMode
    name string
}

func (p *BaseProcessor) Process(input any) any {
    switch p.mode {
    case BatchMode:
        return p.batchProcess(input.(*Tick))
    case StreamMode:
        return p.streamProcess(input.(*Tick))
    case AutoMode:
        return p.adaptiveProcess(input)
    }
    return nil
}
```

## 4. Context-Driven Pattern (Most Flexible)

```go
type ProcessingContext struct {
    Mode      ProcessingMode
    StateKey  string
    Metadata  map[string]any
    
    // Stream-specific
    InputChan  <-chan *Tick
    OutputChan chan<- *Tick
    
    // Batch-specific  
    Series     *Series
    BatchSize  int
    
    // Shared
    Options    map[string]any
}

type ContextualProcessor interface {
    Process(ctx *ProcessingContext, input *Tick) (*Tick, error)
    Reset(ctx *ProcessingContext) error
}

// Usage becomes very clean:
func (sma *SMA) Process(ctx *ProcessingContext, input *Tick) (*Tick, error) {
    // Same logic works for both stream and batch
    sma.window.Add(input.GetField(sma.InputField))
    
    if sma.window.Len() >= sma.Period {
        result := sma.window.Average()
        return input.Clone().SetField(sma.OutputField, result), nil
    }
    
    return NewEmptyTick(), nil
}
```
