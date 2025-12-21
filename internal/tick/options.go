package tick

import (
	"time"
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


func WithTime(timestamp time.Time) TickOptions {
	return func(t *Tick) {
		t.timestamp = timestamp.Unix()
	}
}
