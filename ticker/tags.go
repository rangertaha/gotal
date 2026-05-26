package ticker

import "github.com/rangertaha/gotal"

type tags map[string]string

func (t tags) Get(name string) string {
	return t[name]
}

func (t tags) Set(name, value string) {
	t[name] = value
}

func (t tags) Delete(key string) {
	delete(t, key)
}

func (t tags) Names() []string {
	keys := make([]string, 0, len(t))
	for key := range t {
		keys = append(keys, key)
	}
	return keys
}

func (t tags) Map() map[string]string {
	return t
}

func (t tags) Values() []string {
	values := make([]string, 0, len(t))
	for _, value := range t {
		values = append(values, value)
	}
	return values
}

func (t tags) Filter(names ...string) gotal.Tags {
	filtered := make(tags)
	for _, name := range names {
		filtered[name] = t[name]
	}
	return filtered
}

func (t tags) Len() int {
	return len(t)
}
