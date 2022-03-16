package util

import "encoding/json"

const indentationLevel = "   " // 3 spaces of indentation
const noPrefix = ""            // no prefix

// MarshalIndent marshal json with indentation of 3 spaces.
func MarshalIndent(data interface{}) ([]byte, error) {
	return json.MarshalIndent(data, noPrefix, indentationLevel)
}
