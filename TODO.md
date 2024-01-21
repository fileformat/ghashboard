# To Do

- [ ] security notes page
- [ ] screenshot of warning when enabling Github Actions
- [ ] customization page
- [ ] Fix: `pages-build-deployment` *do* have working flags - [Github Bug](https://support.github.com/ticket/personal/0/2545577)
- [ ] switch template data from array to struct with array as a property, and add more properties (Title, Created, Header)
- [ ] Golang template function `markdown` to escape markdown
- [ ] dogfood GHA: build a ghashboard of my stuff
- [ ] fork-is-current GHA
- [ ] demo GHA: find a couple orgs with a middling number of GHAs for a demo ghashboard
- [ ] support multi-line env vars for `[]String` flags
- [ ] make releases

## Bookmarks

- [pflag](https://pkg.go.dev/github.com/spf13/pflag)
- [github.Repository](https://pkg.go.dev/github.com/google/go-github/v58@v58.0.0/github#Repository)
- [github.Workflow](https://pkg.go.dev/github.com/google/go-github/v58@v58.0.0/github#Workflow)
- Github REST API:
  [Actions](https://docs.github.com/en/rest/actions?apiVersion=2022-11-28)
  | [Workflows](https://docs.github.com/en/rest/actions/workflows?apiVersion=2022-11-28)

## repos.go

builds a list of repos

 * public: boolean 
 * private: boolean
 * forks: boolean
 * archived: boolean
 * repos: a list of repos (or @filename), bypasses all flags

## workflows.go

builds a list of actions for each repo

 * empty flag: include repos w/o workflows in the list (for non-workflow badges)

## badges.go

adds external (non-Github Action) badges to the list
 * Go Template that takes a Repository
 * built-in badges (see below): map of string to template

built-in:
 * https://shields.io/badges/git-hub-last-commit-branch
 * https://shields.io/badges/git-hub-issues
 * https://shields.io/badges/git-hub-forks
 * https://shields.io/badges/git-hub-pull-requests
 * https://badgen.net/github/last-commit/
 * https://badgen.net/github
 * https://img.shields.io/github/license/VectorLogoZone/vectorlogozone.svg

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
 * handle rate limit [library](https://github.com/gofri/go-github-ratelimit).  Would be useful for anonymous runs, since that rate limit is 60 per *hour* [docs](https://docs.github.com/en/rest/using-the-rest-api/rate-limits-for-the-rest-api?apiVersion=2022-11-28)

 * https://github.com/VectorLogoZone/svgzone/workflows/pages-build-deployment/badge.svg
 * https://github.com/VectorLogoZone/svgzone/actions/workflows/pages/pages-build-deployment/badge.svg (website)

 * https://github.com/VectorLogoZone/svgzone/blob/main/dynamic/pages/pages-build-deployment (api)
 * https://github.com/VectorLogoZone/svgzone/actions/workflows/pages/pages-build-deployment

 gh cmd line
       "html_url": "https://github.com/VectorLogoZone/svgzone/blob/main/dynamic/pages/pages-build-deployment",
      "badge_url": "https://github.com/VectorLogoZone/svgzone/workflows/pages-build-deployment/badge.svg"
