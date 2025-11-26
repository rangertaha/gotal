package series

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

// Save series collection to a file.
func (s *Series) Save(paths ...string) error {
	// Map of file extensions to save functions.
	EXTENSIONS := map[string]func(path string) error{
		".csv":   s.SaveCSV,
		".json":  s.SaveJSON,
		".jsonl": s.SaveJSONL,
	}

	In := func(name string, extensions map[string]func(path string) error) (bool, func(path string) error) {
		for extension := range extensions {
			if strings.HasSuffix(name, extension) {
				return true, extensions[extension]
			}
		}
		return false, nil
	}

	if len(paths) == 0 {
		paths = []string{s.Name() + ".csv"}
	}
	errs := make([]error, 0)
	for _, path := range paths {
		path = strings.TrimSpace(path)
		if found, saveFunc := In(path, EXTENSIONS); found {
			errs = append(errs, saveFunc(path))
		} else {
			errs = append(errs, fmt.Errorf("unsupported file format: %s", path))
		}
	}

	return errors.Join(errs...)
}

// columns returns a map of column names to values.
func (s *Series) columns() (columns map[any][]any) {
	columns = make(map[any][]any)

	for _, tick := range s.Ticks() {
		// Add timestamp
		columns["timestamp"] = append(columns["timestamp"], tick.Timestamp().Unix())

		// Add fields
		for name, value := range tick.Fields() {
			columns[name] = append(columns[name], value)
		}

		// Add tags
		for name, value := range tick.Tags() {
			columns[name] = append(columns[name], value)
		}
	}

	if len(columns) == 0 {
		fmt.Println("No data to display")
		return
	}

	return columns
}

// Print series collection to the console.
func (s *Series) Print() {
	columns := s.columns()

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

// SaveCSV series collection to a CSV file.
func (s *Series) SaveCSV(path string) error {
	columns := s.columns()
	// Save the Series as a CSV file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Determine column headers and order
	headers := make([]string, 0, len(columns))
	for col := range columns {
		headers = append(headers, fmt.Sprintf("%v", col))
	}
	sort.Strings(headers) // for a predictable order

	// Write headers
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Determine the number of rows (assume all columns are of same length)
	numRows := 0
	if len(headers) > 0 {
		numRows = len(columns[headers[0]])
	}

	// Write data rows
	for i := 0; i < numRows; i++ {
		row := make([]string, len(headers))
		for j, header := range headers {
			col := columns[header]
			if i < len(col) {
				switch v := col[i].(type) {
				case string:
					row[j] = v
				case fmt.Stringer:
					row[j] = v.String()
				case int:
					row[j] = strconv.Itoa(v)
				case int64:
					row[j] = strconv.FormatInt(v, 10)
				case float64:
					row[j] = strconv.FormatFloat(v, 'f', -1, 64)
				case float32:
					row[j] = strconv.FormatFloat(float64(v), 'f', -1, 32)
				case bool:
					if v {
						row[j] = "true"
					} else {
						row[j] = "false"
					}
				case nil:
					row[j] = ""
				default:
					row[j] = fmt.Sprintf("%v", v)
				}
			} else {
				row[j] = ""
			}
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

// SaveJSON series collection to a JSON file.
func (s *Series) SaveJSON(path string) error {
	return errors.New("JSON format not implemented yet")
}

// SaveJSONL series collection to a JSONL file.
func (s *Series) SaveJSONL(path string) error {
	return errors.New("JSONL format not implemented yet")
}
