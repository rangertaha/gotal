# Gotal Technical Analysis Library

Gotal is a comprehensive technical analysis library for Go, providing tools for market analysis, pattern recognition, and trading strategy development.

## Documentation Structure

### Core Components
- [Patterns](patterns/README.md) - Chart pattern recognition
- [Indicators](indicators/README.md) - Technical indicators
- [Types](types/README.md) - Data structures and types
- [Examples](examples/README.md) - Usage examples

## Quick Start

```go
import "github.com/rangertaha/gotal"

// Create a new analysis instance
analysis := gotal.NewAnalysis()

// Add data
analysis.AddCandles(candles)

// Analyze patterns
patterns := analysis.FindPatterns()

// Calculate indicators
indicators := analysis.CalculateIndicators()
```

## Features

### Pattern Recognition
- Reversal patterns (Head & Shoulders, Double Top/Bottom)
- Continuation patterns (Triangles, Flags & Pennants)
- Candlestick patterns
- Pattern confidence scoring

### Technical Indicators
- Trend indicators
- Momentum indicators
- Volume indicators
- Volatility indicators

### Data Management
- OHLCV data structures
- Time series handling
- Data validation
- Performance optimization

## Installation

```bash
go get github.com/rangertaha/gotal
```

## Usage Examples

See the [examples](examples/README.md) directory for detailed usage examples.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 