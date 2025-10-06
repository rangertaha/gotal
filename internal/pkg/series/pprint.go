package series

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"text/tabwriter"
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

// // Rows returns a slice of strings representing the ticks
// func (s *Series) Rows() [][]string {
// 	rows := [][]string{}
// 	for _, tick := range s.Series() {
// 		row := []string{tick.Timestamp().Format(time.RFC3339)}
// 		for _, field := range s.Fields() {
// 			row = append(row, fmt.Sprintf("%f", field))
// 		}
// 		rows = append(rows, row)
// 	}
// 	return rows
// }
