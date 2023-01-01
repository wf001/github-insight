package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/wf001/github-stat/parse"
	"github.com/wf001/github-stat/render"
	"github.com/wf001/github-stat/util"
)

func main() {
	/*EntryPoint of the CLI*/

	app := &cli.App{
		Name:  "ghi",
		Usage: "Github insight generator",
		Commands: []*cli.Command{
			{
				Name:      "gen",
				Usage:     "generate insight file",
				UsageText: "gen [USERNAME] [REPOSITORY]",
				Action: func(cCtx *cli.Context) error {
					if len(os.Args) < 4 {
						util.Error("Please input [user name] and [repository name]")
					}
					user := os.Args[2]
					repo := os.Args[3]
					_main(user, repo)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		util.Error(fmt.Sprintf("%s", err))
	}
}
func _main(user string, repo string) {
	insight := parse.Insight{}
	parse.GetRepository(user, repo, &insight)
	render.Render(insight)
}
