# Scalping Strategy Backtesting

## Overview

This example demonstrates how to backtest the scalping strategy using historical data. The backtesting framework evaluates the strategy's performance across different market conditions and provides detailed performance metrics.

## Backtesting Components

1. **Data Management**
   - Historical data loading
   - Data preprocessing
   - Time alignment
   - Volume normalization

2. **Strategy Execution**
   - Signal generation
   - Position management
   - Order simulation
   - Slippage modeling

3. **Performance Analysis**
   - Return metrics
   - Risk metrics
   - Trade statistics
   - Visualization

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
type BacktestConfig struct {
    StartTime        time.Time
    EndTime          time.Time
    InitialCapital   float64
    Commission       float64
    Slippage         float64
    DataSource       string
    Timeframe        string
    MaxPositions     int
    RiskPerTrade     float64
}

// Trade record
type Trade struct {
    EntryTime     time.Time
    ExitTime      time.Time
    EntryPrice    float64
    ExitPrice     float64
    Position      string    // "long" or "short"
    Volume        float64
    PnL           float64
    Commission    float64
    Slippage      float64
    ExitReason    string
}

// Performance metrics
type PerformanceMetrics struct {
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
}

// Create new backtest
func NewScalpingBacktest(config BacktestConfig) *ScalpingBacktest {
    return &ScalpingBacktest{
        config: config,
        trades: make([]Trade, 0),
        metrics: &PerformanceMetrics{},
    }
}

// Run backtest
func (b *ScalpingBacktest) Run() error {
    // Load historical data
    candles, err := b.loadData()
    if err != nil {
        return err
    }

    // Initialize strategy
    strategy := NewScalpingStrategy(Config{
        Timeframe:        b.config.Timeframe,
        MaxTradeDuration: time.Minute * 5,
        StopLoss:         0.001,  // 0.1%
        TakeProfit:       0.002,  // 0.2%
        VolumeThreshold:  1.2,    // 120% of average volume
        ATRPeriod:        14,
        RSIPeriod:        14,
        MaxTradesPerDay:  50,
    })

    // Initialize portfolio
    portfolio := &Portfolio{
        Cash:      b.config.InitialCapital,
        Positions: make(map[string]Position),
    }

    // Run simulation
    for i := 1; i < len(candles); i++ {
        // Get current candle
        currentCandle := candles[i]

        // Check for exit conditions
        if b.hasOpenPosition() {
            if b.shouldExit(currentCandle) {
                b.closePosition(currentCandle)
            }
        }

        // Check for entry conditions
        if !b.hasOpenPosition() && b.canOpenNewPosition() {
            signals := strategy.Analyze(candles[:i+1])
            for _, signal := range signals {
                if b.isValidSignal(signal, currentCandle) {
                    b.openPosition(signal, currentCandle)
                }
            }
        }

        // Update portfolio
        b.updatePortfolio(currentCandle)
    }

    // Calculate performance metrics
    b.calculateMetrics()

    return nil
}

// Load historical data
func (b *ScalpingBacktest) loadData() ([]Candle, error) {
    // Implement data loading logic
    return nil, nil
}

// Check if we have an open position
func (b *ScalpingBacktest) hasOpenPosition() bool {
    return len(b.trades) > 0 && b.trades[len(b.trades)-1].ExitTime.IsZero()
}

// Check if we can open a new position
func (b *ScalpingBacktest) canOpenNewPosition() bool {
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
func (b *ScalpingBacktest) shouldExit(candle Candle) bool {
    if !b.hasOpenPosition() {
        return false
    }

    currentTrade := b.trades[len(b.trades)-1]

    // Check stop loss
    if currentTrade.Position == "long" && candle.Low <= currentTrade.EntryPrice*(1-b.config.StopLoss) {
        return true
    }
    if currentTrade.Position == "short" && candle.High >= currentTrade.EntryPrice*(1+b.config.StopLoss) {
        return true
    }

    // Check take profit
    if currentTrade.Position == "long" && candle.High >= currentTrade.EntryPrice*(1+b.config.TakeProfit) {
        return true
    }
    if currentTrade.Position == "short" && candle.Low <= currentTrade.EntryPrice*(1-b.config.TakeProfit) {
        return true
    }

    // Check time-based exit
    if time.Since(currentTrade.EntryTime) > time.Minute*5 {
        return true
    }

    return false
}

// Open new position
func (b *ScalpingBacktest) openPosition(signal Signal, candle Candle) {
    // Calculate position size
    positionSize := b.calculatePositionSize(signal.Price)

    // Calculate slippage
    slippage := signal.Price * b.config.Slippage

    // Calculate commission
    commission := positionSize * signal.Price * b.config.Commission

    // Create trade record
    trade := Trade{
        EntryTime:  candle.Time,
        EntryPrice: signal.Price,
        Position:   signal.Type,
        Volume:     positionSize,
        Commission: commission,
        Slippage:   slippage,
    }

    b.trades = append(b.trades, trade)
}

// Close current position
func (b *ScalpingBacktest) closePosition(candle Candle) {
    if !b.hasOpenPosition() {
        return
    }

    currentTrade := &b.trades[len(b.trades)-1]
    currentTrade.ExitTime = candle.Time
    currentTrade.ExitPrice = candle.Close

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
func (b *ScalpingBacktest) calculatePositionSize(price float64) float64 {
    riskAmount := b.config.InitialCapital * b.config.RiskPerTrade
    return riskAmount / (price * b.config.StopLoss)
}

// Update portfolio
func (b *ScalpingBacktest) updatePortfolio(candle Candle) {
    // Implement portfolio update logic
}

// Calculate performance metrics
func (b *ScalpingBacktest) calculateMetrics() {
    metrics := &PerformanceMetrics{}

    // Calculate basic metrics
    metrics.TotalTrades = len(b.trades)
    for _, trade := range b.trades {
        if trade.PnL > 0 {
            metrics.WinningTrades++
        } else {
            metrics.LosingTrades++
        }
        metrics.TotalPnL += trade.PnL
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

    // Calculate maximum drawdown
    metrics.MaxDrawdown = b.calculateMaxDrawdown()

    // Calculate Sharpe ratio
    metrics.SharpeRatio = b.calculateSharpeRatio()

    b.metrics = metrics
}

// Calculate maximum drawdown
func (b *ScalpingBacktest) calculateMaxDrawdown() float64 {
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
func (b *ScalpingBacktest) calculateSharpeRatio() float64 {
    // Implement Sharpe ratio calculation
    return 0
}
```

## Usage Example

```go
func main() {
    // Create backtest configuration
    config := BacktestConfig{
        StartTime:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
        EndTime:        time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
        InitialCapital: 10000,
        Commission:     0.001,  // 0.1%
        Slippage:       0.0005, // 0.05%
        DataSource:     "binance",
        Timeframe:      "1m",
        MaxPositions:   5,
        RiskPerTrade:   0.01,   // 1%
    }

    // Create backtest instance
    backtest := NewScalpingBacktest(config)

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

### Trade Statistics
- Total Trades
- Winning Trades
- Losing Trades
- Average Holding Time
- Maximum Consecutive Wins
- Maximum Consecutive Losses

## Visualization

### Equity Curve
```go
func plotEquityCurve(trades []Trade) {
    // Implement equity curve plotting
}
```

### Drawdown Chart
```go
func plotDrawdownChart(trades []Trade) {
    // Implement drawdown chart plotting
}
```

### Trade Distribution
```go
func plotTradeDistribution(trades []Trade) {
    // Implement trade distribution plotting
}
```

## Future Improvements

1. **Backtesting Enhancement**
   - Add more market data sources
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