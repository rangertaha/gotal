package main

import (
	"context"
	"fmt"
	"time"
	
	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/tick"
)

// ===== INTEGRATION WITH YOUR EXISTING ARCHITECTURE =====

// Enhance your existing SMA to use the elegant pattern
func (sma *sma) ToUnifiedProcessor() *Processor {
	return NewProcessor("sma", func(input *tick.Tick, state *State) (*tick.Tick, *State) {
		// Your existing stream logic wrapped in functional pattern
		return sma.stream(input), state
	}, func() *State {
		return NewState()
	})
}

// ===== ELEGANT USAGE EXAMPLES =====

func main() {
	fmt.Println("🚀 Elegant Stream/Batch Processing Examples")
	
	// Example 1: Simple usage (same API for both modes)
	demonstrateUnifiedAPI()
	
	// Example 2: Pipeline composition
	demonstratePipelineComposition()
	
	// Example 3: Real-world trading strategy
	demonstrateTradingStrategy()
	
	// Example 4: Performance comparison
	demonstratePerformanceComparison()
}

// ===== EXAMPLE 1: UNIFIED API =====

func demonstrateUnifiedAPI() {
	fmt.Println("\n📊 Example 1: Unified API")
	
	// Create the same processor for both modes
	processor := SMA(20, "close", "sma20")
	
	// Historical analysis (batch)
	historicalData := generateHistoricalData(1000)
	fmt.Printf("Processing %d historical ticks...\n", len(historicalData.Ticks()))
	
	batchResults := processor.Reset().Batch(historicalData)
	fmt.Printf("Batch results: %d output ticks\n", len(batchResults.Ticks()))
	
	// Live trading (stream)
	fmt.Println("Starting live stream processing...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	liveData := generateLiveStream(ctx)
	streamResults := processor.Reset().Stream(ctx, liveData)
	
	count := 0
	for result := range streamResults {
		if result != nil {
			count++
		}
	}
	fmt.Printf("Stream results: %d output ticks\n", count)
}

// ===== EXAMPLE 2: PIPELINE COMPOSITION =====

func demonstratePipelineComposition() {
	fmt.Println("\n🔗 Example 2: Pipeline Composition")
	
	// Build a technical analysis pipeline
	pipeline := NewPipeline("technical_analysis").
		Add(SMA(10, "close", "sma_fast")).
		Add(SMA(20, "close", "sma_slow")).
		Add(EMA(12, "close", "ema12")).
		Add(NewProcessor("signals", generateSignals, func() *State { return NewState() }))
	
	// Same pipeline works for both modes
	data := generateHistoricalData(500)
	
	// Batch mode
	fmt.Println("Running pipeline in batch mode...")
	batchResults := pipeline.Batch(data)
	analyzeBatchResults(batchResults)
	
	// Stream mode
	fmt.Println("Running pipeline in stream mode...")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	liveStream := generateLiveStream(ctx)
	streamResults := pipeline.Stream(ctx, liveStream)
	analyzeStreamResults(streamResults)
}

// ===== EXAMPLE 3: REAL-WORLD TRADING STRATEGY =====

func demonstrateTradingStrategy() {
	fmt.Println("\n💰 Example 3: Real-World Trading Strategy")
	
	// MACD Strategy Pipeline
	macdStrategy := NewPipeline("macd_crossover").
		Add(EMA(12, "close", "ema12")).
		Add(EMA(26, "close", "ema26")).
		Add(NewProcessor("macd", calculateMACD, func() *State { return NewState() })).
		Add(EMA(9, "macd", "signal")).
		Add(NewProcessor("crossover", detectCrossover, func() *State { return NewState() }))
	
	// RSI Filter Pipeline  
	rsiFilter := NewPipeline("rsi_filter").
		Add(NewProcessor("rsi", calculateRSI, func() *State { return NewState() })).
		Add(NewProcessor("rsi_filter", applyRSIFilter, func() *State { return NewState() }))
	
	// Combined Strategy
	strategy := NewPipeline("combined_strategy").
		Add(NewProcessor("parallel", func(input *tick.Tick, state *State) (*tick.Tick, *State) {
			// Run both strategies and combine signals
			macdResult := macdStrategy.Process(input)
			rsiResult := rsiFilter.Process(input)
			
			combined := input.Clone()
			
			// Copy all fields from both results
			if !macdResult.IsEmpty() {
				for _, field := range macdResult.FieldNames() {
					combined.SetField(field, macdResult.GetField(field))
				}
			}
			
			if !rsiResult.IsEmpty() {
				for _, field := range rsiResult.FieldNames() {
					combined.SetField(field, rsiResult.GetField(field))
				}
			}
			
			return combined, state
		}, func() *State { return NewState() }))
	
	// Backtest (batch mode)
	fmt.Println("Running backtest...")
	historicalData := generateHistoricalData(1000)
	backtest := strategy.Batch(historicalData)
	
	signals := 0
	for _, tick := range backtest.Ticks() {
		if tick.HasField("signal") && tick.GetField("signal") != 0 {
			signals++
		}
	}
	fmt.Printf("Backtest generated %d trading signals\n", signals)
	
	// Live trading (stream mode)
	fmt.Println("Starting live trading simulation...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	liveData := generateLiveStream(ctx)
	liveResults := strategy.Stream(ctx, liveData)
	
	liveSignals := 0
	for result := range liveResults {
		if result.HasField("signal") && result.GetField("signal") != 0 {
			liveSignals++
			fmt.Printf("⚡ SIGNAL: %s at price %.2f\n", 
				getSignalType(result.GetField("signal")), 
				result.GetField("close"))
		}
	}
	fmt.Printf("Live trading generated %d signals\n", liveSignals)
}

// ===== EXAMPLE 4: PERFORMANCE COMPARISON =====

func demonstratePerformanceComparison() {
	fmt.Println("\n⚡ Example 4: Performance Comparison")
	
	processor := SMA(20, "close", "sma20")
	data := generateHistoricalData(10000)
	
	// Measure batch processing
	start := time.Now()
	batchResults := processor.Reset().Batch(data)
	batchDuration := time.Since(start)
	
	fmt.Printf("Batch processing: %d ticks in %v (%.0f ticks/sec)\n",
		len(batchResults.Ticks()),
		batchDuration,
		float64(len(batchResults.Ticks()))/batchDuration.Seconds())
	
	// Measure stream processing  
	ctx, cancel := context.WithCancel(context.Background())
	
	start = time.Now()
	streamInput := make(chan *tick.Tick, 1000)
	streamOutput := processor.Reset().Stream(ctx, streamInput)
	
	// Feed data to stream
	go func() {
		defer close(streamInput)
		for _, t := range data.Ticks() {
			streamInput <- t
		}
	}()
	
	streamCount := 0
	for range streamOutput {
		streamCount++
	}
	
	cancel()
	streamDuration := time.Since(start)
	
	fmt.Printf("Stream processing: %d ticks in %v (%.0f ticks/sec)\n",
		streamCount,
		streamDuration,
		float64(streamCount)/streamDuration.Seconds())
}

// ===== HELPER FUNCTIONS =====

// generateSignals creates trading signals based on moving average crossover
func generateSignals(input *tick.Tick, state *State) (*tick.Tick, *State) {
	if !input.HasField("sma_fast") || !input.HasField("sma_slow") {
		return input, state
	}
	
	fast := input.GetField("sma_fast")
	slow := input.GetField("sma_slow")
	
	var signal float64
	if fast > slow {
		signal = 1.0 // Buy signal
	} else if fast < slow {
		signal = -1.0 // Sell signal
	}
	
	result := input.Clone().SetField("signal", signal)
	return result, state
}

func calculateMACD(input *tick.Tick, state *State) (*tick.Tick, *State) {
	if input.HasField("ema12") && input.HasField("ema26") {
		macd := input.GetField("ema12") - input.GetField("ema26")
		return input.Clone().SetField("macd", macd), state
	}
	return input, state
}

func detectCrossover(input *tick.Tick, state *State) (*tick.Tick, *State) {
	if !input.HasField("macd") || !input.HasField("signal") {
		return input, state
	}
	
	macd := input.GetField("macd")
	signal := input.GetField("signal")
	
	var crossover float64
	if macd > signal {
		crossover = 1.0 // Bullish crossover
	} else if macd < signal {
		crossover = -1.0 // Bearish crossover
	}
	
	result := input.Clone().SetField("crossover", crossover)
	return result, state
}

func calculateRSI(input *tick.Tick, state *State) (*tick.Tick, *State) {
	// Simplified RSI calculation
	rsi := 50.0 // Placeholder
	return input.Clone().SetField("rsi", rsi), state
}

func applyRSIFilter(input *tick.Tick, state *State) (*tick.Tick, *State) {
	if input.HasField("rsi") {
		rsi := input.GetField("rsi")
		var filter float64
		if rsi > 70 {
			filter = -1.0 // Overbought
		} else if rsi < 30 {
			filter = 1.0 // Oversold
		}
		return input.Clone().SetField("rsi_filter", filter), state
	}
	return input, state
}

func generateHistoricalData(count int) *series.Series {
	s := series.New("historical")
	basePrice := 100.0
	
	for i := 0; i < count; i++ {
		// Generate realistic price movement
		change := (rand.Float64() - 0.5) * 2.0 // -1 to +1
		basePrice += change
		
		t := tick.New(time.Now().Add(time.Duration(i) * time.Minute))
		t.SetField("close", basePrice)
		t.SetField("volume", rand.Float64()*1000)
		
		s.Add(t)
	}
	
	return s
}

func generateLiveStream(ctx context.Context) <-chan *tick.Tick {
	output := make(chan *tick.Tick, 100)
	
	go func() {
		defer close(output)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		
		basePrice := 100.0
		
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				change := (rand.Float64() - 0.5) * 2.0
				basePrice += change
				
				t := tick.New(time.Now())
				t.SetField("close", basePrice)
				t.SetField("volume", rand.Float64()*1000)
				
				select {
				case output <- t:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	
	return output
}

func analyzeBatchResults(results *series.Series) {
	signals := 0
	for _, tick := range results.Ticks() {
		if tick.HasField("signal") && tick.GetField("signal") != 0 {
			signals++
		}
	}
	fmt.Printf("Batch analysis: %d signals from %d ticks\n", signals, len(results.Ticks()))
}

func analyzeStreamResults(results <-chan *tick.Tick) {
	signals := 0
	total := 0
	for result := range results {
		total++
		if result.HasField("signal") && result.GetField("signal") != 0 {
			signals++
		}
	}
	fmt.Printf("Stream analysis: %d signals from %d ticks\n", signals, total)
}

func getSignalType(signal float64) string {
	if signal > 0 {
		return "BUY"
	} else if signal < 0 {
		return "SELL"
	}
	return "HOLD"
}

// ===== MIGRATION GUIDE FOR YOUR EXISTING CODE =====

/*
MIGRATION GUIDE:

1. KEEP your existing interfaces (Process/Compute) for backward compatibility

2. ENHANCE existing indicators:
   ```go
   // Before (your current approach)
   type sma struct {
       BatchFunc  func(*series.Series) *series.Series
       StreamFunc func(*tick.Tick) *tick.Tick
   }
   
   // After (enhanced)
   type sma struct {
       BatchFunc    func(*series.Series) *series.Series
       StreamFunc   func(*tick.Tick) *tick.Tick
       UnifiedProc  *Processor  // Add this
   }
   
   func (s *sma) ToUnified() *Processor {
       return s.UnifiedProc
   }
   ```

3. GRADUAL adoption:
   - Use unified processors for new indicators
   - Migrate existing indicators over time
   - Both approaches can coexist

4. BENEFITS you'll get:
   - ✅ Same code works for stream and batch
   - ✅ Easy pipeline composition
   - ✅ Better testing (test once, works everywhere)
   - ✅ Cleaner API
   - ✅ Better performance (optimized for each mode)
*/
