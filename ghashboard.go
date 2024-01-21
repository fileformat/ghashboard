package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	github "github.com/google/go-github/v58/github"
	flag "github.com/spf13/pflag"
)

type MetaRepo struct {
	Workflows      []*github.Workflow `json:"workflows"`
	Repo           *github.Repository `json:"repo"`
	ExternalBadges []*ExternalBadge   `json:"external"`
}

func setFlagsFromEnvironment(f *flag.FlagSet) {
	f.VisitAll(func(oneFlag *flag.Flag) {
		flagName := oneFlag.Name
		envName := "GHASHBOARD_" + strings.ToUpper(strings.Replace(flagName, "-", "_", -1))
		value, ok := os.LookupEnv(envName)
		fmt.Printf("DEBUG: env lookup for %s: '%s' (%s)\n", envName, value, oneFlag.Value.Type())
		if !ok {
			// no environment variable set, so nothing to do
			return
		}
		if value == "" && oneFlag.Value.Type() == "bool" {
			// empty string for a bool, just use the default (i.e. not "true")
			return
		}
		err := f.Set(flagName, value)
		if err != nil {
			panic(err)
		}
	})
}

var (
	help    bool
	Verbose bool

	// repo flags
	owners []string
	Empty  bool

	// workflow flags
	Inactive   bool
	includes   []string
	IncludeSet map[string]struct{}
	excludes   []string
	ExcludeSet map[string]struct{}

	// external badge flags
	externals []string

	// output flags
	Output string
)

func usage(f *flag.FlagSet) {
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

func main() {

	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.BoolVar(&help, "help", false, "Show help")
	f.MarkHidden("help")
	f.BoolVar(&Verbose, "verbose", false, "Verbose messages")

	f.StringSliceVar(&owners, "owners", []string{}, "Owners")
	f.BoolVar(&Empty, "empty", false, "Include repos with no eligible workflows")

	f.BoolVar(&Inactive, "inactive", false, "Include inactive workflows")
	f.StringSliceVar(&includes, "include", []string{}, "Actions to include")
	f.StringSliceVar(&excludes, "exclude", []string{"codeql", "pages-build-deployment"}, "Actions to exclude")

	f.StringSliceVar(&externals, "external", []string{}, "External badges (see below)")

	f.StringVar(&Output, "output", "markdown", "Output format [ markdown | json | csv ]")

	setFlagsFromEnvironment(f)

	f.Parse(os.Args[1:])
	if help {
		usage(f)
		os.Exit(0)
	}

	if len(f.Args()) > 1 {
		fmt.Printf("ERROR: Only one output file please!\n")
		usage(f)
		os.Exit(1)
	}

	if len(owners) == 0 {
		fmt.Printf("ERROR: At least one owner please!\n")
		usage(f)
		os.Exit(1)
	}

	IncludeSet = make(map[string]struct{})
	for _, include := range includes {
		IncludeSet[strings.ToLower(include)] = struct{}{}
	}
	ExcludeSet = make(map[string]struct{})
	for _, exclude := range excludes {
		ExcludeSet[strings.ToLower(exclude)] = struct{}{}
	}

	client := github.NewClient(nil)
	token, tokenOk := os.LookupEnv("GHASHBOARD_TOKEN")
	if tokenOk {
		fmt.Printf("INFO: Loading GHASHBOARD_TOKEN\n")
		client = client.WithAuthToken(token)
	} else {
		fmt.Printf("WARNING: No GHASHBOARD_TOKEN not set: will make anonymous (and rate-limited) Github API calls\n")
	}

	fmt.Printf("INFO: owners = %v\n", owners)

	var allRepos []*github.Repository
	for _, owner := range owners {
		fmt.Printf("INFO: loading repos for %s...\n", owner)
		repos, err := GetReposForOwner(client, owner)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("INFO repos for %s: %d\n", owner, len(repos))
		allRepos = append(allRepos, repos...)
	}
	fmt.Printf("INFO: total repos: %d\n", len(allRepos))

	var allData []*MetaRepo
	for _, repo := range allRepos {
		fmt.Printf("INFO: loading workflows for %s...\n", *repo.FullName)
		workflows, err := GetWorkflowsForRepo(client, repo)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("INFO: workflows for %s: %d\n", *repo.FullName, len(workflows))
		if Empty || len(workflows) > 0 {
			SortWorkflowsCaseInsensitive(workflows)
			allData = append(allData, &MetaRepo{Workflows: workflows, Repo: repo})
		}
	}
	fmt.Printf("INFO: repos with workflows: %d\n", len(allData))
	SortReposCaseInsensitive(allData)

	if len(externals) > 0 {
		fmt.Printf("INFO: loading external badges...\n")
		for _, metaRepo := range allData {
			for _, external := range externals {
				xb, xbErr := GenerateExternalBadge(external, metaRepo.Repo)
				if xbErr != nil {
					fmt.Printf("ERROR: unable to expand external badge %s for %s: %v\n", external, *metaRepo.Repo.Name, xbErr)
					os.Exit(1)
				}
				metaRepo.ExternalBadges = append(metaRepo.ExternalBadges, xb)
			}
		}
	}

	filename := f.Arg(0)
	var writer io.Writer
	if filename == "" || filename == "-" {
		fmt.Printf("INFO: writing to stdout\n")
		writer = os.Stdout
	} else {
		fmt.Printf("INFO: writing to %s\n", filename)
		file, openErr := os.Create(filename)
		if openErr != nil {
			fmt.Printf("Error: %v\n", openErr)
			os.Exit(1)
		}
		defer file.Close()
		writer = file
	}

	if Output == "json" {
		jsonStr, jsonErr := json.Marshal(allData)
		if jsonErr != nil {
			fmt.Printf("Error: %v\n", jsonErr)
			os.Exit(1)
		}
		writer.Write(jsonStr)
	} else {
		tmpl, tmplErr := GetStandardTemplate(Output)
		if tmplErr != nil {
			fmt.Printf("Error: %v\n", tmplErr)
			os.Exit(1)
		}
		mergeErr := tmpl.Execute(writer, allData)
		if mergeErr != nil {
			fmt.Printf("Error: %v\n", mergeErr)
			os.Exit(1)
		}
	}
}
