package internal

import "fmt"

type SignalType string
type SignalFunc func(name string, value float64) float64

const (
	// Action signals
	BUY  SignalType = "BUY"
	SELL SignalType = "SELL"
	HOLD SignalType = "HOLD"

	// Momentum signals
	UP   SignalType = "UP"
	DOWN SignalType = "DOWN"
	SIDE SignalType = "SIDE"

	// Strength or speed signals
	STRONG SignalType = "STRONG" // Strong trend: Clear, consistent price movement in one direction, with small pullbacks.
	WEAK   SignalType = "WEAK"   // Weak trend: Frequent reversals, shallow moves, hard to distinguish from sideways action.
	PARA   SignalType = "PARA"   // Parabolic trend: Price accelerates rapidly, often before a sharp reversal.

	// Error signals
	ERROR SignalType = "ERR"
)

type Signal struct {
	typ  SignalType
	weight float64
	fn    SignalFunc
}
func NewSignal(typ SignalType, weight float64, fn SignalFunc) (signal *Signal) {
	signal = &Signal{
		typ:    typ,
		weight: weight,
		fn:    fn,
	}
	return
}

func (s *Signal) Type() SignalType {
	return s.typ
}

func (s *Signal) Weight() float64 {
	return s.weight
}

func (s *Signal) String() string {
	return fmt.Sprintf("%s(%f)", s.typ, s.weight)
}

func (s *Signal) Update(weight float64) {
	s.weight = weight
}

func (s *Signal) Is(typ SignalType) bool {
	return s.typ == typ
}
