package main

import (
	"embed"
	"fmt"
	"strings"
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

func GetStandardTemplates() []string {
	files, readErr := templateFS.ReadDir("templates")
	if readErr != nil {
		panic(readErr)
	}

	var templateNames []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".tmpl") {
			templateNames = append(templateNames, file.Name()[0:len(file.Name())-5])
		}
	}

	return templateNames
}
