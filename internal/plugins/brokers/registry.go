package brokers

import (
	"fmt"
	"strings"

	"github.com/rangertaha/gotal/internal"
)

type BrokerPlugin func(opts ...internal.ConfigOption) (internal.Plugin, error)

var BROKERS = map[string]BrokerPlugin{}


func Add(id string, plugin BrokerPlugin) error {
	id = strings.ToLower(id)

	if _, ok := BROKERS[id]; ok {
		return fmt.Errorf("broker %s already exists", id)
	}
	BROKERS[id] = plugin

	return nil
}

func Get(id string) (BrokerPlugin, error) {
	id = strings.ToLower(id)

	if plugin, ok := BROKERS[id]; ok {
		return plugin, nil
	}
	return nil, fmt.Errorf("broker %s not found", id)
}

func All() (brokerPlugins []BrokerPlugin) {
	for _, plugin := range BROKERS {
		brokerPlugins = append(brokerPlugins, plugin)
	}
	return brokerPlugins
}


func Batch(name string) internal.BatchFunc {
	return func(opts ...internal.ConfigOption) (internal.Series, error) {

		plg, err := Get(name)
		if err != nil {
			return nil, err
		}
		plugin, err := plg(opts...)
		if err != nil {
			return nil,  err
		}

		if processor, ok := plugin.(internal.Processor); ok {
			return processor.Compute(), nil
		}

		return nil, nil
	}
}

func Stream(name string) internal.StreamFunc {
	return func(opts ...internal.ConfigOption) (internal.Stream, error) {


		plg, err := Get(name)
		if err != nil {
			return nil, err
		}
		plugin, err := plg(opts...)
		if err != nil {
			return nil,  err
		}


		if processor, ok := plugin.(internal.Processor); ok {
			return processor.Stream(), nil
		}

		return nil, nil
	}
}