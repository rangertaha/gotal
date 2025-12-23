name = "prices"
start = date("2021/01/01")
end = date("2021/01/02")
interval = minutes(1)

sine "open" {
  periods   = 10            // Number of periods (points)
  amplitude = 4.0            // Amplitude of the sine wave
  frequency = 3.0            // Frequency of the sine wave
  phase     = 0.0            // Phase offset (radians)
  offset    = 0.0            // Value offset
}

square "high" {
  periods   = 10            // Number of periods (points)
  amplitude = 1.0            // Amplitude of the sine wave
  frequency = 1.0            // Frequency of the sine wave
  phase     = 0.0            // Phase offset (radians)
  offset    = 2.0            // Value offset
}

triangle "low" {
  periods   = 10            // Number of periods (points)
  amplitude = 1.0            // Amplitude of the sine wave
  frequency = 1.0            // Frequency of the sine wave
  phase     = 0.0            // Phase offset (radians)
  offset    = 3.0            // Value offset
}

sawtooth "close" {
  periods   = 10            // Number of periods (points)
  amplitude = 1.0            // Amplitude of the sine wave
  frequency = 2.0            // Frequency of the sine wave
  phase     = 0.0            // Phase offset (radians)
  offset    = 4.0            // Value offset
}
