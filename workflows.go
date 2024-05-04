package main

import (
	"context"
	"log/slog"
	"strings"

	github "github.com/google/go-github/v58/github"
	"github.com/spf13/viper"
)

func GetWorkflowsForRepo(client *github.Client, repo *github.Repository) ([]*github.Workflow, error) {

	opt := &github.ListOptions{PerPage: 50}
	ctx := context.Background()
	includeInactive := viper.GetBool("inactive")

	var allWorkflows []*github.Workflow
	for {
		workflows, resp, err := client.Actions.ListWorkflows(ctx, *repo.Owner.Login, *repo.Name, opt)
		if err != nil {
			return nil, err
		}
		for _, workflow := range workflows.Workflows {
			if !includeInactive && *workflow.State != "active" {
				continue
			}
			if len(IncludeSet) > 0 {
				_, ok := IncludeSet[strings.ToLower(*workflow.Name)]
				if !ok {
					continue
				}
			} else if len(ExcludeSet) > 0 {
				_, ok := ExcludeSet[strings.ToLower(*workflow.Name)]
				if ok {
					continue
				}
			}
			allWorkflows = append(allWorkflows, workflow)
		}
		if resp.NextPage == 0 {
			break
		}
		slog.Debug("paging for workflows", "repo", *repo.FullName, "page", resp.NextPage)
		opt.Page = resp.NextPage
	}
	return allWorkflows, nil
}
