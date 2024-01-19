package main

import (
	"fmt"
	"os"
	"text/template"
)

func GetStandardTemplate(templateName string) (*template.Template, error) {

	fileName := fmt.Sprintf("templates/%s.tmpl", templateName)
	rawTmpl, readErr := os.ReadFile(fileName)
	if readErr != nil {
		return nil, readErr
	}

	tmpl, parseErr := template.New(templateName).Parse(string(rawTmpl))
	if parseErr != nil {
		return nil, parseErr
	}

	return tmpl, nil
}
