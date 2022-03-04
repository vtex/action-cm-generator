package gen

// File is the generated file of a given output.
type File struct {
	// Path is the jsonnet key/path.
	Path string
	// Content is the jsonnet content.
	Content string
}

// Reader is responsible for read files.
type Reader interface {
	Read() <-chan File
}
