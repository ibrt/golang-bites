//go:generate go run .

package main

import (
	"bytes"
	_ "embed" // embed
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/ibrt/golang-bites/internal"
)

var (
	//go:embed numeric.go.gotmpl
	numericGoTemplate string

	//go:embed numeric_test.go.gotmpl
	numericTestGoTemplate string

	tpl     = template.Must(template.New("").Parse(numericGoTemplate))
	testTpl = template.Must(template.New("").Parse(numericTestGoTemplate))
)

// TemplateParams describes the template parameters.
type TemplateParams struct {
	Type string
	Size int
}

// Pkg returns the package name.
func (p *TemplateParams) Pkg() string {
	return p.Type + "z"
}

// ParseCall generates the strconv.Parse* call.
func (p *TemplateParams) ParseCall() string {
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
	outDirPath := filepath.Join("..", "..", "..", "numeric")

	internal.MaybePanic(os.RemoveAll(outDirPath))
	internal.MaybePanic(os.MkdirAll(filepath.Join(outDirPath), 0777))

	for _, params := range []*TemplateParams{
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
		buf := &bytes.Buffer{}
		internal.MaybePanic(tpl.Execute(buf, params))
		internal.MaybePanic(os.MkdirAll(filepath.Join(outDirPath, params.Pkg()), 0777))
		internal.MaybePanic(ioutil.WriteFile(filepath.Join(outDirPath, params.Pkg(), params.Pkg()+".go"), buf.Bytes(), 0666))

		buf.Reset()
		internal.MaybePanic(testTpl.Execute(buf, params))
		internal.MaybePanic(ioutil.WriteFile(filepath.Join(outDirPath, params.Pkg(), params.Pkg()+"_test.go"), buf.Bytes(), 0666))
	}
}
