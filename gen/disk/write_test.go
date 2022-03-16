package disk

import (
	"errors"
	"fmt"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/vtex/action-cm-generator/gen"
)

func TestWriteOnDisk(t *testing.T) {
	Convey("Given a file exporter", t, func() {
		inDir := "/tmp/in-jsonnet"
		outDir := "/tmp/out-jsonnet"
		file := "valid.jsonnet"

		expected := fmt.Sprintf("%s/%s", outDir, "valid.json")

		exporter := NewExporter(inDir, outDir)
		Convey("And a stream of file input", func() {
			ch := make(chan gen.Config, 1)
			ch <- gen.Config{
				Schema:  nil,
				Content: map[string]interface{}{},
				Path:    fmt.Sprintf("%s/%s", inDir, file),
			}
			close(ch)
			So(exporter.Export(ch), ShouldBeNil)

			if _, err := os.Stat(expected); errors.Is(err, os.ErrNotExist) {
				panic("it should exists")
			}
		})
	})
}
