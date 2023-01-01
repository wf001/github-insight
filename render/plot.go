/* create instance of go-echarts*/
package render

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/wf001/github-stat/parse"
	"github.com/wf001/github-stat/util"
)

type Charts struct {
	ChartActivitiesBar *charts.Bar
	ChartCommittersPie *charts.Pie
	ChartLanguagePie   *charts.Pie
	ChartIssueStack    *charts.Line
}

func plotDailyCommits(c parse.DailyCommits, y []string) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(
			opts.Initialization{
				Theme:  types.ThemeShine,
				Width:  "100%",
				Height: "100%",
			},
		),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
	)
	barItems := make([]opts.BarData, 0)
	for _, v := range c {
		barItems = append(barItems, opts.BarData{Value: v.Amount})
	}

	bar.SetXAxis(y).
		AddSeries("commits", barItems).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	return bar
}

func plotRepoLang(lang parse.Languages) *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithInitializationOpts(
			opts.Initialization{
				Theme: types.ThemeShine,
				Width: "90%",
			},
		),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
	)
	pieItems := make([]opts.PieData, 0)
	for _, v := range lang {
		pieItems = append(pieItems, opts.PieData{Name: v.Name, Value: v.Amount})
	}

	pie.AddSeries("Language", pieItems)
	return pie
}

func plotCommittersPie(c parse.Committers) *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithInitializationOpts(
			opts.Initialization{
				Theme: types.ThemeShine,
				Width: "100%",
			},
		),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
	)
	pieItems := make([]opts.PieData, 0)
	for _, v := range c {
		pieItems = append(pieItems, opts.PieData{Name: v.Name, Value: v.Amount})
	}

	pie.AddSeries("committer", pieItems)
	return pie
}
func plotIssueStack(is parse.IssueStack) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(
			opts.Initialization{
				Theme: types.ThemeShine,
				Width: "100%",
			},
		),
		charts.WithTooltipOpts(
			opts.Tooltip{Show: true},
		),
	)
	lineAxis := util.GetSortedKeyByKey(is)
	var lineValue []opts.LineData
	for _, v := range lineAxis {
		lineValue = append(lineValue, opts.LineData{Value: is[v]})
	}
	line.SetXAxis(lineAxis).AddSeries("Category", lineValue).
		SetSeriesOptions(
			charts.WithAreaStyleOpts(
				opts.AreaStyle{
					Opacity: 0.2,
				},
			),
			charts.WithLineChartOpts(
				opts.LineChart{
					Smooth: true,
				},
			),
		)
	return line
}
func MakeCharts(r parse.Insight) Charts {
	return Charts{
		ChartActivitiesBar: plotDailyCommits(r.DailyCommits, r.Author),
		ChartCommittersPie: plotCommittersPie(r.Committers),
		ChartLanguagePie:   plotRepoLang(r.Languages),
		ChartIssueStack:    plotIssueStack(r.IssueStack),
	}
}
