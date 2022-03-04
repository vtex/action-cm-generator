package gen

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

// Validator is responsible for validate configuration contents.
type Validator struct{}

// Validate returns a channel of all validated files.
func (v *Validator) Validate(in <-chan Config) (out <-chan Config) {
	ch := make(chan Config)
	logger := log.New(os.Stdout, log.Prefix(), log.Flags())

	go func() {
		for config := range in {
			logger.SetPrefix("[validator]: ")

			ls := gojsonschema.NewGoLoader(config.Schema)
			cl := gojsonschema.NewGoLoader(config.Content)
			result, err := gojsonschema.Validate(ls, cl)

			if err != nil {
				schema, err := json.MarshalIndent(config.Schema, noPrefix, indentationLevel)

				if err != nil {
					logger.Println(err)
				}

				content, err := json.MarshalIndent(config.Content, noPrefix, indentationLevel)

				if err != nil {
					logger.Println(err)
				}

				logger.Fatal(fmt.Errorf("\n>>>>>> Schema\n %s\n>>>>>> Content\n %s\n error when trying to validate the config %s %v", schema, content, config.Path, err))
			}

			resultErrs := result.Errors()

			if len(resultErrs) > 0 || !result.Valid() {
				logger.Fatalln(resultErrs)
			}

			ch <- Config{
				Content: config.Content,
				Path:    config.Path,
			}
		}

		close(ch)
	}()

	return ch
}

// NewValidator creates a new validator instance.
func NewValidator() *Validator {
	return &Validator{}
}
