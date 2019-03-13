package got

import (
	"io"
	"path/filepath"
	"text/template"
)

// ParserTemplateFunc type of parser function
type ParserTemplateFunc func(string, interface{}, template.FuncMap, io.Writer) error

// ParseTemplateFile function
// will parse template from file
func ParseTemplateFile(filePath string, data interface{}, funcMaps template.FuncMap, result io.Writer) error {
	tName := filepath.Base(filePath)
	tmpl, err := template.New(tName).Funcs(funcMaps).ParseFiles(filePath)
	if err != nil {
		return err
	}

	err = tmpl.Execute(result, data)
	if err != nil {
		return err
	}

	return nil
}

// ParseTemplateText function
// will parse template from text
func ParseTemplateText(text string, data interface{}, funcMaps template.FuncMap, result io.Writer) error {
	tmpl, err := template.New("tmpl").Funcs(funcMaps).Parse(text)
	if err != nil {
		return err
	}

	err = tmpl.Execute(result, data)
	if err != nil {
		return err
	}

	return nil
}

// Parse will parse specific input file to io.Writer
// you can put anything to target parameter, as long as target is io.Writer implementation
func Parse(input string, data interface{}, target io.Writer, funcMaps template.FuncMap, f ParserTemplateFunc) error {

	err := f(input, data, funcMaps, target)
	if err != nil {
		return err
	}

	return nil
}
