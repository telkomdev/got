package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/telkomdev/got"
)

// Person struct
type Person struct {
	ID   string
	Name string
}

func main() {

	data := Person{
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

	err := got.Parse(tmpl, data, os.Stdout, funcMaps, got.ParseTemplateText)

	if err != nil {
		fmt.Println(err)
	}

}
