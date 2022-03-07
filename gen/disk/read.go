package disk

import (
	"log"
	"os"
	"path/filepath"

	"github.com/vtex/action-cm-generator/gen"
)

const jsonnetExt = ".jsonnet"

// isJsonnet returns if the file specified on the filePath is a jsonnet file.
func isJsonnet(filePath string) bool {
	return filepath.Ext(filePath) == jsonnetExt
}

// Reader is responsible for read files.
type Reader struct {
	Dir string
}

// Read returns a channel of file.
func (r *Reader) Read() <-chan gen.File {
	ch := make(chan gen.File)

	go func() {
		defer close(ch)

		err := filepath.Walk(r.Dir,
			func(path string, f os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if f.IsDir() {
					return nil
				}

				if !isJsonnet(path) {
					return nil
				}

				ch <- gen.File{
					Path: path,
				}
				return nil
			})

		if err != nil {
			log.Fatal(err)
		}
	}()

	return ch
}

// NewReader creates a new disk reader.
func NewReader(dir string) *Reader {
	return &Reader{
		Dir: dir,
	}
}
