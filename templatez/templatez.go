package templatez

import (
	"bytes"
	htmltemplate "html/template"
	texttemplate "text/template"
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
	if err != nil {
		panic(err)
	}
	return buf
}

// MustExecuteText executes a text template, panics on error.
func MustExecuteText(template *texttemplate.Template, data interface{}) []byte {
	buf := &bytes.Buffer{}
	if err := template.Execute(buf, data); err != nil {
		panic(err)
	}
	return buf.Bytes()
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
	if err != nil {
		panic(err)
	}
	return buf
}

// MustExecuteHTML executes a HTML template, panics on error.
func MustExecuteHTML(template *htmltemplate.Template, data interface{}) []byte {
	buf := &bytes.Buffer{}
	if err := template.Execute(buf, data); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
