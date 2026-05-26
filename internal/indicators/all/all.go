// Package all blank-imports every indicator package so their init()
// registrations run. Real implementations are imported first; stubs are
// imported last so they only fill in IDs that no real implementation took.
package all

import (
	// Real implementations.
	_ "github.com/rangertaha/gotal/internal/indicators/ma"
	_ "github.com/rangertaha/gotal/internal/indicators/mathop"
	_ "github.com/rangertaha/gotal/internal/indicators/mathtransform"
	_ "github.com/rangertaha/gotal/internal/indicators/momentum"
	_ "github.com/rangertaha/gotal/internal/indicators/overlap"
	_ "github.com/rangertaha/gotal/internal/indicators/price"
	_ "github.com/rangertaha/gotal/internal/indicators/signal"
	_ "github.com/rangertaha/gotal/internal/indicators/statistic"
	_ "github.com/rangertaha/gotal/internal/indicators/volatility"
	_ "github.com/rangertaha/gotal/internal/indicators/volume"

	// Stubs (must come last — RegisterStub no-ops on already-registered ids).
	_ "github.com/rangertaha/gotal/internal/indicators/stubs"
)
