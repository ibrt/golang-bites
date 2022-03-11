package enumz_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ibrt/golang-bites/enumz"
	"github.com/ibrt/golang-bites/filez"
)

func TestEnumz(t *testing.T) {
	specs := enumz.ProcessSimpleEnumSpecs(map[string][]string{
		"OutputModes": {
			"Batched",
			"Streaming",
			"BatchedStreaming",
		},
		"Frequencies": {
			"Single",
			"Multi",
		},
	})

	require.Equal(t, []*enumz.EnumSpec{
		{
			EnumNamePlural:     "OutputModes",
			EnumNameSingular:   "OutputMode",
			EnumNameInComments: "output modes",
			FileName:           "output_modes_enum.go",
			Values: []*enumz.EnumSpecValue{
				{
					Name:  "Batched",
					Value: "batched",
				},
				{
					Name:  "Streaming",
					Value: "streaming",
				},
				{
					Name:  "BatchedStreaming",
					Value: "batched_streaming",
				},
			},
		},
		{
			EnumNamePlural:     "Frequencies",
			EnumNameSingular:   "Frequency",
			EnumNameInComments: "frequencies",
			FileName:           "frequencies_enum.go",
			Values: []*enumz.EnumSpecValue{
				{
					Name:  "Single",
					Value: "single",
				},
				{
					Name:  "Multi",
					Value: "multi",
				},
			},
		},
	}, specs)

	require.NotPanics(t, func() {
		filez.WithMustCreateTempDir("golang-bites", func(dirPath string) {
			enumz.MustGenerateEnums(dirPath, false, "enums", specs)
			enumz.MustGenerateEnums(dirPath, true, "enums", specs)
		})
	})
}
