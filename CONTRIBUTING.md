

# Contributing to github-insight

Bug reports and code and documentation patches are welcome. 

## 1. Reporting bugs

It's important that you provide the **full command argument list**
as well as the output of the failing command.

```bash
$ ./bin/ghi <COMPLETE ARGUMENT LIST THAT TRIGGERS THE ERROR>

-- complete output here --
```

## 2. Contributing Code and Docs

Before working on a new feature or a bug, please browse [existing issues](https://github.com/wf001/github-insight/issues)
to see whether it has previously been discussed.

If you are fixing an issue, the first step should be to create a test case that
reproduces the incorrect behaviour. That will also help you to build an
understanding of the issue at hand.

### 2.1 Getting the code

Go to <https://github.com/wf001/github-insight> and fork the project repository.

```bash
# Clone your fork
git clone git@github.com:<YOU>/github-insight.git

# Enter the project directory
cd github-insight

# Create a branch for your changes
git checkout -b my_topical_branch
```

### 2.2. Setup

The [Makefile](https://github.com/wf001/github-insight/blob/main/Makefile) contains a bunch of tasks to get you started.

To get started, run the command below, which:


**install all of dependencies.**
``` bash
go mod tidy
or
go mod tidy -compat=1.17
```

**test using real API**
``` bash
make check
```
**test without real API**
``` bash
make test
```
Please make sure  `make check` passes before create PR.


