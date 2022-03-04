package gen

// Compiled represents a compiled jsonnet file.
type Compiled struct {
	// Path is the jsonnet key/path.
	Path string
	// Content is a JSON compiled representation generated by jsonnet compiler.
	Content string
}

// Compiler is used to compile files.
type Compiler interface {
	// Compile receives a channel of files and returns a channel of compiled files.
	Compile(files <-chan File) <-chan Compiled
}