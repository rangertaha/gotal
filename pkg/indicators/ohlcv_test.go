package indicators

import (
	"testing"
	"time"

	"github.com/rangertaha/gotal/internal/pkg/series"
)

func TestNewOHLC(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		input    *series.Series
		output   *series.Series
	}{
		{
			name:     "prices to ohlc",
			duration: 4 * time.Minute,
			input: series.New("prices", series.WithFields([]map[string]float64{
				{"time": 10.0, "price": 10.0},
				{"time": 15.0, "price": 15.0},
				{"time": 5.0, "price": 5.0},
				{"time": 20.0, "price": 20.0},
				{"time": 12.0, "price": 12.0},
				{"time": 18.0, "price": 18.0},
				{"time": 13.0, "price": 13.0},
				{"time": 17.0, "price": 17.0},
			})),
			output: series.New("ohlc", series.WithFields([]map[string]float64{
				{"time": 10.0, "open": 10.0, "high": 10.0, "low": 10.0, "close": 10.0},
				{"time": 15.0, "open": 15.0, "high": 15.0, "low": 15.0, "close": 15.0},
				{"time": 5.0, "open": 5.0, "high": 5.0, "low": 5.0, "close": 5.0},
				{"time": 20.0, "open": 20.0, "high": 20.0, "low": 20.0, "close": 20.0},
				{"time": 12.0, "open": 12.0, "high": 12.0, "low": 12.0, "close": 12.0},
				{"time": 18.0, "open": 18.0, "high": 18.0, "low": 18.0, "close": 18.0},
				{"time": 13.0, "open": 13.0, "high": 13.0, "low": 13.0, "close": 13.0},
				{"time": 17.0, "open": 17.0, "high": 17.0, "low": 17.0, "close": 17.0},
			})),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			bar := OHLC(tt.input, WithDuration(tt.duration), OnField("price"))

			if bar.Name() != "ohlc" {
				t.Errorf("expected name %s, got %s", "prices", bar.Name())
			}

			if bar.Duration() != tt.duration {
				t.Errorf("expected duration %d, got %d", tt.duration, bar.Duration())
			}

			if !bar.HasField("open") {
				t.Errorf("expected field 'open' not found")
			}

			if !bar.HasField("high") {
				t.Errorf("expected field 'high' not found")
			}

			if !bar.HasField("low") {
				t.Errorf("expected field 'low' not found")
			}

			if !bar.HasField("close") {
				t.Errorf("expected field 'close' not found")
			}

			if bar.Len() != tt.input.Len() {
				t.Errorf("expected length %d, got %d", tt.input.Len(), bar.Len())
			}

			// check max, min, first, last values of open, high, low, close price fields
			for _, field := range []string{"open", "high", "low", "close"} {
				if bar.Max(field) != tt.input.Max(field) {
					t.Errorf("expected max %s price %f, got %f", field, tt.input.Max(field), bar.Max(field))
				}

				if bar.Min(field) != tt.input.Min(field) {
					t.Errorf("expected min %s price %f, got %f", field, tt.input.Min(field), bar.Min(field))
				}

				if bar.Last(field) != tt.input.Last(field) {
					t.Errorf("expected last %s price %f, got %f", field, tt.input.Last(field), bar.Last(field))
				}

				if bar.First(field) != tt.input.First(field) {
					t.Errorf("expected first %s price %f, got %f", field, tt.input.First(field), bar.First(field))
				}
			}
		})
	}
}


func TestNewOHLCV(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		input    *series.Series
		output   *series.Series
	}{
		{
			name:     "prices to ohlcv",
			duration: 4 * time.Minute,
			input: series.New("prices", series.WithFields([]map[string]float64{
				{"time": 10.0, "price": 10.0, "vol": 10.0},
				{"time": 15.0, "price": 15.0, "vol": 15.0},
				{"time": 5.0, "price": 5.0, "vol": 5.0},
				{"time": 20.0, "price": 20.0, "vol": 20.0},
				{"time": 12.0, "price": 12.0, "vol": 12.0},
				{"time": 18.0, "price": 18.0, "vol": 18.0},
				{"time": 13.0, "price": 13.0, "vol": 13.0},
				{"time": 17.0, "price": 17.0, "vol": 17.0},
			})),
			output: series.New("ohlcv", series.WithFields([]map[string]float64{
				{"time": 10.0, "open": 10.0, "high": 10.0, "low": 10.0, "close": 10.0, "vol": 10.0},
				{"time": 15.0, "open": 15.0, "high": 15.0, "low": 15.0, "close": 15.0, "vol": 15.0},
				{"time": 5.0, "open": 5.0, "high": 5.0, "low": 5.0, "close": 5.0, "vol": 5.0},
				{"time": 20.0, "open": 20.0, "high": 20.0, "low": 20.0, "close": 20.0, "vol": 20.0},
				{"time": 12.0, "open": 12.0, "high": 12.0, "low": 12.0, "close": 12.0, "vol": 12.0},
				{"time": 18.0, "open": 18.0, "high": 18.0, "low": 18.0, "close": 18.0, "vol": 18.0},
				{"time": 13.0, "open": 13.0, "high": 13.0, "low": 13.0, "close": 13.0, "vol": 13.0},
				{"time": 17.0, "open": 17.0, "high": 17.0, "low": 17.0, "close": 17.0, "vol": 17.0},
			})),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			bar := OHLCV(tt.input, WithDuration(tt.duration), OnField("price"))

			if bar.Name() != "ohlcv" {
				t.Errorf("expected name %s, got %s", "prices", bar.Name())
			}

			if bar.Duration() != tt.duration {
				t.Errorf("expected duration %d, got %d", tt.duration, bar.Duration())
			}

			if !bar.HasField("open") {
				t.Errorf("expected field 'open' not found")
			}

			if !bar.HasField("high") {
				t.Errorf("expected field 'high' not found")
			}

			if !bar.HasField("low") {
				t.Errorf("expected field 'low' not found")
			}

			if !bar.HasField("close") {
				t.Errorf("expected field 'close' not found")
			}

			if !bar.HasField("vol") {
				t.Errorf("expected field 'vol' not found")
			}

			if bar.Len() != tt.input.Len() {
				t.Errorf("expected length %d, got %d", tt.input.Len(), bar.Len())
			}

			// check max, min, first, last values of open, high, low, close price fields
			for _, field := range []string{"open", "high", "low", "close", "vol"} {
				if bar.Max(field) != tt.input.Max(field) {
					t.Errorf("expected max %s price %f, got %f", field, tt.input.Max(field), bar.Max(field))
				}

				if bar.Min(field) != tt.input.Min(field) {
					t.Errorf("expected min %s price %f, got %f", field, tt.input.Min(field), bar.Min(field))
				}

				if bar.Last(field) != tt.input.Last(field) {
					t.Errorf("expected last %s price %f, got %f", field, tt.input.Last(field), bar.Last(field))
				}

				if bar.First(field) != tt.input.First(field) {
					t.Errorf("expected first %s price %f, got %f", field, tt.input.First(field), bar.First(field))
				}
			}
		})
	}
}
