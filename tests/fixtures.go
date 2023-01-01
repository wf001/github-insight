package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"testing"

	"github.com/google/go-github/v48/github"
	"github.com/wf001/github-stat/parse"
)

type MockAPI struct{}

func (c *MockAPI) FetchContent(user, repo, file string) *github.RepositoryContent {
	var fakeContent *github.RepositoryContent
	jsonString := ReadFile("../tests/assets/test_readme.json")
	DecodeJson(jsonString, &fakeContent)
	return fakeContent
}
func (c *MockAPI) FetchRepoCommits(user string, repo string, opts *github.CommitsListOptions) []*github.RepositoryCommit {
	var fakeCommits []*github.RepositoryCommit
	jsonString := ReadFile("../tests/assets/test_commits.json")
	DecodeJson(jsonString, &fakeCommits)
	return fakeCommits
}
func (c *MockAPI) FetchLanguages(user string, repo string) map[string]int {
	return map[string]int{
		"Python": 1000,
	}
}
func (c *MockAPI) FetchRateLimit() *github.RateLimits {
	return &github.RateLimits{
		Core: &github.Rate{
			Remaining: 4900,
			Limit:     5000,
		},
	}
}
func (c *MockAPI) FetchRepository(user string, repo string) *github.Repository {
	var fakeRepo *github.Repository
	jsonString := ReadFile("../tests/assets/test_repo.json")
	DecodeJson(jsonString, &fakeRepo)
	return fakeRepo
}
func (c *MockAPI) FetchIssues(user string, repo string, opts *github.IssueListByRepoOptions) []*github.Issue {
	var fakeIssue []*github.Issue
	jsonString := ReadFile("../tests/assets/test_issues.json")
	DecodeJson(jsonString, &fakeIssue)
	return fakeIssue
}

func InitTestCaseCommits() parse.Commits {
	cl := parse.Commits{}
	cl = append(cl, parse.NewCommit("user3", "user3@example.com", "2022-01-04 12:34:56+09:00", "s1", "m1", "u1", "a1", "p1"))
	cl = append(cl, parse.NewCommit("user2", "user2@example.com", "2022-01-02 02:34:56+09:00", "s1", "m1", "u1", "a1", "p1"))
	cl = append(cl, parse.NewCommit("user1", "user1@example.com", "2022-01-01 12:34:56+09:00", "s1", "m1", "u1", "a1", "p1"))
	cl = append(cl, parse.NewCommit("user1", "user1@example.com", "2022-01-01 14:34:56+09:00", "s1", "m1", "u1", "a1", "p1"))
	cl = append(cl, parse.NewCommit("user1", "user1@example.com", "2022-01-01 22:34:56+09:00", "s1", "m1", "u1", "a1", "p1"))
	cl = append(cl, parse.NewCommit("user2", "user2@example.com", "2022-01-02 12:34:56+09:00", "s1", "m1", "u1", "a1", "p1"))
	return cl
}
func InitTestCaseCommitsCorner() parse.Commits {
	cl := parse.Commits{}
	cl = append(cl, parse.NewCommit("user1", "user1@example.com", "2022-01-01 12:34:56+09:00", "s1", "m1", "u1", "a1", "p1"))
	return cl
}

func ReadFile(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
func DecodeJson(jsonString string, v interface{}) {
	if err := json.NewDecoder(strings.NewReader(jsonString)).Decode(v); err != nil {
		fmt.Println(err)
		return
	}
}
func AssertEq(act interface{}, expected interface{}, t *testing.T, errmsg string) {
	if act != expected {
		t.Errorf("\033[031m[FAIL]\033[0m [%v] got:\033[033m %v, \033[0mwant: \033[032m %v \033[0m", errmsg, act, expected)
	}
}
func AssertLt(act int, baseline int, t *testing.T, errmsg string) {
	if act >= baseline {
		t.Errorf("\033[031m[FAIL]\033[0m [%v] got:\033[033m %v, \033[0mmust be larger than: \033[032m %v \033[0m", errmsg, act, baseline)
	}
}
func AssertGt(act int, baseline int, t *testing.T, errmsg string) {
	if act <= baseline {
		t.Errorf("\033[031m[FAIL]\033[0m [%v] got:\033[033m %v, \033[0mmust be less than: \033[032m %v \033[0m", errmsg, act, baseline)
	}
}
func AssertTrue(condition bool, t *testing.T, errmsg string) {
	if !condition {
		t.Errorf("\033[031m[FAIL]\033[0m [%v] got:\033[033m %v, \033[0m want:\033[032m %v \033[0m", errmsg, condition, !condition)
	}
}

func Where() {
	pc, file, line, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\nfile:%s:%d:%s\n", file, line, strings.Split(f.Name(), ".")[2])
}
