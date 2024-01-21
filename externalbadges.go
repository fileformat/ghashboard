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
		Description: "License (as found by Github)",
		ImageTmpl:   template.Must(template.New("license_image").Parse(`https://img.shields.io/github/license/{{.FullName}}.svg?style=plastic`)),
		LinkTmpl:    template.Must(template.New("license_link").Parse(`https://github.com/search?q=repo%3A{{.Owner.Login}}%2F{{urlquery .Name}}%20path%3A%2Flicense.*&type=code`)),
	},
	/* NO: I do NOT want to encourage people to chase stars
	"stars": {
		Name:        "Stars",
		Description: "Github stars",
		ImageTmpl:   template.Must(template.New("stars_image").Parse(`https://img.shields.io/github/stars/{{.FullName}}.svg?style=plastic`)),
		LinkTmpl:    template.Must(template.New("stars_link").Parse(`https://github.com/{{.FullName}}/stargazers`)),
	},
	*/
	"forks": {
		Name:        "Forks",
		Description: "Github forks",
		ImageTmpl:   template.Must(template.New("forks_image").Parse(`https://img.shields.io/github/forks/{{.FullName}}.svg?style=plastic`)),
		LinkTmpl:    template.Must(template.New("forks_link").Parse(`https://github.com/{{.FullName}}/forks`)),
	},
	"openissues": {
		Name:        "Open issues",
		Description: "Github open issues",
		ImageTmpl:   template.Must(template.New("openissues_image").Parse(`https://img.shields.io/github/issues-raw/{{.FullName}}.svg?style=plastic`)),
		LinkTmpl:    template.Must(template.New("openissues_link").Parse(`https://github.com/{{.FullName}}/issues`)),
	},
	"lastcommit": {
		Name:        "Last commit",
		Description: "Most recent commit on default branch",
		ImageTmpl:   template.Must(template.New("lastcommit_image").Parse(`https://img.shields.io/github/last-commit/{{.FullName}}/{{.DefaultBranch}}.svg?style=plastic`)),
		LinkTmpl:    template.Must(template.New("lastcommit_link").Parse(`https://github.com/{{.FullName}}/commits/{{.DefaultBranch}}/`)),
	},
	"reposize": {
		Name:        "Repo size",
		Description: "Size of the repository",
		ImageTmpl:   template.Must(template.New("reposize_image").Parse(`https://img.shields.io/github/repo-size/{{.FullName}}.svg?style=plastic`)),
		LinkTmpl:    template.Must(template.New("reposize_link").Parse(`https://stackoverflow.com/a/76995820`)), // really, Github???
	},
}

func getBuiltins() map[string]string {
	retVal := make(map[string]string)
	for key := range builtins {
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
