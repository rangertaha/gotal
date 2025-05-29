package patterns

import (
	"math"
)

// IsHeadAndShoulders checks if the candles form a head and shoulders pattern
func IsHeadAndShoulders(candles []Candle) bool {
	if len(candles) < 7 {
		return false
	}

	// Find the highest point (head)
	headIndex := 0
	headHigh := candles[0].High
	for i := 1; i < len(candles); i++ {
		if candles[i].High > headHigh {
			headHigh = candles[i].High
			headIndex = i
		}
	}

	// Need at least 3 candles before and after the head
	if headIndex < 3 || headIndex > len(candles)-4 {
		return false
	}

	// Find left shoulder (highest point before head)
	leftShoulderIndex := 0
	leftShoulderHigh := candles[0].High
	for i := 1; i < headIndex; i++ {
		if candles[i].High > leftShoulderHigh {
			leftShoulderHigh = candles[i].High
			leftShoulderIndex = i
		}
	}

	// Find right shoulder (highest point after head)
	rightShoulderIndex := headIndex + 1
	rightShoulderHigh := candles[rightShoulderIndex].High
	for i := headIndex + 2; i < len(candles); i++ {
		if candles[i].High > rightShoulderHigh {
			rightShoulderHigh = candles[i].High
			rightShoulderIndex = i
		}
	}

	// Check if shoulders are roughly at the same level (within 3% of each other)
	shoulderDiff := math.Abs(leftShoulderHigh - rightShoulderHigh)
	avgShoulderHeight := (leftShoulderHigh + rightShoulderHigh) / 2
	if shoulderDiff/avgShoulderHeight > 0.03 {
		return false
	}

	// Check if head is higher than shoulders (at least 3% higher)
	if headHigh <= leftShoulderHigh || headHigh <= rightShoulderHigh {
		return false
	}
	headShoulderDiff := math.Min(headHigh-leftShoulderHigh, headHigh-rightShoulderHigh)
	if headShoulderDiff/avgShoulderHeight < 0.03 {
		return false
	}

	// Check if there's a neckline (support level) that connects the troughs
	// Find the lowest points between left shoulder and head, and between head and right shoulder
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

	// Check if the neckline is roughly horizontal (within 3% of each other)
	necklineDiff := math.Abs(leftTroughLow - rightTroughLow)
	avgNeckline := (leftTroughLow + rightTroughLow) / 2
	if necklineDiff/avgNeckline > 0.03 {
		return false
	}

	// Check if the pattern is complete by confirming a break below the neckline
	// Look for a candle that closes below the neckline after the right shoulder
	for i := rightShoulderIndex + 1; i < len(candles); i++ {
		if candles[i].Close < avgNeckline {
			return true
		}
	}

	return false
}

// IsInverseHeadAndShoulders checks if the candles form an inverse head and shoulders pattern
func IsInverseHeadAndShoulders(candles []Candle) bool {
	if len(candles) < 7 {
		return false
	}

	// Find the lowest point (head)
	headIndex := 0
	headLow := candles[0].Low
	for i := 1; i < len(candles); i++ {
		if candles[i].Low < headLow {
			headLow = candles[i].Low
			headIndex = i
		}
	}

	// Need at least 3 candles before and after the head
	if headIndex < 3 || headIndex > len(candles)-4 {
		return false
	}

	// Find left shoulder (lowest point before head)
	leftShoulderIndex := 0
	leftShoulderLow := candles[0].Low
	for i := 1; i < headIndex; i++ {
		if candles[i].Low < leftShoulderLow {
			leftShoulderLow = candles[i].Low
			leftShoulderIndex = i
		}
	}

	// Find right shoulder (lowest point after head)
	rightShoulderIndex := headIndex + 1
	rightShoulderLow := candles[rightShoulderIndex].Low
	for i := headIndex + 2; i < len(candles); i++ {
		if candles[i].Low < rightShoulderLow {
			rightShoulderLow = candles[i].Low
			rightShoulderIndex = i
		}
	}

	// Check if shoulders are roughly at the same level (within 3% of each other)
	shoulderDiff := math.Abs(leftShoulderLow - rightShoulderLow)
	avgShoulderHeight := (leftShoulderLow + rightShoulderLow) / 2
	if shoulderDiff/avgShoulderHeight > 0.03 {
		return false
	}

	// Check if head is lower than shoulders (at least 3% lower)
	if headLow >= leftShoulderLow || headLow >= rightShoulderLow {
		return false
	}
	headShoulderDiff := math.Min(leftShoulderLow-headLow, rightShoulderLow-headLow)
	if headShoulderDiff/avgShoulderHeight < 0.03 {
		return false
	}

	// Check if there's a neckline (resistance level) that connects the peaks
	leftTroughHigh := candles[leftShoulderIndex].High
	for i := leftShoulderIndex + 1; i < headIndex; i++ {
		if candles[i].High > leftTroughHigh {
			leftTroughHigh = candles[i].High
		}
	}

	rightTroughHigh := candles[headIndex].High
	for i := headIndex + 1; i < rightShoulderIndex; i++ {
		if candles[i].High > rightTroughHigh {
			rightTroughHigh = candles[i].High
		}
	}

	// Check if the neckline is roughly horizontal (within 3% of each other)
	necklineDiff := math.Abs(leftTroughHigh - rightTroughHigh)
	avgNeckline := (leftTroughHigh + rightTroughHigh) / 2
	if necklineDiff/avgNeckline > 0.03 {
		return false
	}

	// Check if the pattern is complete by confirming a break above the neckline
	for i := rightShoulderIndex + 1; i < len(candles); i++ {
		if candles[i].Close > avgNeckline {
			return true
		}
	}

	return false
}

// IsDoubleTop checks if the candles form a double top pattern
func IsDoubleTop(candles []Candle) bool {
	if len(candles) < 7 {
		return false
	}

	// Find the two highest points
	var peaks []struct {
		index int
		high  float64
	}

	for i := 1; i < len(candles)-1; i++ {
		if candles[i].High > candles[i-1].High && candles[i].High > candles[i+1].High {
			peaks = append(peaks, struct {
				index int
				high  float64
			}{i, candles[i].High})
		}
	}

	if len(peaks) < 2 {
		return false
	}

	// Find the two highest peaks
	firstPeak := peaks[0]
	secondPeak := peaks[1]
	for _, peak := range peaks[2:] {
		if peak.high > firstPeak.high {
			secondPeak = firstPeak
			firstPeak = peak
		} else if peak.high > secondPeak.high {
			secondPeak = peak
		}
	}

	// Check if peaks are roughly at the same level (within 3% of each other)
	peakDiff := math.Abs(firstPeak.high - secondPeak.high)
	avgPeakHeight := (firstPeak.high + secondPeak.high) / 2
	if peakDiff/avgPeakHeight > 0.03 {
		return false
	}

	// Find the lowest point between the peaks
	troughLow := candles[firstPeak.index].Low
	for i := firstPeak.index + 1; i < secondPeak.index; i++ {
		if candles[i].Low < troughLow {
			troughLow = candles[i].Low
		}
	}

	// Check if the pattern is complete by confirming a break below the trough
	for i := secondPeak.index + 1; i < len(candles); i++ {
		if candles[i].Close < troughLow {
			return true
		}
	}

	return false
}

// IsDoubleBottom checks if the candles form a double bottom pattern
func IsDoubleBottom(candles []Candle) bool {
	if len(candles) < 7 {
		return false
	}

	// Find the two lowest points
	var troughs []struct {
		index int
		low   float64
	}

	for i := 1; i < len(candles)-1; i++ {
		if candles[i].Low < candles[i-1].Low && candles[i].Low < candles[i+1].Low {
			troughs = append(troughs, struct {
				index int
				low   float64
			}{i, candles[i].Low})
		}
	}

	if len(troughs) < 2 {
		return false
	}

	// Find the two lowest troughs
	firstTrough := troughs[0]
	secondTrough := troughs[1]
	for _, trough := range troughs[2:] {
		if trough.low < firstTrough.low {
			secondTrough = firstTrough
			firstTrough = trough
		} else if trough.low < secondTrough.low {
			secondTrough = trough
		}
	}

	// Check if troughs are roughly at the same level (within 3% of each other)
	troughDiff := math.Abs(firstTrough.low - secondTrough.low)
	avgTroughHeight := (firstTrough.low + secondTrough.low) / 2
	if troughDiff/avgTroughHeight > 0.03 {
		return false
	}

	// Find the highest point between the troughs
	peakHigh := candles[firstTrough.index].High
	for i := firstTrough.index + 1; i < secondTrough.index; i++ {
		if candles[i].High > peakHigh {
			peakHigh = candles[i].High
		}
	}

	// Check if the pattern is complete by confirming a break above the peak
	for i := secondTrough.index + 1; i < len(candles); i++ {
		if candles[i].Close > peakHigh {
			return true
		}
	}

	return false
}

// IsTriangle checks if the candles form a triangle pattern
func IsTriangle(candles []Candle) (bool, string) {
	if len(candles) < 10 {
		return false, ""
	}

	// Find local highs and lows
	var highs, lows []float64
	for i := 1; i < len(candles)-1; i++ {
		if candles[i].High > candles[i-1].High && candles[i].High > candles[i+1].High {
			highs = append(highs, candles[i].High)
		}
		if candles[i].Low < candles[i-1].Low && candles[i].Low < candles[i+1].Low {
			lows = append(lows, candles[i].Low)
		}
	}

	if len(highs) < 2 || len(lows) < 2 {
		return false, ""
	}

	// Calculate slopes of highs and lows
	highSlope := (highs[len(highs)-1] - highs[0]) / float64(len(highs)-1)
	lowSlope := (lows[len(lows)-1] - lows[0]) / float64(len(lows)-1)

	// Check for symmetrical triangle
	if math.Abs(highSlope) < 0.001 && math.Abs(lowSlope) < 0.001 {
		return true, "Symmetrical"
	}

	// Check for ascending triangle
	if math.Abs(highSlope) < 0.001 && lowSlope > 0.001 {
		return true, "Ascending"
	}

	// Check for descending triangle
	if math.Abs(lowSlope) < 0.001 && highSlope < -0.001 {
		return true, "Descending"
	}

	return false, ""
}

// IsFlagOrPennant checks if the candles form a flag or pennant pattern
func IsFlagOrPennant(candles []Candle) (bool, string) {
	if len(candles) < 10 {
		return false, ""
	}

	// Find the trend before the pattern
	trendStart := 0
	trendEnd := len(candles) / 2
	trendSlope := (candles[trendEnd].Close - candles[trendStart].Close) / float64(trendEnd-trendStart)

	// Check if there's a strong trend
	if math.Abs(trendSlope) < 0.001 {
		return false, ""
	}

	// Find the consolidation period
	consolidationStart := trendEnd
	consolidationEnd := len(candles) - 1

	// Calculate the range of the consolidation
	consolidationHigh := candles[consolidationStart].High
	consolidationLow := candles[consolidationStart].Low
	for i := consolidationStart + 1; i <= consolidationEnd; i++ {
		if candles[i].High > consolidationHigh {
			consolidationHigh = candles[i].High
		}
		if candles[i].Low < consolidationLow {
			consolidationLow = candles[i].Low
		}
	}

	// Calculate the range of the trend
	trendHigh := candles[trendStart].High
	trendLow := candles[trendStart].Low
	for i := trendStart + 1; i <= trendEnd; i++ {
		if candles[i].High > trendHigh {
			trendHigh = candles[i].High
		}
		if candles[i].Low < trendLow {
			trendLow = candles[i].Low
		}
	}

	// Check if consolidation range is smaller than trend range
	consolidationRange := consolidationHigh - consolidationLow
	trendRange := trendHigh - trendLow
	if consolidationRange >= trendRange*0.5 {
		return false, ""
	}

	// Determine if it's a flag or pennant
	consolidationSlope := (candles[consolidationEnd].Close - candles[consolidationStart].Close) / float64(consolidationEnd-consolidationStart)
	if math.Abs(consolidationSlope) < 0.001 {
		return true, "Pennant"
	}

	// Flag has a slope opposite to the trend
	if (trendSlope > 0 && consolidationSlope < 0) || (trendSlope < 0 && consolidationSlope > 0) {
		return true, "Flag"
	}

	return false, ""
}

// IsRoundingBottom checks if the candles form a rounding bottom pattern
func IsRoundingBottom(candles []Candle) bool {
	if len(candles) < 20 {
		return false
	}

	// Calculate the moving average
	windowSize := 5
	var ma []float64
	for i := windowSize - 1; i < len(candles); i++ {
		sum := 0.0
		for j := 0; j < windowSize; j++ {
			sum += candles[i-j].Close
		}
		ma = append(ma, sum/float64(windowSize))
	}

	// Check if the moving average forms a U-shape
	// First half should be decreasing
	firstHalf := ma[:len(ma)/2]
	for i := 1; i < len(firstHalf); i++ {
		if firstHalf[i] >= firstHalf[i-1] {
			return false
		}
	}

	// Second half should be increasing
	secondHalf := ma[len(ma)/2:]
	for i := 1; i < len(secondHalf); i++ {
		if secondHalf[i] <= secondHalf[i-1] {
			return false
		}
	}

	// Check if the pattern is complete by confirming a break above the highest point
	highestPoint := candles[0].High
	for i := 1; i < len(candles); i++ {
		if candles[i].High > highestPoint {
			highestPoint = candles[i].High
		}
	}

	// Look for a break above the highest point
	for i := len(candles) - 1; i >= len(candles)/2; i-- {
		if candles[i].Close > highestPoint {
			return true
		}
	}

	return false
}

// AnalyzeChartPatterns analyzes a series of candles for chart patterns
func AnalyzeChartPatterns(candles []Candle) []PatternResult {
	if len(candles) < 7 {
		return nil
	}

	var results []PatternResult

	// Check for Head and Shoulders pattern
	if IsHeadAndShoulders(candles) {
		results = append(results, PatternResult{
			Pattern:     "Head and Shoulders",
			Confidence:  0.9,
			Description: "Strong bearish reversal pattern",
		})
	}

	// Check for Inverse Head and Shoulders pattern
	if IsInverseHeadAndShoulders(candles) {
		results = append(results, PatternResult{
			Pattern:     "Inverse Head and Shoulders",
			Confidence:  0.9,
			Description: "Strong bullish reversal pattern",
		})
	}

	// Check for Double Top pattern
	if IsDoubleTop(candles) {
		results = append(results, PatternResult{
			Pattern:     "Double Top",
			Confidence:  0.85,
			Description: "Bearish reversal pattern",
		})
	}

	// Check for Double Bottom pattern
	if IsDoubleBottom(candles) {
		results = append(results, PatternResult{
			Pattern:     "Double Bottom",
			Confidence:  0.85,
			Description: "Bullish reversal pattern",
		})
	}

	// Check for Triangle patterns
	isTriangle, triangleType := IsTriangle(candles)
	if isTriangle {
		results = append(results, PatternResult{
			Pattern:     triangleType + " Triangle",
			Confidence:  0.8,
			Description: "Price compression breakout pattern",
		})
	}

	// Check for Flag and Pennant patterns
	isFlagOrPennant, patternType := IsFlagOrPennant(candles)
	if isFlagOrPennant {
		results = append(results, PatternResult{
			Pattern:     patternType,
			Confidence:  0.8,
			Description: "Small correction in strong trend",
		})
	}

	// Check for Rounding Bottom pattern
	if IsRoundingBottom(candles) {
		results = append(results, PatternResult{
			Pattern:     "Rounding Bottom",
			Confidence:  0.85,
			Description: "Long-term bottom formation",
		})
	}

	return results
}
