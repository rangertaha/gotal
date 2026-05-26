# gotal — Go Technical Analysis

[![Go Report Card](https://goreportcard.com/badge/github.com/rangertaha/gotal?style=flat-square)](https://goreportcard.com/report/github.com/rangertaha/gotal)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/rangertaha/gotal)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/rangertaha/gotal)](https://pkg.go.dev/github.com/rangertaha/gotal)
[![Release](https://img.shields.io/github/release/rangertaha/gotal.svg?style=flat-square)](https://github.com/rangertaha/gotal/releases/latest)

> **Status: UNDER DEVELOPMENT.** API may still change. Educational and research use only — see [Disclaimer](#disclaimer).

`gotal` is a technical-analysis library for Go. It provides indicators that work in two modes:

- **Batch** — run an indicator over a historical `TimeSeries`.
- **Streaming** — pipe live ticks through an indicator using channels.

Both modes share the same registered implementations, so a streaming EMA produces the same values as a batch EMA on the same data.

## Install

```bash
go get github.com/rangertaha/gotal
```

## Package layout

| Import path | Purpose |
| --- | --- |
| `github.com/rangertaha/gotal` | Public interfaces (`TimeSeries`, `Tick`, `Stream`, `Indicator`, `Plot`, `Fields`, `Tags`, `Vector`) and the batch indicator vars (`EMA`, `SMA`, …). |
| `github.com/rangertaha/gotal/ticker` | `*ticker.Ticker` — concrete `TimeSeries` implementation with `Print`, `Save`, `Plot` methods. |
| `github.com/rangertaha/gotal/stream` | `*stream.Stream` — channel-backed live pipeline, plus `stream.Apply` and stream-mode indicator wrappers. |
| `github.com/rangertaha/gotal/io` | Readers: `io.CSV`, `io.JSON`, `io.JSONL`, plus `io.Read(path)` that dispatches by extension. |

## Quick start

### Batch: compute EMA over a CSV

```go
package main

import (
    "log"

    "github.com/rangertaha/gotal"
    "github.com/rangertaha/gotal/io"
)

func main() {
    ts, err := io.Read("prices.csv")
    if err != nil {
        log.Fatal(err)
    }

    ema, err := gotal.EMA(ts,
        gotal.With("name", "ema"),
        gotal.With("source", "close"),
        gotal.With("period", 14),
    )
    if err != nil {
        log.Fatal(err)
    }

    ema.Print()
    _ = ema.Save("ema.csv")
}
```

Or use the typed convenience wrapper:

```go
ema, _ := gotal.Ema(ts, 14, 0) // period=14, alpha=0 means default 2/(period+1)
```

### Streaming: pipe a live source through an indicator

```go
package main

import (
    "fmt"

    "github.com/rangertaha/gotal/io"
    "github.com/rangertaha/gotal/stream"
)

func main() {
    src := io.CSV("prices.csv").Stream().(*stream.Stream)
    out := stream.Ema(src, 14, 0)

    for tick := range out.Ticks() {
        signals, _ := tick.Signals()
        fmt.Printf("%s ema=%.4f\n", tick.Time().Format("2006-01-02"), signals["ema"])
    }
}
```

### Build a Ticker manually

```go
import "github.com/rangertaha/gotal/ticker"

t, _ := ticker.New("prices",
    ticker.WithTags(map[string]string{"symbol": "BTC-USD"}),
    ticker.WithTicks(ticks...),
)
```

## Configuration

Indicators accept functional options through `gotal.ConfigOption`. Construct them with `gotal.With(key, value)`:

```go
sma, _ := gotal.SMA(ts,
    gotal.With("period", 20),
    gotal.With("source", "close"),
    gotal.With("name",   "sma_20"),
)
```

Common keys:

| Key | Type | Used by |
| --- | --- | --- |
| `name` | string | Output signal/field name |
| `source` | string | Name of the input field to read from each tick (default `close`) |
| `period` | int | Window length |
| `alpha` | float | Smoothing factor (EMA) — 0 means default `2/(period+1)` |

## Indicators

Around 330 indicator ids are registered. Implemented today (each callable as `gotal.<NAME>` and most also have a typed convenience wrapper like `gotal.Ema(ts, 14, 0)`):

| Group | Implemented |
| --- | --- |
| **Overlap** | SMA, EMA, WMA, DEMA, TEMA, TRIMA, MIDPOINT, MIDPRICE, BBANDS, HMA |
| **Momentum** | MOM, ROC, ROCP, ROCR, ROCR100, WILLR, RSI, CCI, MACD, APO, PPO, CMO, TRIX, AROON, AROONOSC, STOCH, STOCHF, MFI |
| **Volatility** | TRANGE, ATR, NATR |
| **Volume** | OBV, AD, ADOSC, VWMA, VWAP |
| **Price** | AVGPRICE, MEDPRICE, TYPPRICE, WCLPRICE, HLC3, OHLC4, HEIKINASHI |
| **Statistic** | STDDEV, VARIANCE, LINEARREG, LINEARREG_SLOPE, LINEARREG_ANGLE, LINEARREG_INTERCEPT, TSF, BETA, CORREL |
| **Math operators** | MAX, MIN, SUMWINDOW, ADD, SUB, MULT, DIV, MAXINDEX, MININDEX |
| **Math transforms** | ACOS, ASIN, ATAN, COS, SIN, TAN, COSH, SINH, TANH, CEIL, FLOOR, EXP, LN, LOG10, SQRT |
| **Signals** | RSISIGNAL, MFISIGNAL, CCISIGNAL, WILLRSIGNAL, MOMSIGNAL, ROCSIGNAL, TRIXSIGNAL, PPOSIGNAL, CMOSIGNAL, MACDSIGNAL_CROSS, STOCHSIGNAL |

Unimplemented ids (T3, KAMA, MAMA, ADX family, all `CDL*` patterns, Hilbert transforms, …) are registered as stubs that return a clear `"not implemented"` error at call time. See [`NOTES.md`](NOTES.md) for the full catalogue and per-indicator status.

### Group runners

Apply every implemented indicator in a category at once:

```go
out := gotal.Momentum(ts)   // adds rsi, macd, stoch_k/d, aroon_up/down, mom, roc, …
out  = gotal.Volume(ts)     // obv, ad, adosc, vwma, vwap
out  = gotal.Signal(ts)     // every buy/sell/hold signal generator
out  = gotal.RunAll(ts)     // everything — produces 80+ output fields
```

## Examples

The [`examples/`](examples) directory contains runnable programs:

- `examples/indicators/<name>/` — one runnable program per implemented indicator (SMA, EMA, RSI, MACD, BBANDS, MFI, signals, …).
- `examples/groups/<name>/` — one program per group runner (`overlap`, `momentum`, `volume`, `volatility`, `price`, `statistic`, `signal`, `all`, …).
- `examples/readers/csv/` — read a CSV into a `Ticker` and print it.
- `examples/stream/ema/` — pipe a CSV through a streaming EMA.

Run any of them with `go run ./examples/indicators/rsi`, etc.

## Disclaimer

This software is provided for **educational and research purposes only**. It is not financial advice. Trading and investing involve significant risk of loss — you can lose all of your invested capital, past performance does not guarantee future results, and this software is provided "AS IS" with no warranties. The author(s) accept no liability for any losses. Do your own research and consult qualified professionals before making investment decisions.

## License

GPL-3.0-or-later. See [LICENSE](LICENSE).

## Contributing

Issues and pull requests are welcome. See [CONTRIBUTING.md](.github/CONTRIBUTING.md) if present.

## Author

**Rangertaha** — [rangertaha@gmail.com](mailto:rangertaha@gmail.com)
