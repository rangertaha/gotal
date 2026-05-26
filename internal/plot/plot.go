package plot

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/rangertaha/gotal/internal"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Plot struct {
	title  string
	xLabel string
	yLabel string
	fields []string
	series internal.TimeSeries
	plot   *plot.Plot
}

func New(series internal.TimeSeries, fields ...string) *Plot {
	if len(fields) == 0 {
		fields = series.Fields().Names()
	}
	p := &Plot{
		series: series,
		title:  series.Name(),
		xLabel: "Time",
		yLabel: "Values",
		fields: fields,
		plot:   plot.New(),
	}
	p.plot.Title.Text = p.title
	p.plot.X.Label.Text = p.xLabel
	p.plot.Y.Label.Text = p.yLabel

	lines := []any{}
	for _, field := range fields {
		lines = append(lines, field, p.points(field))
	}
	if len(lines) > 0 {
		if err := plotutil.AddLinePoints(p.plot, lines...); err != nil {
			panic(err)
		}
	}
	return p
}

func (p *Plot) points(field string) plotter.XYs {
	ticks := p.series.Ticks()
	vec := p.series.Fields().Get(field)
	var values []float64
	if vec != nil {
		values = vec.Values()
	}
	pts := make(plotter.XYs, 0, len(ticks))
	for i, tick := range ticks {
		x := float64(tick.Time().Unix())
		var y float64
		if i < len(values) {
			y = values[i]
		} else if vals, ok := tick.Fields(field); ok {
			y = vals[field]
		}
		pts = append(pts, plotter.XY{X: x, Y: y})
	}
	return pts
}

func (p *Plot) Show(width, height int) error {
	tmpfile, err := os.CreateTemp("/tmp", fmt.Sprintf("*-plot-%s.png", p.series.Name()))
	if err != nil {
		return err
	}
	defer tmpfile.Close()

	if err := p.Save(tmpfile.Name(), width, height); err != nil {
		return err
	}

	cmd := exec.Command("xdg-open", tmpfile.Name())
	return cmd.Start()
}

func (p *Plot) Save(path string, width, height int) error {
	return p.plot.Save(vg.Length(width)*vg.Inch, vg.Length(height)*vg.Inch, path)
}
