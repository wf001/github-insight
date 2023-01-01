package tests

import (
	"testing"

	"github.com/google/go-github/v48/github"
	"github.com/wf001/github-stat/parse"
)

func TestNewCommits(t *testing.T) {
	var fakeData []*github.RepositoryCommit
	jsonString := ReadFile("../tests/assets/test_commits.json")
	DecodeJson(jsonString, &fakeData)
	cm, _ := parse.NewCommits(fakeData)
	AssertEq(cm[0].Name, "TestUser", t, "")
	AssertEq(cm[0].Email, "testuser@gmail.com", t, "")
	AssertEq(cm[0].Date, "2022-05-09 06:01:59 +0000 UTC", t, "")
	AssertEq(cm[1].Name, "TestUser", t, "")
	AssertEq(cm[1].Email, "testuser@gmail.com", t, "")
	AssertEq(cm[1].Date, "2022-05-06 17:47:40 +0000 UTC", t, "")
}
func TestNewCommitsCorner(t *testing.T) {
	var fakeData []*github.RepositoryCommit
	cm, _ := parse.NewCommits(fakeData)
	AssertEq(len(cm), 0, t, "")
}
