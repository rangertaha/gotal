# Head and Shoulders Pattern

## Overview

The Head and Shoulders pattern is a bearish reversal pattern that typically forms at the end of an uptrend. It consists of three peaks, with the middle peak (head) being the highest and the two outer peaks (shoulders) being roughly equal in height.

## Pattern Structure

```
    Head
     │
     │
Left │    Right
Shoulder  Shoulder
     │      │
     │      │
     └──────┘
    Neckline
```

## Function Signature

```go
func IsHeadAndShoulders(candles []Candle) bool
```

## Requirements

1. **Minimum Candles**: 7 candles required
2. **Pattern Formation**:
   - Left shoulder: First peak
   - Head: Highest peak in the middle
   - Right shoulder: Third peak at similar level to left shoulder
3. **Level Requirements**:
   - Head must be at least 3% higher than shoulders
   - Shoulders must be within 3% of each other
   - Neckline (support level) must be roughly horizontal
4. **Confirmation**:
   - Pattern is confirmed by break below neckline

## Implementation Details

### Pattern Detection Steps

1. **Find the Head**
   ```go
   headIndex := 0
   headHigh := candles[0].High
   for i := 1; i < len(candles); i++ {
       if candles[i].High > headHigh {
           headHigh = candles[i].High
           headIndex = i
       }
   }
   ```

2. **Find Shoulders**
   ```go
   // Left shoulder (highest point before head)
   leftShoulderIndex := 0
   leftShoulderHigh := candles[0].High
   for i := 1; i < headIndex; i++ {
       if candles[i].High > leftShoulderHigh {
           leftShoulderHigh = candles[i].High
           leftShoulderIndex = i
       }
   }

   // Right shoulder (highest point after head)
   rightShoulderIndex := headIndex + 1
   rightShoulderHigh := candles[rightShoulderIndex].High
   for i := headIndex + 2; i < len(candles); i++ {
       if candles[i].High > rightShoulderHigh {
           rightShoulderHigh = candles[i].High
           rightShoulderIndex = i
       }
   }
   ```

3. **Validate Shoulder Levels**
   ```go
   shoulderDiff := math.Abs(leftShoulderHigh - rightShoulderHigh)
   avgShoulderHeight := (leftShoulderHigh + rightShoulderHigh) / 2
   if shoulderDiff/avgShoulderHeight > 0.03 {
       return false
   }
   ```

4. **Check Head Height**
   ```go
   if headHigh <= leftShoulderHigh || headHigh <= rightShoulderHigh {
       return false
   }
   headShoulderDiff := math.Min(headHigh-leftShoulderHigh, headHigh-rightShoulderHigh)
   if headShoulderDiff/avgShoulderHeight < 0.03 {
       return false
   }
   ```

5. **Validate Neckline**
   ```go
   leftTroughLow := candles[leftShoulderIndex].Low
   for i := leftShoulderIndex + 1; i < headIndex; i++ {
       if candles[i].Low < leftTroughLow {
           leftTroughLow = candles[i].Low
       }
   }

   rightTroughLow := candles[headIndex].Low
   for i := headIndex + 1; i < rightShoulderIndex; i++ {
       if candles[i].Low < rightTroughLow {
           rightTroughLow = candles[i].Low
       }
   }

   necklineDiff := math.Abs(leftTroughLow - rightTroughLow)
   avgNeckline := (leftTroughLow + rightTroughLow) / 2
   if necklineDiff/avgNeckline > 0.03 {
       return false
   }
   ```

6. **Confirm Breakout**
   ```go
   for i := rightShoulderIndex + 1; i < len(candles); i++ {
       if candles[i].Close < avgNeckline {
           return true
       }
   }
   ```

## Trading Considerations

### Entry Points
1. **Conservative**: Wait for close below neckline
2. **Aggressive**: Enter on right shoulder formation

### Stop Loss
- Place stop loss above the right shoulder
- Consider the height of the pattern for position sizing

### Price Target
- Measure the height from head to neckline
- Project downward from the neckline break

## Example Usage

```go
candles := []Candle{...}
if IsHeadAndShoulders(candles) {
    fmt.Println("Head and Shoulders pattern detected")
    // Implement trading logic
}
```

## Confidence Level

The pattern has a confidence level of 0.9 when all requirements are met.

## Limitations

1. **False Signals**
   - Pattern may fail if market conditions change
   - Volume confirmation recommended
   - Consider overall market trend

2. **Timeframe Considerations**
   - More reliable on higher timeframes
   - Minimum 7 candles required
   - Pattern duration varies

## Related Patterns

- [Inverse Head and Shoulders](../reversal/inverse_head_and_shoulders.md)
- [Double Top](../reversal/double_top.md)
- [Triple Top](../reversal/triple_top.md) 