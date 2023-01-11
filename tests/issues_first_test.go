package tests

import (
	"fmt"
	"testing"

	"github.com/google/go-github/v48/github"
	"github.com/wf001/github-insight/parse"
)

func TestNewFirstIssues(t *testing.T) {
	var fakeData []*github.Issue
	jsonString := ReadFile("../tests/assets/test_issues.json")
	DecodeJson(jsonString, &fakeData)
	is, i := parse.NewFirstIssues(fakeData)
	var expect [][]string = [][]string{
		{"Issue1", "103", "2"},
		{"Issue2", "0", "1"},
	}
	AssertEq(len(is), 2, t, "")
	AssertEq(i, 2, t, "")
	for i, v := range expect {
		AssertEq(is[i].Title, v[0], t, "")
		AssertEq(fmt.Sprintf("%d", len(is[i].Body)), v[1], t, "")
		AssertEq(is[i].State, "open", t, "")
		AssertEq(fmt.Sprintf("%d", len(is[i].Label)), v[2], t, "")
	}
}
func TestNewFirstIssuesCorner(t *testing.T) {
	var fakeData []*github.Issue
	is, i := parse.NewFirstIssues(fakeData)
	AssertEq(len(is), 0, t, "")
	AssertEq(i, 0, t, "")
}
