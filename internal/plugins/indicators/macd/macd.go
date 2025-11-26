package macd

import (
	"fmt"
	"math"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/sig"
	"github.com/rangertaha/gotal/internal/pkg/tick"
	"github.com/rangertaha/gotal/internal/plugins/indicators"
)

type macd struct {
	Name         string `hcl:"name,optional"`   // name of the data series
	Input        string `hcl:"input,optional"`  // field to compute the MACD on
	FastPeriod   int    `hcl:"fast,optional"`   // period for fast EMA
	SlowPeriod   int    `hcl:"slow,optional"`   // period for slow EMA
	SignalPeriod int    `hcl:"signal,optional"` // period for signal line EMA

	// series is the series of ticks to compute the MACD on
	series    *series.Series
	fastEma   internal.Indicator
	slowEma   internal.Indicator
	signalEma internal.Indicator

	// MACD values
	macdLine   float64 // MACD line (fast EMA - slow EMA)
	signalLine float64 // Signal line (EMA of MACD line)
	histogram  float64 // Histogram (MACD line - Signal line)
}

func New(opts ...internal.OptFunc) *macd {
	cfg := opt.New(opts...)
	ema, err := indicators.Get("ema")
	if err != nil {
		panic(fmt.Sprintf("ema moving average function not found: %v", err))
	}
	fast := cfg.GetInt("fast", 12)
	slow := cfg.GetInt("slow", 26)
	signal := cfg.GetInt("signal", 9)

	fastEma := ema(opt.WithPeriod(fast), opt.WithInput("value"), opt.WithOutput("fast"))
	slowEma := ema(opt.WithPeriod(slow), opt.WithInput("value"), opt.WithOutput("slow"))
	signalEma := ema(opt.WithPeriod(signal), opt.WithInput("value"), opt.WithOutput("signal"))

	return &macd{
		Name:  cfg.Name("macd"),
		Input: cfg.Field("value"),

		series:    series.New(cfg.Name("macd")),
		fastEma:   fastEma,
		slowEma:   slowEma,
		signalEma: signalEma,
	}
}

// Compute computes for a given series of ticks
func (i *macd) Compute(input *series.Series) (output *series.Series) {
	output = series.New(i.Name)
	i.series.Reset()

	for _, t := range input.Ticks() {
		if t := i.Process(t); !t.IsEmpty() {
			output.Add(t)
		}
	}

	return
}

// Process computes for a given tick
func (i *macd) Process(input *tick.Tick) (output *tick.Tick) {
	// check if the series has the required field
	if !input.HasField(i.Input) {
		panic(fmt.Sprintf("%s series is missing field %v", i.Name, i.Input))
	}

	fastEma := i.fastEma.Process(input)
	slowEma := i.slowEma.Process(input)
	signalEma := i.signalEma.Process(input)
	input.SetField("fast", fastEma.GetField("fast"))
	input.SetField("slow", slowEma.GetField("slow"))
	input.SetField("signal", signalEma.GetField("signal"))

	// add the input tick to the series
	i.series.Push(input)

	// create a new empty tick
	output = tick.New()

	// if the series is not long enough, return false
	if i.series.Len() > i.SlowPeriod {
		// calculate the average while moving the window
		output = i.calculate(i.series.Shift(input))
	}

	return
}

func (i *macd) calculate(input *tick.Tick) (output *tick.Tick) {

	fastValue := input.GetField("fast")
	slowValue := input.GetField("slow")
	signalValue := input.GetField("signal")

	if math.IsNaN(fastValue) || math.IsNaN(slowValue) || math.IsNaN(signalValue) {
		return tick.New() // Return empty tick if no data available
	}

	// Calculate MACD line (fast EMA - slow EMA)
	macdValue := fastValue - slowValue

	// Calculate histogram (MACD line - Signal line)
	histogramValue := macdValue - signalValue

	// Set signals
	if macdValue > signalValue {
		input.SetSignal(sig.BULLISH, sig.MEDIUM)
	} else if macdValue < signalValue {
		input.SetSignal(sig.BEARISH, sig.MEDIUM)
	}

	// Add MACD line to signal EMA series for signal line calculation
	output = tick.New(
		tick.WithTime(input.Time()),
		tick.WithDuration(input.Duration()),
		tick.WithFields(map[string]float64{"macd": macdValue, "histogram": histogramValue, "fast": fastValue, "slow": slowValue, "signal": signalValue}),
		tick.WithTags(input.Tags()),
		tick.WithSignals(input.Signals()),
	)

	return

}

func init() {
	indicators.Add("macd", func(opts ...internal.OptFunc) internal.Indicator {
		return New(opts...)
	}, indicators.TREND)
}
