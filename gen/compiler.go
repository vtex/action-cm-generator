package gen

import (
	"fmt"
	"path/filepath"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
	"github.com/google/go-jsonnet"
)

// File is the generated file of a given output.
type File struct {
	Path    string
	Content string
}

// Compiler compile jsonnet files into json definitions.
type Compiler struct {
	VM       *jsonnet.VM
	Cue      *cue.Context
	cueFiles []string
}

// Cmp compiles all files and return the generated config files.
func (c *Compiler) Cmp(files <-chan File) <-chan File {
	output := make(chan File)
	go func() {
		for file := range files {
			var (
				err error
			)
			path := file.Path
			ext := filepath.Ext(path)
			if ext == ".jsonnet" {
				err = c.FromJSONNET(output, file)
			} else if ext == ".cue" {
				err = c.FromCUE(file)
			}

			if err != nil {
				fmt.Println(err)
				panic(err)
			}
		}
		err := c.CompileCUEs(output)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		close(output)
	}()
	return output
}

// FromJSONNET compiles a file with jsonnet extension.
func (c *Compiler) FromJSONNET(out chan File, file File) error {
	json, err := c.VM.EvaluateFile(file.Path)
	if err != nil {
		fmt.Println(err)
		panic(err)

	}
	out <- File{
		Path:    file.Path,
		Content: json,
	}
	return nil
}

func (c *Compiler) CompileCUEs(out chan File) error {
	fmt.Println("CUE FILES", c.cueFiles)
	bis := load.Instances(c.cueFiles, nil)
	for idx, bi := range bis {
		value := c.Cue.BuildInstance(bi)
		if value.Err() != nil {
			return value.Err()
		}
		err := value.Validate()
		if err != nil {
			return err
		}
		byts, err := value.MarshalJSON()
		if err != nil {
			return err
		}
		out <- File{
			Path:    c.cueFiles[idx],
			Content: string(byts),
		}
	}
	return nil
}

// FromCUE compiles a file with cue extension.
func (c *Compiler) FromCUE(file File) error {
	c.cueFiles = append(c.cueFiles, file.Path)
	return nil
}

// NewCompiler creates a new compiler instance.
func NewCompiler(vm *jsonnet.VM, cue *cue.Context) *Compiler {
	return &Compiler{
		VM:       vm,
		Cue:      cue,
		cueFiles: make([]string, 0),
	}
}
