# To Do

- [ ] custom template (or @file)
- [ ] demo repo (or link to usage)
- [ ] better README
- [ ] sort by full_name, name, pushed_at, updated_at
- [ ] new arg: template, (and @filename)

## Maybe not

- [ ] external badge definitions from a yaml file (but works fine as-is)
- [ ] Golang template function `markdown` to escape markdown (is it really needed?)
- [ ] include/exclude/skip: support for regexes (can workaround with `@repos` file and grep)
- [ ] output file extension based on format
- [ ] progress indicators (if output!=stdout && stdout == tty)
- [ ] custom external badges (could be done by customizing the template)
- [ ] verify that the badges are on the README page (way too much work: maybe a separate utility?)

- [ ] make releases
- [ ] Fix: `pages-build-deployment` *do* have working flags - [Github Bug](https://support.github.com/ticket/personal/0/2545577)

## Bookmarks

- [pflag](https://pkg.go.dev/github.com/spf13/pflag)
- [github.Repository](https://pkg.go.dev/github.com/google/go-github/v58@v58.0.0/github#Repository)
- [github.Workflow](https://pkg.go.dev/github.com/google/go-github/v58@v58.0.0/github#Workflow)
- Github REST API:
  [Actions](https://docs.github.com/en/rest/actions?apiVersion=2022-11-28)
  | [Workflows](https://docs.github.com/en/rest/actions/workflows?apiVersion=2022-11-28)


Potential go libraries:
https://github.com/knadh/koanf
https://pkg.go.dev/github.com/google/go-github/v58/github

## Maybe

 * handle rate limit [library](https://github.com/gofri/go-github-ratelimit).  Would be useful for anonymous runs, since that rate limit is 60 per *hour* [docs](https://docs.github.com/en/rest/using-the-rest-api/rate-limits-for-the-rest-api?apiVersion=2022-11-28)

 * https://gith ub.com/VectorLogoZone/svgzone/workflows/pages-build-deployment/badge.svg
 * https://github.com/VectorLogoZone/svgzone/actions/workflows/pages/pages-build-deployment/badge.svg (website)

 * https://github.com/VectorLogoZone/svgzone/blob/main/dynamic/pages/pages-build-deployment (api)
 * https://github.com/VectorLogoZone/svgzone/actions/workflows/pages/pages-build-deployment

 gh cmd line
       "html_url": "https://github.com/VectorLogoZone/svgzone/blob/main/dynamic/pages/pages-build-deployment",
      "badge_url": "https://github.com/VectorLogoZone/svgzone/workflows/pages-build-deployment/badge.svg"
