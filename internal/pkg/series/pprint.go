package series

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"text/tabwriter"

	"github.com/rangertaha/gotal/internal/pkg/plot"
)

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

	// Convert field map to headers and rows
	fieldMap := s.FieldMap()
	headers := make([]string, 0, len(fieldMap))
	for field := range fieldMap {
		headers = append(headers, field)
	}

	// Sort headers for consistent output
	sort.Strings(headers)

	// Create rows array
	numRows := s.Len()
	rows := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]string, len(headers))
		for j, header := range headers {
			row[j] = fmt.Sprintf("%v", fieldMap[header][i])
		}
		rows[i] = row
	}

	// Observe how the b's and the d's, despite appearing in the
	// second cell of each line, belong to different columns.
	w := tabwriter.NewWriter(os.Stdout, 3, 4, 1, '.', tabwriter.AlignRight|tabwriter.Debug)

	fmt.Fprintln(w, strings.Join(headers, "\t"))
	for _, row := range rows {
		fmt.Fprintln(w, row)
	}
	w.Flush()
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
