package series

import "github.com/rangertaha/gotal/internal/pkg/sig"

// Tags, Signal, Meta methods

// Tags returns the tags of the Series collection.
func (t *Series) Tags() map[string]string {
	t.tags = t.ticks[len(t.ticks)-1].Tags()
	return t.tags
}

func (t *Series) SetTag(name, value string, index ...int) {
	if len(index) == 0 {
		t.tags[name] = value
	}
	for _, idx := range index {
		t.ticks[idx].SetTag(name, value)
	}

}

func (t *Series) GetTag(name string) string {
	return t.tags[name]
}

func (t *Series) HasTag(name string) bool {
	if _, ok := t.tags[name]; ok {
		return true
	}
	return false
}

// // Options returns the options map.
// func (t *Series) Meta() map[string]any {
// 	return t.meta
// }

// // SetMeta sets the meta with the given name and value.
// func (t *Series) SetMeta(name string, value any) {
// 	t.meta[name] = value
// }

// // GetMeta returns the meta with the given name.
// func (t *Series) GetMeta(name string) any {
// 	return t.meta[name]
// }

// // HasMeta returns true if the meta with the given name exists.
// func (t *Series) HasMeta(name string) bool {
// 	if _, ok := t.meta[name]; ok {
// 		return true
// 	}
// 	return false
// }

func (t *Series) Signals() map[sig.Signal]sig.Strength {
	return t.ticks[len(t.ticks)-1].Signals()
}
