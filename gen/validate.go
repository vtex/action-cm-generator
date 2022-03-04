package gen

// Validator is responsible for validate configuration contents.
type Validator interface {
	Validate(in <-chan Config) (out <-chan Config)
}
