package tests

import (
	"testing"

	"github.com/wf001/github-stat/parse"
)

func TestNewCommitter(t *testing.T) {
	cl := InitTestCaseCommits()
	x := parse.NewCommitters(cl)
	var want parse.Committers = []parse.Committer{
		parse.NewCommitter("user1", "user1@example.com", 3, 1, "u1", "p1"),
		parse.NewCommitter("user2", "user2@example.com", 2, 2, "u1", "p1"),
		parse.NewCommitter("user3", "user3@example.com", 1, 3, "u1", "p1"),
	}
	for i := 0; i < 3; i++ {
		AssertEq(x[i], want[i], t, "")
	}
}
func TestNewCommitterCorner(t *testing.T) {
	cl := InitTestCaseCommitsCorner()
	x := parse.NewCommitters(cl)
	AssertEq(x[0].Name, "user1", t, "")
	AssertEq(x[0].Amount, 1, t, "")
	AssertEq(x[0].Row, 1, t, "")
}
