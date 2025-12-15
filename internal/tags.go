package internal

// type Tags map[string]string


// func NewTags(tags map[string]string) Tags {
// 	return tags
// }	


// func (t Tags) Get(key string, defaults ...string) (string, bool) {
// 	if val, ok := t[key]; ok {
// 		return val, true
// 	}
// 	if len(defaults) > 0 {
// 		return defaults[0], true
// 	}
// 	return "", false
// }

// func (t Tags) Set(key string, value string) {
// 	t[key] = value
// }	

// func (t Tags) Clone() *Tags {
// 	clone := make(Tags)
// 	for k, v := range t {
// 		clone[k] = v
// 	}
// 	return &clone
// }


// func (t Tags) Copy(tags ...string) *Tags {
// 	if len(tags) == 0 {
// 		return t.Clone()
// 	}

// 	output := make(Tags)
// 	for _, tag := range tags {
// 		if value, ok := t.Get(tag); ok {
// 			output[tag] = value
// 		}
// 	}
// 	return &output
// }


// func (t Tags) Merge(other Tags) {
// 	for k, v := range other {
// 		t[k] = v
// 	}
// }

// func (t Tags) Keys() []string {
// 	keys := make([]string, 0, len(t))
// 	for k := range t {
// 		keys = append(keys, k)
// 	}
// 	return keys
// }

// func (t Tags) Values() []string {
// 	values := make([]string, 0, len(t))
// 	for _, v := range t {
// 		values = append(values, v)
// 	}
// 	return values
// }

// func (t Tags) Update(key, value string) bool {
// 	if existingValue, exists := t[key]; !exists || existingValue != value {
// 		t[key] = value
// 		return true
// 	}
// 	return false
// }
