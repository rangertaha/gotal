# Gotal - Golang Technical Analysis Library

[![Go Report Card](https://goreportcard.com/badge/github.com/rangertaha/gotal?style=flat-square)](https://goreportcard.com/report/github.com/rangertaha/gotal) [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/rangertaha/gotal) [![PkgGoDev](https://pkg.go.dev/badge/github.com/rangertaha/gotal)](https://pkg.go.dev/github.com/rangertaha/gotal) [![Release](https://img.shields.io/github/release/rangertaha/gotal.svg?style=flat-square)](https://github.com/rangertaha/gotal/releases/latest)

Gotal is a comprehensive technical analysis library for Go that provides a framework for creating, training, testing, and running financial trading bots. It includes a rich set of technical indicators, data providers, broker integrations, and a powerful CLI tool for managing trading strategies.

## Features

- **Technical Indicators**: Moving averages (SMA, EMA, WMA), MACD, OHLC/OHLCV aggregation, and more
- **Data Providers**: Mock data provider for testing, with extensible provider system
- **CLI Tool**: Command-line interface for project management, data filling, strategy training, and execution
- **Time Series Processing**: Efficient time series data structures and operations
- **Extensible Architecture**: Plugin-based system for indicators, providers, brokers, and strategies

## Installation

```bash
go get github.com/rangertaha/gotal
```

## Quick Start

### Using Indicators

```go
package main

import (
    "time"
    ind "github.com/rangertaha/gotal/pkg/indicators"
    "github.com/rangertaha/gotal/examples"
)

func main() {
    // Create a price series
    prices := examples.PricesSeries(1000, time.Second)
    
    // Convert to OHLCV
    ohlcv := ind.OHLCV(prices, ind.WithPeriod(25), ind.OnField("price"))
    
    // Calculate Simple Moving Average
    sma := ind.SMA(ohlcv, ind.WithPeriod(20), ind.OnField("close"))
    sma.Print()
    
    // Calculate MACD
    macd := ind.MACD(ohlcv,
        ind.OnField("close"),
        ind.WithFastPeriod(12),
        ind.WithSlowPeriod(26),
        ind.WithSignalPeriod(9),
    )
    macd.Print()
}
```

### Using Data Providers

```go
package main

import (
    "time"
    "github.com/rangertaha/gotal/pkg/providers"
)

func main() {
    ohlcv := providers.Mock(
        providers.WithDataset("sine"),
        providers.WithSymbol("BTC"),
        providers.WithDuration(1*time.Minute),
        providers.WithStartDate(time.Now().AddDate(-10, 0, 0)),
        providers.WithEndDate(time.Now()),
    )
    ohlcv.Print()
}
```

## CLI Usage

```bash
# Initialize a new project
gota init

# Create a new trading project
gota new myproject

# Fill data from provider
gota fill -p polygon -d 1m -s 2025-01-01

# Train a strategy
gota train -s 2025-01-01 -e 2025-01-02

# Test a strategy
gota test -s 2025-01-01 -e 2025-01-02
```

## Available Indicators

- **SMA** - Simple Moving Average
- **EMA** - Exponential Moving Average
- **WMA** - Weighted Moving Average
- **MACD** - Moving Average Convergence Divergence
- **OHLC/OHLCV** - Open, High, Low, Close (Volume) aggregation

## Examples

See the `examples/` directory for detailed examples:
- `examples/indicators/ma/` - Moving average examples
- `examples/indicators/macd/` - MACD indicator example
- `examples/strategy/macd/` - MACD trading strategy

## Documentation

- [GoDoc](http://godoc.org/github.com/rangertaha/gotal)
- [PkgGoDev](https://pkg.go.dev/github.com/rangertaha/gotal)

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](.github/CONTRIBUTING.md) for details.

## License

See [LICENSE](LICENSE) file for details.

## Author

**Rangertaha** - [rangertaha@gmail.com](mailto:rangertaha@gmail.com)
