# Technical Indicators

This directory contains documentation for various technical indicators implemented in the `indicators` package.

## Overview

Technical indicators are mathematical calculations based on price, volume, or other market data that help traders identify potential trading opportunities and market conditions.

## Indicator Categories

### Trend Indicators
- [Moving Averages](trend/moving_averages.md)
  - Simple Moving Average (SMA)
  - Exponential Moving Average (EMA)
  - Weighted Moving Average (WMA)
- [MACD](trend/macd.md)
- [ADX](trend/adx.md)
- [Parabolic SAR](trend/parabolic_sar.md)

### Momentum Indicators
- [RSI](momentum/rsi.md)
- [Stochastic Oscillator](momentum/stochastic.md)
- [CCI](momentum/cci.md)
- [Williams %R](momentum/williams_r.md)

### Volume Indicators
- [Volume Profile](volume/volume_profile.md)
- [OBV](volume/obv.md)
- [Chaikin Money Flow](volume/cmf.md)
- [Volume Weighted Average Price (VWAP)](volume/vwap.md)

### Volatility Indicators
- [Bollinger Bands](volatility/bollinger_bands.md)
- [ATR](volatility/atr.md)
- [Keltner Channels](volatility/keltner_channels.md)
- [Standard Deviation](volatility/std_dev.md)

## Implementation Details

### Common Features
- Configurable parameters
- Performance optimization
- Error handling
- Data validation

### Data Structures

```go
type IndicatorResult struct {
    Value     float64
    Signal    string    // "buy", "sell", or "neutral"
    Strength  float64   // Signal strength (0.0 to 1.0)
    Time      time.Time
}

type IndicatorConfig struct {
    Period    int
    Threshold float64
    // Additional parameters specific to each indicator
}
```

## Usage Example

```go
import "github.com/rangertaha/gotal/indicators"

// Create indicator configuration
config := indicators.Config{
    Period:    14,
    Threshold: 0.7,
}

// Calculate RSI
rsi := indicators.NewRSI(config)
result := rsi.Calculate(candles)

// Process results
for _, r := range result {
    if r.Signal == "buy" && r.Strength > 0.8 {
        // Implement buy logic
    }
}
```

## Best Practices

1. **Parameter Selection**
   - Use appropriate timeframes
   - Consider market conditions
   - Test different parameters

2. **Signal Confirmation**
   - Use multiple indicators
   - Consider price action
   - Look for confluence

3. **Risk Management**
   - Set appropriate stop losses
   - Consider indicator limitations
   - Monitor for false signals

## Limitations

1. **Lagging Nature**
   - Most indicators are lagging
   - May miss quick market moves
   - Consider leading indicators

2. **False Signals**
   - Indicators can give false signals
   - Use additional confirmation
   - Consider market context

3. **Parameter Sensitivity**
   - Results vary with parameters
   - Test different settings
   - Consider market conditions

## Future Improvements

1. **Machine Learning Integration**
   - Adaptive parameters
   - Pattern recognition
   - Signal optimization

2. **Performance Optimization**
   - Parallel processing
   - Memory optimization
   - Caching strategies

3. **Additional Indicators**
   - Custom indicators
   - Market-specific indicators
   - Advanced statistical indicators 