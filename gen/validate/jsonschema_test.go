package validate

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/vtex/action-cm-generator/gen"
)

func TestJSONSchemaValidator(t *testing.T) {
	Convey("Given a jsonschema validator", t, func() {
		validator := NewJSONSchema()
		Convey("Given an arbitrary config", func() {
			So(validator, ShouldNotBeNil)
			emptyJSON := make(map[string]interface{})
			Convey("With a Schema that validates", func() {
				emptySchema := map[string]interface{}{"type": "object"}
				Convey("Then it should write on output", func() {
					input := make(chan gen.Config, 1)
					input <- gen.Config{
						Schema:  emptySchema,
						Content: emptyJSON,
						Path:    "",
					}
					output := validator.Validate(input)
					So(<-output, ShouldNotBeNil)
				})
			})
		})
	})
}
