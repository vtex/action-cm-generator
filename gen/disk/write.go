package disk

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/vtex/action-cm-generator/gen"
	"github.com/vtex/action-cm-generator/gen/util"
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
	return strings.Replace(absPath, ext, jsonExtension, 1)
}

// Export receives a channel of configuration and write them on disk.
func (e *Exporter) Export(in <-chan gen.Config) error {
	logger := log.New(os.Stdout, "[exporter]: ", log.Flags())

	for config := range in {
		newpath := strings.Replace(config.Path, e.InputDir, e.OutputDir, 1)
		err := os.MkdirAll(filepath.Dir(newpath), os.ModePerm)

		if err != nil {
			logger.Printf("err when trying to ensure dir:%v\n", err)
			continue
		}

		bts, err := util.MarshalIndent(config.Content)

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

// NewExporter creates a new disk exporter instance with the given input/output dir.
func NewExporter(inputDir, outputDir string) *Exporter {
	return &Exporter{
		InputDir:  inputDir,
		OutputDir: outputDir,
	}
}
