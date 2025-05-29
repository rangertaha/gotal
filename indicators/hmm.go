package indicators

import (
	"math"
	"math/rand"
)

// HMMConfig holds the configuration for the Hidden Markov Model
type HMMConfig struct {
	NumStates     int     // Number of hidden states
	NumIterations int     // Number of iterations for EM algorithm
	Convergence   float64 // Convergence threshold
	RandomSeed    int64   // Random seed for initialization
}

// HMMResult contains the results of the HMM analysis
type HMMResult struct {
	States        []int       // Most likely sequence of states
	Probabilities []float64   // State probabilities
	Transitions   [][]float64 // State transition matrix
	Emissions     [][]float64 // Emission probabilities
}

// DefaultHMMConfig returns a default configuration for the HMM
func DefaultHMMConfig() HMMConfig {
	return HMMConfig{
		NumStates:     3,    // Bullish, Bearish, and Sideways states
		NumIterations: 100,  // Maximum number of EM iterations
		Convergence:   1e-6, // Convergence threshold
		RandomSeed:    42,   // Default random seed
	}
}

// HMM implements a Hidden Markov Model for market regime detection
type HMM struct {
	config HMMConfig
	// Model parameters
	initialProbs []float64   // Initial state probabilities
	transitions  [][]float64 // State transition matrix
	emissions    [][]float64 // Emission probabilities
	means        []float64   // Mean returns for each state
	volatilities []float64   // Volatility for each state
}

// NewHMM creates a new HMM instance with the given configuration
func NewHMM(config HMMConfig) *HMM {
	rand.Seed(config.RandomSeed)

	hmm := &HMM{
		config:       config,
		initialProbs: make([]float64, config.NumStates),
		transitions:  make([][]float64, config.NumStates),
		emissions:    make([][]float64, config.NumStates),
		means:        make([]float64, config.NumStates),
		volatilities: make([]float64, config.NumStates),
	}

	// Initialize parameters
	for i := range hmm.transitions {
		hmm.transitions[i] = make([]float64, config.NumStates)
		hmm.emissions[i] = make([]float64, config.NumStates)
	}

	// Initialize with random values
	hmm.initializeRandom()

	return hmm
}

// initializeRandom initializes the HMM parameters with random values
func (hmm *HMM) initializeRandom() {
	// Initialize initial probabilities
	sum := 0.0
	for i := range hmm.initialProbs {
		hmm.initialProbs[i] = rand.Float64()
		sum += hmm.initialProbs[i]
	}
	for i := range hmm.initialProbs {
		hmm.initialProbs[i] /= sum
	}

	// Initialize transition matrix
	for i := range hmm.transitions {
		sum := 0.0
		for j := range hmm.transitions[i] {
			hmm.transitions[i][j] = rand.Float64()
			sum += hmm.transitions[i][j]
		}
		for j := range hmm.transitions[i] {
			hmm.transitions[i][j] /= sum
		}
	}

	// Initialize means and volatilities
	for i := range hmm.means {
		hmm.means[i] = rand.NormFloat64() * 0.01
		hmm.volatilities[i] = 0.01 + rand.Float64()*0.02
	}
}

// Fit trains the HMM on the given price data
func (hmm *HMM) Fit(prices []float64) {
	// Calculate returns
	returns := make([]float64, len(prices)-1)
	for i := 0; i < len(returns); i++ {
		returns[i] = math.Log(prices[i+1] / prices[i])
	}

	// Run EM algorithm
	for iter := 0; iter < hmm.config.NumIterations; iter++ {
		// E-step: Calculate forward and backward probabilities
		alpha := hmm.forward(returns)
		beta := hmm.backward(returns)
		gamma := hmm.calculateGamma(alpha, beta)
		xi := hmm.calculateXi(returns, alpha, beta)

		// M-step: Update parameters
		oldMeans := make([]float64, len(hmm.means))
		copy(oldMeans, hmm.means)

		hmm.updateParameters(returns, gamma, xi)

		// Check convergence
		if hmm.checkConvergence(oldMeans) {
			break
		}
	}
}

// forward calculates the forward probabilities
func (hmm *HMM) forward(returns []float64) [][]float64 {
	T := len(returns)
	N := hmm.config.NumStates
	alpha := make([][]float64, T)

	// Initialize first time step
	alpha[0] = make([]float64, N)
	for i := 0; i < N; i++ {
		alpha[0][i] = hmm.initialProbs[i] * hmm.emissionProb(returns[0], i)
	}

	// Forward recursion
	for t := 1; t < T; t++ {
		alpha[t] = make([]float64, N)
		for j := 0; j < N; j++ {
			sum := 0.0
			for i := 0; i < N; i++ {
				sum += alpha[t-1][i] * hmm.transitions[i][j]
			}
			alpha[t][j] = sum * hmm.emissionProb(returns[t], j)
		}
	}

	return alpha
}

// backward calculates the backward probabilities
func (hmm *HMM) backward(returns []float64) [][]float64 {
	T := len(returns)
	N := hmm.config.NumStates
	beta := make([][]float64, T)

	// Initialize last time step
	beta[T-1] = make([]float64, N)
	for i := 0; i < N; i++ {
		beta[T-1][i] = 1.0
	}

	// Backward recursion
	for t := T - 2; t >= 0; t-- {
		beta[t] = make([]float64, N)
		for i := 0; i < N; i++ {
			sum := 0.0
			for j := 0; j < N; j++ {
				sum += hmm.transitions[i][j] * hmm.emissionProb(returns[t+1], j) * beta[t+1][j]
			}
			beta[t][i] = sum
		}
	}

	return beta
}

// calculateGamma calculates the state probabilities
func (hmm *HMM) calculateGamma(alpha, beta [][]float64) [][]float64 {
	T := len(alpha)
	N := hmm.config.NumStates
	gamma := make([][]float64, T)

	for t := 0; t < T; t++ {
		gamma[t] = make([]float64, N)
		sum := 0.0
		for i := 0; i < N; i++ {
			gamma[t][i] = alpha[t][i] * beta[t][i]
			sum += gamma[t][i]
		}
		for i := 0; i < N; i++ {
			gamma[t][i] /= sum
		}
	}

	return gamma
}

// calculateXi calculates the joint state probabilities
func (hmm *HMM) calculateXi(returns []float64, alpha, beta [][]float64) [][][]float64 {
	T := len(returns)
	N := hmm.config.NumStates
	xi := make([][][]float64, T-1)

	for t := 0; t < T-1; t++ {
		xi[t] = make([][]float64, N)
		for i := 0; i < N; i++ {
			xi[t][i] = make([]float64, N)
			for j := 0; j < N; j++ {
				xi[t][i][j] = alpha[t][i] * hmm.transitions[i][j] *
					hmm.emissionProb(returns[t+1], j) * beta[t+1][j]
			}
		}
	}

	return xi
}

// updateParameters updates the HMM parameters using the EM algorithm
func (hmm *HMM) updateParameters(returns []float64, gamma [][]float64, xi [][][]float64) {
	T := len(returns)
	N := hmm.config.NumStates

	// Update initial probabilities
	for i := 0; i < N; i++ {
		hmm.initialProbs[i] = gamma[0][i]
	}

	// Update transition matrix
	for i := 0; i < N; i++ {
		denominator := 0.0
		for t := 0; t < T-1; t++ {
			denominator += gamma[t][i]
		}
		for j := 0; j < N; j++ {
			numerator := 0.0
			for t := 0; t < T-1; t++ {
				numerator += xi[t][i][j]
			}
			hmm.transitions[i][j] = numerator / denominator
		}
	}

	// Update means and volatilities
	for i := 0; i < N; i++ {
		denominator := 0.0
		for t := 0; t < T; t++ {
			denominator += gamma[t][i]
		}

		// Update mean
		numerator := 0.0
		for t := 0; t < T; t++ {
			numerator += gamma[t][i] * returns[t]
		}
		hmm.means[i] = numerator / denominator

		// Update volatility
		numerator = 0.0
		for t := 0; t < T; t++ {
			numerator += gamma[t][i] * math.Pow(returns[t]-hmm.means[i], 2)
		}
		hmm.volatilities[i] = math.Sqrt(numerator / denominator)
	}
}

// checkConvergence checks if the EM algorithm has converged
func (hmm *HMM) checkConvergence(oldMeans []float64) bool {
	for i := range hmm.means {
		if math.Abs(hmm.means[i]-oldMeans[i]) > hmm.config.Convergence {
			return false
		}
	}
	return true
}

// emissionProb calculates the emission probability for a given return and state
func (hmm *HMM) emissionProb(return_ float64, state int) float64 {
	z := (return_ - hmm.means[state]) / hmm.volatilities[state]
	return math.Exp(-0.5*z*z) / (hmm.volatilities[state] * math.Sqrt(2*math.Pi))
}

// Predict returns the most likely sequence of states and their probabilities
func (hmm *HMM) Predict(prices []float64) HMMResult {
	// Calculate returns
	returns := make([]float64, len(prices)-1)
	for i := 0; i < len(returns); i++ {
		returns[i] = math.Log(prices[i+1] / prices[i])
	}

	// Calculate forward probabilities
	alpha := hmm.forward(returns)
	beta := hmm.backward(returns)
	gamma := hmm.calculateGamma(alpha, beta)

	// Find most likely states
	states := make([]int, len(returns))
	probabilities := make([]float64, len(returns))
	for t := range returns {
		maxProb := 0.0
		maxState := 0
		for i := 0; i < hmm.config.NumStates; i++ {
			if gamma[t][i] > maxProb {
				maxProb = gamma[t][i]
				maxState = i
			}
		}
		states[t] = maxState
		probabilities[t] = maxProb
	}

	return HMMResult{
		States:        states,
		Probabilities: probabilities,
		Transitions:   hmm.transitions,
		Emissions:     hmm.emissions,
	}
}

// GetStateDescription returns a description of the given state
func (hmm *HMM) GetStateDescription(state int) string {
	switch state {
	case 0:
		return "Bullish"
	case 1:
		return "Bearish"
	case 2:
		return "Sideways"
	default:
		return "Unknown"
	}
}



// // Create HMM with default configuration
// config := indicators.DefaultHMMConfig()
// hmm := indicators.NewHMM(config)

// // Fit the model to your price data
// hmm.Fit(prices)

// // Make predictions
// result := hmm.Predict(prices)

// // Access the results
// states := result.States        // Most likely state sequence
// probs := result.Probabilities  // State probabilities
// transitions := result.Transitions  // Transition matrix

