# Mean Reversion Strategy

## Overview

This example demonstrates a mean reversion strategy that identifies and trades price deviations from the mean. The strategy assumes that prices tend to return to their average value over time.

## Strategy Components

1. **Mean Calculation**
   - Simple Moving Average (SMA)
   - Bollinger Bands
   - RSI for overbought/oversold conditions

2. **Entry Signals**
   - Price deviation from mean
   - Volume confirmation
   - RSI extremes

3. **Exit Signals**
   - Price returning to mean
   - Stop loss
   - Take profit

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
    SMAPeriod     int
    BBPeriod      int
    BBStdDev      float64
    RSIPeriod     int
    RSIOverbought float64
    RSIOversold   float64
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
func NewMeanReversionStrategy(config Config) *MeanReversionStrategy {
    return &MeanReversionStrategy{
        config: config,
        state:  &State{Position: "none"},
    }
}

// Main strategy implementation
func (s *MeanReversionStrategy) Analyze(candles []Candle) []Signal {
    var signals []Signal

    // Calculate indicators
    sma := indicators.NewSMA(s.config.SMAPeriod)
    bb := indicators.NewBollingerBands(s.config.BBPeriod, s.config.BBStdDev)
    rsi := indicators.NewRSI(s.config.RSIPeriod)

    smaValues := sma.Calculate(candles)
    bbValues := bb.Calculate(candles)
    rsiValues := rsi.Calculate(candles)

    // Check for mean reversion opportunities
    for i := 1; i < len(candles); i++ {
        // Oversold condition
        if rsiValues[i] < s.config.RSIOversold && 
           candles[i].Close < bbValues[i].Lower {
            // Calculate stop loss and take profit
            stopLoss := candles[i].Low * (1 - s.config.StopLoss)
            takeProfit := smaValues[i] * (1 + s.config.TakeProfit)

            signals = append(signals, Signal{
                Type:       "buy",
                Price:      candles[i].Close,
                StopLoss:   stopLoss,
                TakeProfit: takeProfit,
                Time:       candles[i].Time,
                Strength:   0.8,
            })
        }

        // Overbought condition
        if rsiValues[i] > s.config.RSIOverbought && 
           candles[i].Close > bbValues[i].Upper {
            // Calculate stop loss and take profit
            stopLoss := candles[i].High * (1 + s.config.StopLoss)
            takeProfit := smaValues[i] * (1 - s.config.TakeProfit)

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
func (s *MeanReversionStrategy) ManagePosition(signal Signal) {
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
func (s *MeanReversionStrategy) CheckExit(candle Candle, sma float64) bool {
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

    // Check mean reversion
    if s.state.Position == "long" && candle.Close >= sma {
        return true
    }
    if s.state.Position == "short" && candle.Close <= sma {
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
        SMAPeriod:     20,
        BBPeriod:      20,
        BBStdDev:      2.0,
        RSIPeriod:     14,
        RSIOverbought: 70.0,
        RSIOversold:   30.0,
        StopLoss:      0.02,  // 2%
        TakeProfit:    0.03,  // 3%
    }

    // Create strategy instance
    strategy := NewMeanReversionStrategy(config)

    // Load historical data
    candles := loadCandles("BTC/USD", "1h")

    // Run strategy
    signals := strategy.Analyze(candles)

    // Process signals
    for _, signal := range signals {
        strategy.ManagePosition(signal)
        
        // Check for exit conditions
        if strategy.CheckExit(signal, smaValues[i]) {
            fmt.Printf("Exit signal at %v\n", signal.Time)
        }
    }
}
```

## Strategy Parameters

### Moving Averages
- SMA Period: 20
- Purpose: Identify mean price level

### Bollinger Bands
- Period: 20
- Standard Deviation: 2.0
- Purpose: Identify price deviations

### RSI
- Period: 14
- Overbought: 70
- Oversold: 30
- Purpose: Confirm extreme conditions

### Risk Management
- Stop Loss: 2%
- Take Profit: 3%
- Risk-Reward Ratio: 1:1.5

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