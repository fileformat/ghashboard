package main

import (
	"context"
	"fmt"

	github "github.com/google/go-github/v58/github"
)

func GetWorkflowsForRepo(client *github.Client, repo *github.Repository) ([]*github.Workflow, error) {

	opt := &github.ListOptions{PerPage: 50}
	ctx := context.Background()

	var allWorkflows []*github.Workflow
	for {
		workflows, resp, err := client.Actions.ListWorkflows(ctx, *repo.Owner.Login, *repo.Name, opt)
		if err != nil {
			return nil, err
		}
		allWorkflows = append(allWorkflows, workflows.Workflows...)
		if resp.NextPage == 0 {
			break
		}
		fmt.Printf("Paging...\n")
		opt.Page = resp.NextPage
	}
	return allWorkflows, nil
}
