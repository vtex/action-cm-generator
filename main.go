package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/google/go-jsonnet"
	"github.com/sethvargo/go-githubactions"
	"github.com/vtex/action-cm-generator/gen"
)

const dirIn = "in"
const dirOut = "out"
const jsonExtension = ".json"

// jsonExt sets the file extension to json.
func jsonExt(absPath string) string {
	ext := path.Ext(absPath)
	return absPath[0:len(absPath)-len(ext)] + jsonExtension
}

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
	input := make(chan gen.File)
	go func() {
		err := filepath.Walk(inputDir,
			func(path string, f os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if f.IsDir() {
					return nil
				}
				input <- gen.File{
					Path: path,
				}
				return nil
			})
		close(input)
		if err != nil {
			log.Println(err)
		}
	}()
	vm := jsonnet.MakeVM()
	compiler := gen.NewCompiler(vm)
	err := os.RemoveAll(outputDir)
	if err != nil {
		log.Fatal(err)
	}
	for r := range compiler.Cmp(input) {
		newpath := strings.Replace(r.Path, inputDir, outputDir, 1)
		err := os.MkdirAll(filepath.Dir(newpath), os.ModePerm)
		if err != nil {
			fmt.Printf("err when trying to ensure dir:%v\n", err)
			continue
		}
		outpath := jsonExt(newpath)
		err = ioutil.WriteFile(outpath, []byte(r.Content), 0644)
		if err != nil {
			fmt.Printf("err when writing file %v\n", err)
		}
		fmt.Printf("[cm-generator]: Generating %s\n", outpath)
	}
}
