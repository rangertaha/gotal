package plugins

// import (
// 	"github.com/rangertaha/gotal/internal"
// 	"github.com/rangertaha/gotal/internal/plugins/providers"
// )

// type PluginFunc func(...internal.ConfigOption) (internal.Series, internal.Stream, error)

// func GetProvider(name string) PluginFunc {
// 	return func(opts ...internal.ConfigOption) (internal.Series, internal.Stream, error) {

// 		plg, err := providers.Get(name)
// 		if err != nil {
// 			return nil, nil, err
// 		}
// 		plugin, err := plg(opts...)
// 		if err != nil {
// 			return nil, nil, err
// 		}
// 		if initializer, ok := plugin.(internal.Initializer); ok {
// 			if err := initializer.Init(opts...); err != nil {
// 				return nil, nil, err
// 			}
// 		}
// 		if processor, ok := plugin.(internal.Processor); ok {
// 			return processor.Compute(), processor.Stream(), nil
// 		}

// 		return nil, nil, nil
// 	}
// }

// type Params map[string]interface{}

// type Plugin struct {j
// 	PID      string   `hcl:"id,label"`
// 	Title    string   `hcl:"name,optional"`
// 	Summary  string   `hcl:"description,optional"`
// 	// Fields   []string `hcl:"inputs,optional,default=[value]"` // input field names to compute
// 	Template []byte                               // template to compute the plugin

// 	// HCL configuration of the plugin
// 	Config hcl.Body `hcl:",remain"`
// }

// type PluginFunc internal.PluginFunc

// // New creates a new configuration
// func Get(name string, options ...func(internal.Plugin) error) (internal.Plugin, error) {
// 	s := &Plugin{
// 		PID:     "",
// 		Title:   "",
// 		Summary: "",
// 		// Fields:   []string{"value"},
// 		Template: []byte(""),
// 	}

// 	// Apply config options
// 	for _, opt := range options {
// 		err := opt(s)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return s, nil
// }

// func WithConfigFile(path string, create bool) func(*Plugin) (err error) {
// 	return func(c *Plugin) (err error) {
// 		if path == "" {
// 			return fmt.Errorf("config file is required")
// 		}
// 		// create the config file if it doesn't exist
// 		if create {
// 			if _, err := os.Create(path); err != nil {
// 				return fmt.Errorf("error creating config file: %w", err)
// 			}
// 		}

// 		// parse the config file with the context functions
// 		if err = hclsimple.DecodeFile(path, pkg.CtxFunctions, c); err != nil {
// 			return fmt.Errorf("error parsing config file: %w", err)
// 		}
// 		return nil
// 	}
// }

// func (p *Plugin) ID() string {
// 	return p.PID
// }

// func (p *Plugin) Name() string {
// 	return p.Title
// }

// func (p *Plugin) Description() string {
// 	return p.Summary
// }

// func (p *Plugin) Description() string {
// 	return p.Summary
// }

// func (p *Plugin) Set(key string, value any) {

// }

// func (p *Plugin) Get(key string) any {
// 	return nil
// }

// func (p *Plugin) Load(filename string) error {
// 	content, err := os.ReadFile(filename)
// 	if err != nil {
// 		return fmt.Errorf("failed to read file %s: %w", filename, err)
// 	}

// 	file, diags := hclsyntax.ParseConfig(content, filename, hcl.Pos{Line: 1, Column: 1})
// 	if diags.HasErrors() {
// 		return fmt.Errorf("failed to parse HCL file %s: %v", filename, diags)
// 	}
// 	cfg := Config{
// 		Plugins: []Plugin{*p},
// 	}
// 	diags = gohcl.DecodeBody(file.Body, nil, &cfg)
// 	if diags.HasErrors() {
// 		return diags
// 	}
// 	return nil
// }

// func (p *Plugin) Parse(body hcl.Body) error {
// 	diags := gohcl.DecodeBody(body, nil, p)
// 	if diags.HasErrors() {
// 		return diags
// 	}
// 	return nil
// }

// func WithHCLFile(filename string) func(internal.Plugin) error {
// 	return func(p internal.Plugin) error {
// 		content, err := os.ReadFile(filename)
// 		if err != nil {
// 			return fmt.Errorf("failed to read file %s: %w", filename, err)
// 		}

// 		file, diags := hclsyntax.ParseConfig(content, filename, hcl.Pos{Line: 1, Column: 1})
// 		if diags.HasErrors() {
// 			return fmt.Errorf("failed to parse HCL file %s: %v", filename, diags)
// 		}
// 		if plg, ok := p.(internal.Plugin); ok {
// 			cfg := Config{
// 				Plugins: []Plugin{},
// 			}
// 			diags = gohcl.DecodeBody(file.Body, nil, &cfg)
// 			if diags.HasErrors() {
// 				return diags
// 			}
// 			return nil
// 		}
// 		return fmt.Errorf("plugin is not a Plugin")
// 	}
// }

// func WithBody(body hcl.Body) func(*Plugin) error {
// 	return func(p *Plugin) error {
// 		cfg := Config{
// 			Plugins: []Plugin{*p},
// 		}
// 		diags := gohcl.DecodeBody(body, nil, &cfg)
// 		if diags.HasErrors() {
// 			return diags
// 		}
// 		return nil
// 	}
// }
