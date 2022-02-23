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

func main() {
	input := make(chan gen.File)
	go func() {
		err := filepath.Walk(dirIn,
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
	for r := range compiler.Cmp(input) {
		newpath := strings.Replace(r.Path, dirIn, dirOut, 1)
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
		fmt.Printf(">> Generating %s\n", outpath)
	}
}
