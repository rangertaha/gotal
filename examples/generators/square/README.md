# Square Wave Generator Example

This example demonstrates how to generate a square wave time series using the Gotal library.

## Overview

The square wave generator creates a periodic waveform that alternates between two discrete values (high and low) at regular intervals. Unlike a sine wave which smoothly transitions between values, a square wave has sharp transitions creating a distinctive rectangular pattern.

## Usage

```bash
go run examples/generators/square/main.go
```

## Parameters

The `gen.Square()` function accepts the following parameters:

- `duration`: Time interval between each data point (e.g., `time.Second`)
- `amplitude`: Peak amplitude of the wave (the high value will be 2x amplitude, low is 0)
- `frequency`: Step size for wave generation (controls resolution)
- `offset`: Time offset to apply to the starting time
- `tags`: Map of metadata tags to attach to each data point

## Example Output

The square wave alternates between:
- **High value**: `2 * amplitude` 
- **Low value**: `0`

Sample output with amplitude=100:
```
TIMESTAMP       PRICE          SYMBOL         EXCHANGE       CURRENCY       ASSET
1732468800      200.000000     BTC-USD        BINANCE        USD            BTC
1732468801      200.000000     BTC-USD        BINANCE        USD            BTC
1732468802      200.000000     BTC-USD        BINANCE        USD            BTC
1732468803      0.000000       BTC-USD        BINANCE        USD            BTC
1732468804      0.000000       BTC-USD        BINANCE        USD            BTC
1732468805      0.000000       BTC-USD        BINANCE        USD            BTC
```

## Applications

Square waves are useful for:

1. **Signal Processing**: Testing digital systems and filters
2. **Algorithm Testing**: Creating predictable binary patterns for testing trading strategies
3. **Backtesting**: Simulating markets with distinct bull/bear phases
4. **Pattern Recognition**: Training algorithms to detect sharp market transitions
5. **Control Systems**: Modeling on/off or buy/sell signals

## Mathematical Properties

- **Period**: Complete cycle from high → low → high
- **Duty Cycle**: 50% (equal time high and low)
- **Frequency Response**: Rich harmonic content (odd harmonics)
- **Rise/Fall Time**: Instantaneous (ideal square wave)

## Customization

You can modify the example to:

- Change the amplitude for different price ranges
- Adjust frequency for finer/coarser resolution
- Add noise for more realistic market simulation
- Combine with other generators for complex patterns

## Comparison with Sine Wave

| Feature | Square Wave | Sine Wave |
|---------|-------------|-----------|
| Transitions | Sharp/Instant | Smooth/Gradual |
| Values | Binary (2 states) | Continuous |
| Harmonics | Rich (odd only) | Pure fundamental |
| Use Case | Digital signals | Analog signals |

The square wave generator complements the sine wave generator by providing a different type of periodic signal that's useful for testing algorithms that need to handle sudden market changes or binary conditions.
