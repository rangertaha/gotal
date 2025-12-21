provider "sine" {
  market = "CRYPTO"
  symbols = ["BTC-USD"]

  sine {
    periods   = 100            // Number of periods (points)
    amplitude = 1.0            // Amplitude of the sine wave
    frequency = 1.0            // Frequency of the sine wave
    phase     = 0.0            // Phase offset (radians)
    offset    = 0.0            // Value offset
    interval  = "1m"           // Interval between points
    start     = "2021-01-01T00:00:00Z"
    end       = "2021-01-02T00:00:00Z"
  }

}


broker "coinbase" {
  market = "CRYPTO"
  symbols = ["BTC-USD"]

  sine {
    periods   = 100            // Number of periods (points)
    amplitude = 1.0            // Amplitude of the sine wave
    frequency = 1.0            // Frequency of the sine wave
    phase     = 0.0            // Phase offset (radians)
    offset    = 0.0            // Value offset
    interval  = "1m"           // Interval between points
    start     = "2021-01-01T00:00:00Z"
    end       = "2021-01-02T00:00:00Z"
  }

}