package loaders

import (
	"fmt"
	"strings"

	"github.com/rangertaha/gotal/internal"
)


type NewBrokerFunc func(opts ...internal.OptFunc) internal.Broker

var BROKERS = map[string]NewBrokerFunc{}

func Add(name string, fn NewBrokerFunc) error {
	name = strings.ToLower(name)

	if _, ok := BROKERS[name]; ok {
		return fmt.Errorf("broker %s already exists", name)
	}

	BROKERS[name] = fn

	return nil
}

func Get(name string) (NewBrokerFunc, error) {
	name = strings.ToLower(name)

	if broker, ok := BROKERS[name]; ok {
		return broker, nil
	}
	return nil, fmt.Errorf("broker %s not found", name)
}
