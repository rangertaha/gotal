package series


// Field retrieves a slice of float64 values for the specified field from all ticks.
// Returns an array containing the field values in chronological order.
func (t *Series) SetField(name string, field []float64) {
	if len(field) != t.Len() {
		panic("field length does not match series length")
	}
	for i := range t.ticks {
		t.ticks[i].SetField(name, field[i])
	}

}

// Field retrieves a slice of float64 values for the specified field from all ticks.
// Returns an array containing the field values in chronological order.
func (s *Series) Field(field string) (out []float64) {
	for _, tick := range s.ticks {
		out = append(out, tick.GetField(field))
	}
	return out
}

// Field retrieves a slice of float64 values for the specified field from all ticks.
// Returns an array containing the field values in chronological order.
func (s *Series) GetCol(field string) (column [][]float64) {
	for _, tick := range s.ticks {
		column = append(column, []float64{float64(tick.Timestamp().Unix()), tick.GetField(field)})
	}
	return column
}

// Fields retrieves a map of field names to slices of float64 values from all ticks.
// Returns a map where the keys are field names and the values are slices of field values in chronological order.
func (t *Series) FieldMap() map[string][]float64 {
	fields := make(map[string][]float64)
	for _, tick := range t.ticks {
		for k, v := range tick.Fields() {
			fields[k] = append(fields[k], v)
		}
	}
	return fields
}

func (t *Series) FieldLen(field string) int {
	return len(t.Field(field))
}

// HasField returns true if the Series collection contains the specified field.
func (t *Series) HasField(fields ...string) bool {
	for _, field := range fields {
		for _, tick := range t.ticks {
			if !tick.HasField(field) {
				return false
			}
		}
	}
	return true
}
