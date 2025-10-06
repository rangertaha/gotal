package indicators

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/rangertaha/gotal/internal"
// )

// type GroupType string

// const (
// 	TREND      GroupType = "trend"
// 	MOMENTUM   GroupType = "momentum"
// 	VOLATILITY GroupType = "volatility"
// 	VOLUME     GroupType = "volume"
// 	CYCLE      GroupType = "cycle"
// 	OTHER      GroupType = "other"
// )

// type indicatorFuncTypes struct {
// 	Series SeriesIndicatorFunc
// 	Stream IndicatorStreamFunc
// }

// var INDICATORS = map[string]*indicatorFuncTypes{}

// var GROUPS = map[GroupType][]*indicatorFuncTypes{}

// func Add(name string, indicator internal.Indicator, groups ...GroupType) error {
// 	name = strings.ToLower(name)
	
// 	if _, ok := INDICATORS[name]; ok {
// 		return fmt.Errorf("indicator %s already exists", name)
// 	}
// 	INDICATORS[name] = &indicatorFuncTypes{
// 		Series: seriesFunc,
// 		Stream: streamFunc,
// 	}
// 	for _, group := range groups {
// 		if _, ok := GROUPS[group]; !ok {
// 			GROUPS[group] = []*indicatorFuncTypes{}
// 		}
// 		GROUPS[group] = append(GROUPS[group], INDICATORS[name])
// 	}
// 	return nil
// }

// func Get(name string) (*indicatorFuncTypes, error) {
// 	name = strings.ToLower(name)

// 	if indicator, ok := INDICATORS[name]; ok {
// 		return indicator, nil
// 	}
// 	return nil, fmt.Errorf("indicator %s not found", name)
// }

// func GetSeries(name string) (SeriesIndicatorFunc, error) {
// 	indicator, err := Get(name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return indicator.Series, nil
// }

// func GetStream(name string) (IndicatorStreamFunc, error) {
// 	indicator, err := Get(name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return indicator.Stream, nil
// }

// // Group returns all indicators in a group
// func Group(group GroupType) ([]*indicatorFuncTypes, error) {
// 	return GROUPS[group], nil
// }
