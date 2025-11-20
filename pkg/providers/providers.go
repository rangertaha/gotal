package providers


import (
	"fmt"

	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	p "github.com/rangertaha/gotal/internal/plugins/providers"
	_ "github.com/rangertaha/gotal/internal/plugins/providers/all"
)

var (
	err error

	// Indicator options
	With         = opt.With
	OnField      = opt.WithField
	OnFields     = opt.WithFields
	WithField    = opt.WithField
	WithFields   = opt.WithFields
	WithOutput   = opt.WithOutput
	WithPeriod   = opt.WithPeriod
	WithDuration = opt.WithDuration

	// for MACD,
	WithSlowPeriod   = opt.WithSlowPeriod
	WithFastPeriod   = opt.WithFastPeriod
	WithSignalPeriod = opt.WithSignalPeriod

	WithMAType = opt.WithMAType

	// Mock indicator
	Mock internal.SeriesFunc


)

func init() {

	// Mock indicator
	Mock, err = p.Series("mock")




	if err != nil {
		fmt.Println("Error initializing indicators:", err)
		panic(err)
	}
}
