package tests

import (
	"testing"

	"github.com/wf001/github-stat/parse"
)

func TestGetRepository(t *testing.T) {
	m := &parse.Service{
		Client: &MockAPI{},
	}
	r := &parse.Insight{}
	api := m.Get("o", "r")
	parse.Parse(api, r)
	var want parse.Committers = []parse.Committer{
		parse.NewCommitter("TestUser", "testuser@gmail.com", 2, 1, "e", "a"),
	}
	AssertEq(len(r.Committers), 1, t, "")
	for i := 0; i < 1; i++ {
		AssertEq(r.Committers[i], want[i], t, "")
	}
}

func TestGetRepositoryWithoutMock(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	s := &parse.Service{
		Client: &parse.GhAPI{},
	}
	r := &parse.Insight{}
	api := s.Get("wf001", "gh-insight-cli")
	parse.Parse(api, r)
	AssertEq(r.Info.Name, "wf001", t, "")
	AssertEq(r.Info.Language, "Go", t, "")
}
