# Multi-Timeframe Strategy Backtesting

## Overview

This example demonstrates how to backtest the multi-timeframe strategy using historical data across different timeframes. The backtesting framework evaluates the strategy's performance by analyzing price action and indicators across multiple timeframes, simulating trade execution, and calculating performance metrics.

## Backtesting Components

1. **Data Management**
   - Multi-timeframe data synchronization
   - Data preprocessing
   - Time alignment
   - Volume normalization

2. **Strategy Execution**
   - Signal generation across timeframes
   - Position management
   - Order simulation
   - Slippage modeling

3. **Performance Analysis**
   - Return metrics
   - Risk metrics
   - Trade statistics
   - Timeframe analysis

## Implementation

```go
package main

import (
    "fmt"
    "time"
    "github.com/rangertaha/gotal"
    "github.com/rangertaha/gotal/indicators"
)

// Backtest configuration
type MultiTimeframeBacktestConfig struct {
    StartTime        time.Time
    EndTime          time.Time
    InitialCapital   float64
    Timeframes       []TimeframeConfig
    MaxPositions     int
    StopLoss         float64
    TakeProfit       float64
    Commission       float64
    Slippage         float64
    RiskPerTrade     float64
}

// Timeframe configuration
type TimeframeConfig struct {
    Interval      string
    Weight        float64
    Indicators    []IndicatorConfig
}

// Indicator configuration
type IndicatorConfig struct {
    Name       string
    Parameters map[string]interface{}
}

// Trade record
type MultiTimeframeTrade struct {
    EntryTime       time.Time
    ExitTime        time.Time
    EntryPrice      float64
    ExitPrice       float64
    Position        string    // "long" or "short"
    Volume          float64
    PnL             float64
    Commission      float64
    Slippage        float64
    Timeframes      []string  // Timeframes that generated the signal
    ExitReason      string
}

// Performance metrics
type MultiTimeframeMetrics struct {
    TotalTrades       int
    WinningTrades     int
    LosingTrades      int
    WinRate          float64
    ProfitFactor     float64
    TotalPnL         float64
    MaxDrawdown      float64
    SharpeRatio      float64
    AverageTrade     float64
    AverageWin       float64
    AverageLoss      float64
    MaxConsecWins    int
    MaxConsecLosses  int
    AverageHolding   time.Duration
    TimeframeStats   map[string]TimeframeStats
}

// Timeframe statistics
type TimeframeStats struct {
    SignalsGenerated int
    SuccessfulSignals int
    AverageReturn    float64
    WinRate          float64
}

// Create new backtest
func NewMultiTimeframeBacktest(config MultiTimeframeBacktestConfig) *MultiTimeframeBacktest {
    return &MultiTimeframeBacktest{
        config: config,
        trades: make([]MultiTimeframeTrade, 0),
        metrics: &MultiTimeframeMetrics{
            TimeframeStats: make(map[string]TimeframeStats),
        },
    }
}

// Run backtest
func (b *MultiTimeframeBacktest) Run() error {
    // Load historical data
    candles, err := b.loadData()
    if err != nil {
        return err
    }

    // Initialize strategy
    strategy := NewMultiTimeframeStrategy(Config{
        Timeframes:     b.config.Timeframes,
        StopLoss:       b.config.StopLoss,
        TakeProfit:     b.config.TakeProfit,
    })

    // Initialize portfolio
    portfolio := &Portfolio{
        Cash:      b.config.InitialCapital,
        Positions: make(map[string]Position),
    }

    // Run simulation
    for i := 1; i < len(candles); i++ {
        // Get current candles
        currentCandles := candles[i]

        // Check for exit conditions
        if b.hasOpenPosition() {
            if b.shouldExit(currentCandles) {
                b.closePosition(currentCandles)
            }
        }

        // Check for entry conditions
        if !b.hasOpenPosition() && b.canOpenNewPosition() {
            signals := strategy.Analyze(currentCandles)
            for _, signal := range signals {
                if b.isValidSignal(signal) {
                    b.openPosition(signal, currentCandles)
                }
            }
        }

        // Update portfolio
        b.updatePortfolio(currentCandles)
    }

    // Calculate performance metrics
    b.calculateMetrics()

    return nil
}

// Load historical data
func (b *MultiTimeframeBacktest) loadData() (map[string][]Candle, error) {
    // Implement data loading logic
    return nil, nil
}

// Check if we have an open position
func (b *MultiTimeframeBacktest) hasOpenPosition() bool {
    return len(b.trades) > 0 && b.trades[len(b.trades)-1].ExitTime.IsZero()
}

// Check if we can open a new position
func (b *MultiTimeframeBacktest) canOpenNewPosition() bool {
    // Check maximum positions
    if len(b.trades) >= b.config.MaxPositions {
        return false
    }

    // Check time since last trade
    if len(b.trades) > 0 {
        lastTrade := b.trades[len(b.trades)-1]
        if time.Since(lastTrade.ExitTime) < time.Minute*5 {
            return false
        }
    }

    return true
}

// Check if we should exit current position
func (b *MultiTimeframeBacktest) shouldExit(candles map[string][]Candle) bool {
    if !b.hasOpenPosition() {
        return false
    }

    currentTrade := b.trades[len(b.trades)-1]
    currentPrice := candles["1m"][len(candles["1m"])-1].Close

    // Check stop loss
    if currentTrade.Position == "long" && currentPrice <= currentTrade.EntryPrice*(1-b.config.StopLoss) {
        return true
    }
    if currentTrade.Position == "short" && currentPrice >= currentTrade.EntryPrice*(1+b.config.StopLoss) {
        return true
    }

    // Check take profit
    if currentTrade.Position == "long" && currentPrice >= currentTrade.EntryPrice*(1+b.config.TakeProfit) {
        return true
    }
    if currentTrade.Position == "short" && currentPrice <= currentTrade.EntryPrice*(1-b.config.TakeProfit) {
        return true
    }

    // Check timeframe signals
    for timeframe, tfCandles := range candles {
        if b.hasReversalSignal(timeframe, tfCandles, currentTrade.Position) {
            return true
        }
    }

    return false
}

// Open new position
func (b *MultiTimeframeBacktest) openPosition(signal Signal, candles map[string][]Candle) {
    // Calculate position size
    positionSize := b.calculatePositionSize(signal.Price)

    // Calculate slippage
    slippage := signal.Price * b.config.Slippage

    // Calculate commission
    commission := positionSize * signal.Price * b.config.Commission

    // Create trade record
    trade := MultiTimeframeTrade{
        EntryTime:    candles["1m"][len(candles["1m"])-1].Time,
        EntryPrice:   signal.Price,
        Position:     signal.Type,
        Volume:       positionSize,
        Commission:   commission,
        Slippage:     slippage,
        Timeframes:   signal.Timeframes,
    }

    b.trades = append(b.trades, trade)
}

// Close current position
func (b *MultiTimeframeBacktest) closePosition(candles map[string][]Candle) {
    if !b.hasOpenPosition() {
        return
    }

    currentTrade := &b.trades[len(b.trades)-1]
    currentTrade.ExitTime = candles["1m"][len(candles["1m"])-1].Time
    currentTrade.ExitPrice = candles["1m"][len(candles["1m"])-1].Close

    // Calculate PnL
    if currentTrade.Position == "long" {
        currentTrade.PnL = (currentTrade.ExitPrice - currentTrade.EntryPrice) * currentTrade.Volume
    } else {
        currentTrade.PnL = (currentTrade.EntryPrice - currentTrade.ExitPrice) * currentTrade.Volume
    }

    // Subtract costs
    currentTrade.PnL -= currentTrade.Commission
    currentTrade.PnL -= currentTrade.Slippage
}

// Calculate position size
func (b *MultiTimeframeBacktest) calculatePositionSize(price float64) float64 {
    riskAmount := b.config.InitialCapital * b.config.RiskPerTrade
    return riskAmount / (price * b.config.StopLoss)
}

// Update portfolio
func (b *MultiTimeframeBacktest) updatePortfolio(candles map[string][]Candle) {
    // Implement portfolio update logic
}

// Calculate performance metrics
func (b *MultiTimeframeBacktest) calculateMetrics() {
    metrics := &MultiTimeframeMetrics{
        TimeframeStats: make(map[string]TimeframeStats),
    }

    // Calculate basic metrics
    metrics.TotalTrades = len(b.trades)
    for _, trade := range b.trades {
        if trade.PnL > 0 {
            metrics.WinningTrades++
        } else {
            metrics.LosingTrades++
        }
        metrics.TotalPnL += trade.PnL

        // Update timeframe statistics
        for _, timeframe := range trade.Timeframes {
            stats := metrics.TimeframeStats[timeframe]
            stats.SignalsGenerated++
            if trade.PnL > 0 {
                stats.SuccessfulSignals++
            }
            stats.AverageReturn += trade.PnL
            metrics.TimeframeStats[timeframe] = stats
        }
    }

    // Calculate win rate
    metrics.WinRate = float64(metrics.WinningTrades) / float64(metrics.TotalTrades)

    // Calculate profit factor
    var grossProfit, grossLoss float64
    for _, trade := range b.trades {
        if trade.PnL > 0 {
            grossProfit += trade.PnL
        } else {
            grossLoss -= trade.PnL
        }
    }
    metrics.ProfitFactor = grossProfit / grossLoss

    // Calculate average trade
    metrics.AverageTrade = metrics.TotalPnL / float64(metrics.TotalTrades)

    // Calculate timeframe statistics
    for timeframe, stats := range metrics.TimeframeStats {
        stats.WinRate = float64(stats.SuccessfulSignals) / float64(stats.SignalsGenerated)
        stats.AverageReturn = stats.AverageReturn / float64(stats.SignalsGenerated)
        metrics.TimeframeStats[timeframe] = stats
    }

    // Calculate maximum drawdown
    metrics.MaxDrawdown = b.calculateMaxDrawdown()

    // Calculate Sharpe ratio
    metrics.SharpeRatio = b.calculateSharpeRatio()

    b.metrics = metrics
}

// Calculate maximum drawdown
func (b *MultiTimeframeBacktest) calculateMaxDrawdown() float64 {
    var maxDrawdown float64
    var peak float64
    var currentValue float64

    for _, trade := range b.trades {
        currentValue += trade.PnL
        if currentValue > peak {
            peak = currentValue
        }
        drawdown := (peak - currentValue) / peak
        if drawdown > maxDrawdown {
            maxDrawdown = drawdown
        }
    }

    return maxDrawdown
}

// Calculate Sharpe ratio
func (b *MultiTimeframeBacktest) calculateSharpeRatio() float64 {
    // Implement Sharpe ratio calculation
    return 0
}
```

## Usage Example

```go
func main() {
    // Create backtest configuration
    config := MultiTimeframeBacktestConfig{
        StartTime:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
        EndTime:        time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
        InitialCapital: 10000,
        Timeframes: []TimeframeConfig{
            {
                Interval: "1m",
                Weight:  0.2,
                Indicators: []IndicatorConfig{
                    {
                        Name: "RSI",
                        Parameters: map[string]interface{}{
                            "period": 14,
                        },
                    },
                },
            },
            {
                Interval: "5m",
                Weight:  0.3,
                Indicators: []IndicatorConfig{
                    {
                        Name: "MACD",
                        Parameters: map[string]interface{}{
                            "fastPeriod":   12,
                            "slowPeriod":   26,
                            "signalPeriod": 9,
                        },
                    },
                },
            },
            {
                Interval: "15m",
                Weight:  0.5,
                Indicators: []IndicatorConfig{
                    {
                        Name: "EMA",
                        Parameters: map[string]interface{}{
                            "period": 20,
                        },
                    },
                },
            },
        },
        MaxPositions:   5,
        StopLoss:       0.01,   // 1%
        TakeProfit:     0.02,   // 2%
        Commission:     0.001,  // 0.1%
        Slippage:       0.0005, // 0.05%
        RiskPerTrade:   0.01,   // 1%
    }

    // Create backtest instance
    backtest := NewMultiTimeframeBacktest(config)

    // Run backtest
    err := backtest.Run()
    if err != nil {
        log.Fatal(err)
    }

    // Print results
    fmt.Printf("Total Trades: %d\n", backtest.metrics.TotalTrades)
    fmt.Printf("Win Rate: %.2f%%\n", backtest.metrics.WinRate*100)
    fmt.Printf("Profit Factor: %.2f\n", backtest.metrics.ProfitFactor)
    fmt.Printf("Total PnL: $%.2f\n", backtest.metrics.TotalPnL)
    fmt.Printf("Max Drawdown: %.2f%%\n", backtest.metrics.MaxDrawdown*100)
    fmt.Printf("Sharpe Ratio: %.2f\n", backtest.metrics.SharpeRatio)

    // Print timeframe statistics
    fmt.Println("\nTimeframe Statistics:")
    for timeframe, stats := range backtest.metrics.TimeframeStats {
        fmt.Printf("\n%s:\n", timeframe)
        fmt.Printf("  Signals Generated: %d\n", stats.SignalsGenerated)
        fmt.Printf("  Successful Signals: %d\n", stats.SuccessfulSignals)
        fmt.Printf("  Win Rate: %.2f%%\n", stats.WinRate*100)
        fmt.Printf("  Average Return: $%.2f\n", stats.AverageReturn)
    }
}
```

## Performance Metrics

### Return Metrics
- Total Return
- Annualized Return
- Win Rate
- Profit Factor
- Average Trade
- Average Win
- Average Loss

### Risk Metrics
- Maximum Drawdown
- Sharpe Ratio
- Sortino Ratio
- Calmar Ratio
- Value at Risk (VaR)

### Timeframe Metrics
- Signal Generation Rate
- Signal Success Rate
- Average Return per Signal
- Timeframe Contribution
- Signal Correlation

## Visualization

### Equity Curve
```go
func plotEquityCurve(trades []MultiTimeframeTrade) {
    // Implement equity curve plotting
}
```

### Timeframe Analysis
```go
func plotTimeframeAnalysis(trades []MultiTimeframeTrade) {
    // Implement timeframe analysis plotting
}
```

### Signal Distribution
```go
func plotSignalDistribution(trades []MultiTimeframeTrade) {
    // Implement signal distribution plotting
}
```

## Future Improvements

1. **Backtesting Enhancement**
   - Add more timeframes
   - Implement realistic order execution
   - Add market impact modeling
   - Support multiple assets

2. **Performance Analysis**
   - Add more performance metrics
   - Implement Monte Carlo simulation
   - Add stress testing
   - Support custom metrics

3. **Visualization**
   - Add interactive charts
   - Implement real-time plotting
   - Add trade replay
   - Support custom visualizations 