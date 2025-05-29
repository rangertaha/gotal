# Arbitrage Strategy Backtesting

## Overview

This example demonstrates how to backtest the arbitrage strategy using historical order book data from multiple exchanges. The backtesting framework evaluates the strategy's performance by simulating cross-exchange arbitrage opportunities and accounting for transaction costs and execution delays.

## Backtesting Components

1. **Data Management**
   - Multi-exchange order book data
   - Price synchronization
   - Volume normalization
   - Transaction cost modeling

2. **Strategy Execution**
   - Opportunity detection
   - Order execution simulation
   - Slippage modeling
   - Cross-exchange latency

3. **Performance Analysis**
   - Profit metrics
   - Risk metrics
   - Execution statistics
   - Cost analysis

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
type ArbitrageBacktestConfig struct {
    StartTime        time.Time
    EndTime          time.Time
    InitialCapital   float64
    Exchanges        []ExchangeConfig
    MaxPositions     int
    MinSpread        float64
    MaxHoldingTime   time.Duration
    Latency          map[string]time.Duration  // Exchange latency
    MaxTradeSize     float64
    MinTradeSize     float64
}

// Exchange configuration
type ExchangeConfig struct {
    Name            string
    TradingFee      float64
    WithdrawalFee   float64
    MinTradeSize    float64
    MaxTradeSize    float64
    OrderBookDepth  int
}

// Order book snapshot
type OrderBookSnapshot struct {
    Exchange    string
    Timestamp   time.Time
    Bids        []Order
    Asks        []Order
    Spread      float64
}

// Order structure
type Order struct {
    Price     float64
    Volume    float64
}

// Trade record
type ArbitrageTrade struct {
    EntryTime       time.Time
    ExitTime        time.Time
    BuyExchange     string
    SellExchange    string
    BuyPrice        float64
    SellPrice       float64
    Volume          float64
    PnL             float64
    Fees            float64
    Latency         time.Duration
    ExitReason      string
}

// Performance metrics
type ArbitrageMetrics struct {
    TotalTrades       int
    SuccessfulTrades  int
    FailedTrades      int
    TotalPnL          float64
    TotalFees         float64
    AverageLatency    time.Duration
    MaxDrawdown       float64
    SharpeRatio       float64
    AverageTrade      float64
    BestTrade         float64
    WorstTrade        float64
    TotalVolume       float64
    AverageSpread     float64
}

// Create new backtest
func NewArbitrageBacktest(config ArbitrageBacktestConfig) *ArbitrageBacktest {
    return &ArbitrageBacktest{
        config: config,
        trades: make([]ArbitrageTrade, 0),
        metrics: &ArbitrageMetrics{},
    }
}

// Run backtest
func (b *ArbitrageBacktest) Run() error {
    // Load historical data
    orderBooks, err := b.loadData()
    if err != nil {
        return err
    }

    // Initialize strategy
    strategy := NewArbitrageStrategy(Config{
        Exchanges:        b.config.Exchanges,
        MinSpread:        b.config.MinSpread,
        MaxHoldingTime:   b.config.MaxHoldingTime,
        MaxTradeSize:     b.config.MaxTradeSize,
        MinTradeSize:     b.config.MinTradeSize,
    })

    // Initialize portfolio
    portfolio := &Portfolio{
        Cash:      b.config.InitialCapital,
        Positions: make(map[string]Position),
    }

    // Run simulation
    for i := 1; i < len(orderBooks); i++ {
        // Get current order books
        currentBooks := orderBooks[i]

        // Check for exit conditions
        if b.hasOpenPosition() {
            if b.shouldExit(currentBooks) {
                b.closePosition(currentBooks)
            }
        }

        // Check for entry conditions
        if !b.hasOpenPosition() && b.canOpenNewPosition() {
            opportunities := strategy.FindOpportunities(currentBooks)
            for _, opp := range opportunities {
                if b.isValidOpportunity(opp) {
                    b.openPosition(opp, currentBooks)
                }
            }
        }

        // Update portfolio
        b.updatePortfolio(currentBooks)
    }

    // Calculate performance metrics
    b.calculateMetrics()

    return nil
}

// Load historical data
func (b *ArbitrageBacktest) loadData() ([]OrderBookSnapshot, error) {
    // Implement data loading logic
    return nil, nil
}

// Check if we have an open position
func (b *ArbitrageBacktest) hasOpenPosition() bool {
    return len(b.trades) > 0 && b.trades[len(b.trades)-1].ExitTime.IsZero()
}

// Check if we can open a new position
func (b *ArbitrageBacktest) canOpenNewPosition() bool {
    // Check maximum positions
    if len(b.trades) >= b.config.MaxPositions {
        return false
    }

    // Check time since last trade
    if len(b.trades) > 0 {
        lastTrade := b.trades[len(b.trades)-1]
        if time.Since(lastTrade.ExitTime) < time.Second*5 {
            return false
        }
    }

    return true
}

// Check if we should exit current position
func (b *ArbitrageBacktest) shouldExit(books []OrderBookSnapshot) bool {
    if !b.hasOpenPosition() {
        return false
    }

    currentTrade := b.trades[len(b.trades)-1]

    // Check holding time
    if time.Since(currentTrade.EntryTime) > b.config.MaxHoldingTime {
        return true
    }

    // Check if spread has closed
    for _, book := range books {
        if book.Exchange == currentTrade.BuyExchange || book.Exchange == currentTrade.SellExchange {
            if book.Spread < b.config.MinSpread {
                return true
            }
        }
    }

    return false
}

// Open new position
func (b *ArbitrageBacktest) openPosition(opp Opportunity, books []OrderBookSnapshot) {
    // Calculate position size
    positionSize := b.calculatePositionSize(opp)

    // Calculate fees
    fees := b.calculateFees(opp, positionSize)

    // Calculate latency
    latency := b.calculateLatency(opp)

    // Create trade record
    trade := ArbitrageTrade{
        EntryTime:    books[0].Timestamp,
        BuyExchange:  opp.BuyExchange,
        SellExchange: opp.SellExchange,
        BuyPrice:     opp.BuyPrice,
        SellPrice:    opp.SellPrice,
        Volume:       positionSize,
        Fees:         fees,
        Latency:      latency,
    }

    b.trades = append(b.trades, trade)
}

// Close current position
func (b *ArbitrageBacktest) closePosition(books []OrderBookSnapshot) {
    if !b.hasOpenPosition() {
        return
    }

    currentTrade := &b.trades[len(b.trades)-1]
    currentTrade.ExitTime = books[0].Timestamp

    // Calculate PnL
    currentTrade.PnL = (currentTrade.SellPrice - currentTrade.BuyPrice) * currentTrade.Volume
    currentTrade.PnL -= currentTrade.Fees
}

// Calculate position size
func (b *ArbitrageBacktest) calculatePositionSize(opp Opportunity) float64 {
    // Find minimum trade size across exchanges
    minSize := b.config.MaxTradeSize
    for _, exchange := range b.config.Exchanges {
        if exchange.MinTradeSize < minSize {
            minSize = exchange.MinTradeSize
        }
    }

    // Calculate size based on available liquidity
    size := math.Min(opp.BuyVolume, opp.SellVolume)
    size = math.Min(size, b.config.MaxTradeSize)
    size = math.Max(size, minSize)

    return size
}

// Calculate fees
func (b *ArbitrageBacktest) calculateFees(opp Opportunity, size float64) float64 {
    var totalFees float64

    // Find exchange configs
    var buyExchange, sellExchange ExchangeConfig
    for _, exchange := range b.config.Exchanges {
        if exchange.Name == opp.BuyExchange {
            buyExchange = exchange
        }
        if exchange.Name == opp.SellExchange {
            sellExchange = exchange
        }
    }

    // Calculate trading fees
    totalFees += size * opp.BuyPrice * buyExchange.TradingFee
    totalFees += size * opp.SellPrice * sellExchange.TradingFee

    // Calculate withdrawal fees if applicable
    totalFees += buyExchange.WithdrawalFee
    totalFees += sellExchange.WithdrawalFee

    return totalFees
}

// Calculate latency
func (b *ArbitrageBacktest) calculateLatency(opp Opportunity) time.Duration {
    return b.config.Latency[opp.BuyExchange] + b.config.Latency[opp.SellExchange]
}

// Update portfolio
func (b *ArbitrageBacktest) updatePortfolio(books []OrderBookSnapshot) {
    // Implement portfolio update logic
}

// Calculate performance metrics
func (b *ArbitrageBacktest) calculateMetrics() {
    metrics := &ArbitrageMetrics{}

    // Calculate basic metrics
    metrics.TotalTrades = len(b.trades)
    for _, trade := range b.trades {
        if trade.PnL > 0 {
            metrics.SuccessfulTrades++
        } else {
            metrics.FailedTrades++
        }
        metrics.TotalPnL += trade.PnL
        metrics.TotalFees += trade.Fees
        metrics.TotalVolume += trade.Volume
        metrics.AverageLatency += trade.Latency
    }

    // Calculate averages
    if metrics.TotalTrades > 0 {
        metrics.AverageTrade = metrics.TotalPnL / float64(metrics.TotalTrades)
        metrics.AverageLatency = metrics.AverageLatency / time.Duration(metrics.TotalTrades)
    }

    // Calculate best and worst trades
    for _, trade := range b.trades {
        if trade.PnL > metrics.BestTrade {
            metrics.BestTrade = trade.PnL
        }
        if trade.PnL < metrics.WorstTrade {
            metrics.WorstTrade = trade.PnL
        }
    }

    // Calculate maximum drawdown
    metrics.MaxDrawdown = b.calculateMaxDrawdown()

    // Calculate Sharpe ratio
    metrics.SharpeRatio = b.calculateSharpeRatio()

    b.metrics = metrics
}

// Calculate maximum drawdown
func (b *ArbitrageBacktest) calculateMaxDrawdown() float64 {
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
func (b *ArbitrageBacktest) calculateSharpeRatio() float64 {
    // Implement Sharpe ratio calculation
    return 0
}
```

## Usage Example

```go
func main() {
    // Create backtest configuration
    config := ArbitrageBacktestConfig{
        StartTime:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
        EndTime:        time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
        InitialCapital: 10000,
        Exchanges: []ExchangeConfig{
            {
                Name:           "binance",
                TradingFee:     0.001,  // 0.1%
                WithdrawalFee:  0.0005, // 0.05%
                MinTradeSize:   0.001,
                MaxTradeSize:   1.0,
                OrderBookDepth: 10,
            },
            {
                Name:           "kucoin",
                TradingFee:     0.001,
                WithdrawalFee:  0.0005,
                MinTradeSize:   0.001,
                MaxTradeSize:   1.0,
                OrderBookDepth: 10,
            },
        },
        MaxPositions:     5,
        MinSpread:        0.002,  // 0.2%
        MaxHoldingTime:   time.Minute * 5,
        Latency: map[string]time.Duration{
            "binance": time.Millisecond * 100,
            "kucoin":  time.Millisecond * 150,
        },
        MaxTradeSize:     1.0,
        MinTradeSize:     0.001,
    }

    // Create backtest instance
    backtest := NewArbitrageBacktest(config)

    // Run backtest
    err := backtest.Run()
    if err != nil {
        log.Fatal(err)
    }

    // Print results
    fmt.Printf("Total Trades: %d\n", backtest.metrics.TotalTrades)
    fmt.Printf("Successful Trades: %d\n", backtest.metrics.SuccessfulTrades)
    fmt.Printf("Total PnL: $%.2f\n", backtest.metrics.TotalPnL)
    fmt.Printf("Total Fees: $%.2f\n", backtest.metrics.TotalFees)
    fmt.Printf("Average Latency: %v\n", backtest.metrics.AverageLatency)
    fmt.Printf("Max Drawdown: %.2f%%\n", backtest.metrics.MaxDrawdown*100)
    fmt.Printf("Sharpe Ratio: %.2f\n", backtest.metrics.SharpeRatio)
}
```

## Performance Metrics

### Profit Metrics
- Total PnL
- Net PnL (after fees)
- Average Trade
- Best Trade
- Worst Trade
- Success Rate

### Cost Metrics
- Total Fees
- Average Fees per Trade
- Trading Fees
- Withdrawal Fees
- Slippage Costs

### Execution Metrics
- Average Latency
- Maximum Latency
- Failed Executions
- Partial Fills
- Average Spread

## Visualization

### PnL Distribution
```go
func plotPnLDistribution(trades []ArbitrageTrade) {
    // Implement PnL distribution plotting
}
```

### Latency Analysis
```go
func plotLatencyAnalysis(trades []ArbitrageTrade) {
    // Implement latency analysis plotting
}
```

### Spread Analysis
```go
func plotSpreadAnalysis(trades []ArbitrageTrade) {
    // Implement spread analysis plotting
}
```

## Future Improvements

1. **Backtesting Enhancement**
   - Add more exchanges
   - Implement realistic order execution
   - Add market impact modeling
   - Support multiple trading pairs

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