/* FirstIssue indicate issues that has 'good first issue' 'first-timers-only' 'beginner' 'easy' or 'dificulty:easy' label and not closed */
package parse

import (
	"github.com/google/go-github/v48/github"
	"github.com/wf001/github-insight/types"
	"github.com/wf001/github-insight/util"
)

func NewFirstIssues(repoIss []*github.Issue) (Issues, int) {
	var r []*github.Issue
	var label = []string{"good first issue", "first-timers-only", "beginner", "easy", "difficulty:easy"}
	// extract labeled issue
	for _, v := range repoIss {
		for _, l := range v.Labels {
			if util.Contains(*l.Name, label) {
				r = append(r, v)
				break
			}
		}
	}
	var issues []Issue
	// create Issues instance
	for i := 0; i < util.Min(len(r), types.FIRST_ISSUES_SIZE); i++ {
		i, p := newIssue(r[i])
		if !p {
			issues = append(issues, i)
		}
	}
	return issues, len(issues)
}
