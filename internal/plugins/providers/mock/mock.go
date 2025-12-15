package mock

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	"github.com/rangertaha/gotal/internal/pkg/schema"
	"github.com/rangertaha/gotal/internal/plugins/providers"
)

type mock struct {
	Source string `hcl:"name,optional,default=ema"`
	Period int    `hcl:"period,optional,default=14"`
}

func NewMock(opts ...internal.PluginOptions) internal.Plugin {
	cfg := opt.New(opts...)
	m := &mock{
		Source: cfg.String("source", "ema"),
		Period: cfg.Int("period", 100),
	}

	return m
}

func (m *mock) Name() string {
	return "ema"
}

func (m *mock) Description() string {
	return "Mock provider"
}

func (m *mock) Schema() internal.Schema {
	return schema.Plugin{
		Name: "mock",
		Description: "Mock provider",
		Requires: map[string]schema.Plugin{
			"fast": {
				Name:        "ema",
				Description: "EMA dataset",
				Options: map[string]schema.Option{
					"period": {
						Name:        "period",
						Type:        schema.TypeInt,
						Required:    true,
						Description: "Period",
						Default:     100,
					},
				},
			},
			"slow": {
				Name:        "ema",
				Description: "EMA dataset",
			},
			"sma": {
				Name:        "sma",
				Description: "SMA dataset",
			},
			"rsi": {
				Name:        "rsi",
				Description: "RSI dataset",
			},
			"macd": {
				Name:        "macd",
				Description: "MACD dataset",
			},
		},
		Inputs: map[string]schema.Dataset{
			"prices": { // New dataset name
				Name:   "price",             // Dataset provider name
				Fields: []string{"value"}, // Required fields of the dataset
				Tags:   []string{"*"},     // Required tags of the dataset

			},
		},
		Outputs: map[string]schema.Dataset{
			"fast": { // New dataset name
				Name:   "ema",             // Dataset provider name
				Fields: []string{"value"}, // Required fields of the dataset
				Tags:   []string{"ema"},   // Required tags of the dataset
			},
		},
		Parameters: map[string]schema.Parameter{
			"source": {
				Name:        "source",
				Type:        schema.TypeString,
				Required:    true,
				Description: "Source",
				Default:     "ema",
			},
		},
	}
}

func (m *mock) Init() error {
	return nil
}

func (m *mock) Load() error {
	return nil
}

func (m *mock) Start() error {
	return nil
}

func (m *mock) Save() error {
	return nil
}

func (m *mock) Stop() error {
	return nil
}

func init() {
	providers.Add("mock", func(opts ...internal.PluginOptions) internal.Plugin {
		return NewMock(opts...)
	})
}
