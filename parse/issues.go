/*Issue represents a base issue, which is used by
output html file directly and indirectly */
package parse

import (
	"fmt"

	"github.com/google/go-github/v48/github"
	"github.com/wf001/github-insight/types"
	"github.com/wf001/github-insight/util"
)

//TODO: make namespace match the Commit one
type Issue struct {
	Title     string
	URL       string
	User      string
	UserURL   string
	AvatarURL string
	Label     map[string]string
	State     string
	CreatedAt string
	ClosedAt  string
	Body      string
}

type Issues []Issue

func NewOpenIssues(repoIss []*github.Issue) (Issues, int) {
	var issues Issues
	for i := 0; i < len(repoIss); i++ {
		i, p := newIssue(repoIss[i])
		if !p {
			issues = append(issues, i)
		}
	}
	return issues, len(issues)
}

func newIssue(i *github.Issue) (Issue, bool) {
	var ii Issue
	var isPR bool = true
	if i.PullRequestLinks == nil {
		isPR = false
		ii = Issue{
			Title:     *i.Title,
			URL:       *i.URL,
			User:      *i.User.Login,
			UserURL:   *i.User.HTMLURL,
			AvatarURL: *i.User.AvatarURL,
			State:     *i.State,
			CreatedAt: fmt.Sprintf("%s", *i.CreatedAt),
			ClosedAt:  "",
			Body:      "",
		}
		// FIXME
		if ii.AvatarURL == "" {
			ii.AvatarURL = types.DEFAULT_IMG_URL
		}
		if i.Body != nil {
			ii.Body = (*i.Body)[:util.Min(len(*i.Body), types.ISSUES_BODY_LEN)]
			if len(*i.Body) > types.ISSUES_BODY_LEN {
				ii.Body += "..."
			}
		} else {
			ii.Body = "no any description"
		}
		if i.ClosedAt != nil {
			ii.ClosedAt = fmt.Sprintf("%s", *i.ClosedAt)
		}
		ii.Label = map[string]string{}
		for _, l := range i.Labels {
			ii.Label[*l.Name] = *l.Color
		}
	}
	return ii, isPR
}
