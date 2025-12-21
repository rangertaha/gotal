package config

import (
	"time"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/gocty"
)

const YYYYMMDD = "2006/01/02"

var CtxFunctions *hcl.EvalContext = &hcl.EvalContext{
	Functions: map[string]function.Function{
		"seconds": SecondsFunc,
		"minutes": MinutesFunc,
		"hours":   HoursFunc,
		"days":    DaysFunc,
		"date":    DateFunc,
	},
}

var SecondsFunc = function.New(&function.Spec{
	Description: "Returns the given seconds",
	Params: []function.Parameter{
		{
			Name:             "num",
			Type:             cty.Number,
			AllowDynamicType: true,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		var seconds int64
		if err := gocty.FromCtyValue(args[0], &seconds); err != nil {
			return cty.UnknownVal(cty.String), err
		}

		return cty.NumberIntVal(int64(time.Duration(seconds) * time.Second)), nil
	},
})

var MinutesFunc = function.New(&function.Spec{
	Description: "Returns the given minutes",
	Params: []function.Parameter{
		{
			Name:             "num",
			Type:             cty.Number,
			AllowDynamicType: true,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		var minutes int64
		if err := gocty.FromCtyValue(args[0], &minutes); err != nil {
			return cty.UnknownVal(cty.String), err
		}

		return cty.NumberIntVal(int64(time.Duration(minutes) * time.Minute)), nil
	},
})

var HoursFunc = function.New(&function.Spec{
	Description: "Returns the given hours",
	Params: []function.Parameter{
		{
			Name:             "num",
			Type:             cty.Number,
			AllowDynamicType: true,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		var hours int64
		if err := gocty.FromCtyValue(args[0], &hours); err != nil {
			return cty.UnknownVal(cty.String), err
		}

		return cty.NumberIntVal(int64(time.Duration(hours) * time.Hour)), nil
	},
})

var DaysFunc = function.New(&function.Spec{
	Description: "Returns the given hours",
	Params: []function.Parameter{
		{
			Name:             "num",
			Type:             cty.Number,
			AllowDynamicType: true,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		var days int64
		if err := gocty.FromCtyValue(args[0], &days); err != nil {
			return cty.UnknownVal(cty.String), err
		}

		return cty.NumberIntVal(int64(time.Duration(days) * (time.Hour * 24))), nil
	},
})

var DateFunc = function.New(&function.Spec{
	Description: "Returns a date in int64 format",
	Params: []function.Parameter{
		{
			Name:             "date",
			Type:             cty.String,
			AllowDynamicType: true,
		},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		var date string
		if err := gocty.FromCtyValue(args[0], &date); err != nil {
			return cty.UnknownVal(cty.String), err
		}

		t, err := time.Parse(YYYYMMDD, date)
		if err != nil {
			return cty.UnknownVal(cty.String), err
		}

		return cty.NumberIntVal(t.Unix()), nil
	},
})
