package mock

// import (
// 	"github.com/rangertaha/gotal/internal"
// )

// type plugin struct {
// 	source string
// 	fields []string
// 	tags   []string
// }

// func NewPlugin(source string, fields []string, tags []string) internal.Plugin {
// 	return &plugin{
// 		source: source,
// 		fields: fields,
// 		tags:   tags,
// 	}
// }

// func (p *plugin) Name() string {
// 	return p.source
// }

// func (p *plugin) Description() string {
// 	return "Mock provider"
// }

// func (p *plugin) Inputs() []internal.Dataset {
// 	return []internal.Dataset{
// 		{
// 			Name: p.source,
// 			Fields: p.fields,
// 			Tags:   p.tags,
// 		},
// 	}
// }

// func (p *plugin) Outputs() []internal.Dataset {
// 	return []internal.Dataset{
// 		{
// 			Name: p.source,
// 			Fields: p.fields,
// 			Tags:   p.tags,
// 		},
// 	}
// }

// func (p *plugin) Parameters() []internal.Parameter {
// 	return []internal.Parameter{
// 		{
// 			Name: "source",
// 		},
// 	}
// }

// func (p *plugin) Init() error {
// 	return nil
// }

// func (p *plugin) Load() error {
// 	return nil
// }

// func (p *plugin) Start() error {
// 	return nil
// }
