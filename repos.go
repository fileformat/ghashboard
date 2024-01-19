package main

import (
	"context"
	"fmt"

	github "github.com/google/go-github/v58/github"
)

func GetReposForOwner(client *github.Client, owner string) ([]*github.Repository, error) {

	listOptions := github.ListOptions{PerPage: 100}
	opt := &github.RepositoryListByUserOptions{Type: "owner", Sort: "updated", Direction: "desc", ListOptions: listOptions}
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
		fmt.Printf("Paging...\n")
		opt.Page = resp.NextPage
	}
	return allRepos, nil
}
