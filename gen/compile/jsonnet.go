package compile

import (
	"log"
	"os"

	"github.com/google/go-jsonnet"
	"github.com/vtex/action-cm-generator/gen"
)

// JNCompiler compile jsonnet files into json definitions.
type JNCompiler struct {
	VM *jsonnet.VM
}

// Compile compiles all files and return the generated config files.
func (c *JNCompiler) Compile(files <-chan gen.File) <-chan gen.Compiled {
	output := make(chan gen.Compiled)
	logger := log.New(os.Stdout, "[compiler]: ", log.Flags())

	go func() {
		defer close(output)

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
	}()

	return output
}

// NewJNCompiler creates a new jsonnet compiler instance.
func NewJNCompiler(vm *jsonnet.VM) *JNCompiler {
	return &JNCompiler{
		VM: vm,
	}
}
