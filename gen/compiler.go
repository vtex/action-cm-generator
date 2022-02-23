package gen

import (
	"github.com/google/go-jsonnet"
)

// File is the generated file of a given output.
type File struct {
	Path    string
	Content string
}

// Compiler compile jsonnet files into json definitions.
type Compiler struct {
	VM *jsonnet.VM
}

// Cmp compiles all files and return the generated config files.
func (c *Compiler) Cmp(files <-chan File) <-chan File {
	output := make(chan File)
	go func() {
		for file := range files {
			path := file.Path
			out, err := c.VM.EvaluateFile(file.Path)
			if err != nil {
				continue
			}
			output <- File{
				Path:    path,
				Content: out,
			}
		}
		close(output)
	}()
	return output
}

// NewCompiler creates a new compiler instance.
func NewCompiler(vm *jsonnet.VM) *Compiler {
	return &Compiler{
		VM: vm,
	}
}
