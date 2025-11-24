package series

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Table struct {
	fields map[string][]string
}

func (t *Table) Headers() (headers []string) {
	for header := range t.fields {
		headers = append(headers, header)
	}
	return headers
}

func (t *Table) Rows() (rows [][]string) {
	for _, values := range t.fields {
		rows = append(rows, values)
	}
	return rows
}

// Save saves the Series collection to a file.
func (s *Series) Save(filename string, outputs ...string) error {
	return nil
}

func (s *Series) Print() {
	columns := s.Columns()

	if len(columns) == 0 {
		fmt.Println("No data to display")
		return
	}

	// Build headers with timestamp first
	headers := []string{"timestamp"}
	for name := range columns {
		if name != "timestamp" {
			headers = append(headers, fmt.Sprintf("%v", name))
		}
	}

	// Create table with headers
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	// Convert headers to []interface{} and make uppercase
	headerInterfaces := make([]interface{}, len(headers))
	for i, h := range headers {
		headerInterfaces[i] = strings.ToUpper(h)
	}

	tbl := table.New(headerInterfaces...)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	// Determine number of rows (use timestamp column length)
	timestampCol, exists := columns["timestamp"]
	if !exists {
		return
	}

	numRows := len(timestampCol)

	// Add rows
	for rowIdx := 0; rowIdx < numRows; rowIdx++ {
		row := make([]interface{}, len(headers))

		for colIdx, header := range headers {
			if col, exists := columns[header]; exists {
				if rowIdx < len(col) {
					row[colIdx] = col[rowIdx]
				} else {
					row[colIdx] = ""
				}
			} else {
				row[colIdx] = ""
			}
		}

		tbl.AddRow(row...)
	}

	fmt.Printf("\n")
	tbl.Print()
	fmt.Printf("\n")
}

func (s *Series) Columns() (fields map[any][]any) {
	fields = make(map[any][]any)

	for _, tick := range s.Ticks() {
		// Add timestamp
		fields["timestamp"] = append(fields["timestamp"], tick.Timestamp().Unix())

		// Add fields
		for name, value := range tick.Fields() {
			fields[name] = append(fields[name], value)
		}

		// Add tags
		for name, value := range tick.Tags() {
			fields[name] = append(fields[name], value)
		}
	}

	return fields
}
