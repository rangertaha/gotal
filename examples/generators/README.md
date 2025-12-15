# Signal Generators - Gotal Library

This directory contains examples of various signal generators available in the Gotal technical analysis library. These generators are essential for testing trading algorithms, backtesting strategies, and understanding market behavior patterns.

## Available Generators

### Periodic/Mathematical Signals

#### **Sine Wave** - `gen.Sine()`
- **Pattern**: Smooth sinusoidal oscillation
- **Example**: `examples/generators/sine/`
- **Use Case**: Natural market cycles, seasonal patterns
- **Parameters**: `(duration, amplitude, frequency, offset, tags)`

#### **Square Wave** - `gen.Square()`
- **Pattern**: Binary high/low states with instant transitions  
- **Example**: `examples/generators/square/`
- **Use Case**: Digital signals, binary market conditions
- **Parameters**: `(duration, amplitude, frequency, offset, tags)`

#### **Triangle Wave** - `gen.Triangle()`
- **Pattern**: Linear rise and fall transitions
- **Example**: `examples/generators/triangle/`
- **Use Case**: Linear trends, predictable reversals
- **Parameters**: `(duration, amplitude, frequency, offset, tags)`

#### **Sawtooth Wave** - `gen.Sawtooth()`
- **Pattern**: Linear rise, instant drop (asymmetric)
- **Example**: `examples/generators/sawtooth/`
- **Use Case**: Market corrections, pump-and-dump patterns
- **Parameters**: `(duration, amplitude, frequency, offset, tags)`

#### **Pulse Wave** - `gen.Pulse()`
- **Pattern**: Square wave with adjustable duty cycle
- **Use Case**: Bull/bear markets with different durations
- **Parameters**: `(duration, amplitude, frequency, dutyCycle, offset, tags)`

### Noise Generators

#### **White Noise** - `gen.WhiteNoise()`
- **Pattern**: Random uniform distribution
- **Example**: `examples/generators/noise/`
- **Use Case**: Market noise simulation, stress testing
- **Parameters**: `(duration, amplitude, samples, tags)`

#### **Pink Noise** - `gen.PinkNoise()`
- **Pattern**: 1/f noise (more realistic than white noise)
- **Example**: `examples/generators/noise/`
- **Use Case**: Natural market fluctuations
- **Parameters**: `(duration, amplitude, samples, tags)`

### Financial Market Specific

#### **Random Walk** - `gen.RandomWalk()`
- **Pattern**: Cumulative random steps with optional drift
- **Example**: `examples/generators/random_walk/`
- **Use Case**: Efficient market hypothesis, baseline comparison
- **Parameters**: `(duration, stepSize, drift, samples, tags)`

#### **Geometric Brownian Motion** - `gen.GBM()`
- **Pattern**: Log-normal price movements (Black-Scholes model)
- **Example**: `examples/generators/gbm/`
- **Use Case**: Realistic stock price simulation, options pricing
- **Parameters**: `(duration, initialPrice, drift, volatility, samples, tags)`

#### **Ornstein-Uhlenbeck Process** - `gen.OrnsteinUhlenbeck()`
- **Pattern**: Mean-reverting stochastic process
- **Use Case**: Interest rates, commodity prices, pairs trading
- **Parameters**: `(duration, mean, speed, volatility, samples, tags)`

#### **GARCH Process** - `gen.GARCH()`
- **Pattern**: Volatility clustering
- **Use Case**: Risk management, volatility modeling
- **Parameters**: `(duration, samples, alpha, beta, omega, tags)`

### Technical Analysis Patterns

#### **Linear Trend** - `gen.LinearTrend()`
- **Pattern**: Straight-line trend with optional noise
- **Example**: `examples/generators/trends/`
- **Use Case**: Testing trend-following algorithms
- **Parameters**: `(duration, startPrice, slope, noise, samples, tags)`

#### **Exponential Trend** - `gen.ExponentialTrend()`
- **Pattern**: Exponential growth/decay with optional noise
- **Example**: `examples/generators/trends/`
- **Use Case**: Bubble formation, exponential growth patterns
- **Parameters**: `(duration, startPrice, growthRate, noise, samples, tags)`

#### **Support/Resistance** - `gen.SupportResistance()`
- **Pattern**: Price bouncing between defined levels
- **Use Case**: Testing breakout strategies, range-bound markets
- **Parameters**: `(duration, support, resistance, bounceStrength, noise, samples, tags)`

### System/Control Signals

#### **Step Function** - `gen.Step()`
- **Pattern**: Instant level change at specified time
- **Example**: `examples/generators/step_ramp/`
- **Use Case**: Market regime changes, news events
- **Parameters**: `(duration, lowValue, highValue, stepTime, samples, tags)`

#### **Ramp Function** - `gen.Ramp()`
- **Pattern**: Linear transition between two values
- **Example**: `examples/generators/step_ramp/`
- **Use Case**: Gradual market transitions
- **Parameters**: `(duration, startValue, endValue, samples, tags)`

#### **Exponential Decay** - `gen.ExponentialDecay()`
- **Pattern**: Exponential decrease over time
- **Use Case**: Market corrections, decay patterns
- **Parameters**: `(duration, initial, halfLife, samples, tags)`

#### **Exponential Growth** - `gen.ExponentialGrowth()`
- **Pattern**: Exponential increase over time
- **Use Case**: Bull markets, growth patterns
- **Parameters**: `(duration, initial, growthRate, samples, tags)`

### Advanced Signals

#### **Jump Diffusion** - `gen.JumpDiffusion()`
- **Pattern**: Continuous movement with occasional jumps
- **Use Case**: News-driven market movements, event risk
- **Parameters**: `(duration, diffusion, jumpProb, jumpSize, samples, tags)`

## Usage Examples

### Basic Usage Pattern

```go
package main

import (
    "time"
    "github.com/rangertaha/gotal/pkg/gen"
)

func main() {
    tags := map[string]string{
        "symbol":   "BTC-USD",
        "exchange": "BINANCE",
        "currency": "USD",
        "asset":    "BTC",
    }

    // Generate any signal type
    series := gen.RandomWalk(time.Second, 2.0, 0.1, 100, tags)
    series.Print()
}
```

### Testing Different Market Conditions

```go
// Bull market trend
bullTrend := gen.LinearTrend(time.Hour, 100, 0.5, 1.0, 100, tags)

// Bear market with volatility
bearMarket := gen.ExponentialDecay(time.Hour, 200, 50, 50, tags)

// Sideways market
sideways := gen.SupportResistance(time.Minute, 95, 105, 2.0, 0.5, 200, tags)

// High volatility period  
volatility := gen.GARCH(time.Minute, 100, 0.1, 0.85, 0.01, tags)
```

## Applications by Use Case

### Algorithm Testing
- **Trend Following**: Use `LinearTrend`, `ExponentialTrend`
- **Mean Reversion**: Use `OrnsteinUhlenbeck`, `SupportResistance`
- **Breakout Strategies**: Use `Step`, `SupportResistance`
- **Volatility Trading**: Use `GARCH`, `JumpDiffusion`

### Risk Management
- **Stress Testing**: Use `WhiteNoise`, `JumpDiffusion`
- **Scenario Analysis**: Use `GBM`, `GARCH`
- **Backtesting**: Use `RandomWalk` as baseline

### Pattern Recognition
- **Cycle Detection**: Use `Sine`, `Triangle`
- **Trend Identification**: Use `LinearTrend`, `Sawtooth`
- **Regime Change**: Use `Step`, `ExponentialDecay`

## Signal Characteristics Comparison

| Generator | Randomness | Trend | Cycles | Jumps | Realistic |
|-----------|------------|-------|---------|--------|-----------|
| Sine | None | None | Yes | No | Low |
| Square | None | None | Yes | Yes | Low |
| Triangle | None | Linear | Yes | No | Low |
| Sawtooth | None | Asymmetric | Yes | Yes | Medium |
| RandomWalk | High | Optional | No | No | Medium |
| GBM | Medium | Yes | No | No | High |
| GARCH | Medium | No | No | No | High |
| JumpDiffusion | High | No | No | Yes | High |

## Running Examples

```bash
# Test periodic signals
go run examples/generators/sine/main.go
go run examples/generators/square/main.go
go run examples/generators/triangle/main.go

# Test financial models
go run examples/generators/random_walk/main.go
go run examples/generators/gbm/main.go

# Test noise generators
go run examples/generators/noise/main.go

# Test trend generators
go run examples/generators/trends/main.go

# Test control signals
go run examples/generators/step_ramp/main.go
```

## Best Practices

1. **Choose the Right Generator**: Match the generator to your testing scenario
2. **Combine Signals**: Layer multiple generators for realistic market simulation
3. **Parameter Tuning**: Adjust parameters based on the time frame you're testing
4. **Add Noise**: Use noise generators to make synthetic data more realistic
5. **Validate Results**: Always verify your algorithms work on real market data too

## Implementation Notes

- All generators return `*series.Series` objects compatible with Gotal indicators
- Timestamps are automatically generated based on duration and offset
- Random generators use Go's math/rand package (seed as needed for reproducibility)
- Financial models implement standard academic formulations
- Generators are optimized for performance and memory efficiency

This comprehensive set of signal generators provides everything needed for thorough testing of technical analysis algorithms across all market conditions and scenarios.