# Trend Following Strategy

## Overview

This example demonstrates a trend following strategy using multiple technical indicators and pattern recognition. The strategy aims to identify and trade with the prevailing market trend.

## Strategy Components

1. **Trend Identification**
   - Moving Averages (SMA, EMA)
   - ADX for trend strength
   - Price action patterns

2. **Entry Signals**
   - Moving Average crossovers
   - Pattern breakouts
   - Volume confirmation

3. **Exit Signals**
   - Trend reversal patterns
   - Moving Average crossovers
   - Stop loss and take profit

## Implementation

```go
package main

import (
    "fmt"
    "time"
    "github.com/rangertaha/gotal"
    "github.com/rangertaha/gotal/indicators"
    "github.com/rangertaha/gotal/patterns"
)

// Strategy configuration
type Config struct {
    FastPeriod    int
    SlowPeriod    int
    ADXPeriod     int
    ADXThreshold  float64
    StopLoss      float64
    TakeProfit    float64
}

// Strategy state
type State struct {
    Position     string    // "long", "short", or "none"
    EntryPrice   float64
    StopLoss     float64
    TakeProfit   float64
    EntryTime    time.Time
}

// Create new strategy
func NewTrendFollowingStrategy(config Config) *TrendFollowingStrategy {
    return &TrendFollowingStrategy{
        config: config,
        state:  &State{Position: "none"},
    }
}

// Main strategy implementation
func (s *TrendFollowingStrategy) Analyze(candles []Candle) []Signal {
    var signals []Signal

    // Calculate indicators
    fastMA := indicators.NewEMA(s.config.FastPeriod)
    slowMA := indicators.NewEMA(s.config.SlowPeriod)
    adx := indicators.NewADX(s.config.ADXPeriod)

    fastValues := fastMA.Calculate(candles)
    slowValues := slowMA.Calculate(candles)
    adxValues := adx.Calculate(candles)

    // Check for trend strength
    for i := 1; i < len(candles); i++ {
        // Strong trend condition
        if adxValues[i] > s.config.ADXThreshold {
            // Bullish trend
            if fastValues[i] > slowValues[i] && fastValues[i-1] <= slowValues[i-1] {
                // Check for bullish pattern
                if patterns.IsBullishEngulfing(candles[i-1:i+1]) {
                    // Calculate stop loss and take profit
                    stopLoss := candles[i].Low * (1 - s.config.StopLoss)
                    takeProfit := candles[i].High * (1 + s.config.TakeProfit)

                    signals = append(signals, Signal{
                        Type:       "buy",
                        Price:      candles[i].Close,
                        StopLoss:   stopLoss,
                        TakeProfit: takeProfit,
                        Time:       candles[i].Time,
                        Strength:   0.8,
                    })
                }
            }
            // Bearish trend
            if fastValues[i] < slowValues[i] && fastValues[i-1] >= slowValues[i-1] {
                // Check for bearish pattern
                if patterns.IsBearishEngulfing(candles[i-1:i+1]) {
                    // Calculate stop loss and take profit
                    stopLoss := candles[i].High * (1 + s.config.StopLoss)
                    takeProfit := candles[i].Low * (1 - s.config.TakeProfit)

                    signals = append(signals, Signal{
                        Type:       "sell",
                        Price:      candles[i].Close,
                        StopLoss:   stopLoss,
                        TakeProfit: takeProfit,
                        Time:       candles[i].Time,
                        Strength:   0.8,
                    })
                }
            }
        }
    }

    return signals
}

// Position management
func (s *TrendFollowingStrategy) ManagePosition(signal Signal) {
    switch signal.Type {
    case "buy":
        if s.state.Position == "none" {
            s.state.Position = "long"
            s.state.EntryPrice = signal.Price
            s.state.StopLoss = signal.StopLoss
            s.state.TakeProfit = signal.TakeProfit
            s.state.EntryTime = signal.Time
        }
    case "sell":
        if s.state.Position == "none" {
            s.state.Position = "short"
            s.state.EntryPrice = signal.Price
            s.state.StopLoss = signal.StopLoss
            s.state.TakeProfit = signal.TakeProfit
            s.state.EntryTime = signal.Time
        }
    }
}

// Check for exit conditions
func (s *TrendFollowingStrategy) CheckExit(candle Candle) bool {
    if s.state.Position == "none" {
        return false
    }

    // Check stop loss
    if s.state.Position == "long" && candle.Low <= s.state.StopLoss {
        return true
    }
    if s.state.Position == "short" && candle.High >= s.state.StopLoss {
        return true
    }

    // Check take profit
    if s.state.Position == "long" && candle.High >= s.state.TakeProfit {
        return true
    }
    if s.state.Position == "short" && candle.Low <= s.state.TakeProfit {
        return true
    }

    return false
}
```

## Usage Example

```go
func main() {
    // Create strategy configuration
    config := Config{
        FastPeriod:    9,
        SlowPeriod:    21,
        ADXPeriod:     14,
        ADXThreshold:  25.0,
        StopLoss:      0.02,  // 2%
        TakeProfit:    0.04,  // 4%
    }

    // Create strategy instance
    strategy := NewTrendFollowingStrategy(config)

    // Load historical data
    candles := loadCandles("BTC/USD", "1h")

    // Run strategy
    signals := strategy.Analyze(candles)

    // Process signals
    for _, signal := range signals {
        strategy.ManagePosition(signal)
        
        // Check for exit conditions
        if strategy.CheckExit(signal) {
            fmt.Printf("Exit signal at %v\n", signal.Time)
        }
    }
}
```

## Strategy Parameters

### Moving Averages
- Fast EMA: 9 periods
- Slow EMA: 21 periods
- Purpose: Identify trend direction and crossovers

### ADX
- Period: 14
- Threshold: 25
- Purpose: Confirm trend strength

### Risk Management
- Stop Loss: 2%
- Take Profit: 4%
- Risk-Reward Ratio: 1:2

## Performance Considerations

### Optimization
1. **Calculation Speed**
   - Use rolling windows
   - Pre-allocate slices
   - Minimize allocations

2. **Memory Usage**
   - Reuse indicator instances
   - Clear old data
   - Use fixed-size arrays

### Backtesting
1. **Data Requirements**
   - Sufficient historical data
   - Multiple timeframes
   - Different market conditions

2. **Performance Metrics**
   - Win rate
   - Profit factor
   - Maximum drawdown
   - Sharpe ratio

## Risk Management

### Position Sizing
```go
func calculatePositionSize(capital float64, risk float64, stopLoss float64) float64 {
    riskAmount := capital * risk
    return riskAmount / stopLoss
}
```

### Stop Loss Management
```go
func updateStopLoss(position string, currentPrice float64, atr float64) float64 {
    if position == "long" {
        return currentPrice - (atr * 2)
    }
    return currentPrice + (atr * 2)
}
```

## Future Improvements

1. **Strategy Enhancement**
   - Add more entry/exit conditions
   - Implement trailing stops
   - Add position scaling

2. **Risk Management**
   - Dynamic position sizing
   - Portfolio optimization
   - Correlation analysis

3. **Performance Analysis**
   - Detailed backtesting
   - Monte Carlo simulation
   - Risk metrics calculation 