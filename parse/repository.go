package parse

import (
	"fmt"
	"sort"

	"github.com/google/go-github/v48/github"
)

type Insight struct {
	User            string
	RepoName        string
	Author          []string
	Info            Info
	RateLimits      *github.RateLimits
	GeneratedAt     string
	CommitCount     int
	IssueCount      int
	Languages       Languages
	Committers      Committers
	DailyCommits    DailyCommits
	LatestCommit    Commits
	LatestIssues    Issues
	FirstIssues     Issues
	FirstIssueCount int
	IssueStack      IssueStack
	Readme          string
	Contributing    string
}

type Languages []Language
type Language struct {
	Name   string
	Amount int
}

func (c Languages) Len() int {
	return len(c)
}
func (c Languages) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c Languages) Less(i, j int) bool {
	return c[i].Amount > c[j].Amount
}

func NewLanguages(m map[string]int) Languages {
	var ls Languages
	for k, v := range m {
		var l Language = Language{}
		l.Name = k
		l.Amount = v
		ls = append(ls, l)
	}
	sort.Sort(ls)
	return ls
}

type Info struct {
	Name        string
	AvatarURL   string
	Language    string
	Description string
	HomePage    string
	License     string
	LicenseURL  string
	Stars       int
	Forks       int
	IsFork      bool
	Size        int
	Issues      int
	CreatedAt   string
	UpdatedAt   string
}

func NewInfo(r *github.Repository) Info {
	var i Info
	if r != nil {
		i = Info{
			Name:        *r.Owner.Login,
			AvatarURL:   *r.Owner.AvatarURL,
			Language:    *r.Language,
			Description: "",
			HomePage:    "",
			License:     "",
			LicenseURL:  "",
			Stars:       *r.StargazersCount,
			Forks:       *r.ForksCount,
			IsFork:      *r.Fork,
			Size:        *r.Size,
			Issues:      *r.OpenIssuesCount,
			CreatedAt:   fmt.Sprintf("%s", *r.CreatedAt),
			UpdatedAt:   fmt.Sprintf("%s", *r.PushedAt),
		}
		if r.Description != nil {
			i.Description = *r.Description
		}
		if r.Homepage != nil {
			i.HomePage = *r.Homepage
		}
		if r.License != nil {
			i.License = *r.License.SPDXID
			i.LicenseURL = *r.License.URL
		}
	}
	return i
}
