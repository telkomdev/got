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

	// create file
	output, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}

	// close file
	defer output.Close()

	err = got.Parse(tmpl, data, output, funcMaps, got.ParseTemplateText)

	if err != nil {
		fmt.Println(err)
	}

}
