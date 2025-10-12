package strategies

import (
	"fmt"
	"strings"

	"github.com/rangertaha/gotal/internal"
)

type GroupType string

type NewStrategyFunc func(opts ...internal.OptFunc) internal.Strategy

var STRATEGIES = map[string]NewStrategyFunc{}

func Add(name string, fn NewStrategyFunc) error {
	name = strings.ToLower(name)

	if _, ok := STRATEGIES[name]; ok {
		return fmt.Errorf("strategy %s already exists", name)
	}

	STRATEGIES[name] = fn

	return nil
}

func Get(name string) (NewStrategyFunc, error) {
	name = strings.ToLower(name)

	if strategy, ok := STRATEGIES[name]; ok {
		return strategy, nil
	}
	return nil, fmt.Errorf("strategy %s not found", name)
}
