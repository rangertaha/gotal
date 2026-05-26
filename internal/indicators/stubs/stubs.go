// Package stubs registers every TA-Lib and tamcp-extended indicator ID as a
// placeholder. Real implementations in sibling packages register the same id
// first (init order from all/all.go), so RegisterStub silently no-ops on
// already-registered names. This guarantees gotal.<NAME> returns a clear
// "not implemented" error for ids without real math yet, instead of a panic.
package stubs

import "github.com/rangertaha/gotal/internal/indicators"

func init() {
	stub := indicators.RegisterStub

	// CYCLE
	stub("HT_DCPERIOD", indicators.CYCLE)
	stub("HT_DCPHASE", indicators.CYCLE)
	stub("HT_PHASOR", indicators.CYCLE)
	stub("HT_SINE", indicators.CYCLE)
	stub("HT_TRENDMODE", indicators.CYCLE)

	// MATH OPERATORS
	for _, id := range []string{"ADD", "SUB", "MULT", "DIV", "MAX", "MIN", "MAXINDEX", "MININDEX", "MINMAX", "MINMAXINDEX", "SUMWINDOW"} {
		stub(id, indicators.MATHOP)
	}

	// MATH TRANSFORMS
	for _, id := range []string{"ACOS", "ASIN", "ATAN", "CEIL", "COS", "COSH", "EXP", "FLOOR", "LN", "LOG10", "SIN", "SINH", "SQRT", "TAN", "TANH"} {
		stub(id, indicators.MATHTRANSFORM)
	}

	// MOMENTUM
	for _, id := range []string{"ADX", "ADXR", "APO", "AROON", "AROONOSC", "BOP", "CCI", "CMO", "DX", "MACD", "MACDEXT", "MACDFIX", "MFI", "MINUS_DI", "MINUS_DM", "MOM", "PLUS_DI", "PLUS_DM", "PPO", "ROC", "ROCP", "ROCR", "ROCR100", "RSI", "STOCH", "STOCHF", "STOCHRSI", "TRIX", "ULTOSC", "WILLR"} {
		stub(id, indicators.MOMENTUM)
	}

	// OVERLAP
	for _, id := range []string{"BBANDS", "DEMA", "EMA", "HT_TRENDLINE", "KAMA", "MA", "MAMA", "MAVP", "MIDPOINT", "MIDPRICE", "SAR", "SAREXT", "SMA", "T3", "TEMA", "TRIMA", "WMA"} {
		stub(id, indicators.OVERLAP)
	}

	// CANDLESTICK PATTERNS
	for _, id := range []string{
		"CDL2CROWS", "CDL3BLACKCROWS", "CDL3INSIDE", "CDL3LINESTRIKE", "CDL3OUTSIDE",
		"CDL3STARSINSOUTH", "CDL3WHITESOLDIERS", "CDLABANDONEDBABY", "CDLADVANCEBLOCK",
		"CDLBELTHOLD", "CDLBREAKAWAY", "CDLCLOSINGMARUBOZU", "CDLCONCEALBABYSWALL",
		"CDLCOUNTERATTACK", "CDLDARKCLOUDCOVER", "CDLDOJI", "CDLDOJISTAR", "CDLDRAGONFLYDOJI",
		"CDLENGULFING", "CDLEVENINGDOJISTAR", "CDLEVENINGSTAR", "CDLGAPSIDESIDEWHITE",
		"CDLGRAVESTONEDOJI", "CDLHAMMER", "CDLHANGINGMAN", "CDLHARAMI", "CDLHARAMICROSS",
		"CDLHIGHWAVE", "CDLHIKKAKE", "CDLHIKKAKEMOD", "CDLHOMINGPIGEON", "CDLIDENTICAL3CROWS",
		"CDLINNECK", "CDLINVERTEDHAMMER", "CDLKICKING", "CDLKICKINGBYLENGTH", "CDLLADDERBOTTOM",
		"CDLLONGLEGGEDDOJI", "CDLLONGLINE", "CDLMARUBOZU", "CDLMATCHINGLOW", "CDLMATHOLD",
		"CDLMORNINGDOJISTAR", "CDLMORNINGSTAR", "CDLONNECK", "CDLPIERCING", "CDLRICKSHAWMAN",
		"CDLRISEFALL3METHODS", "CDLSEPARATINGLINES", "CDLSHOOTINGSTAR", "CDLSHORTLINE",
		"CDLSPINNINGTOP", "CDLSTALLEDPATTERN", "CDLSTICKSANDWICH", "CDLTAKURI", "CDLTASUKIGAP",
		"CDLTHRUSTING", "CDLTRISTAR", "CDLUNIQUE3RIVER", "CDLUPSIDEGAP2CROWS", "CDLXSIDEGAP3METHODS",
	} {
		stub(id, indicators.PATTERN)
	}

	// PRICE TRANSFORMS
	for _, id := range []string{"AVGPRICE", "MEDPRICE", "TYPPRICE", "WCLPRICE"} {
		stub(id, indicators.PRICE)
	}

	// STATISTIC
	for _, id := range []string{"BETA", "CORREL", "LINEARREG", "LINEARREG_ANGLE", "LINEARREG_INTERCEPT", "LINEARREG_SLOPE", "STDDEV", "TSF", "VARIANCE"} {
		stub(id, indicators.STATISTIC)
	}

	// VOLATILITY
	for _, id := range []string{"ATR", "NATR", "TRANGE"} {
		stub(id, indicators.VOLATILITY)
	}

	// VOLUME
	for _, id := range []string{"AD", "ADOSC", "OBV"} {
		stub(id, indicators.VOLUME)
	}

	// Extended indicators (Pandas TA / community library). Grouped as OTHER for now.
	for _, id := range []string{
		"ABERRATION", "ACCBANDS", "ADPCT", "ADSMOOTH", "ADXSIGNAL", "ADZSCORE",
		"ALLIGATOR", "ALMA", "AO", "AOACC", "AOBV", "ATRBANDS", "BBP", "BBPSIGNAL",
		"BBSQUEEZE", "BBW", "BIAS", "BRAR", "BSTS", "CAMARILLA", "CCISIGNAL", "CFO",
		"CHAIKINVOL", "CHANDEXIT", "CHOP", "CKSP", "CMF", "CMFSIGNAL", "COPPOCK",
		"CPR", "CRSI", "CSI", "CTI", "CVD", "CYBERCYCLE", "DECAY", "DECREASING",
		"DECYCLER", "DEM", "DEMARKPIV", "DONCHIAN", "DONCHIANPCT", "DPO", "EFI",
		"EMADIFF", "EMAENV", "ENTROPY", "ENVELOPE", "EOM", "ER", "ERI", "FCB",
		"FIBPIVOTS", "FISHER", "FRACTAL", "FRAMA", "FVE", "FWMA", "GARCH", "GATOR",
		"GAUSSIAN", "GD", "HEIKINASHI", "HILO", "HLC3", "HMA", "HMM", "HWC", "HWMA",
		"ICHIMOKU", "IFISHER", "INCREASING", "INERTIA", "KC", "KCB", "KDJ", "KST",
		"KURTOSIS", "KVO", "KVOPCT", "LINREGRESID", "LOGRET", "LSTM", "MACDHIST",
		"MACDPCT", "MACDV", "MACDZL", "MACDZLHIST", "MAD", "MANSFIELD", "MASSI",
		"MCGINLEY", "MEDIAN", "MEDSMOOTH", "MFISIGNAL", "MFISMOOTH", "MFV",
		"MOMSIGNAL", "NVI", "OBVSIGNAL", "OBVSMOOTH", "OBVZSCORE", "OHLC4", "PCTRET",
		"PDIST", "PGO", "PPOSIGNAL", "PSL", "PVI", "PVR", "PVT", "PWMA", "PWO", "QQE",
		"QUANTILE", "RANGEPCT", "RATRPCT", "REFLEX", "RETZSCORE", "RMI", "ROCS",
		"ROCSIGNAL", "RSISMOOTH", "RVI", "SINWMA", "SKEW", "SLOPEPCT", "SMI", "SMMA",
		"SQUEEZE", "SQUEEZEPRO", "SSF", "STC", "STOCHDIFF", "SUPERTREND", "SWINGINDEX",
		"SWMA", "TDI", "THERMO", "TII", "TMF", "TPSMOOTH", "TRANGESMOOTH", "TRENDFLEX",
		"TRENDSCORE", "TRIXSIGNAL", "TRPCT", "TSI", "TTMTREND", "TVI", "ULCER", "VHF",
		"VIDYA", "VO", "VOLRATIO", "VORTEX", "VORTEXDIFF", "VSTOP", "VWAP", "VWAPANCH",
		"VWAPPCT", "VWMA", "WAD", "WCPSMOOTH", "WILLRSIGNAL", "WMADIFF", "WMAENV",
		"WOODIE", "WT", "ZLEMA", "ZLHMA", "ZSCORE",
	} {
		stub(id, indicators.OTHER)
	}
}
