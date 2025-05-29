# Arbitrage Strategy

## Overview

This example demonstrates an arbitrage strategy that aims to profit from price differences between different markets or exchanges. The strategy identifies and executes trades when price discrepancies exceed transaction costs.

## Strategy Components

1. **Price Monitoring**
   - Multiple exchange feeds
   - Order book analysis
   - Price spread calculation
   - Transaction cost analysis

2. **Entry Signals**
   - Price discrepancy threshold
   - Volume availability
   - Market depth check
   - Liquidity verification

3. **Exit Signals**
   - Price convergence
   - Time-based exit
   - Maximum holding time
   - Risk limit reached

## Implementation

```go
package main

import (
    "fmt"
    "time"
    "github.com/rangertaha/gotal"
    "github.com/rangertaha/gotal/indicators"
)

// Exchange configuration
type ExchangeConfig struct {
    Name            string
    TradingFee      float64
    WithdrawalFee   float64
    MinTradeSize    float64
    MaxTradeSize    float64
    PricePrecision  int
    VolumePrecision int
}

// Strategy configuration
type Config struct {
    Exchanges        []ExchangeConfig
    MinSpread        float64
    MaxHoldingTime   time.Duration
    MaxPositions     int
    RiskPerTrade     float64
    UpdateInterval   time.Duration
    PriceThreshold   float64
    VolumeThreshold  float64
}

// Order book structure
type OrderBook struct {
    Bids [][]float64  // [price, volume]
    Asks [][]float64  // [price, volume]
    Time time.Time
}

// Strategy state
type State struct {
    Positions     []Position
    LastUpdate    time.Time
    TotalProfit   float64
    ActiveTrades  int
}

// Position structure
type Position struct {
    Exchange1     string
    Exchange2     string
    Type1         string    // "buy" or "sell"
    Type2         string    // "buy" or "sell"
    Price1        float64
    Price2        float64
    Volume        float64
    EntryTime     time.Time
    Status        string    // "open" or "closed"
    Profit        float64
}

// Create new strategy
func NewArbitrageStrategy(config Config) *ArbitrageStrategy {
    return &ArbitrageStrategy{
        config: config,
        state:  &State{},
    }
}

// Main strategy implementation
func (s *ArbitrageStrategy) Analyze(orderBooks map[string]OrderBook) []Signal {
    var signals []Signal

    // Check if we can take more positions
    if s.state.ActiveTrades >= s.config.MaxPositions {
        return signals
    }

    // Calculate arbitrage opportunities
    opportunities := s.findArbitrageOpportunities(orderBooks)

    // Filter and rank opportunities
    validOpportunities := s.filterOpportunities(opportunities)

    // Generate signals for valid opportunities
    for _, opp := range validOpportunities {
        if s.isValidOpportunity(opp) {
            signals = append(signals, s.createSignal(opp))
        }
    }

    return signals
}

// Find arbitrage opportunities
func (s *ArbitrageStrategy) findArbitrageOpportunities(orderBooks map[string]OrderBook) []Opportunity {
    var opportunities []Opportunity

    // Compare each pair of exchanges
    for i, ex1 := range s.config.Exchanges {
        for j, ex2 := range s.config.Exchanges {
            if i >= j {
                continue
            }

            // Get order books
            book1 := orderBooks[ex1.Name]
            book2 := orderBooks[ex2.Name]

            // Calculate spread
            spread := s.calculateSpread(book1, book2)

            // Check if spread is profitable
            if spread > s.config.MinSpread {
                opportunities = append(opportunities, Opportunity{
                    Exchange1: ex1.Name,
                    Exchange2: ex2.Name,
                    Spread:    spread,
                    Book1:     book1,
                    Book2:     book2,
                })
            }
        }
    }

    return opportunities
}

// Calculate spread between two order books
func (s *ArbitrageStrategy) calculateSpread(book1, book2 OrderBook) float64 {
    // Get best bid and ask prices
    bestBid1 := book1.Bids[0][0]
    bestAsk1 := book1.Asks[0][0]
    bestBid2 := book2.Bids[0][0]
    bestAsk2 := book2.Asks[0][0]

    // Calculate possible spreads
    spread1 := (bestBid1 - bestAsk2) / bestAsk2
    spread2 := (bestBid2 - bestAsk1) / bestAsk1

    return math.Max(spread1, spread2)
}

// Filter opportunities
func (s *ArbitrageStrategy) filterOpportunities(opportunities []Opportunity) []Opportunity {
    var validOpportunities []Opportunity

    for _, opp := range opportunities {
        // Check volume
        if !s.hasSufficientVolume(opp) {
            continue
        }

        // Check transaction costs
        if !s.isProfitableAfterCosts(opp) {
            continue
        }

        // Check market depth
        if !s.hasEnoughDepth(opp) {
            continue
        }

        validOpportunities = append(validOpportunities, opp)
    }

    return validOpportunities
}

// Check if opportunity has sufficient volume
func (s *ArbitrageStrategy) hasSufficientVolume(opp Opportunity) bool {
    // Get available volumes
    volume1 := opp.Book1.Bids[0][1]
    volume2 := opp.Book2.Asks[0][1]

    // Check against minimum trade size
    minVolume := math.Max(s.config.Exchanges[0].MinTradeSize, s.config.Exchanges[1].MinTradeSize)
    return volume1 >= minVolume && volume2 >= minVolume
}

// Check if opportunity is profitable after costs
func (s *ArbitrageStrategy) isProfitableAfterCosts(opp Opportunity) bool {
    // Calculate total costs
    totalCosts := s.calculateTotalCosts(opp)

    // Check if spread covers costs
    return opp.Spread > totalCosts
}

// Calculate total transaction costs
func (s *ArbitrageStrategy) calculateTotalCosts(opp Opportunity) float64 {
    var costs float64

    // Add trading fees
    for _, ex := range s.config.Exchanges {
        costs += ex.TradingFee
    }

    // Add withdrawal fees if applicable
    costs += s.config.Exchanges[0].WithdrawalFee
    costs += s.config.Exchanges[1].WithdrawalFee

    return costs
}

// Check if opportunity has enough market depth
func (s *ArbitrageStrategy) hasEnoughDepth(opp Opportunity) bool {
    // Implement market depth analysis
    return true
}

// Create signal from opportunity
func (s *ArbitrageStrategy) createSignal(opp Opportunity) Signal {
    return Signal{
        Type:       "arbitrage",
        Exchange1:  opp.Exchange1,
        Exchange2:  opp.Exchange2,
        Price1:     opp.Book1.Bids[0][0],
        Price2:     opp.Book2.Asks[0][0],
        Volume:     s.calculateTradeVolume(opp),
        Time:       time.Now(),
        Strength:   0.9,
    }
}

// Calculate trade volume
func (s *ArbitrageStrategy) calculateTradeVolume(opp Opportunity) float64 {
    // Get available volumes
    volume1 := opp.Book1.Bids[0][1]
    volume2 := opp.Book2.Asks[0][1]

    // Use minimum of available volumes
    return math.Min(volume1, volume2)
}

// Position management
func (s *ArbitrageStrategy) ManagePosition(signal Signal) {
    // Create new position
    position := Position{
        Exchange1: signal.Exchange1,
        Exchange2: signal.Exchange2,
        Type1:     "buy",
        Type2:     "sell",
        Price1:    signal.Price1,
        Price2:    signal.Price2,
        Volume:    signal.Volume,
        EntryTime: signal.Time,
        Status:    "open",
    }

    // Add position to state
    s.state.Positions = append(s.state.Positions, position)
    s.state.ActiveTrades++
}

// Check for exit conditions
func (s *ArbitrageStrategy) CheckExit(orderBooks map[string]OrderBook) []Signal {
    var signals []Signal

    for i, pos := range s.state.Positions {
        if pos.Status == "closed" {
            continue
        }

        // Check time-based exit
        if time.Since(pos.EntryTime) > s.config.MaxHoldingTime {
            signals = append(signals, Signal{
                Type:      "close",
                Position:  i,
                Time:      time.Now(),
            })
            continue
        }

        // Check price convergence
        if s.hasPriceConverged(pos, orderBooks) {
            signals = append(signals, Signal{
                Type:      "close",
                Position:  i,
                Time:      time.Now(),
            })
        }
    }

    return signals
}

// Check if prices have converged
func (s *ArbitrageStrategy) hasPriceConverged(pos Position, orderBooks map[string]OrderBook) bool {
    // Get current prices
    currentPrice1 := orderBooks[pos.Exchange1].Bids[0][0]
    currentPrice2 := orderBooks[pos.Exchange2].Asks[0][0]

    // Calculate current spread
    currentSpread := math.Abs(currentPrice1 - currentPrice2) / currentPrice2

    // Check if spread has converged
    return currentSpread < s.config.PriceThreshold
}
```

## Usage Example

```go
func main() {
    // Create exchange configurations
    exchanges := []ExchangeConfig{
        {
            Name:            "exchange1",
            TradingFee:      0.001,  // 0.1%
            WithdrawalFee:   0.0005, // 0.05%
            MinTradeSize:    0.01,
            MaxTradeSize:    1.0,
            PricePrecision:  8,
            VolumePrecision: 8,
        },
        {
            Name:            "exchange2",
            TradingFee:      0.001,  // 0.1%
            WithdrawalFee:   0.0005, // 0.05%
            MinTradeSize:    0.01,
            MaxTradeSize:    1.0,
            PricePrecision:  8,
            VolumePrecision: 8,
        },
    }

    // Create strategy configuration
    config := Config{
        Exchanges:        exchanges,
        MinSpread:        0.002,  // 0.2%
        MaxHoldingTime:   time.Minute * 5,
        MaxPositions:     5,
        RiskPerTrade:     0.01,   // 1%
        UpdateInterval:   time.Second,
        PriceThreshold:   0.001,  // 0.1%
        VolumeThreshold:  0.1,    // 10% of available volume
    }

    // Create strategy instance
    strategy := NewArbitrageStrategy(config)

    // Start order book feeds
    orderBooks := make(map[string]OrderBook)
    go startOrderBookFeed("exchange1", orderBooks)
    go startOrderBookFeed("exchange2", orderBooks)

    // Main loop
    for {
        // Analyze opportunities
        signals := strategy.Analyze(orderBooks)

        // Process signals
        for _, signal := range signals {
            strategy.ManagePosition(signal)
        }

        // Check for exits
        exitSignals := strategy.CheckExit(orderBooks)
        for _, signal := range exitSignals {
            // Close position
            strategy.ClosePosition(signal.Position)
        }

        // Wait for next update
        time.Sleep(config.UpdateInterval)
    }
}
```

## Strategy Parameters

### Exchange Configuration
- Trading Fee: 0.1%
- Withdrawal Fee: 0.05%
- Min Trade Size: 0.01
- Max Trade Size: 1.0

### Strategy Parameters
- Min Spread: 0.2%
- Max Holding Time: 5 minutes
- Max Positions: 5
- Risk Per Trade: 1%
- Update Interval: 1 second

### Entry Conditions
- Price discrepancy > 0.2%
- Sufficient volume
- Profitable after costs
- Enough market depth

### Risk Management
- Maximum positions
- Time-based exit
- Price convergence
- Volume limits

## Performance Considerations

### Optimization
1. **Execution Speed**
   - Minimize latency
   - Use efficient data structures
   - Optimize calculations

2. **Memory Usage**
   - Reuse order book instances
   - Clear old data
   - Use fixed-size arrays

### Backtesting
1. **Data Requirements**
   - Multiple exchange feeds
   - Order book data
   - Transaction history

2. **Performance Metrics**
   - Win rate
   - Profit factor
   - Maximum drawdown
   - Sharpe ratio

## Risk Management

### Position Sizing
```go
func calculatePositionSize(capital float64, risk float64, spread float64) float64 {
    riskAmount := capital * risk
    return riskAmount / spread
}
```

### Cost Analysis
```go
func calculateTotalCosts(exchanges []ExchangeConfig, volume float64) float64 {
    var costs float64
    for _, ex := range exchanges {
        costs += volume * ex.TradingFee
        costs += ex.WithdrawalFee
    }
    return costs
}
```

## Future Improvements

1. **Strategy Enhancement**
   - Add more exchanges
   - Implement triangular arbitrage
   - Add statistical arbitrage

2. **Risk Management**
   - Dynamic position sizing
   - Portfolio optimization
   - Correlation analysis

3. **Performance Analysis**
   - Detailed backtesting
   - Monte Carlo simulation
   - Risk metrics calculation 