provider "generator" "default" {
	# Data parameters (required)
	market = "STOCKS" # Market (e.g. STOCKS, CRYPTO, FX, ETFs)
	symbols = ["AAPL", "MSFT", "GOOGL"] # Symbols (e.g. AAPL, MSFT, GOOGL, etc.)

	# Data parameters (optional)
	// start = "2024-01-01" # Start date (e.g. 2024-01-01)
	// end = "2025-01-01"   # End date (e.g. 2025-01-01)
	// interval = "1m"      # Interval (e.g. 1m, 5m, 15m, 30m, 1h, 4h, 1d)

	// Waveforms
	sine {
		periods = 100
		amplitude = 1.0
		phase = 0.0
	}
	square {
		periods = 100
		amplitude = 1.0
		phase = 0.0
	}
	triangle {
		periods = 100
		amplitude = 1.0	
		phase = 0.0
	}
	sawtooth {
		periods = 100
		amplitude = 1.0
		phase = 0.0
	}
}