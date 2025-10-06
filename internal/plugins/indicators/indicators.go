package indicators

// import (
// 	"github.com/rangertaha/gotal/internal"
// 	"github.com/rangertaha/gotal/internal/pkg/series"
// 	"github.com/rangertaha/gotal/internal/pkg/stream"
// )

// type IndicatorFunc func(ind internal.Indicator, input *series.Series, opts ...internal.OptFunc) *series.Series

// func (i IndicatorFunc) Series(input *series.Series, opts ...internal.OptFunc) *series.Series {
// 	return i(i, input, opts...)
// }

// func (i IndicatorFunc) Stream(input *stream.Stream, opts ...internal.OptFunc) *stream.Stream {
// 	return input
// }

// var GROUP = "indicators"

// var DEFAULT_PERIOD int = 10

// type IndicatorsFunc func(*internal.Ticker, ...internal.OptFunc) *internal.Ticker

// func (i IndicatorsFunc) Stream(stream *internal.Streamer, opts ...internal.OptFunc) *internal.Streamer {
// 	return StreamFunc(stream, opts...)
// }

// func SeriesFunc(input *internal.Ticker, opts ...internal.OptFunc) (output *internal.Ticker) {
// 	cfg := opt.New()
// 	for _, opt := range opts {
// 		opt(cfg)
// 	}
// 	output = input.Copy()
// 	output.Set(cfg.Name(), cfg.Period(), cfg.Suffix())
// 	return output
// }

// // Stream calculates the moving average for a stream of metrics
// func StreamFunc(input *internal.Streamer, opts ...internal.OptFunc) (output *internal.Streamer) {
// 	cfg := opt.New()
// 	for _, opt := range opts {
// 		opt(cfg)
// 	}

// 	return input
// }
