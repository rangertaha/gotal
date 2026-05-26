# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- **Public API redesign.** Four-package layout: root `gotal` for shared types and indicator vars; `ticker` for batch `TimeSeries`; `stream` for channel-based live pipelines; `io` for readers (CSV/JSON/JSONL).
- **Standard buy/sell/hold signal generators** (TA-Lib `+100`/`-100`/`0` convention):
  - Threshold-cross on bounded oscillators: `RSISIGNAL`, `MFISIGNAL`, `CCISIGNAL`, `WILLRSIGNAL`.
  - Zero-cross on unbounded oscillators: `MOMSIGNAL`, `ROCSIGNAL`, `TRIXSIGNAL`, `PPOSIGNAL`, `CMOSIGNAL`.
  - Line crossovers: `MACDSIGNAL_CROSS`, `STOCHSIGNAL`.
  - New `SIGNAL` registry group + `gotal.Signal(ts)` group runner.
  - Reusable building blocks in `internal/indicators/util`: `ThresholdCross`, `ZeroCross`, `LineCrossover`.
- **Per-group runner functions** (`gotal.Overlap`, `gotal.Momentum`, `gotal.Volume`, `gotal.Volatility`, `gotal.Price`, `gotal.Statistic`, `gotal.MathOps`, `gotal.MathTransforms`, `gotal.Signal`, `gotal.Cycle`, `gotal.Pattern`) and a master `gotal.RunAll(ts)` that applies every implemented indicator.
- **Registry introspection**: `gotal.Groups()`, `gotal.Group(name)`, `gotal.List()`, `gotal.Has(id)`, `gotal.Get(id)`.
- **Indicator catalogue.** Registered every TA-Lib function id plus the extended Pandas TA / community catalogue (~330 indicators) — real implementations win over stubs regardless of init order; unimplemented ids return a clear "not implemented" error.
- **Real indicator implementations:**
  - *Overlap*: SMA, EMA, WMA, DEMA, TEMA, TRIMA, MIDPOINT, MIDPRICE, BBANDS, HMA.
  - *Momentum*: MOM, ROC, ROCP, ROCR, ROCR100, WILLR, RSI, CCI, MACD, APO, PPO, CMO, TRIX, AROON, AROONOSC, STOCH, STOCHF.
  - *Volatility*: TRANGE, ATR, NATR.
  - *Volume*: OBV, AD, ADOSC, MFI, VWMA, VWAP.
  - *Price transforms*: AVGPRICE, MEDPRICE, TYPPRICE, WCLPRICE, HLC3, OHLC4, HEIKINASHI.
  - *Statistic*: STDDEV, VARIANCE, LINEARREG, LINEARREG_SLOPE, LINEARREG_ANGLE, LINEARREG_INTERCEPT, TSF, BETA, CORREL.
  - *Math operators*: MAX, MIN, SUMWINDOW, ADD, SUB, MULT, DIV, MAXINDEX, MININDEX.
  - *Math transforms*: ACOS, ASIN, ATAN, COS, SIN, TAN, COSH, SINH, TANH, CEIL, FLOOR, EXP, LN, LOG10, SQRT.
- **Streaming.** Channel-backed `*stream.Stream`, `stream.Apply(in, indicator)`, and `stream.EMA/SMA/Ema/Sma` wrappers. Streaming EMA values match batch EMA on the same data.
- **Tests.** Table-driven `*_test.go` in package `gotal_test` covering every implemented indicator (31 passing tests).
- **Examples.** One runnable program per implemented indicator under `examples/indicators/<name>/`, plus CSV reader and streaming examples.
- `gotal.Get(id)`, `gotal.Has(id)`, `gotal.List()` for runtime indicator discovery.
- `gotal.With(key, value)` config-option helper at the root.
- Per-tick `tick.WithSignals` helper for stream indicators to emit signals on output ticks.
- `gotal.Plot` interface and `ticker.Ticker.Plot()` returning a gonum-backed implementation under `internal/plot`.
- Public `io.Source` interface with `Ticks()` (batch) and `Stream()` (channel); `io.Read(path)` dispatches by file extension.
- Vector interface gained `Values() []float64` and `Mean()`.

### Changed
- Canonical interfaces (`Tick`, `TimeSeries`, `Stream`, `Indicator`, `Configurator`, `ConfigOption`, `IndicatorFunc`, `Plot`, `Fields`, `Tags`, `Vector`) now live in `internal/` and are re-exported as type aliases from root `gotal`, avoiding import cycles between root and `internal/indicators`.
- `*ticker.Ticker` now implements `gotal.TimeSeries`, including `Print`, `Save` (CSV), and `Plot` methods.
- Registry tracks stub vs real entries; real `Add` cleanly replaces a stub for the same id.

### Removed
- Dead packages: `internal/series`, `internal/stream`, `internal/opt`, `internal/io`.
- Stale `internal/indicators/functions.go` (all commented out).
- `internal/tick/options.go` (`TickOptions` had no callers).
- `examples/config/` placeholder.
- Broken `ticker.File` and `ticker.WithIndicator` options (readers now live in `io/`; indicator wiring in `stream/`).
- `scripts/banner.sh`, `scripts/version.sh`.

### Fixed
- `stream/indicators.go` declared `package gotal` (collided with root) and referenced the deleted `gotal.Series` type.
- `examples/indicators/ema/main.go` and `examples/readers/csv/main.go` imported types that no longer existed.
- `internal/indicators/registry.go` and `internal/config/options.go` referenced the deleted `internal.Series` type.
- `internal/plot/plot.go` used methods (`HasField`, `Epock`, `FieldNames`) that no longer existed on the new interfaces.

## [0.0.0] - 2025-10-06
### Initial Code
- Initial code
