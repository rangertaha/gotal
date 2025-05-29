# Moving Averages

## Overview

Moving Averages are trend-following indicators that smooth out price data by creating a constantly updated average price. They help identify the direction of the trend and potential support/resistance levels.

## Types of Moving Averages

### Simple Moving Average (SMA)
```go
type SMA struct {
    Period int
    Data   []float64
}

func NewSMA(period int) *SMA
func (s *SMA) Calculate(prices []float64) []float64
```

**Formula:**
```
SMA = (P1 + P2 + ... + Pn) / n
where:
- P = Price
- n = Period
```

### Exponential Moving Average (EMA)
```go
type EMA struct {
    Period int
    Data   []float64
}

func NewEMA(period int) *EMA
func (e *EMA) Calculate(prices []float64) []float64
```

**Formula:**
```
EMA = (Price - EMA(previous)) × Multiplier + EMA(previous)
where:
- Multiplier = 2 / (Period + 1)
```

### Weighted Moving Average (WMA)
```go
type WMA struct {
    Period int
    Data   []float64
}

func NewWMA(period int) *WMA
func (w *WMA) Calculate(prices []float64) []float64
```

**Formula:**
```
WMA = (P1 × n + P2 × (n-1) + ... + Pn × 1) / (n + (n-1) + ... + 1)
where:
- P = Price
- n = Period
```

## Usage Examples

### Basic SMA Calculation
```go
package main

import (
    "fmt"
    "github.com/rangertaha/gotal/indicators"
)

func main() {
    // Create SMA with 20-period
    sma := indicators.NewSMA(20)
    
    // Calculate SMA
    prices := []float64{100, 101, 102, 103, 104, 105}
    result := sma.Calculate(prices)
    
    fmt.Printf("SMA values: %v\n", result)
}
```

### Multiple Moving Averages
```go
package main

import (
    "fmt"
    "github.com/rangertaha/gotal/indicators"
)

func main() {
    // Create different moving averages
    sma20 := indicators.NewSMA(20)
    ema50 := indicators.NewEMA(50)
    wma200 := indicators.NewWMA(200)
    
    // Calculate all moving averages
    prices := loadPrices("BTC/USD", "1h")
    
    smaValues := sma20.Calculate(prices)
    emaValues := ema50.Calculate(prices)
    wmaValues := wma200.Calculate(prices)
    
    // Generate signals
    for i := range prices {
        if smaValues[i] > emaValues[i] && emaValues[i] > wmaValues[i] {
            fmt.Printf("Bullish trend at %v\n", prices[i])
        }
    }
}
```

### Moving Average Crossover Strategy
```go
package main

import (
    "fmt"
    "github.com/rangertaha/gotal/indicators"
)

func main() {
    // Create fast and slow moving averages
    fastMA := indicators.NewEMA(9)
    slowMA := indicators.NewEMA(21)
    
    // Calculate moving averages
    prices := loadPrices("BTC/USD", "1h")
    fastValues := fastMA.Calculate(prices)
    slowValues := slowMA.Calculate(prices)
    
    // Generate signals
    for i := 1; i < len(prices); i++ {
        // Bullish crossover
        if fastValues[i-1] <= slowValues[i-1] && fastValues[i] > slowValues[i] {
            fmt.Printf("Bullish crossover at %v\n", prices[i])
        }
        // Bearish crossover
        if fastValues[i-1] >= slowValues[i-1] && fastValues[i] < slowValues[i] {
            fmt.Printf("Bearish crossover at %v\n", prices[i])
        }
    }
}
```

## Trading Strategies

### 1. Moving Average Crossover
- Buy when fast MA crosses above slow MA
- Sell when fast MA crosses below slow MA
- Common pairs: 9/21 EMA, 20/50 SMA, 50/200 SMA

### 2. Moving Average Ribbon
- Multiple moving averages of different periods
- Bullish when shorter MAs are above longer MAs
- Bearish when shorter MAs are below longer MAs

### 3. Moving Average Support/Resistance
- Use moving averages as dynamic support/resistance
- Buy near support, sell near resistance
- Consider volume confirmation

## Best Practices

### Parameter Selection
1. **Timeframe Considerations**
   - Shorter periods (5-20) for short-term trading
   - Medium periods (20-50) for swing trading
   - Longer periods (50-200) for long-term trends

2. **Market Conditions**
   - Trending markets: Longer periods
   - Ranging markets: Shorter periods
   - Volatile markets: Consider WMA

### Signal Confirmation
1. **Price Action**
   - Candlestick patterns
   - Support/resistance levels
   - Volume confirmation

2. **Multiple Timeframes**
   - Use higher timeframe for trend direction
   - Use lower timeframe for entry/exit
   - Look for confluence

### Risk Management
1. **Stop Loss**
   - Place below/above moving average
   - Consider ATR for volatility
   - Use trailing stops

2. **Position Sizing**
   - Consider trend strength
   - Account for volatility
   - Use proper risk-reward ratio

## Limitations

1. **Lagging Nature**
   - Moving averages are lagging indicators
   - May miss quick market moves
   - Consider leading indicators

2. **False Signals**
   - Can give false signals in ranging markets
   - May whipsaw in volatile conditions
   - Use additional confirmation

3. **Parameter Sensitivity**
   - Results vary with different periods
   - Need to test different combinations
   - Consider market conditions

## Performance Optimization

### Memory Usage
```go
// Pre-allocate slice for results
result := make([]float64, len(prices))

// Use fixed-size arrays for small periods
type SMA struct {
    Period int
    Data   [20]float64  // For period <= 20
}
```

### Calculation Speed
```go
// Use rolling window for SMA
func (s *SMA) Calculate(prices []float64) []float64 {
    result := make([]float64, len(prices))
    sum := 0.0
    
    // Initial sum
    for i := 0; i < s.Period; i++ {
        sum += prices[i]
    }
    result[s.Period-1] = sum / float64(s.Period)
    
    // Rolling window
    for i := s.Period; i < len(prices); i++ {
        sum = sum - prices[i-s.Period] + prices[i]
        result[i] = sum / float64(s.Period)
    }
    
    return result
}
```

## Future Improvements

1. **Adaptive Moving Averages**
   - Dynamic period adjustment
   - Volatility-based weighting
   - Machine learning optimization

2. **Advanced Features**
   - Volume-weighted calculations
   - Multiple timeframe analysis
   - Pattern recognition

3. **Integration**
   - Real-time data processing
   - Backtesting framework
   - Performance analytics 