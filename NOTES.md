# Indicator Catalogue

A reference of every indicator gotal registers, mirroring the [TA-Lib functions list](https://ta-lib.org/functions/) plus the extended Pandas TA / community catalogue from [tamcp](https://github.com/rangertaha/tamcp/tree/main/internal/tools/indicators). Every listed id is callable via `gotal.<NAME>` or `gotal.Get("name")`; unimplemented ids return a `"not implemented"` error at call time.

Status legend:
- ✅ implemented (batch `Compute`)
- 🟡 stub (registered, returns "not implemented")

## Overlap Studies

| ID | Status | Notes |
| --- | --- | --- |
| SMA | ✅ | period |
| EMA | ✅ | period, alpha (also streaming) |
| WMA | ✅ | period |
| DEMA | ✅ | period |
| TEMA | ✅ | period |
| TRIMA | ✅ | period |
| MIDPOINT | ✅ | period (over close) |
| MIDPRICE | ✅ | period (over high/low) |
| BBANDS | ✅ | period, nbdevup, nbdevdn — outputs `_upper`, `_middle`, `_lower` |
| HMA | ✅ | period (Hull Moving Average) |
| T3 | 🟡 | |
| KAMA | 🟡 | |
| MAMA | 🟡 | |
| MA | 🟡 | dispatcher |
| MAVP | 🟡 | |
| SAR | 🟡 | |
| SAREXT | 🟡 | |
| HT_TRENDLINE | 🟡 | |

## Momentum

| ID | Status | Notes |
| --- | --- | --- |
| MOM | ✅ | period |
| ROC, ROCP, ROCR, ROCR100 | ✅ | period |
| WILLR | ✅ | period (OHLC) |
| RSI | ✅ | period (Wilder's smoothing) |
| CCI | ✅ | period (OHLC) |
| MACD | ✅ | fast, slow, signal — outputs `macd`, `macd_signal`, `macd_hist` |
| APO | ✅ | fast, slow |
| PPO | ✅ | fast, slow |
| CMO | ✅ | period |
| TRIX | ✅ | period |
| AROON | ✅ | period — outputs `aroon_up`, `aroon_down` |
| AROONOSC | ✅ | period |
| STOCH | ✅ | fastk, slowk, slowd — outputs `_k`, `_d` |
| STOCHF | ✅ | fastk, fastd — outputs `_k`, `_d` |
| MFI | ✅ | period (OHLCV) |
| MACDEXT, MACDFIX | 🟡 | |
| STOCHRSI | 🟡 | |
| ULTOSC, BOP | 🟡 | |
| ADX, ADXR, DX | 🟡 | |
| PLUS_DI, MINUS_DI, PLUS_DM, MINUS_DM | 🟡 | |

## Volatility

| ID | Status | Notes |
| --- | --- | --- |
| TRANGE | ✅ | OHLC |
| ATR | ✅ | period, Wilder's smoothing |
| NATR | ✅ | period |

## Volume

| ID | Status | Notes |
| --- | --- | --- |
| OBV | ✅ | close + volume |
| AD | ✅ | Chaikin A/D line, OHLCV |
| ADOSC | ✅ | fast, slow (OHLCV) |
| VWMA | ✅ | period (close + volume) |
| VWAP | ✅ | session cumulative (OHLCV) |

## Price Transform

| ID | Status | Notes |
| --- | --- | --- |
| AVGPRICE | ✅ | (O+H+L+C)/4 |
| MEDPRICE | ✅ | (H+L)/2 |
| TYPPRICE | ✅ | (H+L+C)/3 |
| WCLPRICE | ✅ | (H+L+2C)/4 |
| HLC3 | ✅ | (H+L+C)/3 |
| OHLC4 | ✅ | (O+H+L+C)/4 |
| HEIKINASHI | ✅ | OHLC — outputs `ha_open`, `ha_high`, `ha_low`, `ha_close` |

## Statistic

| ID | Status | Notes |
| --- | --- | --- |
| STDDEV | ✅ | period, nbdev |
| VARIANCE | ✅ | period (population) |
| LINEARREG | ✅ | period — endpoint of the fit |
| LINEARREG_SLOPE | ✅ | period |
| LINEARREG_ANGLE | ✅ | period — degrees |
| LINEARREG_INTERCEPT | ✅ | period |
| TSF | ✅ | period — one-step-ahead forecast |
| BETA | ✅ | source1, source2, period |
| CORREL | ✅ | source1, source2, period |

## Signal

Buy/sell/hold signal generators. Each emits `+100` (buy), `-100` (sell), or `0` (hold) per tick — the same convention TA-Lib uses for candlestick patterns.

| ID | Status | Notes |
| --- | --- | --- |
| RSISIGNAL | ✅ | RSI crosses 30 / 70 |
| MFISIGNAL | ✅ | MFI crosses 20 / 80 |
| CCISIGNAL | ✅ | CCI crosses ±100 |
| WILLRSIGNAL | ✅ | Williams %R crosses -80 / -20 |
| MOMSIGNAL | ✅ | MOM zero-cross |
| ROCSIGNAL | ✅ | ROC zero-cross |
| TRIXSIGNAL | ✅ | TRIX zero-cross |
| PPOSIGNAL | ✅ | PPO zero-cross |
| CMOSIGNAL | ✅ | CMO zero-cross |
| MACDSIGNAL_CROSS | ✅ | MACD line crosses signal line |
| STOCHSIGNAL | ✅ | %K crosses %D |
| ADXSIGNAL | 🟡 | depends on ADX (still stubbed) |

## Cycle

| ID | Status | Notes |
| --- | --- | --- |
| HT_DCPERIOD, HT_DCPHASE, HT_PHASOR, HT_SINE, HT_TRENDMODE | 🟡 | Hilbert transforms |

## Math Operators

| ID | Status | Notes |
| --- | --- | --- |
| MAX, MIN, SUMWINDOW | ✅ | period |
| ADD, SUB, MULT, DIV | ✅ | source1, source2 (binary) |
| MAXINDEX, MININDEX | ✅ | period — returns the position of the extremum |
| MINMAX, MINMAXINDEX | 🟡 | |

## Math Transforms

| ID | Status |
| --- | --- |
| ACOS, ASIN, ATAN | ✅ |
| COS, SIN, TAN | ✅ |
| COSH, SINH, TANH | ✅ |
| CEIL, FLOOR | ✅ |
| EXP, LN, LOG10, SQRT | ✅ |

## Pattern Recognition (Candlestick)

All 61 `CDL*` patterns are registered as stubs (🟡). They return "not implemented" until candlestick recognition is wired up.

## Extended community catalogue

The following non-TA-Lib indicators (~170) from the Pandas TA / community libraries are also registered as stubs and ready to be filled in: ABERRATION, ACCBANDS, ADPCT, ADSMOOTH, ADZSCORE, ALLIGATOR, ALMA, AO, AOACC, AOBV, ATRBANDS, BBP, BBPSIGNAL, BBSQUEEZE, BBW, BIAS, BRAR, BSTS, CAMARILLA, CFO, CHAIKINVOL, CHANDEXIT, CHOP, CKSP, CMF, CMFSIGNAL, COPPOCK, CPR, CRSI, CSI, CTI, CVD, CYBERCYCLE, DECAY, DECREASING, DECYCLER, DEM, DEMARKPIV, DONCHIAN, DONCHIANPCT, DPO, EFI, EMADIFF, EMAENV, ENTROPY, ENVELOPE, EOM, ER, ERI, FCB, FIBPIVOTS, FISHER, FRACTAL, FRAMA, FVE, FWMA, GARCH, GATOR, GAUSSIAN, GD, HILO, HMM, HWC, HWMA, ICHIMOKU, IFISHER, INCREASING, INERTIA, KC, KCB, KDJ, KST, KURTOSIS, KVO, KVOPCT, LINREGRESID, LOGRET, LSTM, MACDHIST, MACDPCT, MACDV, MACDZL, MACDZLHIST, MAD, MANSFIELD, MASSI, MCGINLEY, MEDIAN, MEDSMOOTH, MFISMOOTH, MFV, NVI, OBVSIGNAL, OBVSMOOTH, OBVZSCORE, PCTRET, PDIST, PGO, PSL, PVI, PVR, PVT, PWMA, PWO, QQE, QUANTILE, RANGEPCT, RATRPCT, REFLEX, RETZSCORE, RMI, ROCS, RSISMOOTH, RVI, SINWMA, SKEW, SLOPEPCT, SMI, SMMA, SQUEEZE, SQUEEZEPRO, SSF, STC, STOCHDIFF, SUPERTREND, SWINGINDEX, SWMA, TDI, THERMO, TII, TMF, TPSMOOTH, TRANGESMOOTH, TRENDFLEX, TRENDSCORE, TRPCT, TSI, TTMTREND, TVI, ULCER, VHF, VIDYA, VO, VOLRATIO, VORTEX, VORTEXDIFF, VSTOP, VWAPANCH, VWAPPCT, WAD, WCPSMOOTH, WMADIFF, WMAENV, WOODIE, WT, ZLEMA, ZLHMA, ZSCORE.

## Adding a real implementation

1. Implement the `internal.Indicator` interface in `internal/indicators/<group>/`. Reuse helpers in `internal/indicators/util/`.
2. Register it from `init()`:
   ```go
   util.Must(indicators.Add("RSI", newRSI, indicators.MOMENTUM))
   ```
3. Add a blank import in `internal/indicators/all/all.go` (under "real implementations") if the package isn't already imported. The registry guarantees a real `Add` cleanly replaces any prior stub.
4. Add a typed convenience wrapper in `indicators.go` at the repo root (e.g. `Rsi(ts, 14)`).
5. Add a test in the same package directory as the implementation (e.g. `internal/indicators/momentum/momentum_test.go`, package `momentum_test`). Use the helpers in `internal/indicators/testutil`.
6. Add a runnable example under `examples/indicators/<name>/main.go`.
