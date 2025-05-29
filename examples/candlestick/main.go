package main

import (
	"fmt"

	"github.com/rangertaha/gotal/patterns"
)

func main() {
	// Sample candlestick data
	candles := []patterns.Candle{
		{Open: 100, High: 105, Low: 98, Close: 103, Volume: 1000}, // Bullish candle
		{Open: 103, High: 104, Low: 102, Close: 102, Volume: 800}, // Bearish candle
		{Open: 102, High: 102, Low: 98, Close: 98, Volume: 1200},  // Bearish candle
		{Open: 98, High: 98, Low: 95, Close: 95, Volume: 1500},    // Bearish candle
		{Open: 95, High: 95, Low: 92, Close: 92, Volume: 2000},    // Bearish candle
		{Open: 92, High: 92, Low: 90, Close: 90, Volume: 1800},    // Bearish candle
		{Open: 90, High: 90, Low: 88, Close: 88, Volume: 1600},    // Bearish candle
		{Open: 88, High: 88, Low: 85, Close: 85, Volume: 1400},    // Bearish candle
		{Open: 85, High: 85, Low: 82, Close: 82, Volume: 1300},    // Bearish candle
		{Open: 82, High: 82, Low: 80, Close: 80, Volume: 1100},    // Bearish candle
		{Open: 80, High: 80, Low: 78, Close: 78, Volume: 1000},    // Bearish candle
		{Open: 78, High: 78, Low: 75, Close: 75, Volume: 900},     // Bearish candle
		{Open: 75, High: 75, Low: 72, Close: 72, Volume: 800},     // Bearish candle
		{Open: 72, High: 72, Low: 70, Close: 70, Volume: 700},     // Bearish candle
		{Open: 70, High: 70, Low: 68, Close: 68, Volume: 600},     // Bearish candle
		{Open: 68, High: 68, Low: 65, Close: 65, Volume: 500},     // Bearish candle
		{Open: 65, High: 65, Low: 62, Close: 62, Volume: 400},     // Bearish candle
		{Open: 62, High: 62, Low: 60, Close: 60, Volume: 300},     // Bearish candle
		{Open: 60, High: 60, Low: 58, Close: 58, Volume: 200},     // Bearish candle
		{Open: 58, High: 58, Low: 55, Close: 55, Volume: 100},     // Bearish candle
		{Open: 55, High: 55, Low: 52, Close: 52, Volume: 50},      // Bearish candle
		{Open: 52, High: 52, Low: 50, Close: 50, Volume: 25},      // Bearish candle
		{Open: 50, High: 50, Low: 48, Close: 48, Volume: 10},      // Bearish candle
		{Open: 48, High: 48, Low: 45, Close: 45, Volume: 5},       // Bearish candle
		{Open: 45, High: 45, Low: 42, Close: 42, Volume: 2},       // Bearish candle
		{Open: 42, High: 42, Low: 40, Close: 40, Volume: 1},       // Bearish candle
		{Open: 40, High: 40, Low: 38, Close: 38, Volume: 1},       // Bearish candle
		{Open: 38, High: 38, Low: 35, Close: 35, Volume: 1},       // Bearish candle
		{Open: 35, High: 35, Low: 32, Close: 32, Volume: 1},       // Bearish candle
		{Open: 32, High: 32, Low: 30, Close: 30, Volume: 1},       // Bearish candle
		{Open: 30, High: 30, Low: 28, Close: 28, Volume: 1},       // Bearish candle
		{Open: 28, High: 28, Low: 25, Close: 25, Volume: 1},       // Bearish candle
		{Open: 25, High: 25, Low: 22, Close: 22, Volume: 1},       // Bearish candle
		{Open: 22, High: 22, Low: 20, Close: 20, Volume: 1},       // Bearish candle
		{Open: 20, High: 20, Low: 18, Close: 18, Volume: 1},       // Bearish candle
		{Open: 18, High: 18, Low: 15, Close: 15, Volume: 1},       // Bearish candle
		{Open: 15, High: 15, Low: 12, Close: 12, Volume: 1},       // Bearish candle
		{Open: 12, High: 12, Low: 10, Close: 10, Volume: 1},       // Bearish candle
		{Open: 10, High: 10, Low: 8, Close: 8, Volume: 1},         // Bearish candle
		{Open: 8, High: 8, Low: 5, Close: 5, Volume: 1},           // Bearish candle
		{Open: 5, High: 5, Low: 2, Close: 2, Volume: 1},           // Bearish candle
		{Open: 2, High: 2, Low: 0, Close: 0, Volume: 1},           // Bearish candle
	}

	// Add some pattern-specific candles
	candles = append(candles, []patterns.Candle{
		{Open: 50, High: 55, Low: 48, Close: 52, Volume: 1000}, // Bullish candle
		{Open: 52, High: 54, Low: 51, Close: 53, Volume: 1200}, // Bullish candle
		{Open: 53, High: 56, Low: 52, Close: 55, Volume: 1500}, // Bullish candle (Three White Soldiers)
		{Open: 55, High: 56, Low: 54, Close: 54, Volume: 800},  // Bearish candle
		{Open: 54, High: 55, Low: 53, Close: 53, Volume: 600},  // Bearish candle
		{Open: 53, High: 54, Low: 52, Close: 52, Volume: 400},  // Bearish candle (Three Black Crows)
		{Open: 52, High: 53, Low: 51, Close: 51, Volume: 300},  // Bearish candle
		{Open: 51, High: 52, Low: 50, Close: 50, Volume: 200},  // Bearish candle
		{Open: 50, High: 51, Low: 49, Close: 49, Volume: 100},  // Bearish candle
		{Open: 49, High: 50, Low: 48, Close: 48, Volume: 50},   // Bearish candle
		{Open: 48, High: 49, Low: 47, Close: 47, Volume: 25},   // Bearish candle
		{Open: 47, High: 48, Low: 46, Close: 46, Volume: 10},   // Bearish candle
		{Open: 46, High: 47, Low: 45, Close: 45, Volume: 5},    // Bearish candle
		{Open: 45, High: 46, Low: 44, Close: 44, Volume: 2},    // Bearish candle
		{Open: 44, High: 45, Low: 43, Close: 43, Volume: 1},    // Bearish candle
	}...)

	// Add more pattern-specific candles
	candles = append(candles, []patterns.Candle{
		{Open: 40, High: 42, Low: 38, Close: 41, Volume: 1000}, // Bullish candle
		{Open: 41, High: 43, Low: 40, Close: 40, Volume: 800},  // Bearish candle (Dark Cloud Cover)
		{Open: 40, High: 41, Low: 38, Close: 39, Volume: 600},  // Bearish candle
		{Open: 39, High: 40, Low: 37, Close: 38, Volume: 400},  // Bearish candle
		{Open: 38, High: 39, Low: 36, Close: 37, Volume: 300},  // Bearish candle
		{Open: 37, High: 38, Low: 35, Close: 36, Volume: 200},  // Bearish candle
		{Open: 36, High: 37, Low: 34, Close: 35, Volume: 100},  // Bearish candle
		{Open: 35, High: 36, Low: 33, Close: 34, Volume: 50},   // Bearish candle
		{Open: 34, High: 35, Low: 32, Close: 33, Volume: 25},   // Bearish candle
		{Open: 33, High: 34, Low: 31, Close: 32, Volume: 10},   // Bearish candle
		{Open: 32, High: 33, Low: 30, Close: 31, Volume: 5},    // Bearish candle
		{Open: 31, High: 32, Low: 29, Close: 30, Volume: 2},    // Bearish candle
		{Open: 30, High: 31, Low: 28, Close: 29, Volume: 1},    // Bearish candle
		{Open: 29, High: 30, Low: 27, Close: 28, Volume: 1},    // Bearish candle
		{Open: 28, High: 29, Low: 26, Close: 27, Volume: 1},    // Bearish candle
	}...)

	// Add even more pattern-specific candles
	candles = append(candles, []patterns.Candle{
		{Open: 20, High: 22, Low: 18, Close: 19, Volume: 1000}, // Bearish candle
		{Open: 19, High: 21, Low: 17, Close: 20, Volume: 800},  // Bullish candle (Piercing Pattern)
		{Open: 20, High: 22, Low: 19, Close: 21, Volume: 600},  // Bullish candle
		{Open: 21, High: 23, Low: 20, Close: 22, Volume: 400},  // Bullish candle
		{Open: 22, High: 24, Low: 21, Close: 23, Volume: 300},  // Bullish candle
		{Open: 23, High: 25, Low: 22, Close: 24, Volume: 200},  // Bullish candle
		{Open: 24, High: 26, Low: 23, Close: 25, Volume: 100},  // Bullish candle
		{Open: 25, High: 27, Low: 24, Close: 26, Volume: 50},   // Bullish candle
		{Open: 26, High: 28, Low: 25, Close: 27, Volume: 25},   // Bullish candle
		{Open: 27, High: 29, Low: 26, Close: 28, Volume: 10},   // Bullish candle
		{Open: 28, High: 30, Low: 27, Close: 29, Volume: 5},    // Bullish candle
		{Open: 29, High: 31, Low: 28, Close: 30, Volume: 2},    // Bullish candle
		{Open: 30, High: 32, Low: 29, Close: 31, Volume: 1},    // Bullish candle
		{Open: 31, High: 33, Low: 30, Close: 32, Volume: 1},    // Bullish candle
		{Open: 32, High: 34, Low: 31, Close: 33, Volume: 1},    // Bullish candle
	}...)

	// Add pattern-specific candles for new patterns
	candles = append(candles, []patterns.Candle{
		{Open: 1, High: 1.1, Low: 0.8, Close: 1.1, Volume: 1000},  // Dragonfly Doji
		{Open: 1.1, High: 1.2, Low: 1, Close: 1.2, Volume: 800},   // Bullish candle
		{Open: 1.2, High: 1.3, Low: 1.1, Close: 1.3, Volume: 600}, // Bullish candle
		{Open: 1.3, High: 1.6, Low: 1.3, Close: 1.3, Volume: 400}, // Gravestone Doji
		{Open: 1.3, High: 1.4, Low: 1.2, Close: 1.2, Volume: 300}, // Bearish candle
		{Open: 1.2, High: 1.3, Low: 1.1, Close: 1.1, Volume: 200}, // Bearish candle
	}...)

	// Add Three Inside Up/Down pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 0.8, High: 0.9, Low: 0.7, Close: 0.7, Volume: 1000},   // Bearish candle
		{Open: 0.75, High: 0.85, Low: 0.7, Close: 0.8, Volume: 800},  // Bullish candle (contained)
		{Open: 0.8, High: 0.9, Low: 0.75, Close: 0.85, Volume: 600},  // Bullish candle (Three Inside Up)
		{Open: 0.85, High: 0.95, Low: 0.8, Close: 0.9, Volume: 400},  // Bullish candle
		{Open: 0.9, High: 1, Low: 0.85, Close: 0.85, Volume: 300},    // Bearish candle
		{Open: 0.85, High: 0.9, Low: 0.8, Close: 0.8, Volume: 200},   // Bearish candle (contained)
		{Open: 0.8, High: 0.85, Low: 0.75, Close: 0.75, Volume: 100}, // Bearish candle (Three Inside Down)
	}...)

	// Add Kicking pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 5, High: 5, Low: 4.8, Close: 4.8, Volume: 1000},  // Bearish Marubozu
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 800}, // Bullish Marubozu (Bullish Kicking)
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 600}, // Bullish Marubozu
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 400}, // Bullish Marubozu
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 300}, // Bullish Marubozu
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 200}, // Bullish Marubozu
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 100}, // Bullish Marubozu
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 50},  // Bullish Marubozu
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 25},  // Bullish Marubozu
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 10},  // Bullish Marubozu
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 5},   // Bullish Marubozu
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 2},   // Bullish Marubozu
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 1},   // Bullish Marubozu
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 1},   // Bullish Marubozu
		{Open: 5.2, High: 5.2, Low: 5, Close: 5.2, Volume: 1},   // Bullish Marubozu
	}...)

	// Add Abandoned Baby pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 3, High: 3.2, Low: 2.8, Close: 2.8, Volume: 1000},  // Bearish candle
		{Open: 2.6, High: 2.7, Low: 2.5, Close: 2.6, Volume: 800}, // Doji (gap down)
		{Open: 2.8, High: 3, Low: 2.7, Close: 2.9, Volume: 600},   // Bullish candle (gap up)
		{Open: 2.9, High: 3.1, Low: 2.8, Close: 3, Volume: 400},   // Bullish candle
		{Open: 3, High: 3.2, Low: 2.9, Close: 3.1, Volume: 300},   // Bullish candle
		{Open: 3.1, High: 3.3, Low: 3, Close: 3.2, Volume: 200},   // Bullish candle
		{Open: 3.2, High: 3.4, Low: 3.1, Close: 3.3, Volume: 100}, // Bullish candle
		{Open: 3.3, High: 3.5, Low: 3.2, Close: 3.4, Volume: 50},  // Bullish candle
		{Open: 3.4, High: 3.6, Low: 3.3, Close: 3.5, Volume: 25},  // Bullish candle
		{Open: 3.5, High: 3.7, Low: 3.4, Close: 3.6, Volume: 10},  // Bullish candle
		{Open: 3.6, High: 3.8, Low: 3.5, Close: 3.7, Volume: 5},   // Bullish candle
		{Open: 3.7, High: 3.9, Low: 3.6, Close: 3.8, Volume: 2},   // Bullish candle
		{Open: 3.8, High: 4, Low: 3.7, Close: 3.9, Volume: 1},     // Bullish candle
		{Open: 3.9, High: 4.1, Low: 3.8, Close: 4, Volume: 1},     // Bullish candle
		{Open: 4, High: 4.2, Low: 3.9, Close: 4.1, Volume: 1},     // Bullish candle
	}...)

	// Add Thrusting pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 7, High: 7.2, Low: 6.8, Close: 6.8, Volume: 1000},  // Bearish candle
		{Open: 6.7, High: 7.1, Low: 6.6, Close: 7, Volume: 800},   // Bullish candle (Thrusting)
		{Open: 7, High: 7.3, Low: 6.9, Close: 7.2, Volume: 600},   // Bullish candle
		{Open: 7.2, High: 7.4, Low: 7.1, Close: 7.3, Volume: 400}, // Bullish candle
		{Open: 7.3, High: 7.5, Low: 7.2, Close: 7.4, Volume: 300}, // Bullish candle
	}...)

	// Add Counterattack pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 6, High: 6.2, Low: 5.8, Close: 5.8, Volume: 1000},  // Bearish candle
		{Open: 5.9, High: 6.1, Low: 5.7, Close: 6, Volume: 800},   // Bullish candle (Counterattack)
		{Open: 6, High: 6.3, Low: 5.9, Close: 6.2, Volume: 600},   // Bullish candle
		{Open: 6.2, High: 6.4, Low: 6.1, Close: 6.3, Volume: 400}, // Bullish candle
		{Open: 6.3, High: 6.5, Low: 6.2, Close: 6.4, Volume: 300}, // Bullish candle
	}...)

	// Add Breakaway pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 9, High: 9.5, Low: 8.5, Close: 9.4, Volume: 1000},  // Long bullish candle
		{Open: 9.4, High: 9.6, Low: 9.3, Close: 9.5, Volume: 800}, // Smaller bullish candle
		{Open: 9.5, High: 9.7, Low: 9.4, Close: 9.6, Volume: 600}, // Even smaller bullish candle
		{Open: 9.6, High: 9.8, Low: 9.5, Close: 9.7, Volume: 400}, // Very small bullish candle
		{Open: 9.7, High: 9.9, Low: 9.6, Close: 9.6, Volume: 300}, // Bearish candle (Breakaway)
	}...)

	// Add Mat Hold pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 11, High: 12, Low: 10, Close: 11.8, Volume: 1000},      // Long bullish candle
		{Open: 11.7, High: 11.8, Low: 11.5, Close: 11.5, Volume: 800}, // Bearish candle (contained)
		{Open: 11.6, High: 11.7, Low: 11.4, Close: 11.4, Volume: 600}, // Bearish candle (contained)
		{Open: 11.5, High: 11.6, Low: 11.3, Close: 11.3, Volume: 400}, // Bearish candle (contained)
		{Open: 11.4, High: 12.2, Low: 11.3, Close: 12.1, Volume: 300}, // Bullish candle (above first)
	}...)

	// Add Rising/Falling Three Methods pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 13, High: 14, Low: 12, Close: 13.8, Volume: 1000},      // Long bullish candle
		{Open: 13.7, High: 13.8, Low: 13.5, Close: 13.5, Volume: 800}, // Bearish candle (contained)
		{Open: 13.6, High: 13.7, Low: 13.4, Close: 13.4, Volume: 600}, // Bearish candle (contained)
		{Open: 13.5, High: 13.6, Low: 13.3, Close: 13.3, Volume: 400}, // Bearish candle (contained)
		{Open: 13.4, High: 14.2, Low: 13.3, Close: 14.1, Volume: 300}, // Bullish candle (above first)
		{Open: 14.1, High: 15, Low: 14, Close: 14, Volume: 200},       // Long bearish candle
		{Open: 14.1, High: 14.2, Low: 13.9, Close: 14.1, Volume: 150}, // Bullish candle (contained)
		{Open: 14.2, High: 14.3, Low: 14, Close: 14.2, Volume: 100},   // Bullish candle (contained)
		{Open: 14.3, High: 14.4, Low: 14.1, Close: 14.3, Volume: 50},  // Bullish candle (contained)
		{Open: 14.4, High: 14.5, Low: 13.8, Close: 13.8, Volume: 25},  // Bearish candle (below first)
	}...)

	// Add Separating Lines pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 15, High: 15.2, Low: 14.8, Close: 14.8, Volume: 1000},  // Bearish candle
		{Open: 15, High: 15.3, Low: 14.9, Close: 15.2, Volume: 800},   // Bullish candle (same open)
		{Open: 15.2, High: 15.4, Low: 15.1, Close: 15.1, Volume: 600}, // Bearish candle
		{Open: 15.1, High: 15.3, Low: 15, Close: 15.3, Volume: 400},   // Bullish candle (same open)
	}...)

	// Add On-Neck pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 16, High: 16.2, Low: 15.8, Close: 15.8, Volume: 1000},  // Bearish candle
		{Open: 15.7, High: 15.9, Low: 15.6, Close: 15.8, Volume: 800}, // Bullish candle (closes at previous low)
		{Open: 15.8, High: 16, Low: 15.7, Close: 15.9, Volume: 600},   // Bullish candle
		{Open: 15.9, High: 16.1, Low: 15.8, Close: 16, Volume: 400},   // Bullish candle
		{Open: 16, High: 16.2, Low: 15.9, Close: 16.1, Volume: 300},   // Bullish candle
	}...)

	// Add In-Neck pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 17, High: 17.2, Low: 16.8, Close: 16.8, Volume: 1000},  // Bearish candle
		{Open: 16.7, High: 16.9, Low: 16.6, Close: 16.8, Volume: 800}, // Bullish candle (closes at previous close)
		{Open: 16.8, High: 17, Low: 16.7, Close: 16.9, Volume: 600},   // Bullish candle
		{Open: 16.9, High: 17.1, Low: 16.8, Close: 17, Volume: 400},   // Bullish candle
		{Open: 17, High: 17.2, Low: 16.9, Close: 17.1, Volume: 300},   // Bullish candle
	}...)

	// Add Three Outside Up/Down pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 18, High: 18.2, Low: 17.8, Close: 17.8, Volume: 1000},  // Bearish candle
		{Open: 17.7, High: 18.3, Low: 17.6, Close: 18.2, Volume: 800}, // Bullish candle (engulfs first)
		{Open: 18.2, High: 18.4, Low: 18.1, Close: 18.3, Volume: 600}, // Bullish candle (Three Outside Up)
		{Open: 18.3, High: 18.5, Low: 18.2, Close: 18.4, Volume: 400}, // Bullish candle
		{Open: 18.4, High: 18.6, Low: 18.3, Close: 18.5, Volume: 300}, // Bullish candle
		{Open: 18.5, High: 18.7, Low: 18.4, Close: 18.4, Volume: 200}, // Bearish candle
		{Open: 18.3, High: 18.5, Low: 18.2, Close: 18.2, Volume: 150}, // Bearish candle (engulfs previous)
		{Open: 18.1, High: 18.3, Low: 18, Close: 18, Volume: 100},     // Bearish candle (Three Outside Down)
	}...)

	// Add Stick Sandwich pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 19, High: 19.2, Low: 18.8, Close: 18.8, Volume: 1000}, // Bearish candle
		{Open: 18.9, High: 19, Low: 18.7, Close: 18.9, Volume: 800},  // Bullish candle (contained)
		{Open: 19, High: 19.2, Low: 18.8, Close: 18.8, Volume: 600},  // Bearish candle (same close as first)
		{Open: 18.8, High: 19, Low: 18.6, Close: 18.9, Volume: 400},  // Bullish candle
		{Open: 18.9, High: 19.1, Low: 18.7, Close: 19, Volume: 300},  // Bullish candle
	}...)

	// Add Unique Three River Bottom pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 21, High: 21.2, Low: 20.8, Close: 20.8, Volume: 1000},  // Bearish candle
		{Open: 20.7, High: 20.8, Low: 20.5, Close: 20.6, Volume: 800}, // Bearish hammer (lower low)
		{Open: 20.6, High: 21, Low: 20.5, Close: 21, Volume: 600},     // Bullish candle (above hammer)
		{Open: 21, High: 21.2, Low: 20.8, Close: 21.1, Volume: 400},   // Bullish candle
		{Open: 21.1, High: 21.3, Low: 20.9, Close: 21.2, Volume: 300}, // Bullish candle
	}...)

	// Add Three Stars in the South pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 23, High: 23.2, Low: 22.8, Close: 22.8, Volume: 1000},  // Large bearish candle
		{Open: 22.7, High: 22.9, Low: 22.5, Close: 22.5, Volume: 800}, // Medium bearish candle
		{Open: 22.4, High: 22.6, Low: 22.2, Close: 22.2, Volume: 600}, // Small bearish candle
		{Open: 22.2, High: 22.4, Low: 22, Close: 22.3, Volume: 400},   // Bullish candle
		{Open: 22.3, High: 22.5, Low: 22.1, Close: 22.4, Volume: 300}, // Bullish candle
	}...)

	// Add Tasuki Gap pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 25, High: 26, Low: 24.5, Close: 25.5, Volume: 1000},  // Bullish candle
		{Open: 26, High: 27, Low: 25.5, Close: 26.5, Volume: 800},   // Bullish candle (gap up)
		{Open: 26.5, High: 26.8, Low: 25.8, Close: 26, Volume: 600}, // Bearish candle (fills gap)
		{Open: 26, High: 26.5, Low: 25.5, Close: 26.2, Volume: 400}, // Bullish candle
		{Open: 26.2, High: 26.7, Low: 26, Close: 26.5, Volume: 300}, // Bullish candle
	}...)

	// Add Three Gap Ups pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 27, High: 28, Low: 26.5, Close: 27.5, Volume: 1000}, // Bullish candle
		{Open: 28, High: 29, Low: 27.5, Close: 28.5, Volume: 800},  // Bullish candle (gap up)
		{Open: 29, High: 30, Low: 28.5, Close: 29.5, Volume: 600},  // Bullish candle (gap up)
		{Open: 29.5, High: 30.5, Low: 29, Close: 30, Volume: 400},  // Bullish candle
		{Open: 30, High: 31, Low: 29.5, Close: 30.5, Volume: 300},  // Bullish candle
	}...)

	// Add Three Gap Downs pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 32, High: 32.5, Low: 31, Close: 31.5, Volume: 1000}, // Bearish candle
		{Open: 31, High: 31.5, Low: 30, Close: 30.5, Volume: 800},  // Bearish candle (gap down)
		{Open: 30, High: 30.5, Low: 29, Close: 29.5, Volume: 600},  // Bearish candle (gap down)
		{Open: 29.5, High: 30, Low: 28.5, Close: 29, Volume: 400},  // Bearish candle
		{Open: 29, High: 29.5, Low: 28, Close: 28.5, Volume: 300},  // Bearish candle
	}...)

	// Add Concealing Baby Swallow pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 33, High: 33, Low: 32, Close: 32, Volume: 1000},        // Bearish Marubozu
		{Open: 32, High: 32, Low: 31, Close: 31, Volume: 800},         // Bearish Marubozu
		{Open: 31, High: 31.2, Low: 30.8, Close: 30.9, Volume: 600},   // Small bearish candle
		{Open: 31.2, High: 31.3, Low: 30.7, Close: 30.8, Volume: 400}, // Bearish candle (engulfs third)
		{Open: 30.8, High: 31.5, Low: 30.5, Close: 31.3, Volume: 300}, // Bullish candle
	}...)

	// Add Ladder Bottom pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 34, High: 34.5, Low: 33.5, Close: 33.5, Volume: 1000}, // Bearish candle
		{Open: 33.5, High: 34, Low: 33, Close: 33, Volume: 800},      // Bearish candle (lower low)
		{Open: 33, High: 33.5, Low: 32.5, Close: 32.5, Volume: 600},  // Bearish candle (lower low)
		{Open: 32.5, High: 33, Low: 32, Close: 32, Volume: 400},      // Bearish candle (lower low)
		{Open: 32, High: 34.5, Low: 31.5, Close: 34, Volume: 300},    // Bullish candle (above first high)
	}...)

	// Add Three Line Strike pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 35, High: 36, Low: 34.5, Close: 35.5, Volume: 1000}, // Bullish candle
		{Open: 35.5, High: 36.5, Low: 35, Close: 36, Volume: 800},  // Bullish candle
		{Open: 36, High: 37, Low: 35.5, Close: 36.5, Volume: 600},  // Bullish candle
		{Open: 36.5, High: 37, Low: 34, Close: 34.5, Volume: 400},  // Bearish candle (engulfs all)
		{Open: 34.5, High: 35, Low: 34, Close: 34.5, Volume: 300},  // Bullish candle
	}...)

	// Add Two Crows pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 37, High: 38, Low: 36.5, Close: 37.5, Volume: 1000}, // Bullish candle
		{Open: 38.5, High: 39, Low: 37, Close: 37.5, Volume: 800},  // Bearish candle (opens above)
		{Open: 38, High: 38.5, Low: 36.5, Close: 37, Volume: 600},  // Bearish candle (opens within)
		{Open: 37, High: 37.5, Low: 36.5, Close: 37, Volume: 400},  // Bullish candle
		{Open: 37, High: 37.5, Low: 36.5, Close: 37, Volume: 300},  // Bullish candle
	}...)

	// Add Three Advancing White Soldiers pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 39, High: 40, Low: 38.5, Close: 39.5, Volume: 1000},    // Bullish candle
		{Open: 39.5, High: 41, Low: 39, Close: 40.5, Volume: 800},     // Larger bullish candle
		{Open: 40.5, High: 42, Low: 40, Close: 41.5, Volume: 600},     // Even larger bullish candle
		{Open: 41.5, High: 42, Low: 41, Close: 41.8, Volume: 400},     // Bullish candle
		{Open: 41.8, High: 42.5, Low: 41.5, Close: 42.2, Volume: 300}, // Bullish candle
	}...)

	// Add Three Declining Black Crows pattern candles
	candles = append(candles, []patterns.Candle{
		{Open: 43, High: 43.5, Low: 42, Close: 42.5, Volume: 1000},  // Bearish candle
		{Open: 42.5, High: 43, Low: 41.5, Close: 41.8, Volume: 800}, // Larger bearish candle
		{Open: 41.8, High: 42.5, Low: 41, Close: 41.2, Volume: 600}, // Even larger bearish candle
		{Open: 41.2, High: 41.5, Low: 40.5, Close: 41, Volume: 400}, // Bearish candle
		{Open: 41, High: 41.5, Low: 40.5, Close: 40.8, Volume: 300}, // Bearish candle
	}...)

	// Add Identical Three Crows pattern
	candles = append(candles, []patterns.Candle{
		{Open: 100, High: 102, Low: 98, Close: 99, Volume: 1000}, // Bearish
		{Open: 99, High: 101, Low: 97, Close: 98, Volume: 1000},  // Bearish
		{Open: 98, High: 100, Low: 96, Close: 97, Volume: 1000},  // Bearish
	}...)

	// Add Deliberation pattern
	candles = append(candles, []patterns.Candle{
		{Open: 100, High: 105, Low: 99, Close: 104, Volume: 1000},  // Bullish
		{Open: 104, High: 109, Low: 103, Close: 108, Volume: 1000}, // Bullish
		{Open: 108, High: 110, Low: 107, Close: 109, Volume: 1000}, // Small bullish
	}...)

	// Add Three Strike pattern (Bullish)
	candles = append(candles, []patterns.Candle{
		{Open: 100, High: 105, Low: 99, Close: 104, Volume: 1000},  // Bullish
		{Open: 104, High: 109, Low: 103, Close: 108, Volume: 1000}, // Bullish
		{Open: 109, High: 110, Low: 98, Close: 99, Volume: 1000},   // Bearish
	}...)

	// Add Three Strike pattern (Bearish)
	candles = append(candles, []patterns.Candle{
		{Open: 100, High: 101, Low: 95, Close: 96, Volume: 1000}, // Bearish
		{Open: 96, High: 97, Low: 91, Close: 92, Volume: 1000},   // Bearish
		{Open: 91, High: 102, Low: 90, Close: 101, Volume: 1000}, // Bullish
	}...)

	// Add Upside Gap Two Crows pattern
	candles = append(candles, []patterns.Candle{
		{Open: 100, High: 105, Low: 99, Close: 104, Volume: 1000},  // Bullish
		{Open: 106, High: 107, Low: 102, Close: 103, Volume: 1000}, // Bearish
		{Open: 104, High: 105, Low: 101, Close: 102, Volume: 1000}, // Bearish
	}...)

	// Analyze candlestick patterns
	results := patterns.AnalyzeCandlestickPatterns(candles)

	// Print results
	fmt.Println("Candlestick Pattern Analysis Results:")
	fmt.Println("=====================================")
	for _, result := range results {
		fmt.Printf("Pattern: %s\n", result.Pattern)
		fmt.Printf("Confidence: %.2f\n", result.Confidence)
		fmt.Printf("Description: %s\n", result.Description)
		fmt.Println("-------------------------------------")
	}

	// Print detailed analysis
	fmt.Println("\nDetailed Analysis:")
	fmt.Println("==================")
	for i, candle := range candles {
		fmt.Printf("Candle %d:\n", i+1)
		fmt.Printf("  Open: %.2f\n", candle.Open)
		fmt.Printf("  High: %.2f\n", candle.High)
		fmt.Printf("  Low: %.2f\n", candle.Low)
		fmt.Printf("  Close: %.2f\n", candle.Close)
		fmt.Printf("  Volume: %.2f\n", candle.Volume)

		// Check for single candle patterns
		if patterns.IsDoji(candle) {
			fmt.Println("  Pattern: Doji")
		}
		if patterns.IsHammer(candle) {
			fmt.Println("  Pattern: Hammer")
		}
		if patterns.IsShootingStar(candle) {
			fmt.Println("  Pattern: Shooting Star")
		}
		if patterns.IsSpinningTop(candle) {
			fmt.Println("  Pattern: Spinning Top")
		}
		if patterns.IsMarubozu(candle) {
			fmt.Println("  Pattern: Marubozu")
		}
		if patterns.IsHangingMan(candle) {
			fmt.Println("  Pattern: Hanging Man")
		}
		if patterns.IsInvertedHammer(candle) {
			fmt.Println("  Pattern: Inverted Hammer")
		}

		// Check for two-candle patterns
		if i > 0 {
			isEngulfing, isBullish := patterns.IsEngulfing(candles[i-1], candle)
			if isEngulfing {
				if isBullish {
					fmt.Println("  Pattern: Bullish Engulfing")
				} else {
					fmt.Println("  Pattern: Bearish Engulfing")
				}
			}

			isHarami, isBullish := patterns.IsHarami(candles[i-1], candle)
			if isHarami {
				if isBullish {
					fmt.Println("  Pattern: Bullish Harami")
				} else {
					fmt.Println("  Pattern: Bearish Harami")
				}
			}

			if patterns.IsPiercing(candles[i-1], candle) {
				fmt.Println("  Pattern: Piercing Pattern")
			}

			if patterns.IsDarkCloudCover(candles[i-1], candle) {
				fmt.Println("  Pattern: Dark Cloud Cover")
			}
		}

		// Check for three-candle patterns
		if i > 1 {
			if patterns.IsMorningStar(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Morning Star")
			}
			if patterns.IsEveningStar(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Evening Star")
			}
			if patterns.IsThreeWhiteSoldiers(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Three White Soldiers")
			}
			if patterns.IsThreeBlackCrows(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Three Black Crows")
			}
		}

		// Check for Belt Hold pattern
		isBeltHold, isBullish := patterns.IsBeltHold(candle)
		if isBeltHold {
			if isBullish {
				fmt.Println("  Pattern: Bullish Belt Hold")
			} else {
				fmt.Println("  Pattern: Bearish Belt Hold")
			}
		}

		// Check for Kicking pattern
		if i > 0 {
			isKicking, isBullish := patterns.IsKicking(candles[i-1], candle)
			if isKicking {
				if isBullish {
					fmt.Println("  Pattern: Bullish Kicking")
				} else {
					fmt.Println("  Pattern: Bearish Kicking")
				}
			}
		}

		// Check for Abandoned Baby pattern
		if i > 1 {
			isAbandonedBaby, isBullish := patterns.IsAbandonedBaby(candles[i-2], candles[i-1], candle)
			if isAbandonedBaby {
				if isBullish {
					fmt.Println("  Pattern: Bullish Abandoned Baby")
				} else {
					fmt.Println("  Pattern: Bearish Abandoned Baby")
				}
			}
		}

		// Check for Dragonfly Doji
		if patterns.IsDragonflyDoji(candle) {
			fmt.Println("  Pattern: Dragonfly Doji")
		}

		// Check for Gravestone Doji
		if patterns.IsGravestoneDoji(candle) {
			fmt.Println("  Pattern: Gravestone Doji")
		}

		// Check for Three Inside Up/Down patterns
		if i > 1 {
			if patterns.IsThreeInsideUp(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Three Inside Up")
			}
			if patterns.IsThreeInsideDown(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Three Inside Down")
			}
		}

		// Check for Thrusting pattern
		if i > 0 {
			if patterns.IsThrusting(candles[i-1], candle) {
				fmt.Println("  Pattern: Thrusting Pattern")
			}
		}

		// Check for Counterattack pattern
		if i > 0 {
			isCounterattack, isBullish := patterns.IsCounterattack(candles[i-1], candle)
			if isCounterattack {
				if isBullish {
					fmt.Println("  Pattern: Bullish Counterattack")
				} else {
					fmt.Println("  Pattern: Bearish Counterattack")
				}
			}
		}

		// Check for Breakaway pattern
		if i > 3 {
			isBreakaway, isBullish := patterns.IsBreakaway(candles[i-4], candles[i-3], candles[i-2], candles[i-1], candle)
			if isBreakaway {
				if isBullish {
					fmt.Println("  Pattern: Bullish Breakaway")
				} else {
					fmt.Println("  Pattern: Bearish Breakaway")
				}
			}
		}

		// Check for Mat Hold pattern
		if i > 3 {
			isMatHold, isBullish := patterns.IsMatHold(candles[i-4], candles[i-3], candles[i-2], candles[i-1], candle)
			if isMatHold {
				if isBullish {
					fmt.Println("  Pattern: Bullish Mat Hold")
				} else {
					fmt.Println("  Pattern: Bearish Mat Hold")
				}
			}
		}

		// Check for Rising/Falling Three Methods patterns
		if i > 3 {
			if patterns.IsRisingThreeMethods(candles[i-4], candles[i-3], candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Rising Three Methods")
			}
			if patterns.IsFallingThreeMethods(candles[i-4], candles[i-3], candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Falling Three Methods")
			}
		}

		// Check for Separating Lines pattern
		if i > 0 {
			isSeparatingLines, isBullish := patterns.IsSeparatingLines(candles[i-1], candle)
			if isSeparatingLines {
				if isBullish {
					fmt.Println("  Pattern: Bullish Separating Lines")
				} else {
					fmt.Println("  Pattern: Bearish Separating Lines")
				}
			}
		}

		// Check for On-Neck pattern
		if i > 0 {
			isOnNeck, isBullish := patterns.IsOnNeck(candles[i-1], candle)
			if isOnNeck {
				if isBullish {
					fmt.Println("  Pattern: Bullish On-Neck")
				} else {
					fmt.Println("  Pattern: Bearish On-Neck")
				}
			}
		}

		// Check for In-Neck pattern
		if i > 0 {
			isInNeck, isBullish := patterns.IsInNeck(candles[i-1], candle)
			if isInNeck {
				if isBullish {
					fmt.Println("  Pattern: Bullish In-Neck")
				} else {
					fmt.Println("  Pattern: Bearish In-Neck")
				}
			}
		}

		// Check for Three Outside Up/Down patterns
		if i > 1 {
			if patterns.IsThreeOutsideUp(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Three Outside Up")
			}
			if patterns.IsThreeOutsideDown(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Three Outside Down")
			}
		}

		// Check for Stick Sandwich pattern
		if i > 1 {
			if patterns.IsStickSandwich(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Stick Sandwich")
			}
		}

		// Check for Unique Three River Bottom pattern
		if i > 1 {
			if patterns.IsUniqueThreeRiverBottom(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Unique Three River Bottom")
			}
		}

		// Check for Three Stars in the South pattern
		if i > 1 {
			if patterns.IsThreeStarsInTheSouth(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Three Stars in the South")
			}
		}

		// Check for Tasuki Gap pattern
		if i > 1 {
			isTasukiGap, isBullish := patterns.IsTasukiGap(candles[i-2], candles[i-1], candle)
			if isTasukiGap {
				if isBullish {
					fmt.Println("  Pattern: Bullish Tasuki Gap")
				} else {
					fmt.Println("  Pattern: Bearish Tasuki Gap")
				}
			}
		}

		// Check for Three Gap Ups/Downs patterns
		if i > 1 {
			if patterns.IsThreeGapUps(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Three Gap Ups")
			}
			if patterns.IsThreeGapDowns(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Three Gap Downs")
			}
		}

		// Check for Concealing Baby Swallow pattern
		if i > 2 {
			if patterns.IsConcealingBabySwallow(candles[i-3], candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Concealing Baby Swallow")
			}
		}

		// Check for Ladder Bottom pattern
		if i > 3 {
			if patterns.IsLadderBottom(candles[i-4], candles[i-3], candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Ladder Bottom")
			}
		}

		// Check for Three Line Strike pattern
		if i > 2 {
			isThreeLineStrike, isBullish := patterns.IsThreeLineStrike(candles[i-3], candles[i-2], candles[i-1], candle)
			if isThreeLineStrike {
				if isBullish {
					fmt.Println("  Pattern: Bullish Three Line Strike")
				} else {
					fmt.Println("  Pattern: Bearish Three Line Strike")
				}
			}
		}

		// Check for Two Crows pattern
		if i > 1 {
			if patterns.IsTwoCrows(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Two Crows")
			}
		}

		// Check for Three Advancing White Soldiers pattern
		if i > 1 {
			if patterns.IsThreeAdvancingWhiteSoldiers(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Three Advancing White Soldiers")
			}
		}

		// Check for Three Declining Black Crows pattern
		if i > 1 {
			if patterns.IsThreeDecliningBlackCrows(candles[i-2], candles[i-1], candle) {
				fmt.Println("  Pattern: Three Declining Black Crows")
			}
		}

		// Add checks for new patterns
		if i > 2 {
			if patterns.IsIdenticalThreeCrows(candles[i-2], candles[i-1], candles[i]) {
				fmt.Printf("Identical Three Crows pattern found at index %d\n", i)
			}

			if patterns.IsDeliberation(candles[i-2], candles[i-1], candles[i]) {
				fmt.Printf("Deliberation pattern found at index %d\n", i)
			}

			isThreeStrike, isBullish := patterns.IsThreeStrike(candles[i-2], candles[i-1], candles[i])
			if isThreeStrike {
				pattern := "Bearish Three Strike"
				if isBullish {
					pattern = "Bullish Three Strike"
				}
				fmt.Printf("%s pattern found at index %d\n", pattern, i)
			}

			if patterns.IsUpsideGapTwoCrows(candles[i-2], candles[i-1], candles[i]) {
				fmt.Printf("Upside Gap Two Crows pattern found at index %d\n", i)
			}
		}

		fmt.Println()
	}
}
