package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/google/go-jsonnet"
	"github.com/sethvargo/go-githubactions"
	"github.com/vtex/action-cm-generator/gen"
)

const dirIn = "in"
const dirOut = "out"

// inputOrDefault gets the input value or use default if empty.
func inputOrDefault(name, defaultValue string) string {
	input := githubactions.GetInput(name)
	if len(input) == 0 {
		return defaultValue
	}

	return input
}

func main() {
	inputDir := inputOrDefault(dirIn, dirIn)
	outputDir := inputOrDefault(dirOut, dirOut)
	files := make(chan gen.File)

	go func() {
		err := filepath.Walk(inputDir,
			func(path string, f os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if f.IsDir() {
					return nil
				}
				files <- gen.File{
					Path: path,
				}
				return nil
			})

		close(files)

		if err != nil {
			log.Println(err)
		}
	}()

	vm := jsonnet.MakeVM()

	compiler := gen.NewCompiler(vm)
	parser := gen.NewParser()
	validator := gen.NewValidator()
	exporter := gen.NewExporter(inputDir, outputDir)

	err := os.RemoveAll(outputDir)
	if err != nil {
		log.Fatal(err)
	}

	err = exporter.Export(
		validator.Validate(
			parser.Parse(
				compiler.Compile(files),
			),
		),
	)

	if err != nil {
		log.Fatal(err)
	}
}
