package main

import (
	"context"
	"fmt"
	"strings"

	github "github.com/google/go-github/v58/github"
)

func GetRepos(client *github.Client, owners []string) ([]*github.Repository, error) {
	var allRepos []*github.Repository
	if Public {
		repos, err := getPublicRepos(client, owners)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
	}
	if Private {
		repos, err := getPrivateRepos(client, owners)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
	}

	var filteredRepos []*github.Repository
	for _, repo := range allRepos {
		if !Forks && *repo.Fork {
			fmt.Printf("DEBUG: skipping forked repo %s\n", *repo.FullName)
			continue
		}
		if !Archived && *repo.Archived {
			fmt.Printf("DEBUG: skipping archived repo %s\n", *repo.FullName)
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
		allRepos = append(allRepos, repos...)

		if resp.NextPage == 0 {
			break
		}
		fmt.Printf("INFO: Paging for public repos from %s...\n", owner)
		opt.Page = resp.NextPage
	}
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
			if _, ok := ownerSet[strings.ToLower(*repo.Owner.Login)]; ok == false {
				fmt.Printf("DEBUG: skipping private repo %s\n", *repo.FullName)
				continue
			}
			allRepos = append(allRepos, repo)
		}

		if resp.NextPage == 0 {
			break
		}
		fmt.Printf("INFO: Paging through private repos...\n")
		opt.Page = resp.NextPage
	}
	return allRepos, nil
}
