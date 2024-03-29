package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/tsuyoshiwada/go-gitlog"

	"github.com/joselitofilho/go-conventional-commits/internal/changelogs"
	"github.com/joselitofilho/go-conventional-commits/internal/gitargs"
	"github.com/joselitofilho/go-conventional-commits/internal/transformers"
)

func main() {
	parser := argparse.NewParser("print", "Prints provided string to stdout")

	coverageCmd := parser.String("", "coverageCmd", &argparse.Options{
		Required: false,
		Default:  "",
		Help:     "Specify your code coverage command to get the value",
	})
	latestVersion := parser.String("l", "latestVersion", &argparse.Options{
		Required: false,
		Default:  "",
		Help:     "The name of the git tag with the latest version. For example: v1.2.3",
	})
	newVersion := parser.String("n", "newVersion", &argparse.Options{
		Required: true,
		Help:     "The name of the git tag with new version. For example: v0.1.2",
	})
	projectLink := parser.String("", "projectLink", &argparse.Options{
		Required: false,
		Default:  "",
		Help:     "The base project link that we will concatenate the task ID. For example: https://myproject.domain.com/board/",
	})
	repoPath := parser.String("p", "path", &argparse.Options{
		Required: false,
		Default:  ".",
		Help:     `The repository path`,
	})
	updateChangelog := parser.Flag("u", "updateChangelog", &argparse.Options{
		Required: false,
		Default:  false,
		Help:     "If this flag is true the changelog file will be updated",
	})

	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	// Running git command to get logs
	git := gitlog.New(&gitlog.Config{Path: *repoPath})
	commits, err := git.Log(&gitargs.GitLogArgs{LatestVersion: *latestVersion}, nil)
	if err != nil {
		log.Fatalln(err)
	}

	changeLogs := transformers.TransformChangeLogs(transformers.TransformMessages(commits), *projectLink)

	codeCoverageValue := buildCoverageValue(*repoPath, *coverageCmd)

	// TODO: auto-generate number based on the lastest version and kind of commits.

	changeLog := buildChangeLog(changeLogs, *newVersion, codeCoverageValue)

	if *updateChangelog {
		fmt.Println("Updating changelog file. Please, wait...")
		updateChangelogFile(changeLog, *repoPath)
	}

	fmt.Println("Output:")
	fmt.Println(changeLog)
}

// TODO: Move to other package
func buildChangeLog(changeLogs changelogs.ChangeLogs, newVersion, codeCoverageValue string) (changeLog string) {
	changeLog = fmt.Sprintf("## Release %s", newVersion)
	changeLog += changeLogs.String()
	changeLog += fmt.Sprintf("\n\nCode coverage: %s\n", codeCoverageValue)
	return
}

// TODO: Move to other package
func buildCoverageValue(repoPath, coverageCmd string) string {
	codeCoverageValue := "<put the value here>%"

	if coverageCmd != "" {
		fmt.Println("Running coverage. Please, wait...")

		coverageCmdSlice := strings.Split(coverageCmd, " ")

		cmd := exec.Command(coverageCmdSlice[0], coverageCmdSlice[0:]...)
		cmd.Dir = repoPath
		out, err := cmd.Output()
		if err != nil {
			log.Fatalln(err)
		}
		re := regexp.MustCompile("[+-]?([0-9]*[.])?[0-9]+%")
		found := re.FindAllString(string(out), -1)
		if len(found) > 0 {
			codeCoverageValue = found[len(found)-1]
		}
	}

	return codeCoverageValue
}

// TODO: Move to other package
func updateChangelogFile(changeLog, repoPath string) {
	changeLogLines := strings.Split(changeLog, "\n")

	changeLogFile := fmt.Sprintf("%s/changelog.md", repoPath)

	input, err := ioutil.ReadFile(changeLogFile)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	newLines := make([]string, 0, len(changeLogLines)+len(lines))
	newLines = append(newLines, lines[0:2]...)
	newLines = append(newLines, changeLogLines...)
	newLines = append(newLines, lines[2:]...)

	output := strings.Join(newLines, "\n")
	err = ioutil.WriteFile(changeLogFile, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
