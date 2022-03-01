package templatez

import (
	"bytes"
	"encoding/json"
	"go/format"
	htmltemplate "html/template"
	texttemplate "text/template"

	"github.com/yosssi/gohtml"

	"github.com/ibrt/golang-bites/internal"
)

// ParseAndExecuteText parses and executes a text template.
func ParseAndExecuteText(template string, data interface{}) ([]byte, error) {
	parsedTemplate, err := texttemplate.New("").Parse(template)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if err := parsedTemplate.Execute(buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// MustParseAndExecuteText is like ParseAndExecuteText but panics on error.
func MustParseAndExecuteText(template string, data interface{}) []byte {
	buf, err := ParseAndExecuteText(template, data)
	internal.MaybePanic(err)
	return buf
}

// ExecuteText executes a text template.
func ExecuteText(template *texttemplate.Template, data interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := template.Execute(buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MustExecuteText is like ExecuteText but panics on error.
func MustExecuteText(template *texttemplate.Template, data interface{}) []byte {
	buf, err := ExecuteText(template, data)
	internal.MaybePanic(err)
	return buf
}

// ParseAndExecuteGo parses and executes a text template, formatting the result as Go code.
func ParseAndExecuteGo(template string, data interface{}) ([]byte, error) {
	buf, err := ParseAndExecuteText(template, data)
	if err != nil {
		return nil, err
	}
	return format.Source(buf)
}

// MustParseAndExecuteGo is like ParseAndExecuteGo but panics on error.
func MustParseAndExecuteGo(template string, data interface{}) []byte {
	buf, err := ParseAndExecuteGo(template, data)
	internal.MaybePanic(err)
	return buf
}

// ExecuteGo executes a text template, formatting the result as Go code.
func ExecuteGo(template *texttemplate.Template, data interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := template.Execute(buf, data); err != nil {
		return nil, err
	}
	return format.Source(buf.Bytes())
}

// MustExecuteGo is like ExecuteGo but panics on error.
func MustExecuteGo(template *texttemplate.Template, data interface{}) []byte {
	buf, err := ExecuteGo(template, data)
	internal.MaybePanic(err)
	return buf
}

// ParseAndExecuteJSON parses and executes a text template, formatting the result as JSON code.
func ParseAndExecuteJSON(template, prefix, indent string, data interface{}) ([]byte, error) {
	buf, err := ParseAndExecuteText(template, data)
	if err != nil {
		return nil, err
	}
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, buf, prefix, indent); err != nil {
		return nil, err
	}
	return dst.Bytes(), nil
}

// MustParseAndExecuteJSON is like ParseAndExecuteJSON but panics on error.
func MustParseAndExecuteJSON(template, prefix, indent string, data interface{}) []byte {
	buf, err := ParseAndExecuteJSON(template, prefix, indent, data)
	internal.MaybePanic(err)
	return buf
}

// ExecuteJSON executes a text template, formatting the result as JSON code.
func ExecuteJSON(template *texttemplate.Template, prefix, indent string, data interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := template.Execute(buf, data); err != nil {
		return nil, err
	}
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, buf.Bytes(), prefix, indent); err != nil {
		return nil, err
	}
	return dst.Bytes(), nil
}

// MustExecuteJSON executes a text template, formatting the result as JSON code, panics on error.
func MustExecuteJSON(template *texttemplate.Template, prefix, indent string, data interface{}) []byte {
	buf, err := ExecuteJSON(template, prefix, indent, data)
	internal.MaybePanic(err)
	return buf
}

// ParseAndExecuteHTML parses and executes a HTML template.
func ParseAndExecuteHTML(template string, data interface{}) ([]byte, error) {
	parsedTemplate, err := htmltemplate.New("").Parse(template)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if err := parsedTemplate.Execute(buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// MustParseAndExecuteHTML is like ParseAndExecuteHTML but panics on error.
func MustParseAndExecuteHTML(template string, data interface{}) []byte {
	buf, err := ParseAndExecuteHTML(template, data)
	internal.MaybePanic(err)
	return buf
}

// ExecuteHTML executes a HTML template.
func ExecuteHTML(template *htmltemplate.Template, data interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := template.Execute(buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MustExecuteHTML is like ExecuteHTML but panics on error.
func MustExecuteHTML(template *htmltemplate.Template, data interface{}) []byte {
	buf, err := ExecuteHTML(template, data)
	internal.MaybePanic(err)
	return buf
}

// ParseAndExecuteHTMLTidy parses and executes a HTML template, formatting the result as HTML code.
func ParseAndExecuteHTMLTidy(template string, data interface{}) ([]byte, error) {
	parsedTemplate, err := htmltemplate.New("").Parse(template)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if err := parsedTemplate.Execute(gohtml.NewWriter(buf), data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// MustParseAndExecuteHTMLTidy is like ParseAndExecuteHTMLTidy but panics on error.
func MustParseAndExecuteHTMLTidy(template string, data interface{}) []byte {
	buf, err := ParseAndExecuteHTMLTidy(template, data)
	internal.MaybePanic(err)
	return buf
}

// ExecuteHTMLTidy executes a HTML template, formatting the result as HTML code.
func ExecuteHTMLTidy(template *htmltemplate.Template, data interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := template.Execute(gohtml.NewWriter(buf), data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MustExecuteHTMLTidy is like ExecuteHTMLTidy but panics on error.
func MustExecuteHTMLTidy(template *htmltemplate.Template, data interface{}) []byte {
	buf, err := ExecuteHTMLTidy(template, data)
	internal.MaybePanic(err)
	return buf
}
