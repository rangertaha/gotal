package internal

type Fields map[string]float64

func NewFields(fields map[string]float64) Fields {
	return fields
}


func (f Fields) Get(key string, defaults ...float64) (float64, bool) {
	if val, ok := f[key]; ok {
		return val, true
	}
	if len(defaults) > 0 {
		return defaults[0], true
	}
	return 0, false
}

func (f Fields) Set(key string, value float64) {
	f[key] = value
}

func (f Fields) Clone() *Fields {	
	clone := make(Fields)
	for k, v := range f {
		clone[k] = v
	}
	return &clone
}

func (f Fields) Copy(fields ...string) *Fields {
	if len(fields) == 0 {
		return f.Clone()
	}

	output := make(Fields)
	for _, field := range fields {
		if value, ok := f.Get(field); ok {
			output[field] = value
		}
	}
	return &output
}


func (f Fields) Merge(other Fields) {
	for k, v := range other {
		f[k] = v
	}
}

func (f Fields) Keys() []string {
	keys := make([]string, 0, len(f))
	for k := range f {
		keys = append(keys, k)
	}
	return keys
}

func (f Fields) Values() []float64 {
	values := make([]float64, 0, len(f))
	for _, v := range f {
		values = append(values, v)
	}
	return values
}

func (f Fields) Len() int {
	return len(f)
}	

func (f Fields) IsEmpty() bool {
	return len(f) == 0
}

func (f Fields) Clear() {
	for k := range f {
		delete(f, k)
	}
}

func (f Fields) ForEach(fn func(key string, value float64)) {
	for k, v := range f {
		fn(k, v)
	}
}

func (f Fields) AsMap() map[string]float64 {
	return f
}
