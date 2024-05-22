package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	Verbose bool
)

func intro() {
	fmt.Print(`
Ghashboard is a tool for creating a Github Actions dashboard full of badges.

The dashboard page is static and can be a Github README.md file.

If you include private repos, the page should be hosted on Github (or your badges will get 404ed).

The badge images will update automatically (by Github).

Example usage:
ghashboard --owners=owner1,owner2 --repos=repo1,repo2 --private=true --output=dashboard.md

For detailed usage:
ghashboard --help`)
}

func usage(f *pflag.FlagSet) {
	fmt.Printf("Usage: %s [options] [file]\n", path.Base(os.Args[0]))
	fmt.Printf("\n")
	fmt.Printf("Create a Github Actions dashboard full of badges\n")
	fmt.Printf("\n")
	fmt.Printf("%s\n", f.FlagUsages())
	fmt.Printf("\n")
	fmt.Printf("      file: output file (default: stdout)\n")
	fmt.Printf("\n")
	fmt.Printf("      options can also be set via environment variables\n")
	fmt.Printf("      set GH_TOKEN to use a personal access token and avoid rate limit errors.\n")
	fmt.Printf("\n")
	fmt.Printf("    built-in external badges:\n")
	for key, value := range getBuiltins() {
		fmt.Printf("      %s - %s\n", key, value)
	}
}

func initConfig(args []string) {

	if len(args) == 0 {
		intro()
		os.Exit(0)
	}
	f := pflag.NewFlagSet("config", pflag.ExitOnError)
	help := f.Bool("help", false, "Show help")
	f.MarkHidden("help")

	versionFlag := f.Bool("version", false, "Show version information")

	f.String("log-level", "warn", "Log level [ debug | info | warn | error ]")
	viper.BindPFlag("log-level", f.Lookup("log-level"))
	viper.BindEnv("log-level", "LOG_LEVEL", "INPUT_LOG_LEVEL")

	/*
	 * flags to build the repo list
	 */
	f.String("owners", "", "Owners")
	viper.BindPFlag("owners", f.Lookup("owners"))
	viper.BindEnv("owners", "OWNERS", "INPUT_OWNERS")

	f.String("repos", "", "Repositories")
	viper.BindPFlag("repos", f.Lookup("repos"))
	viper.BindEnv("repos", "REPOS", "INPUT_REPOS")

	// deliberately no flag because security
	viper.BindEnv("github-token", "INPUT_TOKEN", "GITHUB_TOKEN")

	/*
	 * flags for filtering the list
	 */
	f.Bool("empty", false, "Include repos with no eligible workflows")
	viper.BindPFlag("empty", f.Lookup("empty"))
	viper.BindEnv("empty", "EMPTY", "INPUT_EMPTY")

	f.Bool("private", false, "Include private repos")
	viper.BindPFlag("private", f.Lookup("private"))
	viper.BindEnv("private", "PRIVATE", "INPUT_PRIVATE")

	f.Bool("public", true, "Include public repos")
	viper.BindPFlag("public", f.Lookup("public"))
	viper.BindEnv("public", "PUBLIC", "INPUT_PUBLIC")

	f.Bool("forks", false, "Include forks")
	viper.BindPFlag("forks", f.Lookup("forks"))
	viper.BindEnv("forks", "FORKS", "INPUT_FORKS")

	f.Bool("archived", false, "Include archived repos")
	viper.BindPFlag("archived", f.Lookup("archived"))
	viper.BindEnv("archived", "ARCHIVED", "INPUT_ARCHIVED")

	f.Bool("inactive", false, "Include inactive workflows")
	viper.BindPFlag("inactive", f.Lookup("inactive"))
	viper.BindEnv("inactive", "INACTIVE", "INPUT_INACTIVE")

	f.String("include", "", "Actions to include")
	viper.BindPFlag("include", f.Lookup("include"))
	viper.BindEnv("include", "INCLUDE", "INPUT_INCLUDE")

	f.String("exclude", "codeql,pages-build-deployment", "Actions to exclude")
	viper.BindPFlag("exclude", f.Lookup("exclude"))
	viper.BindEnv("exclude", "EXCLUDE", "INPUT_EXCLUDE")

	/*
	 * flags to control the output
	 */
	f.String("externals", "", "External badges")
	viper.BindPFlag("externals", f.Lookup("externals"))
	viper.BindEnv("externals", "EXTERNALS", "INPUT_EXTERNALS")

	f.String("format", "markdown", "Output format [ "+strings.Join(GetStandardTemplates(), " | ")+" | json ]")
	viper.BindPFlag("format", f.Lookup("format"))
	viper.BindEnv("format", "FORMAT", "INPUT_FORMAT")

	f.String("output", "dashboard.md", "File name (or - for stdout)")
	viper.BindPFlag("output", f.Lookup("output"))
	viper.BindEnv("output", "OUTPUT", "INPUT_OUTPUT")

	f.Parse(args)

	if *versionFlag {
		fmt.Printf("ghashboard v%s (%s built on %s by %s)\n", version, commit, date, builtBy)
		os.Exit(0)
	}

	if *help {
		usage(f)
		os.Exit(0)
	}
}

var arraySplitter = regexp.MustCompile(`[\s,]+`)

func viper_GetStringSlice(key string) []string {
	str := viper.GetString(key)
	if str == "" {
		return []string{}
	}
	return arraySplitter.Split(str, -1)
}
