package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/wf001/github-insight/parse"
	"github.com/wf001/github-insight/render"
	"github.com/wf001/github-insight/util"
)

func main() {
	/*EntryPoint of the CLI*/

	app := &cli.App{
		Name:    "ghi",
		Usage:   "Github insight page generator",
		Version: "0.1.1",
		Commands: []*cli.Command{
			{
				Name:      "gen",
				Usage:     "generate insight page",
				UsageText: "gen USERNAME/REPOSITORY",
				Action: func(cCtx *cli.Context) error {
					if len(os.Args) != 3 {
						util.Error("Please input [user name]/[repository name] format.")
					}
					userRepo := os.Args[2]
					user := strings.Split(userRepo, "/")[0]
					repo := strings.Split(userRepo, "/")[1]
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
