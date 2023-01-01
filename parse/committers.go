/*
Committer represents number of commits aggregated by Email.
*/
package parse

import (
	"fmt"
	"sort"

	"github.com/wf001/github-stat/types"
	"github.com/wf001/github-stat/util"
)

type Committer struct {
	Name      string
	Email     string
	Amount    int
	Row       int
	URL       string
	AvatarURL string
}

type Committers []Committer

func (c Committers) String() string {
	s := ""
	for _, v := range c {
		s += fmt.Sprintf("%+v\n", v)
	}
	return s
}

func (c Committers) Len() int {
	return len(c)
}
func (c Committers) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c Committers) Less(i, j int) bool {
	return c[i].Amount > c[j].Amount
}
func NewCommitter(Name string, Email string, Amount int, Row int, URL string, AvatarURL string) Committer {
	return Committer{
		Name:      Name,
		Email:     Email,
		Amount:    Amount,
		Row:       Row,
		URL:       URL,
		AvatarURL: AvatarURL,
	}
}

func NewCommitters(cl Commits) Committers {
	type author struct {
		name      string
		url       string
		avatarURL string
		count     int
	}

	// aggregate by Email
	authors := map[string]*author{}
	for _, ci := range cl {
		if _, ok := authors[ci.Email]; ok {
			authors[ci.Email].count += 1
		} else {
			authors[ci.Email] = &author{
				name:      ci.Name,
				url:       ci.URL,
				avatarURL: ci.AvatarURL,
				count:     1,
			}
			if ci.AvatarURL == "" {
				authors[ci.Email].avatarURL = types.DEFAULT_IMG_URL
			}
		}
	}
	// convert map-based value to struct
	var keys []string
	for k := range authors {
		keys = append(keys, k)
	}
	var origin Committers
	for _, k := range keys {
		origin = append(origin, NewCommitter(authors[k].name, k, authors[k].count, 0, authors[k].url, authors[k].avatarURL))
	}
	// sort by amount desc
	sort.Sort(origin)

	origin_sz := len(origin)
	sz := util.Min(types.COMMITTERS_SIZE, origin_sz)
	committers := origin[:sz]

	// sum all of remaining committers's commit amount
	remaining := 0
	if origin_sz > types.COMMITTERS_SIZE {
		for i := types.COMMITTERS_SIZE; i < origin_sz; i++ {
			remaining += origin[i].Amount
		}
		committers = append(committers, NewCommitter("the others", "", remaining, 0, "", ""))
	}

	// Fix to be 1-indexed list.
	for i := 0; i < sz; i++ {
		committers[i].Row = i + 1
	}
	return committers
}
