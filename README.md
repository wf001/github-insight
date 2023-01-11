![](https://drive.google.com/uc?export=view&id=1D9yvxyLG-Q4YFg7R4Ymgt1q8rP5KTnut)

# github-insight
github-insight is a command-line tool for creating insight page of GitHub repository.

## Features
- Generate GitHub insight page of specified repository.
- The insight include below data, 
	- Basic information
	- Pie chart of distribution of languages
	- Bar chart of daily commit.
	- Latest commits
	- Pie chart of distribution of committers
	- Latest issues
	- Issues of beginners
	- Stack bar chart of issues.

## Getting started
``` sh
git clone https://github.com/wf001/github-insight
```

and create [Github Personal Access token](https://github.com/settings/tokens), run following command.
``` sh
export GIT_KEY="<github personal access token>"
```

``` sh
make build
```

``` sh
./bin/ghi gen <username>/<repository>
```

Bob's your uncle! Finally open `github-insight/generated/output.html` with your browser.

Example is [here](https://wf001.github.io/github-insight/)

```
$ ./bin/ghi -h
NAME:
   ghi - Github insight page generator

USAGE:
   ghi [global options] command [command options] [arguments...]

VERSION:
   0.1.1

COMMANDS:
   gen      generate insight page
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

```

Note: GitHub's personal access token has an API call limit within a unit period.
The maximum fetching amount per calling is 100. Thus, for example, running this CLI against a repository with 500 commits and 1000 issues will consume 15, and 4985 API call is available in the remaining hour.

see also [official doc](https://docs.github.com/en/rest/rate-limit?apiVersion=2022-11-28)(this CLI uses only core API)

## Thanks
- [go-echarts/go-echarts](https://github.com/go-echarts/go-echarts)
- [google/go-github](https://github.com/google/go-github)
- [pterm/pterm](https://github.com/pterm/pterm)

## License
Licensed under the MIT License.
