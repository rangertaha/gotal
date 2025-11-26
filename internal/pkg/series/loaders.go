package series

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rangertaha/gotal/internal/pkg/tick"
)

type LoaderOptions struct {
	Name      string
	Summary   string
	StartTime string
	EndTime   string
	Signals   []string
	Fields    []string
	Tags      []string
}

var Loaders = map[string]LoaderOptions{
	".csv": {
		Fields: []string{"time", "price", "volume"},
		Tags:   []string{"symbol", "exchange", "currency", "asset"},
	},
}

type Loader func(path string) (*Series, error)

var LOADERS = map[string]Loader{
	".csv":   ReadCSV,
	".json":  ReadJSON,
	".jsonl": ReadJSONL,
}

// Load series collection from a file.
func Load(path string, options ...Option) (*Series, error) {
	path = strings.TrimSpace(path)
	errs := make([]error, 0)

	In := func(name string, extensions map[string]Loader) (bool, Loader) {
		for extension := range extensions {
			if strings.HasSuffix(name, extension) {
				return true, extensions[extension]
			}
		}
		return false, nil
	}

	// Find and use a loader for the file extension
	if found, loader := In(path, LOADERS); found {
		// Load the series collection from the file
		ts, err := loader(path)
		if err != nil {
			return nil, err
		}

		// Apply options to the series collection
		for _, opt := range options {
			errs = append(errs, opt(ts))
		}

		// Return the series collection with the applied options or any errors
		return ts, errors.Join(errs...)
	}
	// If no loader is found, return an error
	// Return an error if the file format is not supported
	return nil, fmt.Errorf("unsupported file format: %s", path)
}

// ReadCSV reads a CSV file and returns a series collection.
func ReadCSV(path string) (*Series, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, errors.New("empty CSV file")
	}

	header := records[0]
	fieldIdx := make(map[string]int)
	for i, field := range header {
		fieldIdx[field] = i
	}

	ticks := make([]*tick.Tick, 0, len(records)-1)

	for _, row := range records[1:] {
		t := tick.New()

		for k, idx := range fieldIdx {
			if idx >= len(row) {
				continue
			}
			valueStr := row[idx]
			if k == "time" || k == "timestamp" {
				// Parse as integer seconds since epoch
				ts, err := strconv.ParseInt(valueStr, 10, 64)
				if err != nil {
					// try RFC3339
					tm, err2 := time.Parse(time.RFC3339, valueStr)
					if err2 == nil {
						t.SetTime(tm)
					}
				} else {
					t.SetTime(time.Unix(ts, 0))
				}
			} else {
				v, err := strconv.ParseFloat(valueStr, 64)
				if err == nil {
					t.SetField(k, v)
				}
			}
		}
		ticks = append(ticks, t)
	}

	s := New("series")
	s.ticks = ticks
	return s, nil
}

// ReadJSON reads a JSON file and returns a series collection.
func ReadJSON(path string) (*Series, error) {
	return nil, errors.New("JSON format not implemented yet")
}

// ReadJSONL reads a JSONL file and returns a series collection.
func ReadJSONL(path string) (*Series, error) {
	return nil, errors.New("JSONL format not implemented yet")
}
