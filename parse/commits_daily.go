/*
DailyCommit represents number of commits aggregated by day.
*/
package parse

import (
	"fmt"
	"sort"
)

type DailyCommit struct {
	Amount   int
	Datetime string
}

type DailyCommits []DailyCommit

func (ca DailyCommits) Len() int {
	return len(ca)
}
func (ca DailyCommits) Swap(i, j int) {
	ca[i], ca[j] = ca[j], ca[i]
}
func (ca DailyCommits) Less(i, j int) bool {
	return ca[i].Datetime < ca[i].Datetime
}
func (c DailyCommits) String() string {
	s := ""
	for _, v := range c {
		s += fmt.Sprintf("%+v\n", v)
	}
	return s
}
func NewDailyCommit(Amount int, Datetime string) DailyCommit {
	return DailyCommit{
		Amount:   Amount,
		Datetime: Datetime,
	}
}
func NewDailyCommits(cl Commits) (DailyCommits, []string) {
	dateDesc := map[string]int{}
	for _, ci := range cl {
		d := ci.Date[:10]
		if _, ok := dateDesc[d]; ok {
			dateDesc[d] += 1
		} else {
			dateDesc[d] = 1
		}
	}

	var keys []string
	for key := range dateDesc {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var ca DailyCommits
	for _, k := range keys {
		ca = append(ca, DailyCommit{Amount: dateDesc[k], Datetime: k})
	}
	sort.Sort(ca)
	return ca, keys

}
