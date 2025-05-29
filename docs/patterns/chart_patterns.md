# Chart Pattern Detection Documentation

This document provides detailed information about the chart pattern detection implementation in the `patterns` package.

## Overview

The chart pattern detection system analyzes a series of candlesticks to identify various technical chart patterns. These patterns are used to predict potential price movements and market trends.

## Pattern Types

### 1. Head and Shoulders
**Function:** `IsHeadAndShoulders(candles []Candle) bool`

A bearish reversal pattern consisting of three peaks:
- Left shoulder: First peak
- Head: Highest peak in the middle
- Right shoulder: Third peak at similar level to left shoulder

**Requirements:**
- Minimum 7 candles
- Head must be at least 3% higher than shoulders
- Shoulders must be within 3% of each other
- Neckline (support level) must be roughly horizontal
- Pattern confirmed by break below neckline

**Confidence Level:** 0.9

### 2. Inverse Head and Shoulders
**Function:** `IsInverseHeadAndShoulders(candles []Candle) bool`

A bullish reversal pattern, mirror image of Head and Shoulders:
- Left shoulder: First trough
- Head: Lowest point in the middle
- Right shoulder: Third trough at similar level to left shoulder

**Requirements:**
- Minimum 7 candles
- Head must be at least 3% lower than shoulders
- Shoulders must be within 3% of each other
- Neckline (resistance level) must be roughly horizontal
- Pattern confirmed by break above neckline

**Confidence Level:** 0.9

### 3. Double Top
**Function:** `IsDoubleTop(candles []Candle) bool`

A bearish reversal pattern with two peaks at similar levels:
- Two distinct peaks
- Peaks must be within 3% of each other
- Pattern confirmed by break below the trough between peaks

**Requirements:**
- Minimum 7 candles
- Two clear peaks
- Confirmation through breakdown

**Confidence Level:** 0.85

### 4. Double Bottom
**Function:** `IsDoubleBottom(candles []Candle) bool`

A bullish reversal pattern with two troughs at similar levels:
- Two distinct troughs
- Troughs must be within 3% of each other
- Pattern confirmed by break above the peak between troughs

**Requirements:**
- Minimum 7 candles
- Two clear troughs
- Confirmation through breakout

**Confidence Level:** 0.85

### 5. Triangle Patterns
**Function:** `IsTriangle(candles []Candle) (bool, string)`

Three types of triangle patterns:

#### a. Symmetrical Triangle
- Converging trend lines
- Both highs and lows converge
- Neutral pattern, direction determined by breakout

#### b. Ascending Triangle
- Flat upper trend line
- Rising lower trend line
- Generally bullish

#### c. Descending Triangle
- Flat lower trend line
- Falling upper trend line
- Generally bearish

**Requirements:**
- Minimum 10 candles
- At least 2 highs and 2 lows
- Slope analysis for pattern type determination

**Confidence Level:** 0.8

### 6. Flags and Pennants
**Function:** `IsFlagOrPennant(candles []Candle) (bool, string)`

Continuation patterns that appear after strong trends:

#### a. Flag
- Small rectangular pattern
- Slopes against the trend
- Short duration

#### b. Pennant
- Small symmetrical triangle
- Horizontal consolidation
- Short duration

**Requirements:**
- Minimum 10 candles
- Strong trend in first half
- Consolidation range < 50% of trend range
- Pattern duration shorter than trend

**Confidence Level:** 0.8

### 7. Rounding Bottom
**Function:** `IsRoundingBottom(candles []Candle) bool`

A long-term bottom formation:
- U-shaped price movement
- Gradual transition from bearish to bullish
- Confirmed by break above resistance

**Requirements:**
- Minimum 20 candles
- 5-period moving average for smoothing
- U-shaped price movement
- Confirmation through breakout

**Confidence Level:** 0.85

## Main Analysis Function

**Function:** `AnalyzeChartPatterns(candles []Candle) []PatternResult`

Analyzes a series of candles for all chart patterns and returns a slice of PatternResult containing:
- Pattern name
- Confidence level
- Description of pattern significance

**Usage Example:**
```go
candles := []Candle{...}
patterns := AnalyzeChartPatterns(candles)
for _, pattern := range patterns {
    fmt.Printf("Pattern: %s (Confidence: %.2f)\n", pattern.Pattern, pattern.Confidence)
    fmt.Printf("Description: %s\n", pattern.Description)
}
```

## PatternResult Structure

```go
type PatternResult struct {
    Pattern     string  // Name of the pattern
    Confidence  float64 // Confidence level (0.0 to 1.0)
    Description string  // Description of the pattern
}
```

## Implementation Details

### Common Features
- All patterns use percentage-based thresholds (typically 3%) for level comparisons
- Pattern confirmation requires breakout/breakdown
- Minimum candle requirements for each pattern
- Error checking and validation
- Integration with existing Candle struct

### Candle Structure
```go
type Candle struct {
    Open   float64
    High   float64
    Low    float64
    Close  float64
    Volume float64
    Time   time.Time
}
```

## Best Practices

1. **Pattern Confirmation**
   - Always wait for pattern confirmation through breakout/breakdown
   - Use volume confirmation when available
   - Consider market context and timeframe

2. **Risk Management**
   - Use stop-loss orders based on pattern structure
   - Consider pattern size for position sizing
   - Monitor for pattern failure

3. **Timeframe Considerations**
   - Patterns on higher timeframes are more reliable
   - Consider multiple timeframe analysis
   - Be aware of pattern duration

## Limitations

1. **False Signals**
   - Not all patterns lead to expected outcomes
   - Market conditions can invalidate patterns
   - Need additional confirmation

2. **Timeframe Dependencies**
   - Pattern reliability varies by timeframe
   - Longer timeframes generally more reliable
   - Consider market context

3. **Market Conditions**
   - Patterns work best in trending markets
   - Less reliable in choppy or ranging markets
   - Consider overall market conditions

## Future Improvements

1. **Volume Analysis**
   - Add volume confirmation
   - Volume profile analysis
   - Volume trend analysis

2. **Pattern Combinations**
   - Multiple pattern confirmation
   - Pattern failure detection
   - Pattern strength measurement

3. **Machine Learning**
   - Pattern recognition accuracy improvement
   - False signal reduction
   - Dynamic threshold adjustment 