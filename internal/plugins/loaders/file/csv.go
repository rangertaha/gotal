package file

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type csvLoader struct {
	Name      string   `hcl:"name,optional"`       // name of the data series
	Summary   string   `hcl:"summary,optional"`    // summary of the data series
	StartTime string   `hcl:"start_time,optional"` // start time of the data series
	EndTime   string   `hcl:"end_time,optional"`   // end time of the data series
	Fields    []string `hcl:"fields,optional"`     // fields of the data series
	Tags      []string `hcl:"tags,optional"`       // tags of the data series
	Signals   []string `hcl:"signals,optional"`    // signals of the data series
}

// LoadCSVConfig loads and validates CSV configuration from HCL file
func LoadCSVConfig(filepath string) (*csvLoader, error) {
	var config csvLoader
	err := hclsimple.DecodeFile(filepath, nil, &config)
	if err != nil {
		return nil, fmt.Errorf("error decoding HCL file: %w", err)
	}

	// Validate the loaded configuration
	if err := ValidateCSVConfig(&config); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return &config, nil
}

// ValidateCSVConfig validates the CSV loader configuration
func ValidateCSVConfig(config *csvLoader) error {
	if len(config.Fields) == 0 {
		return fmt.Errorf("at least one field is required")
	}

	// Validate field names contain required fields
	requiredFields := []string{"time"}
	for _, required := range requiredFields {
		found := false
		for _, field := range config.Fields {
			if field == required {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("required field '%s' is missing", required)
		}
	}

	return nil
}
