package gen

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Exporter is used to write config final result somewhere.
type Exporter struct {
	InputDir  string
	OutputDir string
}

const jsonExtension = ".json"

// jsonExt sets the file extension to json.
func jsonExt(absPath string) string {
	ext := path.Ext(absPath)
	return absPath[0:len(absPath)-len(ext)] + jsonExtension
}

const indentationLevel = "   " // 4 spaces of identation
const noPrefix = ""            // no prefix

// Export receives a channel of configuration and write them on disk.
func (e *Exporter) Export(in <-chan Config) error {
	logger := log.New(os.Stdout, log.Prefix(), log.Flags())

	for config := range in {
		logger.SetPrefix("[exporter]: ")

		newpath := strings.Replace(config.Path, e.InputDir, e.OutputDir, 1)
		err := os.MkdirAll(filepath.Dir(newpath), os.ModePerm)

		if err != nil {
			logger.Printf("err when trying to ensure dir:%v\n", err)
			continue
		}

		bts, err := json.MarshalIndent(config.Content, noPrefix, indentationLevel)

		if err != nil {
			logger.Printf("err when marshalling file %v\n", err)
			continue
		}

		outpath := jsonExt(newpath)

		err = ioutil.WriteFile(outpath, bts, 0644) //nolint:gosec,gomnd

		if err != nil {
			logger.Printf("err when writing file %v\n", err)
			continue
		}

		logger.Printf("Exporting %s\n", outpath)
	}

	return nil
}

// NewExporter creates a new exporter instance with the given input/output dir.
func NewExporter(inputDir, outputDir string) *Exporter {
	return &Exporter{
		InputDir:  inputDir,
		OutputDir: outputDir,
	}
}
