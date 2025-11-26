package tick

import (
	"math"
	"time"

	"github.com/rangertaha/gotal/internal/pkg/sig"
)

type IDFunc func(t *Tick) string

// Tick represents a single market event, capturing the most granular form of market data.
type Tick struct {
	uuid      string                      // The unique identifier for the tick
	timestamp int64                       // The time at which the tick was recorded
	duration  time.Duration               // The duration of the tick, typically very short
	fields    map[string]float64          // The numerical fields
	tags      map[string]string           // The classification tags, e.g. market, symbol, exchange, currency, etc.
	signals   map[sig.Signal]sig.Strength // The signals, e.g. buy, sell, etc.
	idFunc    IDFunc
}

func New(opts ...TickOptions) *Tick {
	tick := &Tick{
		uuid:      "",
		timestamp: time.Now().Unix(),
		duration:  0,
		fields:    map[string]float64{},
		tags:      map[string]string{},
		signals:   map[sig.Signal]sig.Strength{},
		idFunc:    nil,
	}
	for _, opt := range opts {
		opt(tick)
	}

	if tick.idFunc == nil {
		tick.SetIDFunc(func(t *Tick) string {
			return t.uuid
		})
	}

	tick.SetID(tick.idFunc(tick))

	return tick
}

func (t *Tick) ID() string {
	return t.uuid
}

func (t *Tick) SetID(id string) {
	t.uuid = id
}

func (t *Tick) SetIDFunc(idFunc IDFunc) {
	t.idFunc = idFunc
}

func (t *Tick) Time() time.Time {
	return time.Unix(t.timestamp, 0)
}

func (t *Tick) Epock() int64 {
	return t.timestamp
}

func (t *Tick) SetEpock(epock int64) {
	t.timestamp = epock
}

func (t *Tick) SetTime(timestamp time.Time) {
	t.timestamp = timestamp.Unix()

	// Truncate the timestamp to the duration
	if t.duration > 0 {
		// Get the timestamp as a time.Time
		timestamp := time.Unix(t.timestamp, 0)

		// Truncate the timestamp to the duration
		t.timestamp = timestamp.Truncate(t.duration).Unix()
	}
}

func (t *Tick) Duration() time.Duration {
	return t.duration
}

func (t *Tick) SetDuration(duration time.Duration) {
	t.duration = duration

	// Truncate the timestamp to the duration
	if t.duration > 0 {
		// Get the timestamp as a time.Time
		timestamp := time.Unix(t.timestamp, 0)
		
		// Truncate the timestamp to the duration
		t.timestamp = timestamp.Truncate(t.duration).Unix()
	}
}

// Field methods
// ------------------------------------------------------------

func (t *Tick) Fields() map[string]float64 {
	return t.fields
}

func (t *Tick) GetField(key string) float64 {
	if val, ok := t.fields[key]; ok {
		return val
	}
	return math.NaN()
}

func (t *Tick) SetField(key string, value float64) {
	t.fields[key] = value
}

func (t *Tick) SetFields(fields map[string]float64) {
	t.fields = fields
}

func (t *Tick) HasField(key string) bool {
	_, ok := t.fields[key]
	return ok
}

func (t *Tick) HasFields(keys ...string) bool {
	for _, key := range keys {
		if !t.HasField(key) {
			return false
		}
	}
	return true
}

func (t *Tick) RemoveField(key string) {
	delete(t.fields, key)
}

func (t *Tick) FieldNames() []string {
	keys := make([]string, 0, len(t.fields))
	for k := range t.fields {
		keys = append(keys, k)
	}
	return keys
}

func (t *Tick) Len() int {
	return len(t.fields)
}

func (t *Tick) IsEmpty() bool {
	if t == nil || t.fields == nil {
		return true
	}
	return len(t.fields) == 0
}

func (t *Tick) Reset() {
	t.fields = map[string]float64{}
}

func (t *Tick) ForEach(fn func(key string, value float64) float64) {
	for k, v := range t.fields {
		t.fields[k] = fn(k, v)
	}
}

// Tag methods
// ------------------------------------------------------------

func (t *Tick) Tags() map[string]string {
	return t.tags
}

func (t *Tick) HasTag(key string) bool {
	_, ok := t.tags[key]
	return ok
}

func (t *Tick) GetTag(key string) string {
	if val, ok := t.tags[key]; ok {
		return val
	}
	return ""
}

func (t *Tick) SetTag(key string, value string) {
	t.tags[key] = value
}

func (t *Tick) SetTags(tags map[string]string) {
	t.tags = tags
}

func (t *Tick) UpdateTags(tags map[string]string) {
	for k, v := range tags {
		t.tags[k] = v
	}
}

func (t *Tick) RemoveTag(key string) {
	delete(t.tags, key)
}

func (t *Tick) TagNames() []string {
	keys := make([]string, 0, len(t.tags))
	for k := range t.tags {
		keys = append(keys, k)
	}
	return keys
}

// Signal methods
// ------------------------------------------------------------

func (t *Tick) Signals() map[sig.Signal]sig.Strength {
	return t.signals
}

func (t *Tick) HasSignal(key sig.Signal) bool {
	_, ok := t.signals[key]
	return ok
}

func (t *Tick) GetSignal(key sig.Signal) sig.Strength {
	if val, ok := t.signals[key]; ok {
		return val
	}
	return 0
}

func (t *Tick) SetSignal(key sig.Signal, value sig.Strength) {
	t.signals[key] = value
}

func (t *Tick) SetSignals(signals map[sig.Signal]sig.Strength) {
	t.signals = signals
}

func (t *Tick) RemoveSignal(key sig.Signal) {
	delete(t.signals, key)
}

func (t *Tick) SignalNames() []sig.Signal {
	keys := make([]sig.Signal, 0, len(t.signals))
	for k := range t.signals {
		keys = append(keys, k)
	}
	return keys
}

// Other methods
// ------------------------------------------------------------

func (t *Tick) Clone() *Tick {
	clone := make(map[string]float64)
	for k, v := range t.fields {
		clone[k] = v
	}
	return &Tick{
		fields:    clone,
		tags:      t.tags,
		signals:   t.signals,
		uuid:      t.uuid,
		timestamp: t.timestamp,
		duration:  t.duration,
	}
}

func (t *Tick) Update(other *Tick) *Tick {
	for k, v := range other.fields {
		t.fields[k] = v
	}
	for k, v := range other.tags {
		t.tags[k] = v
	}
	for k, v := range other.signals {
		t.signals[k] = v
	}
	t.SetEpock(other.timestamp)
	t.SetDuration(other.duration)
	t.SetID(other.uuid)
	t.SetIDFunc(other.idFunc)

	// update the uuid
	t.SetID(t.idFunc(t))

	return t
}

func (t *Tick) Spawn(opts ...TickOptions) *Tick {
	tick := &Tick{
		uuid:      t.uuid,
		timestamp: t.Time().Add(t.duration).Unix(),
		duration:  t.duration,
		fields:    map[string]float64{},
		tags:      t.tags,
		signals:   map[sig.Signal]sig.Strength{},
		idFunc:    t.idFunc,
	}
	for _, opt := range opts {
		opt(tick)
	}
	tick.SetID(tick.idFunc(tick))

	return tick
}
