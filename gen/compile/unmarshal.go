package compile

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/vtex/action-cm-generator/gen"
)

// JSONParser receives a compiled file and parse into a Config Struct.
type JSONParser struct{}

const schemaKey = "__schema"

// Parse receives a channel of compiled files and returns a channel of configuration parsed.
func (p *JSONParser) Parse(in <-chan gen.Compiled) (out <-chan gen.Config) {
	ch := make(chan gen.Config)
	logger := log.New(os.Stdout, "[parser]: ", log.Flags())

	go func() {
		defer close(ch)

		for compiled := range in {
			var config map[string]interface{}
			err := json.Unmarshal([]byte(compiled.Content), &config)

			if err != nil {
				logger.Fatal(err)
			}

			schema, ok := config[schemaKey]
			if !ok {
				logger.Fatal(errors.New("the configuration does not contain __schema property, did you try to use .schema?"))
			}

			delete(config, schemaKey) // remove schema from final result

			ch <- gen.Config{
				Schema:  schema,
				Content: config,
				Path:    compiled.Path,
			}
		}
	}()

	return ch
}

// NewJSONParser creates a new jsonnet parser instance.
func NewJSONParser() *JSONParser {
	return &JSONParser{}
}
