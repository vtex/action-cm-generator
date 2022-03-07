package compile

import (
	"io/ioutil"
	"testing"

	"github.com/google/go-jsonnet"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/vtex/action-cm-generator/gen"
)

func TestJsonnetCompiler(t *testing.T) {
	Convey("Given a jsonnet compiler", t, func() {
		jnCompiler := NewJNCompiler(jsonnet.MakeVM())
		in := make(chan gen.File, 1)
		Convey("Given a valid file path", func() {
			valid := "/tmp/valid.jsonnet"
			invalid := "/tmp/invalid.jsonnet"
			Convey("With valid jsonnet content", func() {
				So(ioutil.WriteFile(valid, []byte(`{}`), 0644), ShouldBeNil)
				Convey("Then it should compile the file", func() {
					in <- gen.File{
						Path:    valid,
						Content: "{}",
					}
					out := jnCompiler.Compile(in)
					So(<-out, ShouldNotBeNil)
				})
			})
			Convey("With invalid jsonnet content", func() {
				So(ioutil.WriteFile(invalid, []byte(`abcde`), 0644), ShouldBeNil)
				Convey("The it should skip the file", func() {
					in <- gen.File{
						Path:    invalid,
						Content: "{}",
					}
					close(in)
					out := jnCompiler.Compile(in)
					So(<-out, ShouldBeZeroValue)
				})
			})
		})
		Convey("Given a invalid file path", func() {
			Convey("Then it should skip", func() {
				in <- gen.File{
					Path:    "/tmp/inexistent.jsonnet",
					Content: "",
				}
				close(in)
				out := jnCompiler.Compile(in)
				So(<-out, ShouldBeZeroValue)
			})
		})
	})
}
