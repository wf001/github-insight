/*IssueStack represent cumulative sum of issue */
package parse

import (
	"github.com/google/go-github/v48/github"
	"github.com/wf001/github-stat/util"
)

type IssueStack map[string]int

func NewIssueStack(repoIss []*github.Issue) IssueStack {
	var issues IssueStack = map[string]int{}
	// memorize daily issue activity (open = +1 / close = -1)
	for i := 0; i < len(repoIss); i++ {
		allIss, isPR := newIssue(repoIss[i])
		if !isPR {
			created := util.ParseToDate(allIss.CreatedAt)
			if _, ok := issues[created]; ok {
				issues[created] += 1
			} else {
				issues[created] = 1
			}
			if allIss.ClosedAt != "" {
				closed := util.ParseToDate(allIss.ClosedAt)
				if _, ok := issues[closed]; ok {
					issues[closed] -= 1
				} else {
					issues[closed] = -1
				}
			}
		}
	}
	// accumulate issue
	keys := util.GetSortedKeyByKey(issues)
	var issStack map[string]int = map[string]int{}
	if len(issues) > 0 {
		issStack[keys[0]] = issues[keys[0]]
		sz := len(keys)
		for i := 1; i < sz; i++ {
			issStack[keys[i]] += issues[keys[i]] + issStack[keys[i-1]]
		}
	}
	return issStack
}
