package sig

type (
	Signal   int
	Strength int
)

const (
	WEAK Strength = iota
	MEDIUM
	STRONG
)

const (
	// Actions
	BUY Signal = iota
	SELL
	HOLD

	// Indicators
	BULLISH
	BEARISH
	NEUTRAL
	CONVERGING
	DIVERGING
	CROSSOVER
	CROSSUNDER
)

func (s Signal) String() string {
	return []string{
		// Actions
		"BUY",
		"SELL",
		"HOLD",

		// Indicators
		"BULLISH",    // Bullish trend
		"BEARISH",    // Bearish trend
		"NEUTRAL",    // Neutral trend
		"CONVERGING", // Converging: two or more indicators are moving closer together
		"DIVERGING",  // Diverging: two or more indicators are moving further apart
		"CROSSOVER",  // Crossover: one indicator crosses over another
		"CROSSUNDER", // Crossunder: one indicator crosses under another

	}[s]
}

func (s Signal) Int() int {
	return int(s)
}
