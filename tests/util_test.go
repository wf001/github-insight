package tests

import (
	"fmt"
	"testing"

	"github.com/wf001/github-insight/util"
)

func TestMin(t *testing.T) {
	AssertEq(util.Min(1, 3), 1, t, "")
	AssertEq(util.Min(0, 10), 0, t, "")
	AssertEq(util.Min(1, 1), 1, t, "")
}
func TestParseDecSep(t *testing.T) {
	AssertEq("1", util.ParseDecSep("1"), t, "")
	AssertEq("10", util.ParseDecSep("10"), t, "")
	AssertEq("100", util.ParseDecSep("100"), t, "")
	AssertEq("1,000", util.ParseDecSep("1000"), t, "")
	AssertEq("10,000", util.ParseDecSep("10000"), t, "")
	AssertEq("100,000", util.ParseDecSep("100000"), t, "")
	AssertEq("1,000,000", util.ParseDecSep("1000000"), t, "")
}

func TestContains(t *testing.T) {
	AssertEq(true, util.Contains("a", []string{"a", "b", "c"}), t, "")
	AssertEq(false, util.Contains("d", []string{"a", "b", "c"}), t, "")
	AssertEq(false, util.Contains("a", []string{}), t, "")
}
func TestGetSortedKeyByValue(t *testing.T) {
	var m map[string]int = map[string]int{}
	m["a"] = 7
	m["b"] = 2
	m["c"] = 5
	m["d"] = 9

	act := util.GetSortedKeyByValue(m)
	var expect []string = []string{"b", "c", "a", "d"}

	for i := range expect {
		AssertEq(expect[i], act[i], t, "")
	}
}
func TestGetSortedKeyByKey(t *testing.T) {
	var m map[string]int = map[string]int{}
	m["2022-01-02"] = 3
	m["2022-01-01"] = 1
	m["2022-02-01"] = 5
	m["2021-02-01"] = 9
	s := util.GetSortedKeyByKey(m)
	expect := [][]string{
		{"2021-02-01", "9"},
		{"2022-01-01", "1"},
		{"2022-01-02", "3"},
		{"2022-02-01", "5"},
	}
	for i, si := range s {
		AssertEq(fmt.Sprintf("%d", m[si]), expect[i][1], t, "")
	}
}
