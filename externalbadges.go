package main

import (
	"fmt"
	"strings"
	"text/template"

	github "github.com/google/go-github/v58/github"
)

type ExternalBadge struct {
	Name        string
	Description string
	Image       string
	Link        string
}

type ExternalBadgeTmpl struct {
	Name        string
	Description string
	ImageTmpl   *template.Template
	LinkTmpl    *template.Template
}

var builtins = map[string]ExternalBadgeTmpl{
	"license": {
		// LATER: maybe make an extra call to get the license: https://pkg.go.dev/github.com/google/go-github/v58@v58.0.0/github#Repository.GetLicense
		// Badge URL looks like: https://img.shields.io/github/license/VectorLogoZone/vectorlogozone.svg
		// Target URL looks like: https://github.com/search?q=repo%3AVectorLogoZone%2Fvectorlogozone%20path%3A%2Flicense.*&type=code
		Name:        "License",
		Description: "License",
		ImageTmpl:   template.Must(template.New("license_image").Parse(`https://img.shields.io/github/license/{{.FullName}}.svg`)),
		LinkTmpl:    template.Must(template.New("license_link").Parse(`https://github.com/search?q=repo%3A{{.Owner.Login}}%2F{{urlquery .Name}}%20path%3A%2Flicense.*&type=code`)),
	},
}

func getBuiltins() map[string]string {
	retVal := make(map[string]string)
	for key, _ := range builtins {
		retVal[key] = builtins[key].Description
	}
	return retVal
}

func GenerateExternalBadge(name string, repo *github.Repository) (*ExternalBadge, error) {

	externalBadgeTmpl, ok := builtins[name]
	if !ok {
		return nil, fmt.Errorf("no standard badge named %q", name)
	}

	var image strings.Builder
	imageErr := externalBadgeTmpl.ImageTmpl.Execute(&image, repo)
	if imageErr != nil {
		return nil, imageErr
	}
	var link strings.Builder
	linkErr := externalBadgeTmpl.LinkTmpl.Execute(&link, repo)
	if linkErr != nil {
		return nil, linkErr
	}

	externalBadge := ExternalBadge{
		Name:        externalBadgeTmpl.Name,
		Description: externalBadgeTmpl.Description,
		Image:       image.String(),
		Link:        link.String(),
	}
	return &externalBadge, nil
}
