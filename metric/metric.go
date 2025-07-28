package metric

import "time"

type Metric struct {
	Name      string
	Timestamp time.Time
	Fields    map[string]float64
	Tags      map[string]string
	Error     []error
}

func NewMetric(name string, timestamp time.Time, fields map[string]float64, tags map[string]string) *Metric {
	return &Metric{
		Name:      name,
		Timestamp: timestamp,
		Fields:    fields,
		Tags:      tags,
		Error:     make([]error, 0),
	}
}

func (m *Metric) AddField(name string, value float64) {
	m.Fields[name] = value
}

func (m *Metric) AddTag(name string, value string) {
	m.Tags[name] = value
}

func (m *Metric) GetField(name string) float64 {
	return m.Fields[name]
}

func (m *Metric) GetTag(name string) string {
	return m.Tags[name]
}

func (m *Metric) GetFields() map[string]float64 {
	return m.Fields
}

func (m *Metric) GetTags() map[string]string {
	return m.Tags
}

func (m *Metric) GetName() string {
	return m.Name
}

func (m *Metric) GetTimestamp() time.Time {
	return m.Timestamp
}
