package tests

import (
	"testing"

	"github.com/google/go-github/v48/github"
	"github.com/wf001/github-insight/parse"
)

func TestNewLanguages(t *testing.T) {
	var m map[string]int = map[string]int{}
	m["a"] = 4
	m["b"] = 2
	m["c"] = 7

	l := parse.NewLanguages(m)
	var expect []int = []int{7, 4, 2}
	for i := range l {
		AssertEq(expect[i], l[i].Amount, t, "")
	}
}
func TestNewInfo(t *testing.T) {
	s := &parse.Service{
		Client: &MockAPI{},
	}
	r := s.Client.FetchRepository("o", "r")
	info := parse.NewInfo(r)
	AssertEq(info.Name, "fake", t, "")

}
func TestNewInfoCorner(t *testing.T) {
	var r *github.Repository
	info := parse.NewInfo(r)
	AssertEq(info, parse.Info{}, t, "")
}
