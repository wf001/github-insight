package tests

import (
	"testing"

	"github.com/google/go-github/v48/github"
	"github.com/wf001/github-insight/parse"
)

func TestNewIssuesStack(t *testing.T) {
	var fakeData []*github.Issue
	jsonString := ReadFile("../tests/assets/test_issues.json")
	DecodeJson(jsonString, &fakeData)
	iStack := parse.NewIssueStack(fakeData)
	var expect map[string]int = map[string]int{}
	expect["2022-12-08"] = 1
	expect["2022-12-09"] = 2
	expect["2022-12-15"] = 2
	expect["2022-12-16"] = 1
	expect["2022-12-19"] = 0

	for k, v := range expect {
		AssertEq(iStack[k], v, t, "")
	}
}
func TestNewIssuesStackCorner(t *testing.T) {
	var fakeData []*github.Issue
	iStack := parse.NewIssueStack(fakeData)
	AssertEq(len(iStack), 0, t, "")
}
