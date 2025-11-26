package series

import (
	"fmt"
	"os"
	"os/exec"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Plot struct {
	title  string
	xLabel string
	yLabel string
	height vg.Length
	width  vg.Length
	fields []string
	series *Series
	plot   *plot.Plot
}

func NewPlot(series *Series) *Plot {
	return &Plot{
		series: series,
		title:  series.Name(),
		xLabel: "Time",
		yLabel: "Values",
		height: 4 * vg.Inch,
		width:  8 * vg.Inch,
		fields: series.FieldNames(),
		plot:   plot.New(),
	}
}

func (p *Plot) SetTitle(title string) {
	p.plot.Title.Text = title
}

func (p *Plot) SetXLabel(xLabel string) {
	p.plot.X.Label.Text = xLabel
}

func (p *Plot) SetYLabel(yLabel string) {
	p.plot.Y.Label.Text = yLabel
}

func (p *Plot) SetFields(fields ...string) {
	p.fields = fields
	lines := []any{}
	for _, field := range fields {
		lines = append(lines, field, p.getPoints(field))
	}
	err := plotutil.AddLinePoints(p.plot, lines...)
	if err != nil {
		panic(err)
	}

}

func (p *Plot) getPoints(field string) plotter.XYs {
	pts := make(plotter.XYs, p.series.Len())
	for i, tick := range p.series.Ticks() {
		if tick.HasField(field) && !tick.IsEmpty() {
			pts[i].X = float64(tick.Epock())
			pts[i].Y = tick.GetField(field)
		}
	}
	return pts
}

func (p *Plot) Show(width, height int) error {
	// Save plot as temporary PNG file
	tmpfile, err := os.CreateTemp("/tmp", fmt.Sprintf("*-plot-%s.png", p.series.Name()))
	if err != nil {
		return err
	}
	defer tmpfile.Close()

	// Save plot to temporary file
	if err := p.Save(tmpfile.Name(), width, height); err != nil {
		return err
	}

	// Open temporary image using 'xdg-open' on Linux (floating window via default program)
	cmd := exec.Command("xdg-open", tmpfile.Name())
	err = cmd.Start()
	if err != nil {
		return err
	}

	return nil
}

// Save plot to file
func (p *Plot) Save(path string, width, height int) error {
	return p.plot.Save(vg.Length(width)*vg.Inch, vg.Length(height)*vg.Inch, path)
}

// Create a new plot
func (s *Series) Plot(fields ...string) (*Plot, error) {
	p := NewPlot(s)

	if len(fields) > 0 {
		p.SetFields(fields...)
	} else {
		p.SetFields(s.FieldNames()...)
	}

	return p, nil
}
