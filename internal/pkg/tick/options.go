package tick

import (
	"time"

	"github.com/rangertaha/gotal/internal/pkg/sig"
)

type TickOptions func(*Tick)

func WithDuration(duration time.Duration) TickOptions {
	return func(t *Tick) {
		t.SetDuration(duration)
	}
}

func WithFields(fields map[string]float64) TickOptions {
	return func(t *Tick) {
		t.fields = fields
	}
}

func WithTags(tags map[string]string) TickOptions {
	return func(t *Tick) {
		t.tags = tags
	}
}

func WithSignals(signals map[sig.Signal]sig.Strength) TickOptions {
	return func(t *Tick) {
		t.signals = signals
	}
}

func WithTimestamp(timestamp time.Time) TickOptions {
	return func(t *Tick) {
		t.timestamp = timestamp
	}
}
