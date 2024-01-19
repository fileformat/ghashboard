package main

import (
	"sort"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"

	github "github.com/google/go-github/v58/github"
)

func SortStringsCaseInsensitive(data []string) {
	collator := collate.New(language.English)
	collator.SortStrings(data)
}

func SortWorkflowsCaseInsensitive(data []*github.Workflow) {
	collator := collate.New(language.English)
	sort.Slice(data, func(i, j int) bool { return collator.CompareString(*data[i].Name, *data[j].Name) < 0 })
}

type MetaRepoSlice struct {
	MetaRepos []*MetaRepo
}

func (mrs MetaRepoSlice) Len() int {
	return len(mrs.MetaRepos)
}

func (mrs MetaRepoSlice) Swap(i, j int) {
	temp := mrs.MetaRepos[i]
	mrs.MetaRepos[i] = mrs.MetaRepos[j]
	mrs.MetaRepos[j] = temp
}

func (mrs MetaRepoSlice) Bytes(i int) []byte {
	// returns the bytes of text at index i
	return []byte(*mrs.MetaRepos[i].Repo.FullName)
}

func SortReposCaseInsensitive(data []*MetaRepo) {

	mrs := MetaRepoSlice{MetaRepos: data}

	collator := collate.New(language.English)
	collator.Sort(mrs)
}
