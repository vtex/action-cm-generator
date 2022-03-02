// dim.cue
package test
import "github.com/vtex/action-cm-generator/lib/other:other"
// Validate JSON configurations embedded strings.
configs: [string]: other.#Dimensions

configs: bed:      { width: 2, height: 0.1, depth: 2 }
configs: table:    { width: 34, height: 23, depth: 0.2 }
configs: painting: { width: 34, height: 12, depth: 0.2 }
