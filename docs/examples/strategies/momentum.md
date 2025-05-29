# Momentum Strategy

## Overview

This example demonstrates a momentum strategy that trades based on price momentum and trend strength. The strategy aims to capture strong directional moves by identifying and trading with the prevailing momentum.

## Strategy Components

1. **Momentum Indicators**
   - RSI for momentum strength
   - MACD for trend direction
   - ADX for trend strength

2. **Entry Signals**
   - Momentum confirmation
   - Trend alignment
   - Volume confirmation

3. **Exit Signals**
   - Momentum reversal
   - Trend exhaustion
   - Stop loss and take profit

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
    RSIPeriod     int
    MACDFast      int
    MACDSlow      int
    MACDSignal    int
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
func NewMomentumStrategy(config Config) *MomentumStrategy {
    return &MomentumStrategy{
        config: config,
        state:  &State{Position: "none"},
    }
}

// Main strategy implementation
func (s *MomentumStrategy) Analyze(candles []Candle) []Signal {
    var signals []Signal

    // Calculate indicators
    rsi := indicators.NewRSI(s.config.RSIPeriod)
    macd := indicators.NewMACD(s.config.MACDFast, s.config.MACDSlow, s.config.MACDSignal)
    adx := indicators.NewADX(s.config.ADXPeriod)

    rsiValues := rsi.Calculate(candles)
    macdValues := macd.Calculate(candles)
    adxValues := adx.Calculate(candles)

    // Check for momentum opportunities
    for i := 1; i < len(candles); i++ {
        // Strong trend condition
        if adxValues[i] > s.config.ADXThreshold {
            // Bullish momentum
            if rsiValues[i] > 50 && 
               macdValues[i].Histogram > 0 && 
               macdValues[i].Histogram > macdValues[i-1].Histogram {
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

            // Bearish momentum
            if rsiValues[i] < 50 && 
               macdValues[i].Histogram < 0 && 
               macdValues[i].Histogram < macdValues[i-1].Histogram {
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
    }

    return signals
}

// Position management
func (s *MomentumStrategy) ManagePosition(signal Signal) {
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
func (s *MomentumStrategy) CheckExit(candle Candle, rsi float64, macd MACD) bool {
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

    // Check momentum reversal
    if s.state.Position == "long" && rsi < 50 && macd.Histogram < 0 {
        return true
    }
    if s.state.Position == "short" && rsi > 50 && macd.Histogram > 0 {
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
        RSIPeriod:     14,
        MACDFast:      12,
        MACDSlow:      26,
        MACDSignal:    9,
        ADXPeriod:     14,
        ADXThreshold:  25.0,
        StopLoss:      0.02,  // 2%
        TakeProfit:    0.04,  // 4%
    }

    // Create strategy instance
    strategy := NewMomentumStrategy(config)

    // Load historical data
    candles := loadCandles("BTC/USD", "1h")

    // Run strategy
    signals := strategy.Analyze(candles)

    // Process signals
    for _, signal := range signals {
        strategy.ManagePosition(signal)
        
        // Check for exit conditions
        if strategy.CheckExit(signal, rsiValues[i], macdValues[i]) {
            fmt.Printf("Exit signal at %v\n", signal.Time)
        }
    }
}
```

## Strategy Parameters

### Momentum Indicators
- RSI Period: 14
- MACD Fast: 12
- MACD Slow: 26
- MACD Signal: 9
- ADX Period: 14

### Entry Conditions
- RSI > 50 for longs, < 50 for shorts
- MACD histogram increasing/decreasing
- ADX > 25 for trend strength

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

### Momentum Management
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
   - Add more momentum indicators
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