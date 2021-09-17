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
