# Github Actions Dashboard (aka ghashboard) [<img alt="Ghashboard Logo" src="docs/favicon.svg" height="96" align="right"/>](https://ghashboard.marcuse.info/)

[![build](https://github.com/fileformat/ghashboard/actions/workflows/build.yaml/badge.svg)](https://github.com/fileformat/ghashboard/actions/workflows/build.yaml)
[![release](https://github.com/fileformat/ghashboard/actions/workflows/release.yaml/badge.svg)](https://github.com/fileformat/ghashboard/actions/workflows/release.yaml)
[![Version](https://img.shields.io/github/v/tag/fileformat/ghashboard?sort=semver&style=plastic&label=Version&labelColor=black&color=blue)](https://github.com/fileformat/ghashboard/releases)


A tool for making dashboard pages with all the badges from your Github Actions.

## Examples

[![VectorLogoZone example](docs/images/thumbnails/08_result.png)](docs/images/screenshots/08_result.png "Example w/defaults")
[![External badges example](docs/images/thumbnails/10_customized_result.png)](docs/images/screenshots/10_customized_result.png "Example w/additional external badges")

## Pronunciation

It is pronounced (in an upper-crust accent) "gosh-board", not "gashboard" like some sort of horror flick.

## Using via Github Actions

Recursion detected:  Stack overflow!  ...just kidding!  This is actually a great way to run it.

It works really well to as a generator for Github profiles (example: [code](https://github.com/VectorLogoZone/.github) and [result](https://github.com/VectorLogoZone)).

## Using via CLI

Download the latest version from the [Github Releases page](https://github.com/fileformat/ghashboard/releases)

A sample command:
```
export GITHUB_TOKEN=$(gh auth token)
ghashboard --owners=google,spf13
```

## Options

Either as inputs for the Github Action or flags for the CLI.

Note: either `owners` or `repos` is required.

| Name       |  Description          |
| ---------- |  -------------------- |
|  `archived` |  include archived repos? (default is `false`) |
|  `empty` |  include repos with no workflows: useful if you have external badges (default is `false`) |
|  `exclude` |  workflows to exclude (comma-separated list) |
|  `externals` |  external badges ([list](https://github.com/fileformat/ghashboard/blob/main/externalbadges.go#L25)) to include.  Note that these usually only work with public repos |
|  `footer` |  footer text |
|  `forks` |  include forked repos? |
|  `format` |  output format: `csv`, `markdown` or `json`.  `json` is good for debugging.  (default is `markdown`) |
|  `header` |  header text |
|  `inactive` |  include inactive workflows? |
|  `include` |  workflows to include: others will be skipped (comma-separated list)) |
|  `log-level` |  log level: `debug`, `info`, `warn` or `error` |
|  `output` |  output file to create (or `-` for stdout) |
|  `owners` |  list of owners  |
|  `private` |  include private repos? (one of public|private must be true) |
|  `public` |  include public repos? (one of public|private must be true) |
|  `repos` |  list of repos |
|  `skip` |  exclude specific repos (use `owner/name`) |
|  `title` |  title text |
|  `token` |  Github token: not absolutely required, but you will run into rate-limits without one |

## Developing

See `run.sh` for how I run it during development.

## Contributing

Contributions welcome!

## License

[MIT](LICENSE.txt)

## Credits

[![Git](https://www.vectorlogo.zone/logos/git-scm/git-scm-ar21.svg)](https://git-scm.com/ "Version control")
[![Github](https://www.vectorlogo.zone/logos/github/github-ar21.svg)](https://github.com/ "Code hosting")
[![golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org/ "Programming language")
[![Google Noto Emoji](https://www.vectorlogo.zone/logos/google/google-ar21.svg)](https://github.com/googlefonts/noto-emoji/ "Logo")
[![Markdown](https://www.vectorlogo.zone/logos/commonmark/commonmark-ar21.svg)](https://commonmark.org/ "CommonMark Markdown")
[![Shields.IO](https://www.vectorlogo.zone/logos/shieldsio/shieldsio-ar21.svg)](http://shields.io/ "README badges")
[![VectorLogoZone](https://www.vectorlogo.zone/logos/vectorlogozone/vectorlogozone-ar21.svg)](https://www.vectorlogo.zone/ "Logos")

* [GoReleaser](https://goreleaser.com/) - packaging for release
* [Steve Francia](https://spf13.com/) - [viper](https://github.com/spf13/viper)
* [Steve Francia](https://spf13.com/) - [cobra](https://github.com/spf13/cobra)
* See [`go.mod`](https://github.com/fileformat/ghashboard/blob/main/go.mod) for other golang modules used
* [jq](https://jqlang.github.io/jq/) - JSON manipulation
* [cb](https://github.com/niedzielski/cb) - clipboard utility
