package gen

// File is the generated file of a given output.
type File struct {
	// Path is the jsonnet key/path.
	Path string
	// Content is the jsonnet content.
	Content string
}

// Retriever is responsible for read files.
type Retriever interface {
	Retrieve() <-chan File
}
