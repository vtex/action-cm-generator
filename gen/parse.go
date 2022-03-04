package gen

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

// Config is the generic representation of a config
type Config struct {
	// Schema is the configuration JSON Schema.
	Schema interface{}
	// Content is the configuration content.
	Content map[string]interface{}
	// Path is the jsonnet key/path.
	Path string
}

// Parser receives a compiled file and parse into a Config Struct.
type Parser struct{}

const schemaKey = "__schema"

// Parse receives a channel of compiled files and returns a channel of configuration parsed.
func (p *Parser) Parse(in <-chan Compiled) (out <-chan Config) {
	ch := make(chan Config)
	logger := log.New(os.Stdout, log.Prefix(), log.Flags())

	go func() {
		for compiled := range in {
			logger.SetPrefix("[parser]: ")

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

			ch <- Config{
				Schema:  schema,
				Content: config,
				Path:    compiled.Path,
			}
		}

		close(ch)
	}()

	return ch
}

// NewParser creates a new parser instance.
func NewParser() *Parser {
	return &Parser{}
}
