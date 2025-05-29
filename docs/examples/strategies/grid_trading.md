# Grid Trading Strategy

## Overview

This example demonstrates a grid trading strategy that places orders at regular price intervals to profit from price oscillations. The strategy aims to generate profits from market volatility by buying low and selling high within a defined price range.

## Strategy Components

1. **Grid Setup**
   - Price range definition
   - Grid level calculation
   - Position sizing

2. **Entry Signals**
   - Price reaching grid levels
   - Volume confirmation
   - Market volatility check

3. **Exit Signals**
   - Take profit at grid levels
   - Stop loss for entire grid
   - Grid adjustment

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
    UpperPrice    float64
    LowerPrice    float64
    GridLevels    int
    PositionSize  float64
    StopLoss      float64
    TakeProfit    float64
    ATRPeriod     int
}

// Grid level structure
type GridLevel struct {
    Price       float64
    Position    string    // "long", "short", or "none"
    EntryPrice  float64
    StopLoss    float64
    TakeProfit  float64
    EntryTime   time.Time
}

// Strategy state
type State struct {
    GridLevels   []GridLevel
    TotalProfit  float64
    LastUpdate   time.Time
}

// Create new strategy
func NewGridTradingStrategy(config Config) *GridTradingStrategy {
    return &GridTradingStrategy{
        config: config,
        state:  &State{},
    }
}

// Initialize grid levels
func (s *GridTradingStrategy) InitializeGrid() {
    priceRange := s.config.UpperPrice - s.config.LowerPrice
    gridSize := priceRange / float64(s.config.GridLevels-1)

    s.state.GridLevels = make([]GridLevel, s.config.GridLevels)
    for i := 0; i < s.config.GridLevels; i++ {
        price := s.config.LowerPrice + (gridSize * float64(i))
        s.state.GridLevels[i] = GridLevel{
            Price:      price,
            Position:   "none",
            EntryPrice: 0,
            StopLoss:   0,
            TakeProfit: 0,
        }
    }
}

// Main strategy implementation
func (s *GridTradingStrategy) Analyze(candles []Candle) []Signal {
    var signals []Signal

    // Calculate ATR for volatility
    atr := indicators.NewATR(s.config.ATRPeriod)
    atrValues := atr.Calculate(candles)

    // Check each grid level
    for i := 1; i < len(candles); i++ {
        currentPrice := candles[i].Close
        volatility := atrValues[i]

        // Check for grid level triggers
        for j, level := range s.state.GridLevels {
            // Buy signal at lower grid levels
            if level.Position == "none" && 
               currentPrice <= level.Price && 
               currentPrice > level.Price - volatility {
                stopLoss := level.Price * (1 - s.config.StopLoss)
                takeProfit := level.Price * (1 + s.config.TakeProfit)

                signals = append(signals, Signal{
                    Type:       "buy",
                    Price:      currentPrice,
                    StopLoss:   stopLoss,
                    TakeProfit: takeProfit,
                    Time:       candles[i].Time,
                    Strength:   0.8,
                })

                s.state.GridLevels[j].Position = "long"
                s.state.GridLevels[j].EntryPrice = currentPrice
                s.state.GridLevels[j].StopLoss = stopLoss
                s.state.GridLevels[j].TakeProfit = takeProfit
                s.state.GridLevels[j].EntryTime = candles[i].Time
            }

            // Sell signal at upper grid levels
            if level.Position == "none" && 
               currentPrice >= level.Price && 
               currentPrice < level.Price + volatility {
                stopLoss := level.Price * (1 + s.config.StopLoss)
                takeProfit := level.Price * (1 - s.config.TakeProfit)

                signals = append(signals, Signal{
                    Type:       "sell",
                    Price:      currentPrice,
                    StopLoss:   stopLoss,
                    TakeProfit: takeProfit,
                    Time:       candles[i].Time,
                    Strength:   0.8,
                })

                s.state.GridLevels[j].Position = "short"
                s.state.GridLevels[j].EntryPrice = currentPrice
                s.state.GridLevels[j].StopLoss = stopLoss
                s.state.GridLevels[j].TakeProfit = takeProfit
                s.state.GridLevels[j].EntryTime = candles[i].Time
            }
        }
    }

    return signals
}

// Check for exit conditions
func (s *GridTradingStrategy) CheckExit(candle Candle) []Signal {
    var signals []Signal

    for i, level := range s.state.GridLevels {
        if level.Position == "none" {
            continue
        }

        // Check stop loss
        if level.Position == "long" && candle.Low <= level.StopLoss {
            signals = append(signals, Signal{
                Type:  "sell",
                Price: candle.Close,
                Time:  candle.Time,
            })
            s.state.GridLevels[i].Position = "none"
        }
        if level.Position == "short" && candle.High >= level.StopLoss {
            signals = append(signals, Signal{
                Type:  "buy",
                Price: candle.Close,
                Time:  candle.Time,
            })
            s.state.GridLevels[i].Position = "none"
        }

        // Check take profit
        if level.Position == "long" && candle.High >= level.TakeProfit {
            signals = append(signals, Signal{
                Type:  "sell",
                Price: candle.Close,
                Time:  candle.Time,
            })
            s.state.GridLevels[i].Position = "none"
        }
        if level.Position == "short" && candle.Low <= level.TakeProfit {
            signals = append(signals, Signal{
                Type:  "buy",
                Price: candle.Close,
                Time:  candle.Time,
            })
            s.state.GridLevels[i].Position = "none"
        }
    }

    return signals
}

// Adjust grid levels based on market conditions
func (s *GridTradingStrategy) AdjustGrid(candles []Candle) {
    // Calculate new price range based on recent volatility
    atr := indicators.NewATR(s.config.ATRPeriod)
    atrValues := atr.Calculate(candles)
    volatility := atrValues[len(atrValues)-1]

    // Adjust upper and lower prices
    s.config.UpperPrice = candles[len(candles)-1].High + volatility
    s.config.LowerPrice = candles[len(candles)-1].Low - volatility

    // Reinitialize grid with new levels
    s.InitializeGrid()
}
```

## Usage Example

```go
func main() {
    // Create strategy configuration
    config := Config{
        UpperPrice:    50000.0,
        LowerPrice:    40000.0,
        GridLevels:    10,
        PositionSize:  0.1,    // 10% of capital per grid level
        StopLoss:      0.02,   // 2%
        TakeProfit:    0.03,   // 3%
        ATRPeriod:     14,
    }

    // Create strategy instance
    strategy := NewGridTradingStrategy(config)

    // Initialize grid levels
    strategy.InitializeGrid()

    // Load historical data
    candles := loadCandles("BTC/USD", "1h")

    // Run strategy
    signals := strategy.Analyze(candles)

    // Process signals
    for _, signal := range signals {
        // Execute trades based on signals
        executeTrade(signal)
        
        // Check for exit conditions
        exitSignals := strategy.CheckExit(signal)
        for _, exitSignal := range exitSignals {
            executeTrade(exitSignal)
        }
    }

    // Periodically adjust grid
    strategy.AdjustGrid(candles)
}
```

## Strategy Parameters

### Grid Setup
- Price Range: Upper and Lower bounds
- Grid Levels: 10
- Position Size: 10% per level

### Entry Conditions
- Price reaching grid levels
- Volume confirmation
- Volatility check using ATR

### Risk Management
- Stop Loss: 2% per level
- Take Profit: 3% per level
- Grid adjustment based on volatility

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
func calculateGridPositionSize(capital float64, gridLevels int) float64 {
    return capital / float64(gridLevels)
}
```

### Grid Adjustment
```go
func adjustGridLevels(currentPrice float64, atr float64) (float64, float64) {
    upperPrice := currentPrice + (atr * 2)
    lowerPrice := currentPrice - (atr * 2)
    return upperPrice, lowerPrice
}
```

## Future Improvements

1. **Strategy Enhancement**
   - Dynamic grid level adjustment
   - Multiple timeframe analysis
   - Position scaling

2. **Risk Management**
   - Dynamic position sizing
   - Portfolio optimization
   - Correlation analysis

3. **Performance Analysis**
   - Detailed backtesting
   - Monte Carlo simulation
   - Risk metrics calculation 