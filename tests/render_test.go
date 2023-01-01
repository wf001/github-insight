package tests

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/wf001/github-stat/parse"
	"github.com/wf001/github-stat/render"
)

func TestRender(t *testing.T) {
	s := &parse.Service{
		Client: &MockAPI{},
	}
	r := &parse.Insight{}
	api := s.Get("r", "u")
	parse.Parse(api, r)
	render.Render(*r)
	b := ReadFile(render.OUTPUT_FILE)
	AssertGt(len(b), 0, t, "html has no element")
	AssertTrue(strings.Contains(b, "<div class=\"item\""), t, "contents  must contains goecharts")
	AssertTrue(strings.Contains(b, "4900"), t, "contents  must contains goecharts")
	AssertTrue(strings.Contains(b, "let goecharts_"), t, "contents size must contains goecharts")
	AssertTrue(strings.Contains(b, "let goecharts_"), t, "contents size must contains goecharts")

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(b))

	table := doc.Find("tbody>tr>td").Text()
	AssertTrue(strings.Contains(table, "TestUser"), t, "contents must contains table")
}
