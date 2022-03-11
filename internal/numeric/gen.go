//go:generate go run .

package main

import (
	_ "embed" // embed
	"path/filepath"

	"github.com/ibrt/golang-bites/filez"
	"github.com/ibrt/golang-bites/internal"
	"github.com/ibrt/golang-bites/templatez"
)

// NumericSpec describes the specification for a numeric type.
type NumericSpec struct {
	Type string
	Size int
}

// Pkg returns the package name.
func (p *NumericSpec) Pkg() string {
	return p.Type + "z"
}

// ParseCall generates the strconv.Parse* call.
func (p *NumericSpec) ParseCall() string {
	switch p.Type[0] {
	case 'i':
		return "strconv.ParseInt(v, 10, BitSize)"
	case 'u':
		return "strconv.ParseUint(v, 10, BitSize)"
	case 'f':
		return "strconv.ParseFloat(v, BitSize)"
	default:
		panic("unknown type")
	}
}

func main() {
	outDirPath := filepath.Join("..", "..", "numeric")
	filez.MustPrepareDir(outDirPath, 0777)

	for _, spec := range []*NumericSpec{
		{"int8", 8},
		{"int16", 16},
		{"int32", 32},
		{"int64", 64},
		{"int", -1},
		{"uint8", 8},
		{"uint16", 16},
		{"uint32", 32},
		{"uint64", 64},
		{"uint", -1},
		{"float32", 32},
		{"float64", 64},
	} {
		filez.MustWriteFile(
			filepath.Join(outDirPath, spec.Pkg(), spec.Pkg()+".go"),
			0777, 0666,
			templatez.MustParseAndExecuteGo(internal.NumericGoTemplateAsset, spec))

		filez.MustWriteFile(
			filepath.Join(outDirPath, spec.Pkg(), spec.Pkg()+"_test.go"),
			0777, 0666,
			templatez.MustParseAndExecuteGo(internal.NumericTestGoTemplateAsset, spec))
	}
}
