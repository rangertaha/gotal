package indicators

import (
	"fmt"
	"strings"

	"github.com/rangertaha/gotal/internal"
)

type GroupType string

const (
	OVERLAP       GroupType = "overlap"
	MOMENTUM      GroupType = "momentum"
	TREND         GroupType = "trend"
	VOLATILITY    GroupType = "volatility"
	VOLUME        GroupType = "volume"
	CYCLE         GroupType = "cycle"
	PATTERN       GroupType = "pattern"
	PRICE         GroupType = "price"
	STATISTIC     GroupType = "statistic"
	MATHOP        GroupType = "mathop"
	MATHTRANSFORM GroupType = "mathtransform"
	SIGNAL        GroupType = "signal"
	OTHER         GroupType = "other"
)

type indicatorPlugin func(opts ...internal.ConfigOption) (internal.Indicator, error)

type entry struct {
	plugin indicatorPlugin
	stub   bool
}

var (
	INDICATORS = map[string]entry{}
	GROUPS     = map[GroupType][]string{}
)

// Add registers a real indicator implementation. A real implementation always
// wins over a stub registration regardless of init() order. Registering two
// real implementations under the same id returns an error.
func Add(id string, plugin indicatorPlugin, groups ...GroupType) error {
	key := strings.ToLower(id)
	if e, ok := INDICATORS[key]; ok && !e.stub {
		return fmt.Errorf("indicator %s already exists", id)
	}
	INDICATORS[key] = entry{plugin: plugin, stub: false}
	addToGroups(key, groups)
	return nil
}

// RegisterStub registers a placeholder for an id only if nothing is registered
// yet. The stub's constructor returns a "not implemented" error. When a real
// impl is later registered via Add, it transparently replaces the stub.
func RegisterStub(id string, groups ...GroupType) {
	key := strings.ToLower(id)
	if _, ok := INDICATORS[key]; ok {
		return
	}
	INDICATORS[key] = entry{
		plugin: func(opts ...internal.ConfigOption) (internal.Indicator, error) {
			return nil, fmt.Errorf("indicators: %q is not implemented", id)
		},
		stub: true,
	}
	addToGroups(key, groups)
}

func addToGroups(key string, groups []GroupType) {
	for _, g := range groups {
		members := GROUPS[g]
		exists := false
		for _, m := range members {
			if m == key {
				exists = true
				break
			}
		}
		if !exists {
			GROUPS[g] = append(members, key)
		}
	}
}

// Get returns the registered constructor for id. Panics if unknown.
func Get(id string) indicatorPlugin {
	key := strings.ToLower(id)
	if e, ok := INDICATORS[key]; ok {
		return e.plugin
	}
	panic(fmt.Errorf("indicator %s not found", id))
}

// Has reports whether an indicator with the given id is registered.
func Has(id string) bool {
	_, ok := INDICATORS[strings.ToLower(id)]
	return ok
}

// IsStub reports whether an indicator is registered as a stub (not implemented).
func IsStub(id string) bool {
	e, ok := INDICATORS[strings.ToLower(id)]
	return ok && e.stub
}

// Group returns the ids in a group.
func Group(g GroupType) []string {
	return append([]string(nil), GROUPS[g]...)
}

// All returns every registered id. Order is not guaranteed.
func All() []string {
	ids := make([]string, 0, len(INDICATORS))
	for id := range INDICATORS {
		ids = append(ids, id)
	}
	return ids
}

func Func(id string) internal.IndicatorFunc {
	pluginFunc := Get(id)
	return func(input internal.TimeSeries, opts ...internal.ConfigOption) (internal.TimeSeries, error) {
		plugin, err := pluginFunc(opts...)
		if err != nil {
			return nil, err
		}
		return plugin.Compute(input), nil
	}
}
