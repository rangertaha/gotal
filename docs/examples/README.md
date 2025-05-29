# Usage Examples

This directory contains practical examples of how to use the Gotal library for technical analysis.

## Overview

The examples demonstrate various use cases and features of the library, from basic usage to advanced trading strategies.

## Example Categories

### Basic Usage
- [Getting Started](basic/getting_started.md)
- [Data Loading](basic/data_loading.md)
- [Simple Analysis](basic/simple_analysis.md)

### Pattern Recognition
- [Head and Shoulders](patterns/head_and_shoulders.md)
- [Double Top/Bottom](patterns/double_top_bottom.md)
- [Triangle Patterns](patterns/triangles.md)

### Technical Indicators
- [Moving Averages](indicators/moving_averages.md)
- [RSI Strategy](indicators/rsi_strategy.md)
- [MACD Crossover](indicators/macd_crossover.md)

### Trading Strategies
- [Trend Following](strategies/trend_following.md)
- [Mean Reversion](strategies/mean_reversion.md)
- [Breakout Trading](strategies/breakout_trading.md)

## Code Examples

### Basic Analysis
```go
package main

import (
    "fmt"
    "github.com/rangertaha/gotal"
)

func main() {
    // Create analysis instance
    analysis := gotal.NewAnalysis()

    // Load data
    candles := loadCandles("BTC/USD", "1h")
    analysis.AddCandles(candles)

    // Find patterns
    patterns := analysis.FindPatterns()
    for _, p := range patterns {
        fmt.Printf("Found %s pattern (confidence: %.2f)\n", 
            p.Type, p.Confidence)
    }

    // Calculate indicators
    rsi := analysis.CalculateRSI(14)
    macd := analysis.CalculateMACD(12, 26, 9)

    // Generate signals
    signals := analysis.GenerateSignals()
    for _, s := range signals {
        fmt.Printf("Signal: %s at %.2f (strength: %.2f)\n",
            s.Type, s.Price, s.Strength)
    }
}
```

### Pattern Recognition
```go
package main

import (
    "fmt"
    "github.com/rangertaha/gotal/patterns"
)

func main() {
    // Load candlestick data
    candles := loadCandles("BTC/USD", "1h")

    // Check for Head and Shoulders pattern
    if patterns.IsHeadAndShoulders(candles) {
        fmt.Println("Head and Shoulders pattern detected")
    }

    // Check for Double Top
    if patterns.IsDoubleTop(candles) {
        fmt.Println("Double Top pattern detected")
    }

    // Analyze all patterns
    results := patterns.AnalyzeChartPatterns(candles)
    for _, r := range results {
        fmt.Printf("Pattern: %s (Confidence: %.2f)\n",
            r.Pattern, r.Confidence)
    }
}
```

### Technical Indicators
```go
package main

import (
    "fmt"
    "github.com/rangertaha/gotal/indicators"
)

func main() {
    // Create indicator configurations
    rsiConfig := indicators.Config{
        Period:    14,
        Threshold: 0.7,
    }

    macdConfig := indicators.Config{
        FastPeriod:   12,
        SlowPeriod:   26,
        SignalPeriod: 9,
    }

    // Calculate indicators
    rsi := indicators.NewRSI(rsiConfig)
    macd := indicators.NewMACD(macdConfig)

    // Process results
    rsiResults := rsi.Calculate(candles)
    macdResults := macd.Calculate(candles)

    // Generate trading signals
    for i := range rsiResults {
        if rsiResults[i].Signal == "buy" && 
           macdResults[i].Signal == "buy" {
            fmt.Printf("Strong buy signal at %v\n", 
                rsiResults[i].Time)
        }
    }
}
```

## Best Practices

1. **Error Handling**
   - Always check for errors
   - Implement proper logging
   - Handle edge cases

2. **Performance**
   - Use appropriate timeframes
   - Optimize calculations
   - Consider memory usage

3. **Trading Logic**
   - Implement proper risk management
   - Use multiple confirmations
   - Backtest strategies

## Running Examples

1. Clone the repository
```bash
git clone https://github.com/rangertaha/gotal.git
cd gotal
```

2. Install dependencies
```bash
go mod download
```

3. Run examples
```bash
go run examples/basic/getting_started.go
```

## Contributing Examples

1. Follow the code style
2. Include comments
3. Add error handling
4. Provide test data
5. Update documentation

## Future Examples

1. **Advanced Strategies**
   - Machine learning integration
   - Portfolio optimization
   - Risk management

2. **Real-time Analysis**
   - WebSocket integration
   - Real-time signals
   - Live trading

3. **Backtesting**
   - Strategy backtesting
   - Performance analysis
   - Optimization 