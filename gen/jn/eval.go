package jn

import (
	"log"
	"os"

	"github.com/google/go-jsonnet"
	"github.com/vtex/action-cm-generator/gen"
)

// Compiler compile jsonnet files into json definitions.
type Compiler struct {
	VM *jsonnet.VM
}

// Compile compiles all files and return the generated config files.
func (c *Compiler) Compile(files <-chan gen.File) <-chan gen.Compiled {
	output := make(chan gen.Compiled)
	logger := log.New(os.Stdout, "[compiler]: ", log.Flags())

	go func() {
		for file := range files {
			path := file.Path

			out, err := c.VM.EvaluateFile(path)
			if err != nil {
				logger.Println(err)
				continue
			}
			output <- gen.Compiled{
				Path:    path,
				Content: out,
			}
		}

		close(output)
	}()

	return output
}

// NewCompiler creates a new jsonnet compiler instance.
func NewCompiler(vm *jsonnet.VM) *Compiler {
	return &Compiler{
		VM: vm,
	}
}
