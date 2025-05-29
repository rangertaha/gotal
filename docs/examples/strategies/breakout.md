# Breakout Strategy

## Overview

This example demonstrates a breakout strategy that identifies and trades price breakouts from consolidation patterns. The strategy aims to capture strong directional moves after periods of range-bound trading.

## Strategy Components

1. **Pattern Identification**
   - Support and Resistance levels
   - Volume analysis
   - ATR for volatility measurement

2. **Entry Signals**
   - Price breakout with volume confirmation
   - False breakout protection
   - Trend confirmation

3. **Exit Signals**
   - Trailing stop loss
   - Take profit targets
   - Breakout failure

## Implementation

```go
package main

import (
    "fmt"
    "time"
    "github.com/rangertaha/gotal"
    "github.com/rangertaha/gotal/indicators"
)

// Strategy configuration
type Config struct {
    ConsolidationPeriod int
    ATRPeriod          int
    VolumeMA           int
    BreakoutThreshold  float64
    StopLoss           float64
    TakeProfit         float64
    TrailingStop       float64
}

// Strategy state
type State struct {
    Position     string    // "long", "short", or "none"
    EntryPrice   float64
    StopLoss     float64
    TakeProfit   float64
    EntryTime    time.Time
    HighestPrice float64
    LowestPrice  float64
}

// Create new strategy
func NewBreakoutStrategy(config Config) *BreakoutStrategy {
    return &BreakoutStrategy{
        config: config,
        state:  &State{Position: "none"},
    }
}

// Main strategy implementation
func (s *BreakoutStrategy) Analyze(candles []Candle) []Signal {
    var signals []Signal

    // Calculate indicators
    atr := indicators.NewATR(s.config.ATRPeriod)
    volumeMA := indicators.NewSMA(s.config.VolumeMA)

    atrValues := atr.Calculate(candles)
    volumeMAValues := volumeMA.Calculate(candles)

    // Find consolidation period
    for i := s.config.ConsolidationPeriod; i < len(candles); i++ {
        // Calculate consolidation range
        high := candles[i-s.config.ConsolidationPeriod:i].High()
        low := candles[i-s.config.ConsolidationPeriod:i].Low()
        range := high - low

        // Check for breakout
        if candles[i].Close > high && 
           candles[i].Volume > volumeMAValues[i] * 1.5 {
            // Calculate stop loss and take profit
            stopLoss := candles[i].Low * (1 - s.config.StopLoss)
            takeProfit := candles[i].Close * (1 + s.config.TakeProfit)

            signals = append(signals, Signal{
                Type:       "buy",
                Price:      candles[i].Close,
                StopLoss:   stopLoss,
                TakeProfit: takeProfit,
                Time:       candles[i].Time,
                Strength:   0.8,
            })
        }

        if candles[i].Close < low && 
           candles[i].Volume > volumeMAValues[i] * 1.5 {
            // Calculate stop loss and take profit
            stopLoss := candles[i].High * (1 + s.config.StopLoss)
            takeProfit := candles[i].Close * (1 - s.config.TakeProfit)

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

    return signals
}

// Position management
func (s *BreakoutStrategy) ManagePosition(signal Signal) {
    switch signal.Type {
    case "buy":
        if s.state.Position == "none" {
            s.state.Position = "long"
            s.state.EntryPrice = signal.Price
            s.state.StopLoss = signal.StopLoss
            s.state.TakeProfit = signal.TakeProfit
            s.state.EntryTime = signal.Time
            s.state.HighestPrice = signal.Price
        }
    case "sell":
        if s.state.Position == "none" {
            s.state.Position = "short"
            s.state.EntryPrice = signal.Price
            s.state.StopLoss = signal.StopLoss
            s.state.TakeProfit = signal.TakeProfit
            s.state.EntryTime = signal.Time
            s.state.LowestPrice = signal.Price
        }
    }
}

// Check for exit conditions
func (s *BreakoutStrategy) CheckExit(candle Candle, atr float64) bool {
    if s.state.Position == "none" {
        return false
    }

    // Update trailing stop for long position
    if s.state.Position == "long" {
        if candle.High > s.state.HighestPrice {
            s.state.HighestPrice = candle.High
            s.state.StopLoss = s.state.HighestPrice * (1 - s.config.TrailingStop)
        }
        if candle.Low <= s.state.StopLoss {
            return true
        }
    }

    // Update trailing stop for short position
    if s.state.Position == "short" {
        if candle.Low < s.state.LowestPrice {
            s.state.LowestPrice = candle.Low
            s.state.StopLoss = s.state.LowestPrice * (1 + s.config.TrailingStop)
        }
        if candle.High >= s.state.StopLoss {
            return true
        }
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
        ConsolidationPeriod: 20,
        ATRPeriod:          14,
        VolumeMA:           20,
        BreakoutThreshold:  0.02,  // 2%
        StopLoss:           0.02,  // 2%
        TakeProfit:         0.04,  // 4%
        TrailingStop:       0.02,  // 2%
    }

    // Create strategy instance
    strategy := NewBreakoutStrategy(config)

    // Load historical data
    candles := loadCandles("BTC/USD", "1h")

    // Run strategy
    signals := strategy.Analyze(candles)

    // Process signals
    for _, signal := range signals {
        strategy.ManagePosition(signal)
        
        // Check for exit conditions
        if strategy.CheckExit(signal, atrValues[i]) {
            fmt.Printf("Exit signal at %v\n", signal.Time)
        }
    }
}
```

## Strategy Parameters

### Pattern Identification
- Consolidation Period: 20
- ATR Period: 14
- Volume MA: 20

### Breakout Conditions
- Volume Threshold: 1.5x average
- Breakout Threshold: 2%
- False Breakout Protection: ATR-based

### Risk Management
- Initial Stop Loss: 2%
- Take Profit: 4%
- Trailing Stop: 2%
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

### Trailing Stop Management
```go
func updateTrailingStop(position string, currentPrice float64, atr float64) float64 {
    if position == "long" {
        return currentPrice - (atr * 2)
    }
    return currentPrice + (atr * 2)
}
```

## Future Improvements

1. **Strategy Enhancement**
   - Add more pattern recognition
   - Implement multiple timeframe analysis
   - Add position scaling

2. **Risk Management**
   - Dynamic position sizing
   - Portfolio optimization
   - Correlation analysis

3. **Performance Analysis**
   - Detailed backtesting
   - Monte Carlo simulation
   - Risk metrics calculation 