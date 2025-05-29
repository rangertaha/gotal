# Scalping Strategy

## Overview

This example demonstrates a scalping strategy that aims to profit from small price movements by executing multiple trades with quick entries and exits. The strategy focuses on high-frequency trading with tight stop losses and quick profit targets.

## Strategy Components

1. **Entry Signals**
   - Price action patterns
   - Order flow analysis
   - Volume profile
   - Micro support/resistance

2. **Exit Signals**
   - Quick take profit
   - Tight stop loss
   - Time-based exit
   - Volume exhaustion

3. **Risk Management**
   - Position sizing
   - Maximum drawdown
   - Time-based filters
   - Volatility filters

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
    Timeframe        string
    MaxTradeDuration time.Duration
    StopLoss         float64
    TakeProfit       float64
    VolumeThreshold  float64
    ATRPeriod        int
    RSIPeriod        int
    MaxTradesPerDay  int
}

// Strategy state
type State struct {
    Position     string    // "long", "short", or "none"
    EntryPrice   float64
    StopLoss     float64
    TakeProfit   float64
    EntryTime    time.Time
    TradesToday  int
    LastTrade    time.Time
}

// Create new strategy
func NewScalpingStrategy(config Config) *ScalpingStrategy {
    return &ScalpingStrategy{
        config: config,
        state:  &State{Position: "none"},
    }
}

// Main strategy implementation
func (s *ScalpingStrategy) Analyze(candles []Candle) []Signal {
    var signals []Signal

    // Calculate indicators
    atr := indicators.NewATR(s.config.ATRPeriod)
    rsi := indicators.NewRSI(s.config.RSIPeriod)
    volumeMA := indicators.NewSMA(20)

    atrValues := atr.Calculate(candles)
    rsiValues := rsi.Calculate(candles)
    volumeMAValues := volumeMA.Calculate(candles)

    // Check for scalping opportunities
    for i := 1; i < len(candles); i++ {
        // Check if we can trade
        if !s.canTrade(candles[i].Time) {
            continue
        }

        volatility := atrValues[i]
        currentVolume := candles[i].Volume
        avgVolume := volumeMAValues[i]

        // Check for long setup
        if s.isLongSetup(candles[i], rsiValues[i], currentVolume, avgVolume) {
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

        // Check for short setup
        if s.isShortSetup(candles[i], rsiValues[i], currentVolume, avgVolume) {
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

// Check if we can trade
func (s *ScalpingStrategy) canTrade(currentTime time.Time) bool {
    // Check if we've exceeded daily trade limit
    if s.state.TradesToday >= s.config.MaxTradesPerDay {
        return false
    }

    // Check if enough time has passed since last trade
    if currentTime.Sub(s.state.LastTrade) < time.Minute*5 {
        return false
    }

    return true
}

// Check for long setup
func (s *ScalpingStrategy) isLongSetup(candle Candle, rsi float64, volume float64, avgVolume float64) bool {
    // RSI oversold
    if rsi < 30 {
        return false
    }

    // Volume confirmation
    if volume < avgVolume*1.2 {
        return false
    }

    // Price action
    if candle.Close < candle.Open {
        return false
    }

    // Check for micro support
    if !s.isMicroSupport(candle) {
        return false
    }

    return true
}

// Check for short setup
func (s *ScalpingStrategy) isShortSetup(candle Candle, rsi float64, volume float64, avgVolume float64) bool {
    // RSI overbought
    if rsi > 70 {
        return false
    }

    // Volume confirmation
    if volume < avgVolume*1.2 {
        return false
    }

    // Price action
    if candle.Close > candle.Open {
        return false
    }

    // Check for micro resistance
    if !s.isMicroResistance(candle) {
        return false
    }

    return true
}

// Check for micro support
func (s *ScalpingStrategy) isMicroSupport(candle Candle) bool {
    // Implement micro support detection logic
    return true
}

// Check for micro resistance
func (s *ScalpingStrategy) isMicroResistance(candle Candle) bool {
    // Implement micro resistance detection logic
    return true
}

// Position management
func (s *ScalpingStrategy) ManagePosition(signal Signal) {
    switch signal.Type {
    case "buy":
        if s.state.Position == "none" {
            s.state.Position = "long"
            s.state.EntryPrice = signal.Price
            s.state.StopLoss = signal.StopLoss
            s.state.TakeProfit = signal.TakeProfit
            s.state.EntryTime = signal.Time
            s.state.TradesToday++
            s.state.LastTrade = signal.Time
        }
    case "sell":
        if s.state.Position == "none" {
            s.state.Position = "short"
            s.state.EntryPrice = signal.Price
            s.state.StopLoss = signal.StopLoss
            s.state.TakeProfit = signal.TakeProfit
            s.state.EntryTime = signal.Time
            s.state.TradesToday++
            s.state.LastTrade = signal.Time
        }
    }
}

// Check for exit conditions
func (s *ScalpingStrategy) CheckExit(candle Candle) bool {
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

    // Check time-based exit
    if time.Since(s.state.EntryTime) > s.config.MaxTradeDuration {
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
        Timeframe:        "1m",
        MaxTradeDuration: time.Minute * 5,
        StopLoss:         0.001,  // 0.1%
        TakeProfit:       0.002,  // 0.2%
        VolumeThreshold:  1.2,    // 120% of average volume
        ATRPeriod:        14,
        RSIPeriod:        14,
        MaxTradesPerDay:  50,
    }

    // Create strategy instance
    strategy := NewScalpingStrategy(config)

    // Load historical data
    candles := loadCandles("BTC/USD", "1m")

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

### Time and Risk
- Timeframe: 1 minute
- Max Trade Duration: 5 minutes
- Stop Loss: 0.1%
- Take Profit: 0.2%
- Max Trades Per Day: 50

### Entry Conditions
- RSI extremes (30/70)
- Volume > 120% average
- Price action confirmation
- Micro support/resistance

### Risk Management
- Quick stop loss
- Quick take profit
- Time-based exit
- Daily trade limit

## Performance Considerations

### Optimization
1. **Execution Speed**
   - Minimize latency
   - Use efficient data structures
   - Optimize calculations

2. **Memory Usage**
   - Reuse indicator instances
   - Clear old data
   - Use fixed-size arrays

### Backtesting
1. **Data Requirements**
   - High-frequency data
   - Accurate volume data
   - Order book data

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

### Time Management
```go
func checkTimeFilter(currentTime time.Time, lastTrade time.Time) bool {
    return currentTime.Sub(lastTrade) >= time.Minute*5
}
```

## Future Improvements

1. **Strategy Enhancement**
   - Add order flow analysis
   - Implement market depth analysis
   - Add more entry/exit conditions

2. **Risk Management**
   - Dynamic position sizing
   - Portfolio optimization
   - Correlation analysis

3. **Performance Analysis**
   - Detailed backtesting
   - Monte Carlo simulation
   - Risk metrics calculation 