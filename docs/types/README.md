# Data Types

This directory contains documentation for the core data types used throughout the Gotal library.

## Overview

The types package provides essential data structures and interfaces for technical analysis, including price data, indicators, and pattern recognition.

## Core Types

### Market Data
- [Candle](market/candle.md) - OHLCV data structure
- [TimeSeries](market/timeseries.md) - Time-ordered data collection
- [MarketData](market/market_data.md) - Complete market information

### Analysis Types
- [Pattern](analysis/pattern.md) - Chart pattern structure
- [Indicator](analysis/indicator.md) - Technical indicator interface
- [Signal](analysis/signal.md) - Trading signal structure

### Utility Types
- [TimeFrame](utility/timeframe.md) - Time interval definitions
- [Price](utility/price.md) - Price-related utilities
- [Volume](utility/volume.md) - Volume-related utilities

## Data Structures

### Candle
```go
type Candle struct {
    Open   float64
    High   float64
    Low    float64
    Close  float64
    Volume float64
    Time   time.Time
}
```

### TimeSeries
```go
type TimeSeries struct {
    Data     []Candle
    TimeFrame TimeFrame
    Symbol    string
}
```

### Pattern
```go
type Pattern struct {
    Type        string
    Confidence  float64
    StartTime   time.Time
    EndTime     time.Time
    Description string
}
```

### Signal
```go
type Signal struct {
    Type      string    // "buy", "sell", "neutral"
    Strength  float64   // 0.0 to 1.0
    Time      time.Time
    Price     float64
    StopLoss  float64
    TakeProfit float64
}
```

## Usage Examples

### Creating a Candle
```go
candle := types.Candle{
    Open:   100.0,
    High:   105.0,
    Low:    98.0,
    Close:  103.0,
    Volume: 1000.0,
    Time:   time.Now(),
}
```

### Working with TimeSeries
```go
series := types.NewTimeSeries("BTC/USD", types.TimeFrame1H)
series.AddCandle(candle)
```

### Creating a Pattern
```go
pattern := types.Pattern{
    Type:        "Head and Shoulders",
    Confidence:  0.9,
    StartTime:   startTime,
    EndTime:     endTime,
    Description: "Bearish reversal pattern",
}
```

## Best Practices

1. **Data Validation**
   - Validate input data
   - Check for missing values
   - Ensure time sequence

2. **Memory Management**
   - Use appropriate data structures
   - Consider memory usage
   - Implement cleanup

3. **Type Safety**
   - Use strong typing
   - Implement interfaces
   - Handle errors properly

## Implementation Details

### Error Handling
```go
type Error struct {
    Code    int
    Message string
    Time    time.Time
}
```

### Interfaces
```go
type Analyzer interface {
    Analyze(data []Candle) ([]Pattern, error)
}

type Indicator interface {
    Calculate(data []Candle) ([]float64, error)
}
```

## Future Improvements

1. **Data Structures**
   - Custom data types
   - Optimized storage
   - Better memory management

2. **Type System**
   - Generic types
   - Better type safety
   - More interfaces

3. **Performance**
   - Faster operations
   - Less memory usage
   - Better concurrency 