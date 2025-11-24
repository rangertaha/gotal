package gen

import (
	"math"
	"math/rand"
	"time"

	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/tick"
	_ "github.com/rangertaha/gotal/internal/plugins/providers/all"
)

func Sine(duration time.Duration, amplitude float64, frequency float64, offset float64, tags map[string]string) *series.Series {
	ticks := series.New("sine")
	t := time.Now()

	// Apply the offset to the starting time
	t = t.Add(time.Duration(offset) * duration)

	// Generate one complete sine wave cycle (2π radians)
	for i := 0.0; i <= 2*math.Pi; i += frequency {
		value := (amplitude * math.Sin(i)) + amplitude
		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// Square generates a square wave series
func Square(duration time.Duration, amplitude float64, frequency float64, offset float64, tags map[string]string) *series.Series {
	ticks := series.New("square")
	t := time.Now()

	// Apply the offset to the starting time
	t = t.Add(time.Duration(offset) * duration)

	// Generate one complete square wave cycle (2π radians)
	stepCount := 0
	stepsPerHalfCycle := int(math.Pi / frequency)

	for i := 0.0; i <= 2*math.Pi; i += frequency {
		// Square wave: alternates between high and low every half cycle
		var value float64
		if (stepCount/stepsPerHalfCycle)%2 == 0 {
			value = amplitude + amplitude // High value (2 * amplitude)
		} else {
			value = 0 // Low value (0)
		}

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
		stepCount++
	}

	return ticks
}

// Triangle generates a triangle wave series
func Triangle(duration time.Duration, amplitude float64, frequency float64, offset float64, tags map[string]string) *series.Series {
	ticks := series.New("triangle")
	t := time.Now()

	// Apply the offset to the starting time
	t = t.Add(time.Duration(offset) * duration)

	// Generate one complete triangle wave cycle (2π radians)
	stepCount := 0
	totalSteps := int(2 * math.Pi / frequency)
	stepsPerHalfCycle := totalSteps / 2

	for i := 0.0; i <= 2*math.Pi; i += frequency {
		var value float64

		// Calculate position within the cycle
		cyclePos := stepCount % totalSteps

		if cyclePos <= stepsPerHalfCycle {
			// Rising phase: linear increase from 0 to 2*amplitude
			progress := float64(cyclePos) / float64(stepsPerHalfCycle)
			value = progress * (2 * amplitude)
		} else {
			// Falling phase: linear decrease from 2*amplitude to 0
			progress := float64(cyclePos-stepsPerHalfCycle) / float64(stepsPerHalfCycle)
			value = (2 * amplitude) * (1 - progress)
		}

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
		stepCount++
	}

	return ticks
}

// Sawtooth generates a sawtooth wave series (linear rise, instant drop)
func Sawtooth(duration time.Duration, amplitude float64, frequency float64, offset float64, tags map[string]string) *series.Series {
	ticks := series.New("sawtooth")
	t := time.Now()
	t = t.Add(time.Duration(offset) * duration)

	stepCount := 0
	totalSteps := int(2 * math.Pi / frequency)

	for i := 0.0; i <= 2*math.Pi; i += frequency {
		// Sawtooth: linear rise from 0 to peak, then instant drop
		cyclePos := stepCount % totalSteps
		progress := float64(cyclePos) / float64(totalSteps)
		value := progress * (2 * amplitude)

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
		stepCount++
	}

	return ticks
}

// Pulse generates a pulse wave series with adjustable duty cycle
func Pulse(duration time.Duration, amplitude float64, frequency float64, dutyCycle float64, offset float64, tags map[string]string) *series.Series {
	ticks := series.New("pulse")
	t := time.Now()
	t = t.Add(time.Duration(offset) * duration)

	stepCount := 0
	totalSteps := int(2 * math.Pi / frequency)
	highSteps := int(float64(totalSteps) * dutyCycle)

	for i := 0.0; i <= 2*math.Pi; i += frequency {
		cyclePos := stepCount % totalSteps
		var value float64
		if cyclePos < highSteps {
			value = 2 * amplitude // High value
		} else {
			value = 0 // Low value
		}

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
		stepCount++
	}

	return ticks
}

// WhiteNoise generates white noise series
func WhiteNoise(duration time.Duration, amplitude float64, samples int, tags map[string]string) *series.Series {
	ticks := series.New("white_noise")
	t := time.Now()

	for i := 0; i < samples; i++ {
		// White noise: uniform random values
		value := (rand.Float64() - 0.5) * 2 * amplitude

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// PinkNoise generates pink noise series (1/f noise)
func PinkNoise(duration time.Duration, amplitude float64, samples int, tags map[string]string) *series.Series {
	ticks := series.New("pink_noise")
	t := time.Now()

	// Simple pink noise approximation using multiple octaves
	var buffer [7]float64

	for i := 0; i < samples; i++ {
		white := rand.Float64() - 0.5

		buffer[0] = 0.99886*buffer[0] + white*0.0555179
		buffer[1] = 0.99332*buffer[1] + white*0.0750759
		buffer[2] = 0.96900*buffer[2] + white*0.1538520
		buffer[3] = 0.86650*buffer[3] + white*0.3104856
		buffer[4] = 0.55000*buffer[4] + white*0.5329522
		buffer[5] = -0.7616*buffer[5] - white*0.0168980

		value := (buffer[0] + buffer[1] + buffer[2] + buffer[3] + buffer[4] + buffer[5] + buffer[6] + white*0.5362) * amplitude
		buffer[6] = white * 0.115926

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// RandomWalk generates a random walk series
func RandomWalk(duration time.Duration, stepSize float64, drift float64, samples int, tags map[string]string) *series.Series {
	ticks := series.New("random_walk")
	t := time.Now()

	value := 100.0 // Starting price

	for i := 0; i < samples; i++ {
		// Random walk: current value + drift + random step
		step := (rand.Float64() - 0.5) * stepSize
		value += drift + step

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// GBM generates a Geometric Brownian Motion series (Black-Scholes model)
func GBM(duration time.Duration, initialPrice float64, drift float64, volatility float64, samples int, tags map[string]string) *series.Series {
	ticks := series.New("gbm")
	t := time.Now()

	value := initialPrice
	dt := 1.0 / 252.0 // Daily time step (252 trading days per year)

	for i := 0; i < samples; i++ {
		// GBM: S(t+1) = S(t) * exp((μ - σ²/2)*dt + σ*sqrt(dt)*Z)
		z := rand.NormFloat64() // Standard normal random variable
		logReturn := (drift-0.5*volatility*volatility)*dt + volatility*math.Sqrt(dt)*z
		value *= math.Exp(logReturn)

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// OrnsteinUhlenbeck generates a mean-reverting process
func OrnsteinUhlenbeck(duration time.Duration, mean float64, speed float64, volatility float64, samples int, tags map[string]string) *series.Series {
	ticks := series.New("ornstein_uhlenbeck")
	t := time.Now()

	value := mean     // Start at the mean
	dt := 1.0 / 252.0 // Daily time step

	for i := 0; i < samples; i++ {
		// OU process: dx = θ(μ - x)dt + σdW
		z := rand.NormFloat64()
		dValue := speed*(mean-value)*dt + volatility*math.Sqrt(dt)*z
		value += dValue

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// GARCH generates a GARCH process with volatility clustering
func GARCH(duration time.Duration, samples int, alpha float64, beta float64, omega float64, tags map[string]string) *series.Series {
	ticks := series.New("garch")
	t := time.Now()

	value := 100.0   // Starting price
	variance := 0.01 // Initial variance

	for i := 0; i < samples; i++ {
		// GARCH(1,1): σ²(t) = ω + α*ε²(t-1) + β*σ²(t-1)
		z := rand.NormFloat64()
		epsilon := math.Sqrt(variance) * z
		value += epsilon

		// Update variance for next period
		variance = omega + alpha*epsilon*epsilon + beta*variance

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// LinearTrend generates a linear trend with optional noise
func LinearTrend(duration time.Duration, startPrice float64, slope float64, noise float64, samples int, tags map[string]string) *series.Series {
	ticks := series.New("linear_trend")
	t := time.Now()

	for i := 0; i < samples; i++ {
		// Linear trend: y = mx + b + noise
		value := startPrice + slope*float64(i)
		if noise > 0 {
			value += (rand.Float64() - 0.5) * noise
		}

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// ExponentialTrend generates an exponential trend with optional noise
func ExponentialTrend(duration time.Duration, startPrice float64, growthRate float64, noise float64, samples int, tags map[string]string) *series.Series {
	ticks := series.New("exponential_trend")
	t := time.Now()

	for i := 0; i < samples; i++ {
		// Exponential trend: y = a * e^(rt)
		value := startPrice * math.Exp(growthRate*float64(i))
		if noise > 0 {
			value += (rand.Float64() - 0.5) * noise
		}

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// SupportResistance generates price bouncing between support and resistance levels
func SupportResistance(duration time.Duration, support float64, resistance float64, bounceStrength float64, noise float64, samples int, tags map[string]string) *series.Series {
	ticks := series.New("support_resistance")
	t := time.Now()

	value := (support + resistance) / 2 // Start in the middle
	velocity := 1.0                     // Price movement direction

	for i := 0; i < samples; i++ {
		// Bounce off support/resistance levels
		if value >= resistance {
			velocity = -bounceStrength
		} else if value <= support {
			velocity = bounceStrength
		}

		value += velocity
		if noise > 0 {
			value += (rand.Float64() - 0.5) * noise
		}

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// Step generates a step function (instant level change)
func Step(duration time.Duration, lowValue float64, highValue float64, stepTime int, samples int, tags map[string]string) *series.Series {
	ticks := series.New("step")
	t := time.Now()

	for i := 0; i < samples; i++ {
		var value float64
		if i < stepTime {
			value = lowValue
		} else {
			value = highValue
		}

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// Ramp generates a linear transition between two values
func Ramp(duration time.Duration, startValue float64, endValue float64, samples int, tags map[string]string) *series.Series {
	ticks := series.New("ramp")
	t := time.Now()

	step := (endValue - startValue) / float64(samples-1)

	for i := 0; i < samples; i++ {
		value := startValue + step*float64(i)

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// ExponentialDecay generates exponential decay
func ExponentialDecay(duration time.Duration, initial float64, halfLife float64, samples int, tags map[string]string) *series.Series {
	ticks := series.New("exponential_decay")
	t := time.Now()

	decayConstant := math.Log(2) / halfLife

	for i := 0; i < samples; i++ {
		value := initial * math.Exp(-decayConstant*float64(i))

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// ExponentialGrowth generates exponential growth
func ExponentialGrowth(duration time.Duration, initial float64, growthRate float64, samples int, tags map[string]string) *series.Series {
	ticks := series.New("exponential_growth")
	t := time.Now()

	for i := 0; i < samples; i++ {
		value := initial * math.Exp(growthRate*float64(i))

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}

// JumpDiffusion generates a process with occasional jumps
func JumpDiffusion(duration time.Duration, diffusion float64, jumpProb float64, jumpSize float64, samples int, tags map[string]string) *series.Series {
	ticks := series.New("jump_diffusion")
	t := time.Now()

	value := 100.0 // Starting price

	for i := 0; i < samples; i++ {
		// Normal diffusion
		value += (rand.Float64() - 0.5) * diffusion

		// Occasional jumps
		if rand.Float64() < jumpProb {
			jump := (rand.Float64() - 0.5) * jumpSize
			value += jump
		}

		tick := tick.New(
			tick.WithTimestamp(t),
			tick.WithDuration(duration),
			tick.WithFields(map[string]float64{"price": value}),
			tick.WithTags(tags),
		)
		ticks.Add(tick)
		t = t.Add(duration)
	}

	return ticks
}
