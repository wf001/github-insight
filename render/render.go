/* write out html*/
package render

import (
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	chartrender "github.com/go-echarts/go-echarts/v2/render"
	"github.com/wf001/github-stat/parse"
	"github.com/wf001/github-stat/util"
)

const (
	TEMPLATE_PATH           = "./assets/template/"
	TEMPLATE_FILE           = TEMPLATE_PATH + "base.html"
	OUTPUT_PATH             = "./generated/"
	OUTPUT_FILE             = OUTPUT_PATH + "output.html"
	TMP_DIR                 = OUTPUT_PATH + "tmp/"
	DAILY_COMMIT_FILE_HTML  = TMP_DIR + "daily_commit.html"
	DAILY_COMMIT_FILE_JS    = TMP_DIR + "daily_commit_js.html"
	COMMITTERS_FILE_HTML    = TMP_DIR + "committers.html"
	COMMITTERS_FILE_JS      = TMP_DIR + "committers_js.html"
	REPO_LANGUAGE_FILE_HTML = TMP_DIR + "repo_language.html"
	REPO_LANGUAGE_FILE_JS   = TMP_DIR + "repo_language_js.html"
	ISSUE_FILE_HTML         = TMP_DIR + "issue.html"
	ISSUE_FILE_JS           = TMP_DIR + "issue_js.html"
)

type GICRenderer struct {
	c      interface{}
	before []func()
	t      string
}

func NewGICRenderer(c interface{}, partial string, before ...func()) chartrender.Renderer {
	return &GICRenderer{c: c, before: before, t: partial}
}

func RenderIndex(r parse.Insight) {
	t := template.Must(
		template.ParseFiles(
			TEMPLATE_FILE,
			DAILY_COMMIT_FILE_HTML,
			DAILY_COMMIT_FILE_JS,
			COMMITTERS_FILE_HTML,
			COMMITTERS_FILE_JS,
			REPO_LANGUAGE_FILE_HTML,
			REPO_LANGUAGE_FILE_JS,
			ISSUE_FILE_HTML,
			ISSUE_FILE_JS,
		),
	)

	f, _ := os.Create(OUTPUT_FILE)
	if err := t.Execute(f, r); err != nil {
		util.Error(fmt.Sprintf("%s", err))
	}

}

func (r *GICRenderer) Render(w io.Writer) error {
	const tplName = ""
	for _, fn := range r.before {
		fn()
	}

	tpl := template.
		Must(template.New(tplName).
			Funcs(template.FuncMap{
				"safeJS": func(s interface{}) template.JS {
					return template.JS(fmt.Sprint(s))
				},
			}).
			Parse(r.t),
		)

	err := tpl.ExecuteTemplate(w, tplName, r.c)
	return err
}

// TODO: simplify
func RenderBar(bar *charts.Bar, part string, file string) {
	bar.Renderer = NewGICRenderer(bar, part, bar.Validate)
	if f, err := os.Create(file); err != nil {
		util.Error(fmt.Sprintf("%s", err))
	} else {
		bar.Render(f)
	}
}
func RenderPie(pie *charts.Pie, part string, file string) {
	pie.Renderer = NewGICRenderer(pie, part, pie.Validate)
	if f, err := os.Create(file); err != nil {
		util.Error(fmt.Sprintf("%s", err))
	} else {
		pie.Render(f)
	}
}
func RenderLine(line *charts.Line, part string, file string) {
	line.Renderer = NewGICRenderer(line, part, line.Validate)
	if f, err := os.Create(file); err != nil {
		util.Error(fmt.Sprintf("%s", err))
	} else {
		line.Render(f)
	}
}
func Render(r parse.Insight) {
	if err := os.MkdirAll(TMP_DIR, 0644); err != nil {
		util.Error(fmt.Sprintf("%s", err))
	}
	r.Info.CreatedAt = util.ParseToDate(r.Info.CreatedAt)
	r.Info.UpdatedAt = util.ParseToDate(r.Info.UpdatedAt)
	c := MakeCharts(r)
	RenderBar(c.ChartActivitiesBar, EchartsCommitActivityHTML, DAILY_COMMIT_FILE_HTML)
	RenderBar(c.ChartActivitiesBar, EchartsCommonJS, DAILY_COMMIT_FILE_JS)

	RenderPie(c.ChartLanguagePie, EchartLangHTML, REPO_LANGUAGE_FILE_HTML)
	RenderPie(c.ChartLanguagePie, EchartsCommonJS, REPO_LANGUAGE_FILE_JS)

	RenderPie(c.ChartCommittersPie, EchartCommittersHTML, COMMITTERS_FILE_HTML)
	RenderPie(c.ChartCommittersPie, EchartsCommonJS, COMMITTERS_FILE_JS)

	RenderLine(c.ChartIssueStack, EchartIssueHTML, ISSUE_FILE_HTML)
	RenderLine(c.ChartIssueStack, EchartsCommonJS, ISSUE_FILE_JS)
	// RenderIndex embed the template which is created above.
	RenderIndex(r)
	if err := os.RemoveAll(TMP_DIR); err != nil {
		util.Error(fmt.Sprintf("%s", err))
	}
	fmt.Printf("\033[032mResult outputed on %s.\033[0m\n", OUTPUT_FILE)
}
