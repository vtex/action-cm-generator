package validate

import (
	"fmt"
	"log"
	"os"

	"github.com/vtex/action-cm-generator/gen"
	"github.com/vtex/action-cm-generator/gen/util"
	"github.com/xeipuuv/gojsonschema"
)

// Validator is responsible for validate configuration contents.
type Validator struct{}

// Validate returns a channel of all validated files.
func (v *Validator) Validate(in <-chan gen.Config) (out <-chan gen.Config) {
	ch := make(chan gen.Config)
	logger := log.New(os.Stdout, "[validator]: ", log.Flags())

	go func() {
		defer close(ch)

		for config := range in {
			ls := gojsonschema.NewGoLoader(config.Schema)
			cl := gojsonschema.NewGoLoader(config.Content)
			result, err := gojsonschema.Validate(ls, cl)

			if err != nil {
				schema, err := util.MarshalIndent(config.Schema)

				if err != nil {
					logger.Println(err)
				}

				content, err := util.MarshalIndent(config.Content)

				if err != nil {
					logger.Println(err)
				}

				logger.Fatal(fmt.Errorf("\n>Schema\n %s\n>Content\n %s\n error when trying to validate the config %s %v", schema, content, config.Path, err))
			}

			resultErrs := result.Errors()

			if len(resultErrs) > 0 || !result.Valid() {
				logger.Fatalf("error when validating %s: %s", config.Path, resultErrs)
			}

			ch <- config
		}
	}()

	return ch
}

// NewJSONSchema creates a new validator for jsonschema instance.
func NewJSONSchema() *Validator {
	return &Validator{}
}
