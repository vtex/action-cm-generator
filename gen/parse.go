package gen

// Config is the generic representation of a config
type Config struct {
	// Schema is the configuration JSON Schema.
	Schema interface{}
	// Content is the configuration content.
	Content map[string]interface{}
	// Path is the jsonnet key/path.
	Path string
}

// Parser receives a compiled file and parse into a Config Struct.
type Parser interface {
	Parse(in <-chan Compiled) (out <-chan Config)
}
