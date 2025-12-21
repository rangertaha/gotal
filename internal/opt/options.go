package opt

import (
	"errors"
	"fmt"
	"log"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/config"
)

type Option struct {
	body   hcl.Body
	target interface{}
	kv     map[string]any
	series map[string]*internal.Series
}

func New(target interface{}, opts ...internal.ConfigOption) internal.Configurator {
	o := &Option{
		target: target,
		body:   hcl.EmptyBody(),
		kv:     make(map[string]any),
		series: make(map[string]*internal.Series),
	}

	for _, optFunc := range opts {
		if err := optFunc(o); err != nil {
			return nil
		}
	}

	return o
}

// func (o *Option) SetSeries(key string, series internal.Series) {
// 	o.series[key] = series
// }

// func (o *Option) GetSeries(key string) internal.Series {
// 	if series, ok := o.series[key]; ok {
// 		return series
// 	}
// 	return nil
// }

func (o *Option) Get(key string, defaults ...any) any {
	if value, ok := o.kv[key]; ok {
		return value
	}
	return nil
}

func (o *Option) Set(key string, value any) {
	o.kv[key] = value

	configBody := fmt.Sprintf("%s = %s", key, value)

	if strVal, ok := value.(string); ok {
		configBody = fmt.Sprintf("%s = %s", key, strVal)
	}

	if intVal, ok := value.(int); ok {
		configBody = fmt.Sprintf("%s = %d", key, intVal)
	}

	if floatVal, ok := value.(float64); ok {
		configBody = fmt.Sprintf("%s = %f", key, floatVal)
	}

	if boolVal, ok := value.(bool); ok {
		configBody = fmt.Sprintf("%s = %t", key, boolVal)
	}

	// if the value is a series, add it to the series map
	if series, ok := value.(internal.Series); ok {
		o.series[key] = &series
	}


	o.Decode(config.CtxFunctions, "config.hcl", configBody)
}

func (o *Option) Decode(ctx *hcl.EvalContext, cfgs ...string) error {
	// path, body, json string

	// if the path
	if len(cfgs) == 1 {
		return hclsimple.DecodeFile(cfgs[0], ctx, o.target)
	}

	// if the path, body
	if len(cfgs) == 2 {
		return hclsimple.Decode(cfgs[0], []byte(cfgs[1]), config.CtxFunctions, o.target)
	}

	// if the path, body, json
	if len(cfgs) == 3 {
		parser := hclparse.NewParser()
		file, diags := parser.ParseJSON([]byte(cfgs[2]), cfgs[0])
		if diags.HasErrors() {
			return diags
		}

		// merge the body with the o.body
		o.body = hcl.MergeBodies([]hcl.Body{o.body, file.Body})

		if err := gohcl.DecodeBody(file.Body, config.CtxFunctions, o.target); err != nil {
			log.Println("Failed to decode json body", err)
			return errors.New(err.Error())
		}
		return nil
	}
	return errors.New("invalid configuration: expected 1, 2 or 3 arguments")
}

func (o *Option) Merge(body hcl.Body) {
	o.body = hcl.MergeBodies([]hcl.Body{o.body, body})
}

func WithFile(path string) internal.ConfigOption {
	return func(conf internal.Configurator) error {

		if err := conf.Decode(config.CtxFunctions, path); err != nil {
			log.Println("Failed to decode file", err)
			return err
		}

		return nil
	}
}

func WithHCL(hclBody string) internal.ConfigOption {
	return func(conf internal.Configurator) error {

		if err := conf.Decode(config.CtxFunctions, "config.hcl", hclBody); err != nil {
			log.Println("Failed to decode hcl body", err)
			return err
		}

		return nil
	}
}

func WithJSON(jsonBody string) internal.ConfigOption {
	return func(conf internal.Configurator) error {

		if err := conf.Decode(config.CtxFunctions, "config.json", "", jsonBody); err != nil {
			log.Println("Failed to decode json", err)
			return err
		}

		return nil
	}
}

func WithParams(params ...any) internal.ConfigOption {
	return func(conf internal.Configurator) error {
		return nil
	}
}

func With(key string, value string) internal.ConfigOption {
	return func(conf internal.Configurator) error {
		conf.Set(key, value)
		return nil
	}
}
