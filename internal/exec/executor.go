package exec

import (
	"github.com/rangertaha/gotal/internal"
	"github.com/rangertaha/gotal/internal/dag"
	"github.com/rangertaha/gotal/internal/diag"
	"github.com/rangertaha/gotal/internal/series"
	"github.com/rangertaha/gotal/internal/stream"
)

type Executor struct {
	dag *dag.DAG
}

func New() *Executor {
	dag := dag.New()
	return &Executor{dag: dag}
}

func (e *Executor) Add(plugin internal.Plugin, opts ...internal.OptionsFunc) {
	e.dag.Add(plugin)
}

func (e *Executor) Execute() *series.Series {
	return e.dag.Execute()
}

func (e *Executor) Stream() *stream.Stream {
	return e.dag.Stream()
}

func (e *Executor) Dia() *diag.Diagnostic {
	return e.dag.Diagnostics()
}
