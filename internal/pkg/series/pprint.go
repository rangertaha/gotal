package series

import (
	"fmt"
	"sort"

	"github.com/fatih/color"
	"github.com/rangertaha/gotal/internal/pkg/plot"
	"github.com/rodaine/table"
)

type Table struct {
	headers []string
	rows    [][]string
}

func (t *Table) AddRow(row ...string) {
	t.rows = append(t.rows, row)
}

// Save saves the Series collection to a file.
func (s *Series) Save(filename string, outputs ...string) error {
	return nil
}

// Print prints the Series collection to the console.
func (s *Series) Print() {
	fmt.Println("Series: ", s.Name())
	fmt.Println("Duration: ", s.Duration())
	fmt.Println("Timestamp: ", s.Timestamp())
	fmt.Println("Ticks: ", len(s.Ticks()))

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	t := s.Table()

	// Convert []string to []interface{} for table.New
	headers := make([]interface{}, len(t.headers))
	for i, h := range t.headers {
		headers[i] = h
	}
	
	tbl := table.New(headers...)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	// Add rows from the series data
	for _, row := range t.rows {
		rowInterface := make([]interface{}, len(row))
		for i, cell := range row {
			rowInterface[i] = cell
		}
		tbl.AddRow(rowInterface...)
	}

	tbl.Print()
}

func (s *Series) Table() (t *Table) {
	t = &Table{}

	// Get all fields and sort them for determinism
	fieldMap := s.FieldMap()
	headers := make([]string, 0, len(fieldMap))
	headers = append(headers, "timestamp")
	for f := range fieldMap {
		headers = append(headers, f)
	}
	sort.Strings(headers)
	t.headers = headers

	// Prepare rows
	numRows := s.Len()
	for i := 0; i < numRows; i++ {
		row := make([]string, len(headers))
		for j, header := range headers {
			// Prevent index out of range
			vals := fieldMap[header]
			if i >= len(vals) {
				row[j] = ""
			} else {
				row[j] = fmt.Sprintf("%v", vals[i])
			}
		}
		t.AddRow(row...)
	}
	return t

}

func (s *Series) Plot() *plot.Plot {
	dimensions := 2
	persist := false
	debug := false
	style := "lines"
	p, _ := plot.NewPlot(dimensions, persist, debug)

	for _, field := range s.Pop().FieldNames() {
		for _, c := range s.GetCol(field) {
			p.AddFunc2d(field, style, c, plot.Func2d(func(x float64) float64 {
				return x * x
			}))
		}
	}

	// plot.ResetPointGroupStyle("Sample1", "points")
	// plot.SavePlot("1.png")

	return p

}
