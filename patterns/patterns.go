package patterns

// Candle represents a single candlestick
type Candle struct {
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

// PatternResult contains the pattern recognition results
type PatternResult struct {
	Pattern     string
	Confidence  float64
	Description string
}
