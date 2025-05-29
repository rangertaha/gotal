# Multi-Timeframe Strategy

## Overview

This example demonstrates a multi-timeframe strategy that analyzes price action across different timeframes to identify high-probability trading opportunities. The strategy combines signals from multiple timeframes to confirm trends and generate trading signals.

## Strategy Components

1. **Timeframe Analysis**
   - Primary timeframe (e.g., 4h)
   - Secondary timeframe (e.g., 1h)
   - Tertiary timeframe (e.g., 15m)
   - Trend alignment check

2. **Entry Signals**
   - Trend confirmation
   - Support/resistance levels
   - Volume analysis
   - Pattern recognition

3. **Exit Signals**
   - Trend reversal
   - Multiple timeframe divergence
   - Stop loss and take profit
   - Time-based exit

## Implementation

```go
package main

import (
    "fmt"
    "time"
    "github.com/rangertaha/gotal"
    "github.com/rangertaha/gotal/indicators"
)

// Timeframe configuration
type TimeframeConfig struct {
    Period      string
    SMAFast     int
    SMASlow     int
    RSIPeriod   int
    ATRPeriod   int
    VolumeMA    int
}

// Strategy configuration
type Config struct {
    Timeframes       []TimeframeConfig
    StopLoss         float64
    TakeProfit       float64
    MaxHoldingTime   time.Duration
    MinVolume        float64
    TrendThreshold   float64
    DivergencePeriod int
}

// Strategy state
type State struct {
    Position     string    // "long", "short", or "none"
    EntryPrice   float64
    StopLoss     float64
    TakeProfit   float64
    EntryTime    time.Time
    Timeframes   map[string]TimeframeState
}

// Timeframe state
type TimeframeState struct {
    Trend       string    // "bullish", "bearish", or "neutral"
    SMAFast     []float64
    SMASlow     []float64
    RSI         []float64
    ATR         []float64
    VolumeMA    []float64
    LastUpdate  time.Time
}

// Create new strategy
func NewMultiTimeframeStrategy(config Config) *MultiTimeframeStrategy {
    return &MultiTimeframeStrategy{
        config: config,
        state:  &State{
            Position: "none",
            Timeframes: make(map[string]TimeframeState),
        },
    }
}

// Main strategy implementation
func (s *MultiTimeframeStrategy) Analyze(candles map[string][]Candle) []Signal {
    var signals []Signal

    // Update timeframe states
    s.updateTimeframeStates(candles)

    // Check for trading opportunities
    if s.state.Position == "none" {
        // Check for entry signals
        if signal := s.checkEntrySignals(); signal != nil {
            signals = append(signals, *signal)
        }
    } else {
        // Check for exit signals
        if signal := s.checkExitSignals(); signal != nil {
            signals = append(signals, *signal)
        }
    }

    return signals
}

// Update timeframe states
func (s *MultiTimeframeStrategy) updateTimeframeStates(candles map[string][]Candle) {
    for _, tf := range s.config.Timeframes {
        if candleData, ok := candles[tf.Period]; ok {
            // Calculate indicators
            smaFast := indicators.NewSMA(tf.SMAFast)
            smaSlow := indicators.NewSMA(tf.SMASlow)
            rsi := indicators.NewRSI(tf.RSIPeriod)
            atr := indicators.NewATR(tf.ATRPeriod)
            volumeMA := indicators.NewSMA(tf.VolumeMA)

            // Update state
            s.state.Timeframes[tf.Period] = TimeframeState{
                SMAFast:    smaFast.Calculate(candleData),
                SMASlow:    smaSlow.Calculate(candleData),
                RSI:        rsi.Calculate(candleData),
                ATR:        atr.Calculate(candleData),
                VolumeMA:   volumeMA.Calculate(candleData),
                LastUpdate: time.Now(),
            }

            // Update trend
            s.updateTrend(tf.Period)
        }
    }
}

// Update trend for timeframe
func (s *MultiTimeframeStrategy) updateTrend(timeframe string) {
    state := s.state.Timeframes[timeframe]
    if len(state.SMAFast) == 0 || len(state.SMASlow) == 0 {
        return
    }

    // Get latest values
    fastMA := state.SMAFast[len(state.SMAFast)-1]
    slowMA := state.SMASlow[len(state.SMASlow)-1]
    rsi := state.RSI[len(state.RSI)-1]

    // Determine trend
    if fastMA > slowMA && rsi > 50 {
        state.Trend = "bullish"
    } else if fastMA < slowMA && rsi < 50 {
        state.Trend = "bearish"
    } else {
        state.Trend = "neutral"
    }

    s.state.Timeframes[timeframe] = state
}

// Check for entry signals
func (s *MultiTimeframeStrategy) checkEntrySignals() *Signal {
    // Check trend alignment
    if !s.isTrendAligned() {
        return nil
    }

    // Get primary timeframe
    primary := s.config.Timeframes[0]
    state := s.state.Timeframes[primary.Period]

    // Check for bullish setup
    if s.isBullishSetup(primary.Period) {
        // Calculate stop loss and take profit
        atr := state.ATR[len(state.ATR)-1]
        stopLoss := state.SMAFast[len(state.SMAFast)-1] - (atr * 2)
        takeProfit := state.SMAFast[len(state.SMAFast)-1] + (atr * 3)

        return &Signal{
            Type:       "buy",
            Price:      state.SMAFast[len(state.SMAFast)-1],
            StopLoss:   stopLoss,
            TakeProfit: takeProfit,
            Time:       time.Now(),
            Strength:   0.8,
        }
    }

    // Check for bearish setup
    if s.isBearishSetup(primary.Period) {
        // Calculate stop loss and take profit
        atr := state.ATR[len(state.ATR)-1]
        stopLoss := state.SMAFast[len(state.SMAFast)-1] + (atr * 2)
        takeProfit := state.SMAFast[len(state.SMAFast)-1] - (atr * 3)

        return &Signal{
            Type:       "sell",
            Price:      state.SMAFast[len(state.SMAFast)-1],
            StopLoss:   stopLoss,
            TakeProfit: takeProfit,
            Time:       time.Now(),
            Strength:   0.8,
        }
    }

    return nil
}

// Check if trends are aligned
func (s *MultiTimeframeStrategy) isTrendAligned() bool {
    var trends []string
    for _, tf := range s.config.Timeframes {
        trends = append(trends, s.state.Timeframes[tf.Period].Trend)
    }

    // Check if all trends are the same
    for i := 1; i < len(trends); i++ {
        if trends[i] != trends[0] || trends[i] == "neutral" {
            return false
        }
    }

    return true
}

// Check for bullish setup
func (s *MultiTimeframeStrategy) isBullishSetup(timeframe string) bool {
    state := s.state.Timeframes[timeframe]
    if len(state.SMAFast) < 2 || len(state.SMASlow) < 2 {
        return false
    }

    // Check moving average crossover
    if state.SMAFast[len(state.SMAFast)-1] <= state.SMASlow[len(state.SMASlow)-1] {
        return false
    }

    // Check RSI
    if state.RSI[len(state.RSI)-1] < 50 {
        return false
    }

    // Check volume
    if state.VolumeMA[len(state.VolumeMA)-1] < s.config.MinVolume {
        return false
    }

    return true
}

// Check for bearish setup
func (s *MultiTimeframeStrategy) isBearishSetup(timeframe string) bool {
    state := s.state.Timeframes[timeframe]
    if len(state.SMAFast) < 2 || len(state.SMASlow) < 2 {
        return false
    }

    // Check moving average crossover
    if state.SMAFast[len(state.SMAFast)-1] >= state.SMASlow[len(state.SMASlow)-1] {
        return false
    }

    // Check RSI
    if state.RSI[len(state.RSI)-1] > 50 {
        return false
    }

    // Check volume
    if state.VolumeMA[len(state.VolumeMA)-1] < s.config.MinVolume {
        return false
    }

    return true
}

// Check for exit signals
func (s *MultiTimeframeStrategy) checkExitSignals() *Signal {
    // Check stop loss
    if s.isStopLossHit() {
        return &Signal{
            Type:  "close",
            Time:  time.Now(),
            Price: s.state.StopLoss,
        }
    }

    // Check take profit
    if s.isTakeProfitHit() {
        return &Signal{
            Type:  "close",
            Time:  time.Now(),
            Price: s.state.TakeProfit,
        }
    }

    // Check time-based exit
    if time.Since(s.state.EntryTime) > s.config.MaxHoldingTime {
        return &Signal{
            Type:  "close",
            Time:  time.Now(),
        }
    }

    // Check trend reversal
    if s.isTrendReversal() {
        return &Signal{
            Type:  "close",
            Time:  time.Now(),
        }
    }

    return nil
}

// Check if stop loss is hit
func (s *MultiTimeframeStrategy) isStopLossHit() bool {
    if s.state.Position == "none" {
        return false
    }

    primary := s.config.Timeframes[0]
    state := s.state.Timeframes[primary.Period]
    currentPrice := state.SMAFast[len(state.SMAFast)-1]

    if s.state.Position == "long" && currentPrice <= s.state.StopLoss {
        return true
    }
    if s.state.Position == "short" && currentPrice >= s.state.StopLoss {
        return true
    }

    return false
}

// Check if take profit is hit
func (s *MultiTimeframeStrategy) isTakeProfitHit() bool {
    if s.state.Position == "none" {
        return false
    }

    primary := s.config.Timeframes[0]
    state := s.state.Timeframes[primary.Period]
    currentPrice := state.SMAFast[len(state.SMAFast)-1]

    if s.state.Position == "long" && currentPrice >= s.state.TakeProfit {
        return true
    }
    if s.state.Position == "short" && currentPrice <= s.state.TakeProfit {
        return true
    }

    return false
}

// Check for trend reversal
func (s *MultiTimeframeStrategy) isTrendReversal() bool {
    if s.state.Position == "none" {
        return false
    }

    // Check if trends have reversed
    if s.state.Position == "long" && !s.isTrendAligned() {
        return true
    }
    if s.state.Position == "short" && !s.isTrendAligned() {
        return true
    }

    return false
}

// Position management
func (s *MultiTimeframeStrategy) ManagePosition(signal Signal) {
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
    case "close":
        s.state.Position = "none"
    }
}
```

## Usage Example

```go
func main() {
    // Create timeframe configurations
    timeframes := []TimeframeConfig{
        {
            Period:      "4h",
            SMAFast:     20,
            SMASlow:     50,
            RSIPeriod:   14,
            ATRPeriod:   14,
            VolumeMA:    20,
        },
        {
            Period:      "1h",
            SMAFast:     20,
            SMASlow:     50,
            RSIPeriod:   14,
            ATRPeriod:   14,
            VolumeMA:    20,
        },
        {
            Period:      "15m",
            SMAFast:     20,
            SMASlow:     50,
            RSIPeriod:   14,
            ATRPeriod:   14,
            VolumeMA:    20,
        },
    }

    // Create strategy configuration
    config := Config{
        Timeframes:       timeframes,
        StopLoss:         0.02,  // 2%
        TakeProfit:       0.04,  // 4%
        MaxHoldingTime:   time.Hour * 24,
        MinVolume:        1000,
        TrendThreshold:   0.02,  // 2%
        DivergencePeriod: 14,
    }

    // Create strategy instance
    strategy := NewMultiTimeframeStrategy(config)

    // Load historical data for each timeframe
    candles := make(map[string][]Candle)
    candles["4h"] = loadCandles("BTC/USD", "4h")
    candles["1h"] = loadCandles("BTC/USD", "1h")
    candles["15m"] = loadCandles("BTC/USD", "15m")

    // Run strategy
    signals := strategy.Analyze(candles)

    // Process signals
    for _, signal := range signals {
        strategy.ManagePosition(signal)
    }
}
```

## Strategy Parameters

### Timeframe Configuration
- Primary: 4h
- Secondary: 1h
- Tertiary: 15m

### Indicator Parameters
- SMA Fast: 20
- SMA Slow: 50
- RSI Period: 14
- ATR Period: 14
- Volume MA: 20

### Risk Management
- Stop Loss: 2%
- Take Profit: 4%
- Max Holding Time: 24 hours
- Min Volume: 1000

### Entry Conditions
- Trend alignment across timeframes
- Moving average crossover
- RSI confirmation
- Volume confirmation

### Exit Conditions
- Stop loss hit
- Take profit hit
- Time-based exit
- Trend reversal

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
   - Multiple timeframe data
   - Sufficient historical data
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

### Trend Analysis
```go
func analyzeTrend(timeframes []TimeframeState) string {
    var trends []string
    for _, tf := range timeframes {
        trends = append(trends, tf.Trend)
    }

    // Check if all trends are the same
    for i := 1; i < len(trends); i++ {
        if trends[i] != trends[0] {
            return "neutral"
        }
    }

    return trends[0]
}
```

## Future Improvements

1. **Strategy Enhancement**
   - Add more timeframes
   - Implement pattern recognition
   - Add market regime detection

2. **Risk Management**
   - Dynamic position sizing
   - Portfolio optimization
   - Correlation analysis

3. **Performance Analysis**
   - Detailed backtesting
   - Monte Carlo simulation
   - Risk metrics calculation 