name = "prices"
start = date("2021/01/01")
end = date("2021/01/02")
interval = minutes(1)

sine "open" {
  periods   = 100            // Number of periods (points)
  amplitude = 1.0            // Amplitude of the sine wave
  frequency = 1.0            // Frequency of the sine wave
  phase     = 0.0            // Phase offset (radians)
  offset    = 0.0            // Value offset
}

sine "high" {
  periods   = 100            // Number of periods (points)
  amplitude = 1.0            // Amplitude of the sine wave
  frequency = 1.0            // Frequency of the sine wave
  phase     = 0.0            // Phase offset (radians)
  offset    = 0.0            // Value offset
}

sine "low" {
  periods   = 100            // Number of periods (points)
  amplitude = 1.0            // Amplitude of the sine wave
  frequency = 1.0            // Frequency of the sine wave
  phase     = 0.0            // Phase offset (radians)
  offset    = 0.0            // Value offset
}

sine "close" {
  periods   = 100            // Number of periods (points)
  amplitude = 1.0            // Amplitude of the sine wave
  frequency = 1.0            // Frequency of the sine wave
  phase     = 0.0            // Phase offset (radians)
  offset    = 0.0            // Value offset
}
