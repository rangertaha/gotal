package gen

import (
	"math"
	"math/rand"
	"time"

	"github.com/rangertaha/gotal/internal/pkg/dag"
	"github.com/rangertaha/gotal/internal/pkg/opt"
	"github.com/rangertaha/gotal/internal/pkg/series"
	"github.com/rangertaha/gotal/internal/pkg/tick"
	"github.com/rangertaha/gotal/internal/plugins/providers"
	_ "github.com/rangertaha/gotal/internal/plugins/providers/all"
)

func Func(name string, opts ...opt.OptFunc) (*series.Series, error) {
	dag := dag.New()
	fn, err := providers.Get(name, opt.WithDag(dag), opts...)
	if err != nil {
		return nil, err
	}
	dag.Add(fn)

	return dag.Execute()
}

func Sine(opts ...opt.Option) *series.Series {
	return series.New("square", opts...)
}
