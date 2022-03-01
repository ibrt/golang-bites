package templatez_test

import (
	htmltemplate "html/template"
	"testing"
	texttemplate "text/template"

	"github.com/ibrt/golang-bites/templatez"

	"github.com/stretchr/testify/require"
)

func TestParseAndExecuteText(t *testing.T) {
	buf, err := templatez.ParseAndExecuteText("{{ . }}", "<a>&</a>")
	require.NoError(t, err)
	require.Equal(t, "<a>&</a>", string(buf))

	buf, err = templatez.ParseAndExecuteText("{{ bad }}", "<a>&</a>")
	require.EqualError(t, err, `template: :1: function "bad" not defined`)
	require.Nil(t, buf)

	buf, err = templatez.ParseAndExecuteText(`{{ template "x" }}`, nil)
	require.EqualError(t, err, `template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)
	require.Nil(t, buf)
}

func TestMustParseAndExecuteText(t *testing.T) {
	require.NotPanics(t, func() {
		buf := templatez.MustParseAndExecuteText("{{ . }}", "<a>&</a>")
		require.Equal(t, "<a>&</a>", string(buf))
	})

	require.PanicsWithError(t, `template: :1: function "bad" not defined`, func() {
		templatez.MustParseAndExecuteText("{{ bad }}", "<a>&</a>")
	})

	require.PanicsWithError(t, `template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`, func() {
		templatez.MustParseAndExecuteText(`{{ template "x" }}`, nil)
	})
}

func TestExecuteText(t *testing.T) {
	okTpl, err := texttemplate.New("").Parse("{{ . }}")
	require.NoError(t, err)

	errTpl, err := texttemplate.New("").Parse(`{{ template "x" }}`)
	require.NoError(t, err)

	buf, err := templatez.ExecuteText(okTpl, "<a>&</a>")
	require.NoError(t, err)
	require.Equal(t, "<a>&</a>", string(buf))

	buf, err = templatez.ExecuteText(errTpl, nil)
	require.EqualError(t, err, `template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)
	require.Nil(t, buf)
}

func TestMustExecuteText(t *testing.T) {
	okTpl, err := texttemplate.New("").Parse("{{ . }}")
	require.NoError(t, err)

	errTpl, err := texttemplate.New("").Parse(`{{ template "x" }}`)
	require.NoError(t, err)

	require.NotPanics(t, func() {
		buf := templatez.MustExecuteText(okTpl, "<a>&</a>")
		require.Equal(t, "<a>&</a>", string(buf))
	})

	require.PanicsWithError(t, `template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`, func() {
		templatez.MustExecuteText(errTpl, nil)
	})
}

func TestParseAndExecuteGo(t *testing.T) {
	buf, err := templatez.ParseAndExecuteGo("package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"{{ . }}\") }", "Hello World")
	require.NoError(t, err)
	require.Equal(t, "package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"Hello World\") }\n", string(buf))

	buf, err = templatez.ParseAndExecuteGo("{{ bad }}", "Hello World")
	require.EqualError(t, err, `template: :1: function "bad" not defined`)
	require.Nil(t, buf)

	buf, err = templatez.ParseAndExecuteGo(`{{ template "x" }}`, nil)
	require.EqualError(t, err, `template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)
	require.Nil(t, buf)

	buf, err = templatez.ParseAndExecuteGo("package main\nfuncmain() { fmt.Println(\"{{ . }}\") }", "Hello World")
	require.EqualError(t, err, "2:1: expected declaration, found funcmain")
	require.Nil(t, buf)
}

func TestMustParseAndExecuteGo(t *testing.T) {
	require.NotPanics(t, func() {
		buf := templatez.MustParseAndExecuteGo("package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"{{ . }}\") }", "Hello World")
		require.Equal(t, "package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"Hello World\") }\n", string(buf))
	})

	require.PanicsWithError(t, `template: :1: function "bad" not defined`, func() {
		templatez.MustParseAndExecuteGo("{{ bad }}", "Hello World")
	})

	require.PanicsWithError(t, `template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`, func() {
		templatez.MustParseAndExecuteGo(`{{ template "x" }}`, nil)
	})

	require.PanicsWithError(t, "2:1: expected declaration, found funcmain", func() {
		templatez.MustParseAndExecuteGo("package main\nfuncmain() { fmt.Println(\"{{ . }}\") }", "Hello World")
	})
}

func TestExecuteGo(t *testing.T) {
	okTpl, err := texttemplate.New("").Parse("package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"{{ . }}\") }")
	require.NoError(t, err)

	errTpl, err := texttemplate.New("").Parse(`{{ template "x" }}`)
	require.NoError(t, err)

	goErrTpl, err := texttemplate.New("").Parse("package main\nfuncmain() { fmt.Println(\"{{ . }}\") }")
	require.NoError(t, err)

	buf, err := templatez.ExecuteGo(okTpl, "Hello World")
	require.NoError(t, err)
	require.Equal(t, "package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"Hello World\") }\n", string(buf))

	buf, err = templatez.ExecuteGo(errTpl, nil)
	require.EqualError(t, err, `template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)
	require.Nil(t, buf)

	buf, err = templatez.ExecuteGo(goErrTpl, "Hello World")
	require.EqualError(t, err, "2:1: expected declaration, found funcmain")
	require.Nil(t, buf)
}

func TestMustExecuteGo(t *testing.T) {
	okTpl, err := texttemplate.New("").Parse("package main\nimport \"fmt\"\nfunc main() { fmt.Println(\"{{ . }}\") }")
	require.NoError(t, err)

	errTpl, err := texttemplate.New("").Parse(`{{ template "x" }}`)
	require.NoError(t, err)

	goErrTpl, err := texttemplate.New("").Parse("package main\nfuncmain() { fmt.Println(\"{{ . }}\") }")
	require.NoError(t, err)

	require.NotPanics(t, func() {
		buf := templatez.MustExecuteGo(okTpl, "Hello World")
		require.Equal(t, "package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Println(\"Hello World\") }\n", string(buf))
	})

	require.PanicsWithError(t, `template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`, func() {
		templatez.MustExecuteGo(errTpl, nil)
	})

	require.PanicsWithError(t, "2:1: expected declaration, found funcmain", func() {
		templatez.MustExecuteGo(goErrTpl, "Hello World")
	})
}

func TestParseAndExecuteJSON(t *testing.T) {
	buf, err := templatez.ParseAndExecuteJSON(`{ "value": "{{ . }}" }`, "|", "  ", "Hello World")
	require.NoError(t, err)
	require.Equal(t, "{\n|  \"value\": \"Hello World\"\n|}", string(buf))

	buf, err = templatez.ParseAndExecuteJSON("{{ bad }}", "|", "  ", "Hello World")
	require.EqualError(t, err, `template: :1: function "bad" not defined`)
	require.Nil(t, buf)

	buf, err = templatez.ParseAndExecuteJSON(`{{ template "x" }}`, "|", "  ", nil)
	require.EqualError(t, err, `template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)
	require.Nil(t, buf)

	buf, err = templatez.ParseAndExecuteJSON(`{ "value": {{ . }} }`, "|", "  ", "Hello World")
	require.EqualError(t, err, "invalid character 'H' looking for beginning of value")
	require.Nil(t, buf)
}

func TestMustParseAndExecuteJSON(t *testing.T) {
	require.NotPanics(t, func() {
		buf := templatez.MustParseAndExecuteJSON(`{ "value": "{{ . }}" }`, "|", "  ", "Hello World")
		require.Equal(t, "{\n|  \"value\": \"Hello World\"\n|}", string(buf))
	})

	require.PanicsWithError(t, `template: :1: function "bad" not defined`, func() {
		templatez.MustParseAndExecuteJSON("{{ bad }}", "|", "  ", "Hello World")
	})

	require.PanicsWithError(t, `template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`, func() {
		templatez.MustParseAndExecuteJSON(`{{ template "x" }}`, "|", "  ", nil)
	})

	require.PanicsWithError(t, "invalid character 'H' looking for beginning of value", func() {
		templatez.MustParseAndExecuteJSON(`{ "value": {{ . }} }`, "|", "  ", "Hello World")
	})
}

func TestExecuteJSON(t *testing.T) {
	okTpl, err := texttemplate.New("").Parse(`{ "value": "{{ . }}" }`)
	require.NoError(t, err)

	errTpl, err := texttemplate.New("").Parse(`{{ template "x" }}`)
	require.NoError(t, err)

	jsonErrTpl, err := texttemplate.New("").Parse(`{ "value": {{ . }} }`)
	require.NoError(t, err)

	buf, err := templatez.ExecuteJSON(okTpl, "|", "  ", "Hello World")
	require.NoError(t, err)
	require.Equal(t, "{\n|  \"value\": \"Hello World\"\n|}", string(buf))

	buf, err = templatez.ExecuteJSON(errTpl, "|", "  ", nil)
	require.EqualError(t, err, `template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`)
	require.Nil(t, buf)

	buf, err = templatez.ExecuteJSON(jsonErrTpl, "|", "  ", "Hello World")
	require.EqualError(t, err, "invalid character 'H' looking for beginning of value")
	require.Nil(t, buf)
}

func TestMustExecuteJSON(t *testing.T) {
	okTpl, err := texttemplate.New("").Parse(`{ "value": "{{ . }}" }`)
	require.NoError(t, err)

	errTpl, err := texttemplate.New("").Parse(`{{ template "x" }}`)
	require.NoError(t, err)

	jsonErrTpl, err := texttemplate.New("").Parse(`{ "value": {{ . }} }`)
	require.NoError(t, err)

	require.NotPanics(t, func() {
		buf := templatez.MustExecuteJSON(okTpl, "|", "  ", "Hello World")
		require.Equal(t, "{\n|  \"value\": \"Hello World\"\n|}", string(buf))
	})

	require.PanicsWithError(t, `template: :1:12: executing "" at <{{template "x"}}>: template "x" not defined`, func() {
		templatez.MustExecuteJSON(errTpl, "|", "  ", nil)
	})

	require.PanicsWithError(t, "invalid character 'H' looking for beginning of value", func() {
		templatez.MustExecuteJSON(jsonErrTpl, "|", "  ", "Hello World")
	})
}

func TestParseAndExecuteHTML(t *testing.T) {
	buf, err := templatez.ParseAndExecuteHTML("{{ . }}", "<a>&</a>")
	require.NoError(t, err)
	require.Equal(t, "&lt;a&gt;&amp;&lt;/a&gt;", string(buf))

	buf, err = templatez.ParseAndExecuteHTML("{{ bad }}", "<a>&</a>")
	require.EqualError(t, err, `template: :1: function "bad" not defined`)
	require.Nil(t, buf)

	buf, err = templatez.ParseAndExecuteHTML(`{{ template "x" }}`, nil)
	require.EqualError(t, err, `html/template::1:12: no such template "x"`)
	require.Nil(t, buf)
}

func TestMustParseAndExecuteHTML(t *testing.T) {
	require.NotPanics(t, func() {
		buf := templatez.MustParseAndExecuteHTML("{{ . }}", "<a>&</a>")
		require.Equal(t, "&lt;a&gt;&amp;&lt;/a&gt;", string(buf))
	})

	require.PanicsWithError(t, `template: :1: function "bad" not defined`, func() {
		templatez.MustParseAndExecuteHTML("{{ bad }}", "<a>&</a>")
	})

	require.PanicsWithError(t, `html/template::1:12: no such template "x"`, func() {
		templatez.MustParseAndExecuteHTML(`{{ template "x" }}`, nil)
	})
}

func TestExecuteHTML(t *testing.T) {
	okTpl, err := htmltemplate.New("").Parse("{{ . }}")
	require.NoError(t, err)

	errTpl, err := htmltemplate.New("").Parse(`{{ template "x" }}`)
	require.NoError(t, err)

	buf, err := templatez.ExecuteHTML(okTpl, "<a>&</a>")
	require.NoError(t, err)
	require.Equal(t, "&lt;a&gt;&amp;&lt;/a&gt;", string(buf))

	buf, err = templatez.ExecuteHTML(errTpl, nil)
	require.EqualError(t, err, `html/template::1:12: no such template "x"`)
	require.Nil(t, buf)
}

func TestMustExecuteHTML(t *testing.T) {
	okTpl, err := htmltemplate.New("").Parse("{{ . }}")
	require.NoError(t, err)

	errTpl, err := htmltemplate.New("").Parse(`{{ template "x" }}`)
	require.NoError(t, err)

	require.NotPanics(t, func() {
		buf := templatez.MustExecuteHTML(okTpl, "<a>&</a>")
		require.Equal(t, "&lt;a&gt;&amp;&lt;/a&gt;", string(buf))
	})

	require.PanicsWithError(t, `html/template::1:12: no such template "x"`, func() {
		templatez.MustExecuteHTML(errTpl, nil)
	})
}

func TestParseAndExecuteHTMLTidy(t *testing.T) {
	buf, err := templatez.ParseAndExecuteHTMLTidy("<html><body>{{ . }}</body></html>", "<a>&</a>")
	require.NoError(t, err)
	require.Equal(t, "<html>\n  <body>\n    &lt;a&gt;&amp;&lt;/a&gt;\n  </body>\n</html>\n", string(buf))

	buf, err = templatez.ParseAndExecuteHTMLTidy("{{ bad }}", "<a>&</a>")
	require.EqualError(t, err, `template: :1: function "bad" not defined`)
	require.Nil(t, buf)

	buf, err = templatez.ParseAndExecuteHTMLTidy(`{{ template "x" }}`, nil)
	require.EqualError(t, err, `html/template::1:12: no such template "x"`)
	require.Nil(t, buf)
}

func TestMustParseAndExecuteHTMLTidy(t *testing.T) {
	require.NotPanics(t, func() {
		buf := templatez.MustParseAndExecuteHTMLTidy("<html><body>{{ . }}</body></html>", "<a>&</a>")
		require.Equal(t, "<html>\n  <body>\n    &lt;a&gt;&amp;&lt;/a&gt;\n  </body>\n</html>\n", string(buf))
	})

	require.PanicsWithError(t, `template: :1: function "bad" not defined`, func() {
		templatez.MustParseAndExecuteHTMLTidy("{{ bad }}", "<a>&</a>")
	})

	require.PanicsWithError(t, `html/template::1:12: no such template "x"`, func() {
		templatez.MustParseAndExecuteHTMLTidy(`{{ template "x" }}`, nil)
	})
}

func TestExecuteHTMLTidy(t *testing.T) {
	okTpl, err := htmltemplate.New("").Parse("<html><body>{{ . }}</body></html>")
	require.NoError(t, err)

	errTpl, err := htmltemplate.New("").Parse(`{{ template "x" }}`)
	require.NoError(t, err)

	buf, err := templatez.ExecuteHTMLTidy(okTpl, "<a>&</a>")
	require.NoError(t, err)
	require.Equal(t, "<html>\n  <body>\n    &lt;a&gt;&amp;&lt;/a&gt;\n  </body>\n</html>\n", string(buf))

	buf, err = templatez.ExecuteHTMLTidy(errTpl, nil)
	require.EqualError(t, err, `html/template::1:12: no such template "x"`)
	require.Nil(t, buf)
}

func TestMustExecuteHTMLTidy(t *testing.T) {
	okTpl, err := htmltemplate.New("").Parse("<html><body>{{ . }}</body></html>")
	require.NoError(t, err)

	errTpl, err := htmltemplate.New("").Parse(`{{ template "x" }}`)
	require.NoError(t, err)

	require.NotPanics(t, func() {
		buf := templatez.MustExecuteHTMLTidy(okTpl, "<a>&</a>")
		require.Equal(t, "<html>\n  <body>\n    &lt;a&gt;&amp;&lt;/a&gt;\n  </body>\n</html>\n", string(buf))
	})

	require.PanicsWithError(t, `html/template::1:12: no such template "x"`, func() {
		templatez.MustExecuteHTMLTidy(errTpl, nil)
	})
}
