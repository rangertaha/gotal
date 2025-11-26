# Gotal - Golang Technical Analysis Library

[![Go Report Card](https://goreportcard.com/badge/github.com/rangertaha/gotal?style=flat-square)](https://goreportcard.com/report/github.com/rangertaha/gotal) [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/rangertaha/gotal) [![PkgGoDev](https://pkg.go.dev/badge/github.com/rangertaha/gotal)](https://pkg.go.dev/github.com/rangertaha/gotal) [![Release](https://img.shields.io/github/release/rangertaha/gotal.svg?style=flat-square)](https://github.com/rangertaha/gotal/releases/latest)

Gotal is a technical analysis library for Go that provides a framework for creating, training, testing, and running algorithmic trading stratagies. It includes technical indicators, data providers, broker integrations, and a CLI tool for managing the workflow.

## ⚠️ **IMPORTANT DISCLAIMER - READ BEFORE USE** ⚠️

**THIS SOFTWARE IS FOR EDUCATIONAL AND RESEARCH PURPOSES ONLY. NOT FINANCIAL ADVICE.**

- 🚨 **Trading and investing involve significant risk of loss**
- 🚨 **You can lose ALL of your invested capital**
- 🚨 **Past performance does not guarantee future results**
- 🚨 **This software is provided "AS IS" without any warranties**
- 🚨 **The author(s) are NOT LIABLE for any financial losses incurred**

**By using this software, you acknowledge that:**
- You understand the risks involved in trading/investing
- You are using this software at your own risk
- Any trading decisions are your own responsibility
- You should consult with qualified financial professionals before making investment decisions
- This software may contain bugs or produce incorrect results

**Use this software only for learning, research, and testing purposes. Never risk money you cannot afford to lose.**

## Features

- **Technical Indicators**: Moving averages (SMA, EMA, WMA), MACD, aggregation, and more
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



## Risk Disclaimer & Legal Notice

**TRADING DISCLAIMER**: This software and any associated documentation, examples, or strategies are provided for educational and informational purposes only. They do not constitute investment advice, financial advice, trading advice, or any other sort of advice. The use of this software does not guarantee profits, and trading/investing involves substantial risk of loss.

**NO LIABILITY**: Under no circumstances shall the author(s), contributors, or distributors of this software be liable for any direct, indirect, incidental, special, consequential, or punitive damages that result from the use of, or inability to use, this software or any associated strategies, even if the author(s) have been advised of the possibility of such damages.

**ASSUMPTION OF RISK**: By using this software, you acknowledge and agree that you have read this disclaimer, understand its contents, and assume all risks associated with the use of this software. You acknowledge that trading and investing carry inherent risks and that you may lose some or all of your invested capital.

**INDEPENDENT RESEARCH**: Before making any trading or investment decisions, you should conduct your own research and analysis and/or consult with qualified financial professionals.


## License

See [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](.github/CONTRIBUTING.md) for details.


## Author

**Rangertaha** - [rangertaha@gmail.com](mailto:rangertaha@gmail.com)
