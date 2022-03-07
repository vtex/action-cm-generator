package disk

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadFromDisk(t *testing.T) {
	Convey("Given a disk retriever", t, func() {
		dir := "/tmp/test-jsonnet"
		filePath := "/tmp/test-jsonnet/valid.jsonnet"
		notJsonnet := "/tmp/test-jsonnet/valid.jsonnet"

		retriever := NewRetriever(dir)

		So(os.MkdirAll(filepath.Dir(filePath), os.ModePerm), ShouldBeNil)
		Convey("And a valid jsonnet", func() {
			So(ioutil.WriteFile(filePath, []byte(`{}`), 0644), ShouldBeNil)
			So(ioutil.WriteFile(notJsonnet, []byte(`{}`), 0644), ShouldBeNil)
			Convey("Retriever should read only jsonnet file", func() {
				out := retriever.Read()
				So(<-out, ShouldNotBeNil)
				So(<-out, ShouldBeZeroValue)
			})
		})
	})
}
