package util

import "encoding/json"

const indentationLevel = "   " // 4 spaces of identation
const noPrefix = ""            // no prefix

// MarshalIndent marshal json with indentation of 4 spaces.
func MarshalIndent(data interface{}) ([]byte, error) {
	return json.MarshalIndent(data, noPrefix, indentationLevel)
}
