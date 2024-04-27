package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"strings"

	github "github.com/google/go-github/v58/github"
	"github.com/spf13/viper"
)

type MetaRepo struct {
	Workflows      []*github.Workflow `json:"workflows"`
	Repo           *github.Repository `json:"repo"`
	ExternalBadges []*ExternalBadge   `json:"external"`
}

var (

	// workflow flags
	IncludeSet map[string]struct{}
	ExcludeSet map[string]struct{}

	// external badge flags
	Externals []string
)

func main() {

	initConfig(os.Args[1:])
	initLogger()

	repos := viper_GetStringSlice("repos")
	owners := viper_GetStringSlice("owners")

	includes := viper_GetStringSlice("include")
	IncludeSet = make(map[string]struct{})
	for _, include := range includes {
		IncludeSet[strings.ToLower(include)] = struct{}{}
	}

	excludes := viper_GetStringSlice("exclude")
	ExcludeSet = make(map[string]struct{})
	for _, exclude := range excludes {
		ExcludeSet[strings.ToLower(exclude)] = struct{}{}
	}

	client := github.NewClient(nil)
	token := viper.GetString("github-token")
	if token != "" {
		slog.Info("using a token to make authorized requests")
		client = client.WithAuthToken(token)
	} else {
		slog.Warn("token not set: will make anonymous (and rate-limited) Github API calls")
	}

	slog.Debug("owners", "owners", owners, "repos", repos)

	var allRepos []*github.Repository

	if len(repos) > 0 {
		for _, repo := range repos {
			theRepo, err := GetRepo(client, repo)
			if err != nil {
				slog.Error("unable to get repo", "error", err, "repo", repo)
				os.Exit(1)
			}
			allRepos = append(allRepos, theRepo)
		}
	}

	if len(owners) > 0 {
		ownerRepos, err := GetRepos(client, owners)
		if err != nil {
			slog.Error("unable to get repos for owners", "error", err, "owners", owners)
			os.Exit(1)
		}
		allRepos = append(allRepos, ownerRepos...)
	}
	if len(allRepos) == 0 {
		slog.Error("no repos found", "repos", repos, "owners", owners)
		os.Exit(1)
	}

	slog.Debug("repos found", "count", len(allRepos))

	var allData []*MetaRepo
	empty := viper.GetBool("empty")
	for _, repo := range allRepos {
		slog.Info("loading workflows", "repo", *repo.FullName)
		workflows, err := GetWorkflowsForRepo(client, repo)
		if err != nil {
			slog.Error("unable to load workflows", "error", err, "repo", *repo.FullName)
			os.Exit(1)
		}
		slog.Info("workflows found", "repo", *repo.FullName, "count", len(workflows))
		if empty || len(workflows) > 0 {
			SortWorkflowsCaseInsensitive(workflows)
			allData = append(allData, &MetaRepo{Workflows: workflows, Repo: repo})
		}
	}
	slog.Info("repos with workflows", "count", len(allData))
	//SortReposCaseInsensitive(allData)

	externals := viper_GetStringSlice("externals")
	if len(externals) > 0 {
		slog.Info("loading external badges", "count", len(externals))
		for _, metaRepo := range allData {
			for _, external := range externals {
				xb, xbErr := GenerateExternalBadge(external, metaRepo.Repo)
				if xbErr != nil {
					slog.Error("unable to expand external badge", "error", xbErr, "external", external, "repo", *metaRepo.Repo.Name)
					os.Exit(1)
				}
				metaRepo.ExternalBadges = append(metaRepo.ExternalBadges, xb)
			}
		}
	}

	filename := viper.GetString("file")
	var writer io.Writer
	if filename == "" || filename == "-" {
		slog.Info("writing to stdout")
		writer = os.Stdout
	} else {
		slog.Info("writing to file", "filename", filename)
		file, openErr := os.Create(filename)
		if openErr != nil {
			slog.Error("unable to open file", "error", openErr, "filename", filename)
			os.Exit(1)
		}
		defer file.Close()
		writer = file
	}

	format := viper.GetString("format")
	if format == "json" {
		jsonStr, jsonErr := json.Marshal(allData)
		if jsonErr != nil {
			slog.Error("unable to marshal json", "error", jsonErr)
			os.Exit(1)
		}
		writer.Write(jsonStr)
	} else {
		tmpl, tmplErr := GetStandardTemplate(format)
		if tmplErr != nil {
			slog.Error("unable to open template", "error", tmplErr, "format", format)
			os.Exit(1)
		}
		mergeErr := tmpl.Execute(writer, allData)
		if mergeErr != nil {
			slog.Error("unable to merge template", "error", mergeErr, "format", format)
			os.Exit(1)
		}
	}

	slog.Info("done")
}
