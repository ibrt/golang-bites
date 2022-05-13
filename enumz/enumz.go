package enumz

import (
	_ "embed" // embed
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/flect"
	"github.com/iancoleman/strcase"

	"github.com/ibrt/golang-bites/filez"
	"github.com/ibrt/golang-bites/internal"
	"github.com/ibrt/golang-bites/templatez"
)

// EnumSpec describes the specification for a generated enum.
type EnumSpec struct {
	EnumNamePlural     string
	EnumNameSingular   string
	EnumNameInComments string
	FileName           string
	Values             []*EnumSpecValue
}

// EnumSpecValue describes the specification for a generated enum value.
type EnumSpecValue struct {
	Name  string
	Value string
	Label string
}

type fullEnumSpec struct {
	PackageName string
	*EnumSpec
}

// ProcessSimpleEnumSpecs implements a standard, simplified specification for enum specs.
//
// - Each key in the map should be an upper camel, plural enum name.
// - Each entry in the values slice should be an upper camel, singular value name.
func ProcessSimpleEnumSpecs(simpleSpecs map[string][]string) []*EnumSpec {
	specs := make([]*EnumSpec, 0, len(simpleSpecs))

	for name, values := range simpleSpecs {
		specs = append(specs, &EnumSpec{
			EnumNamePlural:     name,
			EnumNameSingular:   flect.Singularize(name),
			EnumNameInComments: strcase.ToDelimited(name, ' '),
			FileName:           strcase.ToSnake(name) + "_enum.go",
			Values:             make([]*EnumSpecValue, len(values)),
		})

		for i, value := range values {
			specs[len(specs)-1].Values[i] = &EnumSpecValue{
				Name:  value,
				Value: strcase.ToSnake(value),
				Label: value,
			}
		}

	}

	return specs
}

// MustGenerateEnums generates enums from the given specs.
func MustGenerateEnums(outDirPath string, wipe bool, packageName string, specs []*EnumSpec) {
	if wipe {
		filez.MustPrepareDir(outDirPath, 0777)
	} else {
		internal.MaybePanic(os.MkdirAll(outDirPath, 0777))
	}

	for _, spec := range specs {
		filez.MustWriteFile(
			filepath.Join(outDirPath, strings.TrimSuffix(spec.FileName, ".go")+".go"),
			0777, 0666,
			templatez.MustParseAndExecuteGo(internal.EnumGoTemplateAsset, &fullEnumSpec{
				PackageName: packageName,
				EnumSpec:    spec,
			}))

		filez.MustWriteFile(
			filepath.Join(outDirPath, strings.TrimSuffix(spec.FileName, ".go")+"_test.go"),
			0777, 0666,
			templatez.MustParseAndExecuteGo(internal.EnumTestGoTemplateAsset, &fullEnumSpec{
				PackageName: packageName,
				EnumSpec:    spec,
			}))
	}
}
