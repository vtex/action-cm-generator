package main

import (
	"log"
	"os"

	"github.com/google/go-jsonnet"
	"github.com/sethvargo/go-githubactions"
	"github.com/vtex/action-cm-generator/gen"
	"github.com/vtex/action-cm-generator/gen/config"
	"github.com/vtex/action-cm-generator/gen/disk"
	"github.com/vtex/action-cm-generator/gen/jn"
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

	err := os.RemoveAll(outputDir)
	if err != nil {
		log.Fatal(err)
	}

	runner := gen.Runner{
		Reader:    disk.NewReader(inputDir),
		Compiler:  jn.NewCompiler(jsonnet.MakeVM()),
		Parser:    jn.NewParser(),
		Validator: config.NewValidator(),
		Exporter:  disk.NewExporter(inputDir, outputDir),
	}

	if err = runner.Run(); err != nil {
		log.Fatal(err)
	}
}
