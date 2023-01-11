/*Commit represents a base commit, which is used by
output html file directly and indirectly */
package parse

import (
	"fmt"
	"strings"

	"github.com/google/go-github/v48/github"
	"github.com/wf001/github-insight/types"
)

type Commit struct {
	Name      string
	Email     string
	Date      string
	CommitID  string
	Message   string
	URL       string
	AuthorURL string
	AvatarURL string
}

type Commits []Commit

func NewCommit(Name string, Email string, Date string, CommitID string, Message string, URL string, AuthorURL string, AvatarURL string) Commit {
	// TODO: format the argument of this to match that of newIssue
	c := Commit{
		Name:      Name,
		Email:     Email,
		Date:      Date,
		CommitID:  CommitID,
		Message:   Message,
		URL:       URL,
		AuthorURL: AuthorURL,
		AvatarURL: AvatarURL,
	}
	if c.AvatarURL == "" {
		c.AvatarURL = types.DEFAULT_IMG_URL
	}
	return c
}

func (c Commits) String() string {
	s := ""
	for _, v := range c {
		s += fmt.Sprintf("%+v\n", v)
	}
	return s
}
func NewCommits(rc []*github.RepositoryCommit) (Commits, int) {
	var commits Commits
	for _, k := range rc {
		c := NewCommit(
			*k.Commit.Author.Name,
			*k.Commit.Author.Email,
			fmt.Sprintf("%s", *k.Commit.Author.Date),
			*k.SHA,
			strings.Split(*k.Commit.Message, "\n")[0],
			*k.HTMLURL,
			"",
			"",
		)
		// FIXME
		if k.Author != nil {
			c.AuthorURL = *k.Author.URL
			c.AvatarURL = *k.Author.AvatarURL
		}
		commits = append(commits, c)
	}
	return commits, len(commits)
}
