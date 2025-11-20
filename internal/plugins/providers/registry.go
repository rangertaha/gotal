package providers

import (
	"fmt"
	"strings"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/series"
)

type NewProviderFunc func(opts ...internal.OptFunc) internal.Provider

var PROVIDERS = map[string]NewProviderFunc{}

func Add(name string, fn NewProviderFunc) error {
	name = strings.ToLower(name)

	if _, ok := PROVIDERS[name]; ok {
		return fmt.Errorf("provider %s already exists", name)
	}

	PROVIDERS[name] = fn

	return nil
}

func Get(name string) (NewProviderFunc, error) {
	name = strings.ToLower(name)

	if provider, ok := PROVIDERS[name]; ok {
		return provider, nil
	}
	return nil, fmt.Errorf("provider %s not found", name)
}

func Series(name string) (internal.SeriesFunc, error) {
	name = strings.ToLower(name)

	if provider, ok := PROVIDERS[name]; ok {
		providerFn := internal.SeriesFunc(func(input *series.Series, opts ...internal.OptFunc) (output *series.Series) {
			return provider(opts...).Compute(input)
		})
		return providerFn, nil
	}
	return nil, fmt.Errorf("provider %s not found", name)
}
