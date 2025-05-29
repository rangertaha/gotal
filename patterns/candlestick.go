package patterns

import (
	"math"
)

// IsDoji checks if the candle is a doji
func IsDoji(candle Candle) bool {
	bodySize := math.Abs(candle.Close - candle.Open)
	totalSize := candle.High - candle.Low

	if totalSize == 0 {
		return false
	}

	// Doji has a small body relative to the total size
	return bodySize/totalSize < 0.1
}

// IsHammer checks if the candle is a hammer
func IsHammer(candle Candle) bool {
	bodySize := math.Abs(candle.Close - candle.Open)
	upperWick := candle.High - math.Max(candle.Open, candle.Close)
	lowerWick := math.Min(candle.Open, candle.Close) - candle.Low
	totalSize := candle.High - candle.Low

	if totalSize == 0 {
		return false
	}

	// Hammer has a small upper wick and a long lower wick
	return upperWick/totalSize < 0.3 && lowerWick/totalSize > 0.6 && bodySize/totalSize < 0.3
}

// IsShootingStar checks if the candle is a shooting star
func IsShootingStar(candle Candle) bool {
	bodySize := math.Abs(candle.Close - candle.Open)
	upperWick := candle.High - math.Max(candle.Open, candle.Close)
	lowerWick := math.Min(candle.Open, candle.Close) - candle.Low
	totalSize := candle.High - candle.Low

	if totalSize == 0 {
		return false
	}

	// Shooting star has a small lower wick and a long upper wick
	return lowerWick/totalSize < 0.3 && upperWick/totalSize > 0.6 && bodySize/totalSize < 0.3
}

// IsEngulfing checks if the two candles form an engulfing pattern
func IsEngulfing(prev, curr Candle) (bool, bool) {
	prevBodySize := math.Abs(prev.Close - prev.Open)
	currBodySize := math.Abs(curr.Close - curr.Open)

	// Check if current candle's body is larger than previous
	if currBodySize <= prevBodySize {
		return false, false
	}

	// Bullish engulfing
	if prev.Close < prev.Open && curr.Close > curr.Open &&
		curr.Open < prev.Close && curr.Close > prev.Open {
		return true, true
	}

	// Bearish engulfing
	if prev.Close > prev.Open && curr.Close < curr.Open &&
		curr.Open > prev.Close && curr.Close < prev.Open {
		return true, false
	}

	return false, false
}

// IsMorningStar checks if the three candles form a morning star pattern
func IsMorningStar(first, second, third Candle) bool {
	// First candle is bearish
	if first.Close >= first.Open {
		return false
	}

	// Second candle has a small body
	secondBodySize := math.Abs(second.Close - second.Open)
	secondTotalSize := second.High - second.Low
	if secondTotalSize == 0 || secondBodySize/secondTotalSize > 0.3 {
		return false
	}

	// Third candle is bullish
	if third.Close <= third.Open {
		return false
	}

	// Second candle gaps down from first
	if second.High >= first.Close {
		return false
	}

	// Third candle closes above midpoint of first candle
	firstMidpoint := (first.Open + first.Close) / 2
	return third.Close > firstMidpoint
}

// IsEveningStar checks if the three candles form an evening star pattern
func IsEveningStar(first, second, third Candle) bool {
	// First candle is bullish
	if first.Close <= first.Open {
		return false
	}

	// Second candle has a small body
	secondBodySize := math.Abs(second.Close - second.Open)
	secondTotalSize := second.High - second.Low
	if secondTotalSize == 0 || secondBodySize/secondTotalSize > 0.3 {
		return false
	}

	// Third candle is bearish
	if third.Close >= third.Open {
		return false
	}

	// Second candle gaps up from first
	if second.Low <= first.Close {
		return false
	}

	// Third candle closes below midpoint of first candle
	firstMidpoint := (first.Open + first.Close) / 2
	return third.Close < firstMidpoint
}

// IsHarami checks if the two candles form a harami pattern
func IsHarami(prev, curr Candle) (bool, bool) {
	prevBodySize := math.Abs(prev.Close - prev.Open)
	currBodySize := math.Abs(curr.Close - curr.Open)

	// Check if current candle's body is smaller than previous
	if currBodySize >= prevBodySize {
		return false, false
	}

	// Bullish harami
	if prev.Close < prev.Open && curr.Close > curr.Open &&
		curr.Open > prev.Close && curr.Close < prev.Open {
		return true, true
	}

	// Bearish harami
	if prev.Close > prev.Open && curr.Close < curr.Open &&
		curr.Open < prev.Close && curr.Close > prev.Open {
		return true, false
	}

	return false, false
}

// IsThreeWhiteSoldiers checks if the three candles form a three white soldiers pattern
func IsThreeWhiteSoldiers(first, second, third Candle) bool {
	// All candles must be bullish
	if first.Close <= first.Open || second.Close <= second.Open || third.Close <= third.Open {
		return false
	}

	// Each candle should open within the previous candle's body
	if second.Open <= first.Open || third.Open <= second.Open {
		return false
	}

	// Each candle should close higher than the previous candle
	if second.Close <= first.Close || third.Close <= second.Close {
		return false
	}

	// Each candle should have a significant body
	firstBodySize := first.Close - first.Open
	secondBodySize := second.Close - second.Open
	thirdBodySize := third.Close - third.Open

	return firstBodySize > 0 && secondBodySize > 0 && thirdBodySize > 0
}

// IsThreeBlackCrows checks if the three candles form a three black crows pattern
func IsThreeBlackCrows(first, second, third Candle) bool {
	// All candles must be bearish
	if first.Close >= first.Open || second.Close >= second.Open || third.Close >= third.Open {
		return false
	}

	// Each candle should open within the previous candle's body
	if second.Open >= first.Open || third.Open >= second.Open {
		return false
	}

	// Each candle should close lower than the previous candle
	if second.Close >= first.Close || third.Close >= second.Close {
		return false
	}

	// Each candle should have a significant body
	firstBodySize := first.Open - first.Close
	secondBodySize := second.Open - second.Close
	thirdBodySize := third.Open - third.Close

	return firstBodySize > 0 && secondBodySize > 0 && thirdBodySize > 0
}

// IsSpinningTop checks if the candle is a spinning top
func IsSpinningTop(candle Candle) bool {
	bodySize := math.Abs(candle.Close - candle.Open)
	upperWick := candle.High - math.Max(candle.Open, candle.Close)
	lowerWick := math.Min(candle.Open, candle.Close) - candle.Low
	totalSize := candle.High - candle.Low

	if totalSize == 0 {
		return false
	}

	// Spinning top has small body and roughly equal upper and lower wicks
	return bodySize/totalSize < 0.3 &&
		math.Abs(upperWick-lowerWick)/totalSize < 0.1 &&
		upperWick/totalSize > 0.3 && lowerWick/totalSize > 0.3
}

// IsMarubozu checks if the candle is a marubozu
func IsMarubozu(candle Candle) bool {
	bodySize := math.Abs(candle.Close - candle.Open)
	upperWick := candle.High - math.Max(candle.Open, candle.Close)
	lowerWick := math.Min(candle.Open, candle.Close) - candle.Low
	totalSize := candle.High - candle.Low

	if totalSize == 0 {
		return false
	}

	// Marubozu has no or very small wicks
	return bodySize/totalSize > 0.9 &&
		upperWick/totalSize < 0.05 &&
		lowerWick/totalSize < 0.05
}

// IsPiercing checks if the two candles form a piercing pattern
func IsPiercing(prev, curr Candle) bool {
	// First candle must be bearish
	if prev.Close >= prev.Open {
		return false
	}

	// Second candle must be bullish
	if curr.Close <= curr.Open {
		return false
	}

	// Second candle must open below the previous low
	if curr.Open >= prev.Low {
		return false
	}

	// Second candle must close above the midpoint of the first candle's body
	firstMidpoint := (prev.Open + prev.Close) / 2
	return curr.Close > firstMidpoint && curr.Close < prev.Open
}

// IsDarkCloudCover checks if the two candles form a dark cloud cover pattern
func IsDarkCloudCover(prev, curr Candle) bool {
	// First candle must be bullish
	if prev.Close <= prev.Open {
		return false
	}

	// Second candle must be bearish
	if curr.Close >= curr.Open {
		return false
	}

	// Second candle must open above the previous high
	if curr.Open <= prev.High {
		return false
	}

	// Second candle must close below the midpoint of the first candle's body
	firstMidpoint := (prev.Open + prev.Close) / 2
	return curr.Close < firstMidpoint && curr.Close > prev.Open
}

// IsHangingMan checks if the candle is a hanging man
func IsHangingMan(candle Candle) bool {
	bodySize := math.Abs(candle.Close - candle.Open)
	upperWick := candle.High - math.Max(candle.Open, candle.Close)
	lowerWick := math.Min(candle.Open, candle.Close) - candle.Low
	totalSize := candle.High - candle.Low

	if totalSize == 0 {
		return false
	}

	// Hanging man has a small body and a long lower wick
	return bodySize/totalSize < 0.3 &&
		upperWick/totalSize < 0.2 &&
		lowerWick/totalSize > 0.6
}

// IsInvertedHammer checks if the candle is an inverted hammer
func IsInvertedHammer(candle Candle) bool {
	bodySize := math.Abs(candle.Close - candle.Open)
	upperWick := candle.High - math.Max(candle.Open, candle.Close)
	lowerWick := math.Min(candle.Open, candle.Close) - candle.Low
	totalSize := candle.High - candle.Low

	if totalSize == 0 {
		return false
	}

	// Inverted hammer has a small body and a long upper wick
	return bodySize/totalSize < 0.3 &&
		upperWick/totalSize > 0.6 &&
		lowerWick/totalSize < 0.2
}

// IsBeltHold checks if the candle is a belt hold pattern
func IsBeltHold(candle Candle) (bool, bool) {
	bodySize := math.Abs(candle.Close - candle.Open)
	upperWick := candle.High - math.Max(candle.Open, candle.Close)
	lowerWick := math.Min(candle.Open, candle.Close) - candle.Low
	totalSize := candle.High - candle.Low

	if totalSize == 0 {
		return false, false
	}

	// Bullish Belt Hold
	if candle.Close > candle.Open && // Bullish candle
		upperWick/totalSize < 0.1 && // Small upper wick
		lowerWick/totalSize < 0.1 && // Small lower wick
		bodySize/totalSize > 0.8 { // Large body
		return true, true
	}

	// Bearish Belt Hold
	if candle.Close < candle.Open && // Bearish candle
		upperWick/totalSize < 0.1 && // Small upper wick
		lowerWick/totalSize < 0.1 && // Small lower wick
		bodySize/totalSize > 0.8 { // Large body
		return true, false
	}

	return false, false
}

// IsKicking checks if the two candles form a kicking pattern
func IsKicking(prev, curr Candle) (bool, bool) {
	// Check for gap between candles
	if prev.Close >= curr.Open || prev.Open <= curr.Close {
		return false, false
	}

	// Both candles should be Marubozu
	prevIsMarubozu := IsMarubozu(prev)
	currIsMarubozu := IsMarubozu(curr)

	if !prevIsMarubozu || !currIsMarubozu {
		return false, false
	}

	// Bullish Kicking
	if prev.Close < prev.Open && curr.Close > curr.Open {
		return true, true
	}

	// Bearish Kicking
	if prev.Close > prev.Open && curr.Close < curr.Open {
		return true, false
	}

	return false, false
}

// IsAbandonedBaby checks if the three candles form an abandoned baby pattern
func IsAbandonedBaby(first, second, third Candle) (bool, bool) {
	// Check for gaps
	if second.High >= first.Close || second.Low <= first.Close ||
		second.High >= third.Open || second.Low <= third.Open {
		return false, false
	}

	// Second candle should be a Doji
	if !IsDoji(second) {
		return false, false
	}

	// Bullish Abandoned Baby
	if first.Close < first.Open && // First candle is bearish
		third.Close > third.Open { // Third candle is bullish
		return true, true
	}

	// Bearish Abandoned Baby
	if first.Close > first.Open && // First candle is bullish
		third.Close < third.Open { // Third candle is bearish
		return true, false
	}

	return false, false
}

// IsDragonflyDoji checks if the candle is a dragonfly doji
func IsDragonflyDoji(candle Candle) bool {
	bodySize := math.Abs(candle.Close - candle.Open)
	upperWick := candle.High - math.Max(candle.Open, candle.Close)
	lowerWick := math.Min(candle.Open, candle.Close) - candle.Low
	totalSize := candle.High - candle.Low

	if totalSize == 0 {
		return false
	}

	// Dragonfly doji has a small body and a long lower wick
	return bodySize/totalSize < 0.1 &&
		upperWick/totalSize < 0.1 &&
		lowerWick/totalSize > 0.8
}

// IsGravestoneDoji checks if the candle is a gravestone doji
func IsGravestoneDoji(candle Candle) bool {
	bodySize := math.Abs(candle.Close - candle.Open)
	upperWick := candle.High - math.Max(candle.Open, candle.Close)
	lowerWick := math.Min(candle.Open, candle.Close) - candle.Low
	totalSize := candle.High - candle.Low

	if totalSize == 0 {
		return false
	}

	// Gravestone doji has a small body and a long upper wick
	return bodySize/totalSize < 0.1 &&
		upperWick/totalSize > 0.8 &&
		lowerWick/totalSize < 0.1
}

// IsThreeInsideUp checks if the three candles form a three inside up pattern
func IsThreeInsideUp(first, second, third Candle) bool {
	// First candle must be bearish
	if first.Close >= first.Open {
		return false
	}

	// Second candle must be bullish and contained within first candle
	if second.Close <= second.Open ||
		second.Open >= first.Open ||
		second.Close <= first.Close {
		return false
	}

	// Third candle must be bullish and close above second candle
	if third.Close <= third.Open ||
		third.Close <= second.Close {
		return false
	}

	return true
}

// IsThreeInsideDown checks if the three candles form a three inside down pattern
func IsThreeInsideDown(first, second, third Candle) bool {
	// First candle must be bullish
	if first.Close <= first.Open {
		return false
	}

	// Second candle must be bearish and contained within first candle
	if second.Close >= second.Open ||
		second.Open <= first.Open ||
		second.Close >= first.Close {
		return false
	}

	// Third candle must be bearish and close below second candle
	if third.Close >= third.Open ||
		third.Close >= second.Close {
		return false
	}

	return true
}

// IsThrusting checks if the two candles form a thrusting pattern
func IsThrusting(prev, curr Candle) bool {
	// First candle must be bearish
	if prev.Close >= prev.Open {
		return false
	}

	// Second candle must be bullish
	if curr.Close <= curr.Open {
		return false
	}

	// Second candle must open below the previous low
	if curr.Open >= prev.Low {
		return false
	}

	// Second candle must close above the midpoint of the first candle's body
	firstMidpoint := (prev.Open + prev.Close) / 2
	return curr.Close > firstMidpoint && curr.Close < prev.Open
}

// IsCounterattack checks if the two candles form a counterattack pattern
func IsCounterattack(prev, curr Candle) (bool, bool) {
	// Check if both candles have similar body sizes
	prevBodySize := math.Abs(prev.Close - prev.Open)
	currBodySize := math.Abs(curr.Close - curr.Open)

	if math.Abs(prevBodySize-currBodySize)/prevBodySize > 0.1 {
		return false, false
	}

	// Bullish Counterattack
	if prev.Close < prev.Open && curr.Close > curr.Open &&
		math.Abs(prev.Close-curr.Close) < 0.1*prevBodySize {
		return true, true
	}

	// Bearish Counterattack
	if prev.Close > prev.Open && curr.Close < curr.Open &&
		math.Abs(prev.Close-curr.Close) < 0.1*prevBodySize {
		return true, false
	}

	return false, false
}

// IsBreakaway checks if the five candles form a breakaway pattern
func IsBreakaway(first, second, third, fourth, fifth Candle) (bool, bool) {
	// First candle must be long
	firstBodySize := math.Abs(first.Close - first.Open)
	firstTotalSize := first.High - first.Low
	if firstTotalSize == 0 || firstBodySize/firstTotalSize < 0.6 {
		return false, false
	}

	// Second candle must be in the same direction as first
	if (first.Close > first.Open && second.Close <= second.Open) ||
		(first.Close < first.Open && second.Close >= second.Open) {
		return false, false
	}

	// Third candle must be in the same direction
	if (first.Close > first.Open && third.Close <= third.Open) ||
		(first.Close < first.Open && third.Close >= third.Open) {
		return false, false
	}

	// Fourth candle must be in the same direction
	if (first.Close > first.Open && fourth.Close <= fourth.Open) ||
		(first.Close < first.Open && fourth.Close >= fourth.Open) {
		return false, false
	}

	// Fifth candle must be in the opposite direction
	if (first.Close > first.Open && fifth.Close >= fifth.Open) ||
		(first.Close < first.Open && fifth.Close <= fifth.Open) {
		return false, false
	}

	// Check for decreasing body sizes
	secondBodySize := math.Abs(second.Close - second.Open)
	thirdBodySize := math.Abs(third.Close - third.Open)
	fourthBodySize := math.Abs(fourth.Close - fourth.Open)
	fifthBodySize := math.Abs(fifth.Close - fifth.Open)

	if !(firstBodySize > secondBodySize && secondBodySize > thirdBodySize &&
		thirdBodySize > fourthBodySize && fourthBodySize > fifthBodySize) {
		return false, false
	}

	// Determine if bullish or bearish
	return true, first.Close > first.Open
}

// IsMatHold checks if the five candles form a mat hold pattern
func IsMatHold(first, second, third, fourth, fifth Candle) (bool, bool) {
	// First candle must be long and bullish
	if first.Close <= first.Open {
		return false, false
	}
	firstBodySize := first.Close - first.Open
	firstTotalSize := first.High - first.Low
	if firstTotalSize == 0 || firstBodySize/firstTotalSize < 0.6 {
		return false, false
	}

	// Second, third, and fourth candles must be bearish and contained within first candle
	for _, candle := range []Candle{second, third, fourth} {
		if candle.Close >= candle.Open || // Must be bearish
			candle.High > first.High || // Must be contained
			candle.Low < first.Low {
			return false, false
		}
	}

	// Fifth candle must be bullish and close above first candle's high
	if fifth.Close <= fifth.Open || fifth.Close <= first.High {
		return false, false
	}

	return true, true
}

// IsRisingThreeMethods checks if the five candles form a rising three methods pattern
func IsRisingThreeMethods(first, second, third, fourth, fifth Candle) bool {
	// First candle must be long and bullish
	if first.Close <= first.Open {
		return false
	}
	firstBodySize := first.Close - first.Open
	firstTotalSize := first.High - first.Low
	if firstTotalSize == 0 || firstBodySize/firstTotalSize < 0.6 {
		return false
	}

	// Second, third, and fourth candles must be bearish and contained within first candle
	for _, candle := range []Candle{second, third, fourth} {
		if candle.Close >= candle.Open || // Must be bearish
			candle.High > first.High || // Must be contained
			candle.Low < first.Low {
			return false
		}
	}

	// Fifth candle must be bullish and close above first candle's high
	return fifth.Close > fifth.Open && fifth.Close > first.High
}

// IsFallingThreeMethods checks if the five candles form a falling three methods pattern
func IsFallingThreeMethods(first, second, third, fourth, fifth Candle) bool {
	// First candle must be long and bearish
	if first.Close >= first.Open {
		return false
	}
	firstBodySize := first.Open - first.Close
	firstTotalSize := first.High - first.Low
	if firstTotalSize == 0 || firstBodySize/firstTotalSize < 0.6 {
		return false
	}

	// Second, third, and fourth candles must be bullish and contained within first candle
	for _, candle := range []Candle{second, third, fourth} {
		if candle.Close <= candle.Open || // Must be bullish
			candle.High > first.High || // Must be contained
			candle.Low < first.Low {
			return false
		}
	}

	// Fifth candle must be bearish and close below first candle's low
	return fifth.Close < fifth.Open && fifth.Close < first.Low
}

// IsSeparatingLines checks if the two candles form a separating lines pattern
func IsSeparatingLines(prev, curr Candle) (bool, bool) {
	// Check if both candles have similar body sizes
	prevBodySize := math.Abs(prev.Close - prev.Open)
	currBodySize := math.Abs(curr.Close - curr.Open)

	if math.Abs(prevBodySize-currBodySize)/prevBodySize > 0.1 {
		return false, false
	}

	// Bullish Separating Lines
	if prev.Close < prev.Open && curr.Close > curr.Open &&
		math.Abs(prev.Open-curr.Open) < 0.1*prevBodySize {
		return true, true
	}

	// Bearish Separating Lines
	if prev.Close > prev.Open && curr.Close < curr.Open &&
		math.Abs(prev.Open-curr.Open) < 0.1*prevBodySize {
		return true, false
	}

	return false, false
}

// IsOnNeck checks if the two candles form an on-neck pattern
func IsOnNeck(prev, curr Candle) (bool, bool) {
	// First candle must be bearish
	if prev.Close >= prev.Open {
		return false, false
	}

	// Second candle must be bullish
	if curr.Close <= curr.Open {
		return false, false
	}

	// Second candle must open below the previous low
	if curr.Open >= prev.Low {
		return false, false
	}

	// Second candle must close near the previous low
	lowDiff := math.Abs(curr.Close - prev.Low)
	bodySize := math.Abs(prev.Close - prev.Open)
	return lowDiff/bodySize < 0.1, true
}

// IsInNeck checks if the two candles form an in-neck pattern
func IsInNeck(prev, curr Candle) (bool, bool) {
	// First candle must be bearish
	if prev.Close >= prev.Open {
		return false, false
	}

	// Second candle must be bullish
	if curr.Close <= curr.Open {
		return false, false
	}

	// Second candle must open below the previous low
	if curr.Open >= prev.Low {
		return false, false
	}

	// Second candle must close near the previous close
	closeDiff := math.Abs(curr.Close - prev.Close)
	bodySize := math.Abs(prev.Close - prev.Open)
	return closeDiff/bodySize < 0.1, true
}

// IsThreeOutsideUp checks if the three candles form a three outside up pattern
func IsThreeOutsideUp(first, second, third Candle) bool {
	// First candle must be bearish
	if first.Close >= first.Open {
		return false
	}

	// Second candle must be bullish and engulf the first candle
	if second.Close <= second.Open ||
		second.Open > first.Open ||
		second.Close < first.Close {
		return false
	}

	// Third candle must be bullish and close above second candle
	if third.Close <= third.Open ||
		third.Close <= second.Close {
		return false
	}

	return true
}

// IsThreeOutsideDown checks if the three candles form a three outside down pattern
func IsThreeOutsideDown(first, second, third Candle) bool {
	// First candle must be bullish
	if first.Close <= first.Open {
		return false
	}

	// Second candle must be bearish and engulf the first candle
	if second.Close >= second.Open ||
		second.Open < first.Open ||
		second.Close > first.Close {
		return false
	}

	// Third candle must be bearish and close below second candle
	if third.Close >= third.Open ||
		third.Close >= second.Close {
		return false
	}

	return true
}

// IsStickSandwich checks if the three candles form a stick sandwich pattern
func IsStickSandwich(first, second, third Candle) bool {
	// First and third candles must be bearish
	if first.Close >= first.Open || third.Close >= third.Open {
		return false
	}

	// Second candle must be bullish
	if second.Close <= second.Open {
		return false
	}

	// First and third candles must have similar closes
	closeDiff := math.Abs(first.Close - third.Close)
	bodySize := math.Abs(first.Close - first.Open)
	if closeDiff/bodySize > 0.1 {
		return false
	}

	// Second candle must be contained within the range of first and third
	return second.High <= first.Open && second.Low >= third.Open
}

// IsUniqueThreeRiverBottom checks if the three candles form a unique three river bottom pattern
func IsUniqueThreeRiverBottom(first, second, third Candle) bool {
	// First candle must be bearish
	if first.Close >= first.Open {
		return false
	}

	// Second candle must be bearish and have a lower low than first
	if second.Close >= second.Open || second.Low >= first.Low {
		return false
	}

	// Third candle must be bullish and close above second candle's high
	if third.Close <= third.Open || third.Close <= second.High {
		return false
	}

	// Second candle must be a hammer
	return IsHammer(second)
}

// IsThreeStarsInTheNorth checks if the three candles form a three stars in the north pattern
func IsThreeStarsInTheNorth(first, second, third Candle) bool {
	// All candles must be bullish
	if first.Close <= first.Open || second.Close <= second.Open || third.Close <= third.Open {
		return false
	}

	// Each candle must have a smaller body than the previous
	firstBodySize := first.Close - first.Open
	secondBodySize := second.Close - second.Open
	thirdBodySize := third.Close - third.Open

	if !(firstBodySize > secondBodySize && secondBodySize > thirdBodySize) {
		return false
	}

	// Each candle must have a higher high than the previous
	if !(first.High < second.High && second.High < third.High) {
		return false
	}

	// Each candle must have a higher low than the previous
	if !(first.Low < second.Low && second.Low < third.Low) {
		return false
	}

	return true
}

// IsThreeStarsInTheSouth checks if the three candles form a three stars in the south pattern
func IsThreeStarsInTheSouth(first, second, third Candle) bool {
	// All candles must be bearish
	if first.Close >= first.Open || second.Close >= second.Open || third.Close >= third.Open {
		return false
	}

	// Each candle must have a smaller body than the previous
	firstBodySize := first.Open - first.Close
	secondBodySize := second.Open - second.Close
	thirdBodySize := third.Open - third.Close

	if !(firstBodySize > secondBodySize && secondBodySize > thirdBodySize) {
		return false
	}

	// Each candle must have a lower low than the previous
	if !(first.Low > second.Low && second.Low > third.Low) {
		return false
	}

	// Each candle must have a lower high than the previous
	if !(first.High > second.High && second.High > third.High) {
		return false
	}

	return true
}

// IsTasukiGap checks if the three candles form a tasuki gap pattern
func IsTasukiGap(first, second, third Candle) (bool, bool) {
	// First two candles must be in the same direction
	if first.Close > first.Open && second.Close > second.Open {
		// Bullish Tasuki Gap
		if third.Close < third.Open && // Third candle is bearish
			third.Open > second.Close && // Opens above previous close
			third.Close > first.Close && // Closes above first candle's close
			third.Open < second.High { // Opens below second candle's high
			return true, true
		}
	} else if first.Close < first.Open && second.Close < second.Open {
		// Bearish Tasuki Gap
		if third.Close > third.Open && // Third candle is bullish
			third.Open < second.Close && // Opens below previous close
			third.Close < first.Close && // Closes below first candle's close
			third.Open > second.Low { // Opens above second candle's low
			return true, false
		}
	}
	return false, false
}

// IsThreeGapUps checks if the three candles form a three gap ups pattern
func IsThreeGapUps(first, second, third Candle) bool {
	// All candles must be bullish
	if first.Close <= first.Open || second.Close <= second.Open || third.Close <= third.Open {
		return false
	}

	// Each candle should have a gap up from the previous candle
	if second.Open <= first.High || third.Open <= second.High {
		return false
	}

	// Each candle should close higher than the previous candle
	if second.Close <= first.Close || third.Close <= second.Close {
		return false
	}

	return true
}

// IsThreeGapDowns checks if the three candles form a three gap downs pattern
func IsThreeGapDowns(first, second, third Candle) bool {
	// All candles must be bearish
	if first.Close >= first.Open || second.Close >= second.Open || third.Close >= third.Open {
		return false
	}

	// Each candle should have a gap down from the previous candle
	if second.Open >= first.Low || third.Open >= second.Low {
		return false
	}

	// Each candle should close lower than the previous candle
	if second.Close >= first.Close || third.Close >= second.Close {
		return false
	}

	return true
}

// IsConcealingBabySwallow checks if the four candles form a concealing baby swallow pattern
func IsConcealingBabySwallow(first, second, third, fourth Candle) bool {
	// First candle must be bearish
	if first.Close >= first.Open {
		return false
	}

	// Second candle must be bearish and have a lower low than first
	if second.Close >= second.Open || second.Low >= first.Low {
		return false
	}

	// Third candle must be bearish and have a lower low than second
	if third.Close >= third.Open || third.Low >= second.Low {
		return false
	}

	// Fourth candle must be bearish and have a lower low than third
	if fourth.Close >= fourth.Open || fourth.Low >= third.Low {
		return false
	}

	// Each candle should have a smaller body than the previous
	firstBodySize := first.Open - first.Close
	secondBodySize := second.Open - second.Close
	thirdBodySize := third.Open - third.Close
	fourthBodySize := fourth.Open - fourth.Close

	return firstBodySize > secondBodySize && secondBodySize > thirdBodySize && thirdBodySize > fourthBodySize
}

// IsLadderBottom checks if the five candles form a ladder bottom pattern
func IsLadderBottom(first, second, third, fourth, fifth Candle) bool {
	// First four candles must be bearish
	if first.Close >= first.Open || second.Close >= second.Open || third.Close >= third.Open || fourth.Close >= fourth.Open {
		return false
	}

	// Fifth candle must be bullish
	if fifth.Close <= fifth.Open {
		return false
	}

	// Each bearish candle should have a lower low than the previous
	if second.Low >= first.Low || third.Low >= second.Low || fourth.Low >= third.Low {
		return false
	}

	// Fifth candle should close above the high of the first candle
	if fifth.Close <= first.High {
		return false
	}

	// Each bearish candle should have a smaller body than the previous
	firstBodySize := first.Open - first.Close
	secondBodySize := second.Open - second.Close
	thirdBodySize := third.Open - third.Close
	fourthBodySize := fourth.Open - fourth.Close

	return firstBodySize > secondBodySize && secondBodySize > thirdBodySize && thirdBodySize > fourthBodySize
}

// IsThreeLineStrike checks if the four candles form a three line strike pattern
func IsThreeLineStrike(first, second, third, fourth Candle) (bool, bool) {
	// First three candles must be in the same direction
	if first.Close > first.Open && second.Close > second.Open && third.Close > third.Open {
		// Bullish Three Line Strike
		if fourth.Close < fourth.Open && // Fourth candle is bearish
			fourth.Open > third.Close && // Opens above previous close
			fourth.Close < first.Open { // Closes below first candle's open
			return true, true
		}
	} else if first.Close < first.Open && second.Close < second.Open && third.Close < third.Open {
		// Bearish Three Line Strike
		if fourth.Close > fourth.Open && // Fourth candle is bullish
			fourth.Open < third.Close && // Opens below previous close
			fourth.Close > first.Open { // Closes above first candle's open
			return true, false
		}
	}
	return false, false
}

// IsTwoCrows checks if the three candles form a two crows pattern
func IsTwoCrows(first, second, third Candle) bool {
	// First candle must be bullish
	if first.Close <= first.Open {
		return false
	}

	// Second candle must be bearish and open above first candle's high
	if second.Close >= second.Open || second.Open <= first.High {
		return false
	}

	// Third candle must be bearish and open within second candle's body
	if third.Close >= third.Open || third.Open >= second.Open || third.Open <= second.Close {
		return false
	}

	// Third candle must close below first candle's close
	return third.Close < first.Close
}

// IsThreeAdvancingWhiteSoldiers checks if the three candles form a three advancing white soldiers pattern
func IsThreeAdvancingWhiteSoldiers(first, second, third Candle) bool {
	// All candles must be bullish
	if first.Close <= first.Open || second.Close <= second.Open || third.Close <= third.Open {
		return false
	}

	// Each candle should open within the previous candle's body
	if second.Open <= first.Open || third.Open <= second.Open {
		return false
	}

	// Each candle should close higher than the previous candle
	if second.Close <= first.Close || third.Close <= second.Close {
		return false
	}

	// Each candle should have a significant body
	firstBodySize := first.Close - first.Open
	secondBodySize := second.Close - second.Open
	thirdBodySize := third.Close - third.Open

	// Each candle should have a larger body than the previous
	return firstBodySize > 0 && secondBodySize > firstBodySize && thirdBodySize > secondBodySize
}

// IsThreeDecliningBlackCrows checks if the three candles form a three declining black crows pattern
func IsThreeDecliningBlackCrows(first, second, third Candle) bool {
	// All candles must be bearish
	if first.Close >= first.Open || second.Close >= second.Open || third.Close >= third.Open {
		return false
	}

	// Each candle should open within the previous candle's body
	if second.Open >= first.Open || third.Open >= second.Open {
		return false
	}

	// Each candle should close lower than the previous candle
	if second.Close >= first.Close || third.Close >= second.Close {
		return false
	}

	// Each candle should have a significant body
	firstBodySize := first.Open - first.Close
	secondBodySize := second.Open - second.Close
	thirdBodySize := third.Open - third.Close

	// Each candle should have a larger body than the previous
	return firstBodySize > 0 && secondBodySize > firstBodySize && thirdBodySize > secondBodySize
}

// IsIdenticalThreeCrows checks if the three candles form an identical three crows pattern
func IsIdenticalThreeCrows(first, second, third Candle) bool {
	// All candles must be bearish
	if first.Close >= first.Open || second.Close >= second.Open || third.Close >= third.Open {
		return false
	}

	// All candles should have similar body sizes
	firstBodySize := first.Open - first.Close
	secondBodySize := second.Open - second.Close
	thirdBodySize := third.Open - third.Close

	// Check if body sizes are similar (within 10% of each other)
	avgBodySize := (firstBodySize + secondBodySize + thirdBodySize) / 3
	if math.Abs(firstBodySize-avgBodySize)/avgBodySize > 0.1 ||
		math.Abs(secondBodySize-avgBodySize)/avgBodySize > 0.1 ||
		math.Abs(thirdBodySize-avgBodySize)/avgBodySize > 0.1 {
		return false
	}

	// Each candle should open near the previous candle's close
	if math.Abs(second.Open-first.Close)/firstBodySize > 0.1 ||
		math.Abs(third.Open-second.Close)/secondBodySize > 0.1 {
		return false
	}

	return true
}

// IsDeliberation checks if the three candles form a deliberation pattern
func IsDeliberation(first, second, third Candle) bool {
	// First two candles must be bullish
	if first.Close <= first.Open || second.Close <= second.Open {
		return false
	}

	// Third candle must be bullish but with a smaller body
	thirdBodySize := third.Close - third.Open
	secondBodySize := second.Close - second.Open

	if third.Close <= third.Open || thirdBodySize >= secondBodySize {
		return false
	}

	// Each candle should open within the previous candle's body
	if second.Open <= first.Open || third.Open <= second.Open {
		return false
	}

	// Each candle should close higher than the previous candle
	if second.Close <= first.Close || third.Close <= second.Close {
		return false
	}

	// Third candle should have a significantly smaller body
	return thirdBodySize < 0.5*secondBodySize
}

// IsThreeStrike checks if the three candles form a three strike pattern
func IsThreeStrike(first, second, third Candle) (bool, bool) {
	// First two candles must be in the same direction
	if first.Close > first.Open && second.Close > second.Open {
		// Bullish Three Strike
		if third.Close < third.Open && // Third candle is bearish
			third.Open > second.Close && // Opens above previous close
			third.Close < first.Open { // Closes below first candle's open
			return true, true
		}
	} else if first.Close < first.Open && second.Close < second.Open {
		// Bearish Three Strike
		if third.Close > third.Open && // Third candle is bullish
			third.Open < second.Close && // Opens below previous close
			third.Close > first.Open { // Closes above first candle's open
			return true, false
		}
	}
	return false, false
}

// IsUpsideGapTwoCrows checks if the three candles form an upside gap two crows pattern
func IsUpsideGapTwoCrows(first, second, third Candle) bool {
	// First candle must be bullish
	if first.Close <= first.Open {
		return false
	}

	// Second candle must be bearish and open above first candle's high
	if second.Close >= second.Open || second.Open <= first.High {
		return false
	}

	// Third candle must be bearish and open within second candle's body
	if third.Close >= third.Open || third.Open >= second.Open || third.Open <= second.Close {
		return false
	}

	// Third candle must close below first candle's close
	return third.Close < first.Close
}

// IsThreeRiverBottom checks if the three candles form a three river bottom pattern
func IsThreeRiverBottom(first, second, third Candle) bool {
	// First candle must be bearish
	if first.Close >= first.Open {
		return false
	}

	// Second candle must be bearish and have a lower low than first
	if second.Close >= second.Open || second.Low >= first.Low {
		return false
	}

	// Third candle must be bullish and close above second candle's high
	if third.Close <= third.Open || third.Close <= second.High {
		return false
	}

	// Second candle must be a hammer
	return IsHammer(second)
}

// AnalyzeCandlestickPatterns analyzes a series of candles for patterns
func AnalyzeCandlestickPatterns(candles []Candle) []PatternResult {
	if len(candles) < 3 {
		return nil
	}

	var results []PatternResult

	// Analyze single candle patterns
	for i := 0; i < len(candles); i++ {
		if IsDoji(candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Doji",
				Confidence:  0.7,
				Description: "Indecision in the market",
			})
		}

		if IsHammer(candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Hammer",
				Confidence:  0.8,
				Description: "Potential bullish reversal",
			})
		}

		if IsShootingStar(candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Shooting Star",
				Confidence:  0.8,
				Description: "Potential bearish reversal",
			})
		}

		if IsSpinningTop(candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Spinning Top",
				Confidence:  0.7,
				Description: "Indecision in the market",
			})
		}

		if IsMarubozu(candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Marubozu",
				Confidence:  0.9,
				Description: "Strong trend continuation",
			})
		}

		if IsHangingMan(candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Hanging Man",
				Confidence:  0.8,
				Description: "Potential bearish reversal",
			})
		}

		if IsInvertedHammer(candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Inverted Hammer",
				Confidence:  0.8,
				Description: "Potential bullish reversal",
			})
		}
	}

	// Analyze two-candle patterns
	for i := 1; i < len(candles); i++ {
		isEngulfing, isBullish := IsEngulfing(candles[i-1], candles[i])
		if isEngulfing {
			pattern := "Bearish Engulfing"
			description := "Potential bearish reversal"
			if isBullish {
				pattern = "Bullish Engulfing"
				description = "Potential bullish reversal"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.85,
				Description: description,
			})
		}

		isHarami, isBullish := IsHarami(candles[i-1], candles[i])
		if isHarami {
			pattern := "Bearish Harami"
			description := "Potential bearish reversal"
			if isBullish {
				pattern = "Bullish Harami"
				description = "Potential bullish reversal"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.75,
				Description: description,
			})
		}

		if IsPiercing(candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Piercing Pattern",
				Confidence:  0.85,
				Description: "Strong bullish reversal",
			})
		}

		if IsDarkCloudCover(candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Dark Cloud Cover",
				Confidence:  0.85,
				Description: "Strong bearish reversal",
			})
		}
	}

	// Analyze three-candle patterns
	for i := 2; i < len(candles); i++ {
		if IsMorningStar(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Morning Star",
				Confidence:  0.9,
				Description: "Strong bullish reversal pattern",
			})
		}

		if IsEveningStar(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Evening Star",
				Confidence:  0.9,
				Description: "Strong bearish reversal pattern",
			})
		}

		if IsThreeWhiteSoldiers(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Three White Soldiers",
				Confidence:  0.85,
				Description: "Strong bullish continuation pattern",
			})
		}

		if IsThreeBlackCrows(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Three Black Crows",
				Confidence:  0.85,
				Description: "Strong bearish continuation pattern",
			})
		}
	}

	// Add Belt Hold pattern analysis
	for i := 0; i < len(candles); i++ {
		isBeltHold, isBullish := IsBeltHold(candles[i])
		if isBeltHold {
			pattern := "Bearish Belt Hold"
			description := "Strong bearish continuation"
			if isBullish {
				pattern = "Bullish Belt Hold"
				description = "Strong bullish continuation"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.85,
				Description: description,
			})
		}
	}

	// Add Kicking pattern analysis
	for i := 1; i < len(candles); i++ {
		isKicking, isBullish := IsKicking(candles[i-1], candles[i])
		if isKicking {
			pattern := "Bearish Kicking"
			description := "Strong bearish reversal"
			if isBullish {
				pattern = "Bullish Kicking"
				description = "Strong bullish reversal"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.9,
				Description: description,
			})
		}
	}

	// Add Abandoned Baby pattern analysis
	for i := 2; i < len(candles); i++ {
		isAbandonedBaby, isBullish := IsAbandonedBaby(candles[i-2], candles[i-1], candles[i])
		if isAbandonedBaby {
			pattern := "Bearish Abandoned Baby"
			description := "Strong bearish reversal"
			if isBullish {
				pattern = "Bullish Abandoned Baby"
				description = "Strong bullish reversal"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.95,
				Description: description,
			})
		}
	}

	// Add Dragonfly Doji pattern analysis
	for i := 0; i < len(candles); i++ {
		if IsDragonflyDoji(candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Dragonfly Doji",
				Confidence:  0.75,
				Description: "Potential bullish reversal",
			})
		}
	}

	// Add Gravestone Doji pattern analysis
	for i := 0; i < len(candles); i++ {
		if IsGravestoneDoji(candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Gravestone Doji",
				Confidence:  0.75,
				Description: "Potential bearish reversal",
			})
		}
	}

	// Add Three Inside Up/Down pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsThreeInsideUp(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Three Inside Up",
				Confidence:  0.85,
				Description: "Strong bullish reversal pattern",
			})
		}

		if IsThreeInsideDown(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Three Inside Down",
				Confidence:  0.85,
				Description: "Strong bearish reversal pattern",
			})
		}
	}

	// Add Thrusting pattern analysis
	for i := 1; i < len(candles); i++ {
		if IsThrusting(candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Thrusting Pattern",
				Confidence:  0.8,
				Description: "Potential bullish reversal",
			})
		}
	}

	// Add Counterattack pattern analysis
	for i := 1; i < len(candles); i++ {
		isCounterattack, isBullish := IsCounterattack(candles[i-1], candles[i])
		if isCounterattack {
			pattern := "Bearish Counterattack"
			description := "Potential bearish reversal"
			if isBullish {
				pattern = "Bullish Counterattack"
				description = "Potential bullish reversal"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.75,
				Description: description,
			})
		}
	}

	// Add Breakaway pattern analysis
	for i := 4; i < len(candles); i++ {
		isBreakaway, isBullish := IsBreakaway(candles[i-4], candles[i-3], candles[i-2], candles[i-1], candles[i])
		if isBreakaway {
			pattern := "Bearish Breakaway"
			description := "Strong bearish reversal"
			if isBullish {
				pattern = "Bullish Breakaway"
				description = "Strong bullish reversal"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.9,
				Description: description,
			})
		}
	}

	// Add Mat Hold pattern analysis
	for i := 4; i < len(candles); i++ {
		isMatHold, isBullish := IsMatHold(candles[i-4], candles[i-3], candles[i-2], candles[i-1], candles[i])
		if isMatHold {
			pattern := "Bearish Mat Hold"
			description := "Strong bearish continuation"
			if isBullish {
				pattern = "Bullish Mat Hold"
				description = "Strong bullish continuation"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.85,
				Description: description,
			})
		}
	}

	// Add Rising/Falling Three Methods pattern analysis
	for i := 4; i < len(candles); i++ {
		if IsRisingThreeMethods(candles[i-4], candles[i-3], candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Rising Three Methods",
				Confidence:  0.85,
				Description: "Strong bullish continuation pattern",
			})
		}

		if IsFallingThreeMethods(candles[i-4], candles[i-3], candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Falling Three Methods",
				Confidence:  0.85,
				Description: "Strong bearish continuation pattern",
			})
		}
	}

	// Add Separating Lines pattern analysis
	for i := 1; i < len(candles); i++ {
		isSeparatingLines, isBullish := IsSeparatingLines(candles[i-1], candles[i])
		if isSeparatingLines {
			pattern := "Bearish Separating Lines"
			description := "Potential bearish continuation"
			if isBullish {
				pattern = "Bullish Separating Lines"
				description = "Potential bullish continuation"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.75,
				Description: description,
			})
		}
	}

	// Add On-Neck pattern analysis
	for i := 1; i < len(candles); i++ {
		isOnNeck, isBullish := IsOnNeck(candles[i-1], candles[i])
		if isOnNeck {
			pattern := "Bearish On-Neck"
			description := "Potential bearish continuation"
			if isBullish {
				pattern = "Bullish On-Neck"
				description = "Potential bullish continuation"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.7,
				Description: description,
			})
		}
	}

	// Add In-Neck pattern analysis
	for i := 1; i < len(candles); i++ {
		isInNeck, isBullish := IsInNeck(candles[i-1], candles[i])
		if isInNeck {
			pattern := "Bearish In-Neck"
			description := "Potential bearish continuation"
			if isBullish {
				pattern = "Bullish In-Neck"
				description = "Potential bullish continuation"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.7,
				Description: description,
			})
		}
	}

	// Add Three Outside Up/Down pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsThreeOutsideUp(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Three Outside Up",
				Confidence:  0.85,
				Description: "Strong bullish reversal pattern",
			})
		}

		if IsThreeOutsideDown(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Three Outside Down",
				Confidence:  0.85,
				Description: "Strong bearish reversal pattern",
			})
		}
	}

	// Add Stick Sandwich pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsStickSandwich(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Stick Sandwich",
				Confidence:  0.8,
				Description: "Potential bullish reversal pattern",
			})
		}
	}

	// Add Unique Three River Bottom pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsUniqueThreeRiverBottom(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Unique Three River Bottom",
				Confidence:  0.85,
				Description: "Strong bullish reversal pattern",
			})
		}
	}

	// Add Three Stars in the North pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsThreeStarsInTheNorth(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Three Stars in the North",
				Confidence:  0.85,
				Description: "Strong bearish reversal pattern",
			})
		}
	}

	// Add Three Stars in the South pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsThreeStarsInTheSouth(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Three Stars in the South",
				Confidence:  0.85,
				Description: "Strong bullish reversal pattern",
			})
		}
	}

	// Add Tasuki Gap pattern analysis
	for i := 2; i < len(candles); i++ {
		isTasukiGap, isBullish := IsTasukiGap(candles[i-2], candles[i-1], candles[i])
		if isTasukiGap {
			pattern := "Bearish Tasuki Gap"
			description := "Potential bearish continuation"
			if isBullish {
				pattern = "Bullish Tasuki Gap"
				description = "Potential bullish continuation"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.8,
				Description: description,
			})
		}
	}

	// Add Three Gap Ups pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsThreeGapUps(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Three Gap Ups",
				Confidence:  0.85,
				Description: "Strong bullish continuation pattern",
			})
		}
	}

	// Add Three Gap Downs pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsThreeGapDowns(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Three Gap Downs",
				Confidence:  0.85,
				Description: "Strong bearish continuation pattern",
			})
		}
	}

	// Add Concealing Baby Swallow pattern analysis
	for i := 3; i < len(candles); i++ {
		if IsConcealingBabySwallow(candles[i-3], candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Concealing Baby Swallow",
				Confidence:  0.9,
				Description: "Strong bullish reversal pattern",
			})
		}
	}

	// Add Ladder Bottom pattern analysis
	for i := 4; i < len(candles); i++ {
		if IsLadderBottom(candles[i-4], candles[i-3], candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Ladder Bottom",
				Confidence:  0.85,
				Description: "Strong bullish reversal pattern",
			})
		}
	}

	// Add Three Line Strike pattern analysis
	for i := 3; i < len(candles); i++ {
		isThreeLineStrike, isBullish := IsThreeLineStrike(candles[i-3], candles[i-2], candles[i-1], candles[i])
		if isThreeLineStrike {
			pattern := "Bearish Three Line Strike"
			description := "Strong bearish reversal"
			if isBullish {
				pattern = "Bullish Three Line Strike"
				description = "Strong bullish reversal"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.85,
				Description: description,
			})
		}
	}

	// Add Two Crows pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsTwoCrows(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Two Crows",
				Confidence:  0.8,
				Description: "Potential bearish reversal",
			})
		}
	}

	// Add Three Advancing White Soldiers pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsThreeAdvancingWhiteSoldiers(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Three Advancing White Soldiers",
				Confidence:  0.9,
				Description: "Strong bullish continuation pattern",
			})
		}
	}

	// Add Three Declining Black Crows pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsThreeDecliningBlackCrows(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Three Declining Black Crows",
				Confidence:  0.9,
				Description: "Strong bearish continuation pattern",
			})
		}
	}

	// Add Identical Three Crows pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsIdenticalThreeCrows(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Identical Three Crows",
				Confidence:  0.85,
				Description: "Strong bearish continuation pattern",
			})
		}
	}

	// Add Deliberation pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsDeliberation(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Deliberation",
				Confidence:  0.8,
				Description: "Potential bearish reversal",
			})
		}
	}

	// Add Three Strike pattern analysis
	for i := 2; i < len(candles); i++ {
		isThreeStrike, isBullish := IsThreeStrike(candles[i-2], candles[i-1], candles[i])
		if isThreeStrike {
			pattern := "Bearish Three Strike"
			description := "Strong bearish reversal"
			if isBullish {
				pattern = "Bullish Three Strike"
				description = "Strong bullish reversal"
			}
			results = append(results, PatternResult{
				Pattern:     pattern,
				Confidence:  0.85,
				Description: description,
			})
		}
	}

	// Add Upside Gap Two Crows pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsUpsideGapTwoCrows(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Upside Gap Two Crows",
				Confidence:  0.8,
				Description: "Potential bearish reversal",
			})
		}
	}

	// Add Three River Bottom pattern analysis
	for i := 2; i < len(candles); i++ {
		if IsThreeRiverBottom(candles[i-2], candles[i-1], candles[i]) {
			results = append(results, PatternResult{
				Pattern:     "Three River Bottom",
				Confidence:  0.85,
				Description: "Strong bullish reversal pattern",
			})
		}
	}

	return results
}
