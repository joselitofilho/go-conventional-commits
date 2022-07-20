# go-conventional-commits

This is a tool to parser your git commit messages into a change log message base on [Conventional Commits](#conventional-commits) specification.

## Conventional Commits

We are following the [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) specification for adding human and machine readable meaning to commit messages. You can see the full specificantion [here](https://www.conventionalcommits.org/en/v1.0.0/#specification).

If you are using VSCode, you can follow [this](https://pawelgrzybek.com/multi-paragraph-git-commit-messages-cli-and-visual-studio-code/) to set up your environment.

Commit message example:
```
feat: Super cool new feature

Body (optional)

Refs #CDP-123
```

## Requirements

### Golang
This package requires the Go programming language extension for language support. It also requires you to have golang installed on your machine. To install, follow these instructions

## Installation

```bash
$ go install -v github.com/joselitofilho/go-conventional-commits/cmd/go-conventional-commits@latest
```

## Usage
```bash
usage: print [-h|--help] [--coverageCmd "<value>"] [-l|--latestVersion
             "<value>"] -n|--newVersion "<value>" [--projectLink "<value>"]
             [-p|--path "<value>"] [-u|--updateChangelog]

             Prints provided string to stdout

Arguments:

  -h  --help             Print help information
      --coverageCmd      Specify your code coverage command to get the value.
                         Default: 
  -l  --latestVersion    The name of the git tag with the latest version. For
                         example: v1.2.3. Default: 
  -n  --newVersion       The name of the git tag with new version. For example:
                         v0.1.2
      --projectLink      The base project link that we will concatenate the
                         task ID. For example:
                         https://myproject.domain.com/board/. Default: 
  -p  --path             The repository path. Default: .
  -u  --updateChangelog  If this flag is true the changelog file will be
                         updated. Default: false
```

## Features

### Generate a change log message

For example:
```bash
$ go-conventional-commits --latestVersion v0.1.2 --newVersion v1.0.0 --path $HOME/dev/myProject --projectLink https://myproject.domain.com/tasks/
```

Output:
```bash
## Release v1.0.0

### Fixes
[GCC-007](https://myproject.domain.com/tasks/GCC-007): <put the task title here>

### Features
[GCC-123](https://myproject.domain.com/tasks/GCC-123): <put the task title here>

Code coverage: <put the value here>%
```

### Specify your code coverage command to fill the value in the changelog

For example:
```bash
$ go-conventional-commits --latestVersion v0.1.2 --newVersion v1.0.0 --path $HOME/dev/myProject --projectLink https://myproject.domain.com/tasks/ --coverageCmd "make coverage"
```

Output:
```bash
## Release v1.0.0

### Fixes
[GCC-007](https://myproject.domain.com/tasks/GCC-007): <put the task title here>

### Features
[GCC-123](https://myproject.domain.com/tasks/GCC-123): <put the task title here>

Code coverage: 99.9%
```

### Update the changelog.md file in your repository

For example:
```bash
$ go-conventional-commits --latestVersion v0.1.2 --newVersion v1.0.0 --path $HOME/dev/myProject --projectLink https://myproject.domain.com/tasks/ --coverageCmd "make coverage" --updateChangelog
```

Output:
```bash
## Release v1.0.0

### Fixes
[GCC-007](https://myproject.domain.com/tasks/GCC-007): <put the task title here>

### Features
[GCC-123](https://myproject.domain.com/tasks/GCC-123): <put the task title here>

Code coverage: 99.9%
```

See that output in the top of your `$HOME/dev/myProject/changelog.md` file.
