package data

import "github.com/rangertaha/gotal/internal/pkg/series"

func Load(name string) (*series.Series, error) {
	return series.New(name), nil
}
