package compile

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/vtex/action-cm-generator/gen"
)

func TestJsonParser(t *testing.T) {
	Convey("Given a new jsonparser", t, func() {
		parser := NewJSONParser()
		Convey("Given a valid config", func() {
			in := make(chan gen.Compiled, 1)
			in <- gen.Compiled{
				Path:    "",
				Content: "{ \"__schema\":{ }}",
			}
			Convey("Then it should parse", func() {
				out := parser.Parse(in)
				So(<-out, ShouldNotBeNil)
			})
		})
	})
}
