package main

import (
	"fmt"
	"time"

	"github.com/rangertaha/gotal"
	"github.com/rangertaha/gotal/indicators"
)

var metrics []*gotal.Metric = []*gotal.Metric{
	{
		Value:     10.0,
		Timestamp: time.Now(),
	},
	{
		Value:     11.0,
		Timestamp: time.Now(),
	},
	{
		Value:     12.0,
		Timestamp: time.Now(),
	},
	{
		Value:     13.0,
		Timestamp: time.Now(),
	},
	{
		Value:     14.0,
		Timestamp: time.Now(),
	},
}

func main() {

	// Batch functions operating on a metric series
	{
		result := indicators.Example(metrics, indicators.Period(3))

		fmt.Println(result)
	}

	// Stream functions operating on a metric stream
	{
		mc := make(chan *gotal.Metric)
		go sendMetrics(mc)
		result := indicators.ExampleStream(mc, indicators.Period(3))
		for metric := range result {
			fmt.Println(metric)
		}

		fmt.Println(result)
	}

}

func sendMetrics(mc chan *gotal.Metric) {
	for _, metric := range metrics {
		mc <- metric
	}
	close(mc)
}
