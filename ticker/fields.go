package ticker

import "github.com/rangertaha/gotal"

type fields map[string]vector

func NewFields(fs ...fields) fields {
	fields := make(fields)
	for _, f := range fs {
		for key, value := range f {
			fields[key] = value
		}
	}
	return fields
}

func (f fields) Get(key string) gotal.Vector {
	return f[key]
}

func (f fields) Set(key string, value []float64) {
	f[key] = value
}

func (f fields) Delete(key string) {
	delete(f, key)
}

func (f fields) Len() int {
	return len(f)
}

func (f fields) Names() []string {
	keys := make([]string, 0, len(f))
	for key := range f {
		keys = append(keys, key)
	}
	return keys
}

func (f fields) Filter(names ...string) gotal.Fields {
	filtered := make(fields)
	for _, name := range names {
		filtered[name] = f[name]
	}
	return filtered
}
