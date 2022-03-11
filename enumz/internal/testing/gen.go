//go:generate go run .

package main

import (
	"github.com/ibrt/golang-bites/enumz"
)

func main() {
	enumz.MustGenerateEnums(
		"enums", true, "enums",
		enumz.ProcessSimpleEnumSpecs(map[string][]string{
			"OutputModes": {
				"Batched",
				"Streaming",
				"BatchedStreaming",
			},
			"Frequencies": {
				"Single",
				"Multi",
			},
		}))
}
