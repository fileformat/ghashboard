# To Do

- [ ] Something is wrong with the links
- [ ] dogfood GHA: build a ghashboard of my stuff
- [ ] demo GHA: find a couple orgs with a middling number of GHAs for a demo ghashboard
- [ ] switch template data from array to struct with array as a property, and add more properties (Created, Header)
- [ ] support multi-line env vars for `[]String` flags
- [ ] make releases

## repos.go

builds a list of repos

 * public: boolean 
 * private: boolean
 * forks: boolean
 * archived: boolean
 * repos: a list of repos (or @filename), bypasses all flags

## workflows.go

builds a list of actions for each repo

 * exclude: list of workflows to exclude (default to `codeql,pages-build-deployment`)
 * inactive: boolean, including inactive 
 * include: list of workflows to include
 * 

## badges.go

adds non-action badges to the list

 * flag to include repos w/o GHAs
 * https://shields.io/badges/git-hub-last-commit-branch
 * https://shields.io/badges/git-hub-issues
 * https://shields.io/badges/git-hub-forks
 * https://shields.io/badges/git-hub-pull-requests
 * https://badgen.net/github/last-commit/
 * https://badgen.net/github

## template.go

 * hyperlink the repo name
 * better indenting in markdown.tmpl
 * if only one owner, don't use FullName
 * header/footer: markdown blobs to include in template
 * output: html?
 * template: custom template file
 * template-string: load entire template from env (for customizing forks with a GHA variable)

Potential go libraries:
https://github.com/knadh/koanf
https://pkg.go.dev/github.com/google/go-github/v58/github

## Maybe

 * verify that the badges are on the README page.
