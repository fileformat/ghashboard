# To Do

- [ ] dogfood GHA: build a ghashboard of my stuff
- [ ] demo GHA: find a couple orgs with a middling number of GHAs for a demo ghashboard

## repos.go

builds a list of repos

 * public: boolean 
 * private: boolean
 * forks: boolean
 * archived: boolean
 * repos: a list of repos (or @filename), bypasses all flags

## workflows.go

builds a list of actions for each repo

 * ghpages: skip `pages-build-deployment` since the badges don't work
 * namefilter: regex on name
 * filefilter: regex on filename (from path)
 * status: `[active|all]`
 * actions: a JSON file with a custom-generated (or saved) list of repos/actions

## badges.go

adds non-action badges to the list

 * https://shields.io/badges/git-hub-last-commit-branch
 * https://shields.io/badges/git-hub-issues
 * https://shields.io/badges/git-hub-forks
 * https://shields.io/badges/git-hub-pull-requests
 * https://badgen.net/github/last-commit/
 * https://badgen.net/github

## render.go

builds page

 * output: `[markdown|html|csv|json|demo]`
 * template: custom template file (also load entire template from env for customizing forks)

Potential go libraries:
https://github.com/knadh/koanf
https://pkg.go.dev/github.com/google/go-github/v58/github

## Maybe

 * verify that the badges are on the README page.
