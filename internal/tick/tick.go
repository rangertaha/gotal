package tick

import (
	"time"

	"github.com/rangertaha/gotal/internal"
)

type IDFunc func(t *Tick) string

// Tick represents a single market event, capturing the most granular form of market data.
type Tick struct {
	timestamp int64              // The time at which the tick was recorded
	duration  time.Duration      // The duration of the tick, typically very short
	fields    map[string]float64 // The numerical fields
	signals   map[string]float64 // The signals
	tags      map[string]string  // The tags
}

// WithSignals returns a new Tick that wraps src and merges in the given signals.
// The original tick's fields/signals/tags are preserved (signals override on key collisions).
func WithSignals(src internal.Tick, signals map[string]float64) *Tick {
	if src == nil {
		return nil
	}
	fields, _ := src.Fields()
	tags, _ := src.Tags()
	existing, _ := src.Signals()

	mergedSignals := make(map[string]float64, len(existing)+len(signals))
	for k, v := range existing {
		mergedSignals[k] = v
	}
	for k, v := range signals {
		mergedSignals[k] = v
	}

	clonedFields := make(map[string]float64, len(fields))
	for k, v := range fields {
		clonedFields[k] = v
	}
	clonedTags := make(map[string]string, len(tags))
	for k, v := range tags {
		clonedTags[k] = v
	}

	return NewTick(src.Time(), src.Duration(), clonedFields, mergedSignals, clonedTags)
}

func NewTick(timestamp time.Time, duration time.Duration, fields map[string]float64, signals map[string]float64, tags map[string]string) *Tick {
	// Truncate the timestamp to the duration
	if duration > 0 {
		timestamp = timestamp.Truncate(duration)
	}

	tick := &Tick{
		timestamp: timestamp.Unix(),
		duration:  duration,
		fields:    fields,
		signals:   signals,
		tags:      tags,
	}
	return tick
}

func (t *Tick) Time() time.Time {
	return time.Unix(t.timestamp, 0)
}

func (t *Tick) Epock() int64 {
	return t.timestamp
}

func (t *Tick) Duration() time.Duration {
	return t.duration
}

func (t *Tick) Fields(names ...string) (fields map[string]float64, ok bool) {
	if len(names) == 0 {
		return t.fields, true
	}
	fields = make(map[string]float64)
	for _, name := range names {
		if val, ok := t.fields[name]; ok {
			fields[name] = val
		} else {
			return fields, false
		}
	}
	return fields, true
}

func (t *Tick) Signals(names ...string) (signals map[string]float64, ok bool) {
	if len(names) == 0 {
		return t.signals, true
	}
	signals = make(map[string]float64)
	for _, name := range names {
		if val, ok := t.signals[name]; ok {
			signals[name] = val
		} else {
			return signals, false
		}
	}
	return signals, true
}

func (t *Tick) Tags(names ...string) (tags map[string]string, ok bool) {
	if len(names) == 0 {
		return t.tags, true
	}
	tags = make(map[string]string)
	for _, name := range names {
		if val, ok := t.tags[name]; ok {
			tags[name] = val
		} else {
			return tags, false
		}
	}
	return tags, true
}
