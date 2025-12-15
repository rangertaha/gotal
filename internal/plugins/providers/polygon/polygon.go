package macd

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/opt"
	"github.com/rangertaha/gotal/internal/plugins/providers"
	"github.com/rangertaha/gotal/internal/series"
	"github.com/rangertaha/gotal/internal/tick"
)

const PluginID = "POLYGON"
const PluginName = "Polygon"
const PluginDescription = "Polygon is a provider of financial data."
const PluginHCL = `
provider "polygon" {
  api_key = "YOUR_API_KEY"              // your Polygon API key
  base_url = "https://api.polygon.io"   // your Polygon base URL
  version = "v1"                        // your Polygon version
  stock_api_url = "https://api.polygon.io/v1/stocks" // your Polygon stock API URL
  stock_quotes_api_url = "https://api.polygon.io/v1/stock/quotes" // your Polygon stock quotes API URL
  stock_quotes_api_url = "https://api.polygon.io/v1/stock/quotes" // your Polygon stock quotes API URL
  stock_quotes_api_url = "https://api.polygon.io/v1/stock/quotes" // your Polygon stock quotes API URL
}
`

type polygon struct {
	Name  string `hcl:"name,optional"`  // name of the data series
	Input string `hcl:"input,optional"` // field to compute the MACD on

}

func New(opts ...internal.PluginOptions) *polygon {
	cfg := opt.New(opts...)

	return &polygon{
		Name:  cfg.Name("macd"),
		Input: cfg.Field("value"),
	}
}

func (i *polygon) Compute(input *series.Series) (output *series.Series) {
	output = series.New(i.Name)

	return
}

func (i *polygon) Process(input *tick.Tick) (output *tick.Tick) {

	return
}

func init() {
	providers.Add("polygon", func(opts ...internal.OptFunc) internal.Provider {
		return New(opts...)
	})
}
