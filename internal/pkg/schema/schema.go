package schema


type Schema struct {
	Fields  []string
	Tags    []string
	Signals []string
}

func NewSchema(fields []string, tags []string, signals []string) *Schema {
	return &Schema{Fields: fields, Tags: tags, Signals: signals}
}
