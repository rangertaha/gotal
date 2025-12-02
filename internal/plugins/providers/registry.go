package providers

import (
	"fmt"
	"strings"

	"github.com/rangertaha/gotal/internal"
)

type NewPluginFunc func(opts ...internal.PluginOption) internal.Plugin

var PLUGINS = map[string]NewPluginFunc{}

func Add(name string, fn NewPluginFunc) error {
	name = strings.ToLower(name)

	if _, ok := PLUGINS[name]; ok {
		return fmt.Errorf("provider %s already exists", name)
	}

	PLUGINS[name] = fn

	return nil
}

func Get(name string) (NewPluginFunc, error) {
	name = strings.ToLower(name)

	if plugin, ok := PLUGINS[name]; ok {
		return plugin, nil
	}
	return nil, fmt.Errorf("plugin %s not found", name)
}

