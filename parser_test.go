package got

import (
	"os"
	"strings"
	"testing"
	"text/template"
)

func TestTemplateParser(t *testing.T) {

	t.Run("should fail parse template text to strings.Builder", func(t *testing.T) {
		data := struct {
			ID   string
			Name string
		}{
			ID:   "1",
			Name: "Wuriyanto",
		}

		tmpl := `
		Hello
		Your ID : {{ .NO }}
		Your Name : {{ .FullName }}
		`

		expectedTmpl := `
		Hello
		Your ID : 1
		Your Name : Wuriyanto
		`

		var target strings.Builder

		err := Parse(tmpl, data, &target, nil, ParseTemplateText)

		if err == nil {
			t.Error("should error parsing text template")
		}

		if target.String() == expectedTmpl {
			t.Error("result should not equal to expected")
		}
	})

	t.Run("should success parse template text to os.Stdout", func(t *testing.T) {
		data := struct {
			ID   string
			Name string
		}{
			ID:   "1",
			Name: "Wuriyanto",
		}

		tmpl := `
		Hello
		Your ID : {{ .ID }}
		Your Name : {{ toUpper .Name }}
		`

		funcMaps := template.FuncMap{
			"toUpper": func(v string) string {
				return strings.ToUpper(v)
			},
		}

		err := Parse(tmpl, data, os.Stdout, funcMaps, ParseTemplateText)

		if err != nil {
			t.Error(err)
		}

	})

	t.Run("should success parse template text to strings.Builder", func(t *testing.T) {
		data := struct {
			ID   string
			Name string
		}{
			ID:   "1",
			Name: "Wuriyanto",
		}

		tmpl := `
		Hello
		Your ID : {{ .ID }}
		Your Name : {{ toUpper .Name }}
		`

		expectedTmpl := `
		Hello
		Your ID : 1
		Your Name : WURIYANTO
		`

		funcMaps := template.FuncMap{
			"toUpper": func(v string) string {
				return strings.ToUpper(v)
			},
		}

		var target strings.Builder

		err := Parse(tmpl, data, &target, funcMaps, ParseTemplateText)

		if err != nil {
			t.Error(err)
		}

		if target.String() != expectedTmpl {
			t.Error("error result is not equal to expected")
		}
	})
}
