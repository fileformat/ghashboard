package main

import (
	"embed"
	"fmt"
	"text/template"
)

//go:embed templates
var templateFS embed.FS

func GetStandardTemplate(templateName string) (*template.Template, error) {

	fileName := fmt.Sprintf("templates/%s.tmpl", templateName)
	rawTmpl, readErr := templateFS.ReadFile(fileName)
	if readErr != nil {
		return nil, readErr
	}

	tmpl, parseErr := template.New(templateName).Parse(string(rawTmpl))
	if parseErr != nil {
		return nil, parseErr
	}

	return tmpl, nil
}
