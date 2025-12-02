package schema

const (
	TypeString AttrType = iota
	TypeInt
	TypeFloat
	TypeBool
	TypeList
	TypeMap
	TypeTime
	TypeDuration
	TypeSeries
	TypeStream
)

type AttrType uint8

type Plugin struct {
	Name        string
	Description string

	Inputs     map[string]Dataset
	Outputs    map[string]Dataset
	Requires   map[string]Plugin
	Parameters map[string]Parameter
}

type Dataset struct {
	Name   string
	Source string
	Fields []string
	Tags   []string
}

type Parameter struct {
	Name        string
	Type        AttrType
	Required    bool
	Description string
	Default     any
	Nested      map[string]Parameter
}
