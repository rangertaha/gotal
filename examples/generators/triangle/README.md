# Triangle Wave Generator Example

This example demonstrates how to generate a triangle wave time series using the Gotal library.

## Overview

The triangle wave generator creates a periodic waveform that linearly rises from zero to a peak value, then linearly falls back to zero, creating a triangular pattern. Unlike sine waves which have curved transitions, or square waves which have instant transitions, triangle waves have smooth linear transitions that create a sawtooth appearance.

## Usage

```bash
go run examples/generators/triangle/main.go
```

## Parameters

The `gen.Triangle()` function accepts the following parameters:

- `duration`: Time interval between each data point (e.g., `time.Second`)
- `amplitude`: Peak amplitude of the wave (the peak value will be 2x amplitude)
- `frequency`: Step size for wave generation (controls resolution)
- `offset`: Time offset to apply to the starting time
- `tags`: Map of metadata tags to attach to each data point

## Example Output

The triangle wave linearly transitions between:
- **Minimum value**: `0`
- **Maximum value**: `2 * amplitude` 
- **Linear transitions**: Smooth rise and fall

Sample output with amplitude=100:
```
TIMESTAMP       PRICE          SYMBOL         EXCHANGE       CURRENCY       ASSET
1732468800      0.000000       BTC-USD        BINANCE        USD            BTC
1732468801      20.000000      BTC-USD        BINANCE        USD            BTC
1732468802      40.000000      BTC-USD        BINANCE        USD            BTC
1732468803      60.000000      BTC-USD        BINANCE        USD            BTC
1732468804      80.000000      BTC-USD        BINANCE        USD            BTC
1732468805      100.000000     BTC-USD        BINANCE        USD            BTC
1732468806      120.000000     BTC-USD        BINANCE        USD            BTC
1732468807      140.000000     BTC-USD        BINANCE        USD            BTC
1732468808      160.000000     BTC-USD        BINANCE        USD            BTC
1732468809      180.000000     BTC-USD        BINANCE        USD            BTC
1732468810      200.000000     BTC-USD        BINANCE        USD            BTC
1732468811      180.000000     BTC-USD        BINANCE        USD            BTC
1732468812      160.000000     BTC-USD        BINANCE        USD            BTC
1732468813      140.000000     BTC-USD        BINANCE        USD            BTC
1732468814      120.000000     BTC-USD        BINANCE        USD            BTC
```

## Applications

Triangle waves are useful for:

1. **Linear Trends**: Simulating markets with consistent upward and downward trends
2. **Technical Analysis**: Testing algorithms that detect trend reversals
3. **Backtesting**: Creating predictable linear price movements for strategy validation
4. **Signal Processing**: Testing filters and smoothing algorithms
5. **Ramp Functions**: Modeling gradual price increases/decreases
6. **Volatility Testing**: Creating controlled linear volatility patterns

## Mathematical Properties

- **Waveform**: Symmetric triangular shape
- **Rise Time**: Linear increase over half cycle
- **Fall Time**: Linear decrease over half cycle
- **Peak Value**: 2 × amplitude
- **Minimum Value**: 0
- **Duty Cycle**: 50% rising, 50% falling
- **Harmonic Content**: Rich in odd harmonics (but lower than square wave)

## Customization

You can modify the example to:

- Change amplitude for different price ranges
- Adjust frequency for steeper/gentler slopes
- Add offset for time-shifted patterns
- Combine with noise for realistic market simulation
- Layer multiple triangle waves for complex patterns

## Comparison with Other Waves

| Feature | Triangle Wave | Sine Wave | Square Wave |
|---------|---------------|-----------|-------------|
| Transitions | Linear | Curved | Instant |
| Slope | Constant | Variable | Infinite |
| Peak Value | 2×amplitude | 2×amplitude | 2×amplitude |
| Min Value | 0 | 0 | 0 |
| Use Case | Linear trends | Natural cycles | Binary states |
| Harmonics | Odd harmonics | Pure tone | Rich odd harmonics |

## Trading Applications

Triangle waves are particularly useful for:

- **Trend Analysis**: Testing how indicators respond to consistent upward/downward trends
- **Support/Resistance**: Creating predictable levels for testing breakout strategies
- **Moving Averages**: Generating data where moving averages will show clear directional bias
- **Momentum Indicators**: Testing how oscillators respond to linear price movements
- **Reversal Patterns**: Creating clear trend reversal points for pattern recognition algorithms

## Signal Processing Context

In signal processing, triangle waves are:
- **Bandwidth Limited**: Unlike square waves, easier on filters
- **Continuous**: No discontinuities (unlike square waves)
- **Linear Rate**: Constant rate of change (unlike sine waves)
- **Predictable**: Mathematical certainty for testing algorithms

The triangle wave generator provides a middle ground between the smooth sine wave and the sharp square wave, making it ideal for testing algorithms that need to handle linear trends and predictable reversals.
