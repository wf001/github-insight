package tests

import (
	"testing"

	"github.com/wf001/github-stat/parse"
)

func TestNewDailyCommits(t *testing.T) {
	cl := InitTestCaseCommits()
	x, y := parse.NewDailyCommits(cl)
	var want parse.DailyCommits = []parse.DailyCommit{
		parse.NewDailyCommit(3, "2022-01-01"),
		parse.NewDailyCommit(2, "2022-01-02"),
		parse.NewDailyCommit(1, "2022-01-04"),
	}
	for i := 0; i < 3; i++ {
		AssertEq(x[i], want[i], t, "")
		AssertEq(y[i], want[i].Datetime, t, "")
	}
}
func TestNewDailyCommitsCorner(t *testing.T) {
	cl := InitTestCaseCommitsCorner()
	x, y := parse.NewDailyCommits(cl)
	AssertEq(x[0].Amount, 1, t, "")
	AssertEq(y[0], "2022-01-01", t, "")
}
func TestNewDailyCommitsSort(t *testing.T) {
	cl := InitTestCaseCommits()
	x, _ := parse.NewDailyCommits(cl)
	var expect []string = []string{
		"2022-01-01",
		"2022-01-02",
		"2022-01-04",
	}
	for i, xi := range x {
		AssertEq(xi.Datetime, expect[i], t, "")
	}
}
