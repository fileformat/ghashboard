package main

import (
	"context"
	"log/slog"
	"strings"

	github "github.com/google/go-github/v58/github"
	"github.com/spf13/viper"
)

func GetRepo(client *github.Client, repo string) (*github.Repository, error) {
	ctx := context.Background()
	parts := strings.Split(repo, "/")
	theRepo, _, err := client.Repositories.Get(ctx, parts[0], parts[1])
	return theRepo, err
}

func GetRepos(client *github.Client, owners []string) ([]*github.Repository, error) {
	var allRepos []*github.Repository
	if viper.GetBool("public") {
		repos, err := getPublicRepos(client, owners)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
	}
	if viper.GetBool("private") {
		repos, err := getPrivateRepos(client, owners)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
	}

	forks := viper.GetBool("forks")
	archived := viper.GetBool("archived")
	var filteredRepos []*github.Repository
	for _, repo := range allRepos {
		if !forks && *repo.Fork {
			slog.Debug("skipping forked repo", "repo", *repo.FullName)
			continue
		}
		if !archived && *repo.Archived {
			slog.Debug("skipping archived repo", "repo", *repo.FullName)
			continue
		}
		filteredRepos = append(filteredRepos, repo)
	}

	return filteredRepos, nil
}

func getPublicRepos(client *github.Client, owners []string) ([]*github.Repository, error) {
	var allRepos []*github.Repository
	for _, owner := range owners {
		repos, err := getPublicReposForOwner(client, owner)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
	}
	slog.Debug("total public repos", "count", len(allRepos))
	return allRepos, nil
}

func getPublicReposForOwner(client *github.Client, owner string) ([]*github.Repository, error) {

	listOptions := github.ListOptions{PerPage: 100}
	opt := &github.RepositoryListByUserOptions{Type: "all", Sort: "updated", Direction: "desc", ListOptions: listOptions}
	ctx := context.Background()

	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByUser(ctx, owner, opt)
		if err != nil {
			return nil, err
		}
		for _, repo := range repos {
			// needed since getting a user's repos will return all repos, even if under a different owner.
			if !strings.EqualFold(*repo.Owner.Login, owner) {
				continue
			}
			// needed since getting a user's repos will include private repos
			if *repo.Private {
				continue
			}
			_, skip := SkipSet[strings.ToLower(*repo.FullName)]
			if skip {
				slog.Debug("explicitly skipping repo", "repo", *repo.FullName)
				continue
			}
			allRepos = append(allRepos, repo)
		}

		if resp.NextPage == 0 {
			break
		}
		slog.Debug("paging for public repos", "owner", owner, "page", resp.NextPage)
		opt.Page = resp.NextPage
	}
	slog.Debug("public repos", "owner", owner, "count", len(allRepos))

	return allRepos, nil
}

func getPrivateRepos(client *github.Client, owners []string) ([]*github.Repository, error) {

	listOptions := github.ListOptions{PerPage: 100}
	opt := &github.RepositoryListByAuthenticatedUserOptions{Visibility: "private", Sort: "updated", Direction: "desc", ListOptions: listOptions}
	ctx := context.Background()

	ownerSet := make(map[string]struct{})
	for _, owner := range owners {
		ownerSet[strings.ToLower(owner)] = struct{}{}
	}

	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByAuthenticatedUser(ctx, opt)
		if err != nil {
			return nil, err
		}
		for _, repo := range repos {
			_, ok := ownerSet[strings.ToLower(*repo.Owner.Login)]
			if !ok {
				// these are private repos you can see, but aren't in the list of owners
				slog.Debug("skipping private repo", "repo", *repo.FullName)
				continue
			}
			if !*repo.Private {
				slog.Warn("repo marked as public but returned from query", "repo", *repo.FullName)
				continue
			}
			_, skip := SkipSet[strings.ToLower(*repo.FullName)]
			if skip {
				slog.Debug("explicitly skipping repo", "repo", *repo.FullName)
				continue
			}
			allRepos = append(allRepos, repo)
		}

		if resp.NextPage == 0 {
			break
		}
		slog.Debug("paging through private repos...")
		opt.Page = resp.NextPage
	}
	slog.Debug("total private repos", "count", len(allRepos))
	return allRepos, nil
}
