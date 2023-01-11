/*
Fetch data via Github API wrapper and convert it to GhAPI.
*/
package parse

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v48/github"
	"github.com/pterm/pterm"
	"github.com/wf001/github-insight/types"
	"github.com/wf001/github-insight/util"
	"golang.org/x/oauth2"
)

type Service struct {
	Client I
}
type I interface {
	FetchRepoCommits(string, string, *github.CommitsListOptions) []*github.RepositoryCommit
	FetchLanguages(string, string) map[string]int
	FetchRateLimit() *github.RateLimits
	FetchRepository(string, string) *github.Repository
	FetchIssues(string, string, *github.IssueListByRepoOptions) []*github.Issue
	FetchContent(string, string, string) *github.RepositoryContent
}
type GhAPI struct {
	RepoCommit   []*github.RepositoryCommit
	RepoLang     map[string]int
	RateLimit    *github.RateLimits
	RepoInfo     *github.Repository
	RepoAllIss   []*github.Issue
	RepoOpenIss  []*github.Issue
	Readme       *github.RepositoryContent
	Contributing *github.RepositoryContent
}

func GetRepository(user string, repo string, r *Insight) {
	fmt.Printf("\033[032mGet: https://github.com/%s/%s\033[0m\n", user, repo)
	sc := &Service{
		Client: &GhAPI{},
	}
	api := sc.Get(user, repo)
	Parse(api, r)
	r.User = user
	r.RepoName = repo
}
func (s *Service) Get(user string, repo string) GhAPI {
	/* Fetch entrypoint */
	api := GhAPI{}
	api.RepoCommit = s.Client.FetchRepoCommits(
		user,
		repo,
		&github.CommitsListOptions{
			ListOptions: github.ListOptions{PerPage: 100},
		},
	)
	api.RepoLang = s.Client.FetchLanguages(user, repo)
	api.RepoInfo = s.Client.FetchRepository(user, repo)
	api.RateLimit = s.Client.FetchRateLimit()

	api.RepoAllIss = s.Client.FetchIssues(
		user,
		repo,
		&github.IssueListByRepoOptions{
			ListOptions: github.ListOptions{PerPage: 100},
			State:       "all",
		},
	)
	api.RepoOpenIss = s.Client.FetchIssues(
		user,
		repo,
		&github.IssueListByRepoOptions{
			ListOptions: github.ListOptions{PerPage: 100},
			State:       "open",
		},
	)
	api.Readme = s.Client.FetchContent(user, repo, "README.md")
	api.Contributing = s.Client.FetchContent(user, repo, "CONTRIBUTING.md")
	return api
}
func Parse(api GhAPI, r *Insight) {
	/* Parse entrypoint */
	rc := api.RepoCommit
	rl := api.RepoLang
	rrl := api.RateLimit
	rr := api.RepoInfo
	repoAllIss := api.RepoAllIss
	repoOpenIss := api.RepoOpenIss
	cl, clen := NewCommits(rc)
	c, y := NewDailyCommits(cl)
	r.Committers = NewCommitters(cl)
	r.DailyCommits = c
	r.LatestCommit = cl[:util.Min(len(cl), types.COMMIT_HISTORY)]
	r.CommitCount = clen

	r.Author = y
	r.Languages = NewLanguages(rl)
	r.RateLimits = rrl
	r.Info = NewInfo(rr)

	openIssue, il := NewOpenIssues(repoOpenIss)
	r.LatestIssues = openIssue[:util.Min(len(openIssue), types.ISSUES_SIZE)]
	r.IssueCount = il
	r.IssueStack = NewIssueStack(repoAllIss)
	firstIss, firstIssCount := NewFirstIssues(repoOpenIss)
	r.FirstIssueCount = firstIssCount
	r.FirstIssues = firstIss[:util.Min(len(firstIss), types.FIRST_ISSUES_SIZE)]
	r.GeneratedAt = time.Now().Format(time.RFC3339)
	if api.Readme != nil {
		r.Readme = *api.Readme.HTMLURL
	} else {
		r.Readme = ""
	}
	if api.Contributing != nil {
		r.Contributing = *api.Contributing.HTMLURL
	} else {
		r.Contributing = ""
	}

}

func GetClient() *github.Client {
	key := os.Getenv("GIT_KEY")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: key},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return github.NewClient(tc)
}

func (c *GhAPI) FetchContent(user string, repo string, file string) *github.RepositoryContent {
	client := GetClient()
	content, _, _, _ := client.Repositories.GetContents(
		context.Background(),
		user,
		repo,
		file,
		&github.RepositoryContentGetOptions{},
	)
	return content
}

func (c *GhAPI) FetchRepoCommits(user string, repo string, opts *github.CommitsListOptions) []*github.RepositoryCommit {
	client := GetClient()
	var rc []*github.RepositoryCommit
	var p *pterm.ProgressbarPrinter
	for {
		commits, resp, _ := client.Repositories.ListCommits(
			context.Background(),
			user,
			repo,
			opts,
		)
		if p == nil {
			var lastPage int = util.Max(resp.LastPage, 1)
			p, _ = pterm.DefaultProgressbar.WithTotal(lastPage).WithTitle("Getting all commits").Start()
		}
		if resp.Response.StatusCode >= 400 {
			util.Error(fmt.Sprintf("%s", resp.Response.Body))
		}

		for _, k := range commits {
			rc = append(rc, k)
		}
		//progress update
		p.Increment()

		if resp.NextPage == 0 {
			pterm.Success.Println("Getting all commits")
			break
		}
		opts.Page = resp.NextPage

	}
	return rc
}
func (c *GhAPI) FetchLanguages(user string, repo string) map[string]int {
	client := GetClient()
	lang, _, _ := client.Repositories.ListLanguages(context.Background(), user, repo)
	return lang
}

func (c *GhAPI) FetchRateLimit() *github.RateLimits {
	client := GetClient()
	r, _, _ := client.RateLimits(context.Background())
	return r
}
func (c *GhAPI) FetchRepository(user string, repo string) *github.Repository {
	client := GetClient()
	r, _, _ := client.Repositories.Get(context.Background(), user, repo)
	return r
}

func (c *GhAPI) FetchIssues(user string, repo string, opts *github.IssueListByRepoOptions) []*github.Issue {
	client := GetClient()
	var ri []*github.Issue
	var p *pterm.ProgressbarPrinter
	// TODO: commonize with FetchRepoCommits?
	for {
		r, resp, _ := client.Issues.ListByRepo(
			context.Background(),
			user,
			repo,
			opts,
		)
		if p == nil {
			var lastPage int = util.Max(resp.LastPage, 1)
			/* for progress bar */
			p, _ = pterm.
				DefaultProgressbar.
				WithTotal(lastPage).
				WithTitle("Getting " + opts.State + " issues").
				Start()
		}
		if resp.Response.StatusCode >= 400 {
			util.Error(fmt.Sprintf("%s", resp.Response.Body))
		}

		for _, k := range r {
			ri = append(ri, k)
		}
		p.Increment()

		if resp.NextPage == 0 {
			pterm.Success.Println("Getting " + opts.State + " issues")
			break
		}
		opts.Page = resp.NextPage
	}
	return ri
}
