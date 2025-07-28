package metric

type Series struct {
	Name    string
	Periods int
	Metrics []*Metric
}

func NewSeries(options ...*Series) *Series {
	if len(options) == 0 {
		return &Series{
			Name:    "",
			Periods: 0,
			Metrics: make([]*Metric, 0),
		}
	}

	return &Series{
		Name:    options[0].Name,
		Periods: options[0].Periods,
		Metrics: options[0].Metrics,
	}
}

func (s *Series) Push(metric *Metric) {
	if metric.Name == "" {
		metric.Name = s.Name
	}
	s.Metrics = append(s.Metrics, metric)

	if len(s.Metrics) >= s.Periods {
		s.Metrics = s.Metrics[len(s.Metrics)-s.Periods-1:]
	}
}

func (s *Series) Pop() *Metric {
	if len(s.Metrics) > 0 {
		metric := s.Metrics[len(s.Metrics)-1]
		s.Metrics = s.Metrics[:len(s.Metrics)-1]
		return metric
	}
	return nil
}

func (s *Series) GetMetrics() []*Metric {
	return s.Metrics
}

func (s *Series) GetName() string {
	return s.Name
}

func (s *Series) GetMetricsByTag(tag string) []*Metric {
	var metrics []*Metric
	for _, metric := range s.Metrics {
		if metric.Tags[tag] != "" {
			metrics = append(metrics, metric)
		}
	}
	return metrics
}

func (s *Series) GetMetricsByTagValue(tag string, value string) []*Metric {
	var metrics []*Metric
	for _, metric := range s.Metrics {
		if metric.Tags[tag] == value {
			metrics = append(metrics, metric)
		}
	}
	return metrics
}

func (s *Series) GetMetricsByTag(tag string) []*Metric {
	var metrics []*Metric
	for _, metric := range s.Metrics {
		if metric.Tags[tag] != "" {
			metrics = append(metrics, metric)
		}
	}
	return metrics
}

