package types

// OHLCV represents a single candlestick with Open, High, Low, Close prices and Volume
type OHLCV struct {
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

// TimeSeries represents a series of OHLCV data points
type TimeSeries struct {
	Data []OHLCV
}

// IndicatorResult represents the result of a technical indicator calculation
type IndicatorResult struct {
	Values []float64
	Error  error
}

// MovingAverageType defines the type of moving average calculation
type MovingAverageType int

const (
	// SMA represents Simple Moving Average
	SMA MovingAverageType = iota
	// EMA represents Exponential Moving Average
	EMA
	// WMA represents Weighted Moving Average
	WMA
)

// NewTimeSeries creates a new TimeSeries from a slice of OHLCV data
func NewTimeSeries(data []OHLCV) *TimeSeries {
	return &TimeSeries{
		Data: data,
	}
}

// Length returns the number of data points in the TimeSeries
func (ts *TimeSeries) Length() int {
	return len(ts.Data)
}

// GetClosePrices returns a slice of close prices from the TimeSeries
func (ts *TimeSeries) GetClosePrices() []float64 {
	prices := make([]float64, ts.Length())
	for i, data := range ts.Data {
		prices[i] = data.Close
	}
	return prices
}

// GetHighPrices returns a slice of high prices from the TimeSeries
func (ts *TimeSeries) GetHighPrices() []float64 {
	prices := make([]float64, ts.Length())
	for i, data := range ts.Data {
		prices[i] = data.High
	}
	return prices
}

// GetLowPrices returns a slice of low prices from the TimeSeries
func (ts *TimeSeries) GetLowPrices() []float64 {
	prices := make([]float64, ts.Length())
	for i, data := range ts.Data {
		prices[i] = data.Low
	}
	return prices
}

// GetOpenPrices returns a slice of open prices from the TimeSeries
func (ts *TimeSeries) GetOpenPrices() []float64 {
	prices := make([]float64, ts.Length())
	for i, data := range ts.Data {
		prices[i] = data.Open
	}
	return prices
}

// GetVolumes returns a slice of volumes from the TimeSeries
func (ts *TimeSeries) GetVolumes() []float64 {
	volumes := make([]float64, ts.Length())
	for i, data := range ts.Data {
		volumes[i] = data.Volume
	}
	return volumes
}
