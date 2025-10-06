package indicators

// import "fmt"

// type Options map[string]any

// func NewOptions() Options {
// 	return make(Options)
// }

// func (o Options) Has(key string) bool {
// 	_, ok := o[key]
// 	return ok
// }

// func (o Options) Set(key string, value any) Options {
// 	o[key] = value
// 	return o
// }


// func (o Options) Get(key string, defaultValue ...any) (any, error) {	
// 	if _, ok := o[key]; !ok {
// 		if len(defaultValue) > 0 {
// 			return defaultValue[0], nil // return the first default value
// 		}
// 		return nil, fmt.Errorf("key %s not found", key)
// 	}
// 	return o[key], nil
// }


// func (o Options) GetInt(key string, defaultValue ...int) (int, error) {
// 	// Convert []int to []any
// 	var defaults []any
// 	for _, v := range defaultValue {
// 		defaults = append(defaults, v)
// 	}
	
// 	value, err := o.Get(key, defaults...)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return value.(int), nil
// }

// func (o Options) GetString(key string, defaultValue ...string) (string, error) {
// 	// Convert []string to []any
// 	var defaults []any
// 	for _, v := range defaultValue {
// 		defaults = append(defaults, v)
// 	}
	
// 	value, err := o.Get(key, defaults...)
// 	if err != nil {
// 		return "", err
// 	}
// 	return value.(string), nil
// }
