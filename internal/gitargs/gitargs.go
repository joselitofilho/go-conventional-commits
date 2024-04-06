package gitargs

import "fmt"

const Head = "HEAD"

// GitLogArgs a struct to help us to build a git log command arguments
type GitLogArgs struct {
	latestVersion  string
	currentVersion string
}

func NewGitLogArgs(latestVersion, currentVersion string) *GitLogArgs {
	return &GitLogArgs{latestVersion: latestVersion, currentVersion: currentVersion}
}

// Args return a list of git log command arguments to get logs from the latest version
func (a *GitLogArgs) Args() []string {
	args := []string{"-s"}

	if a.latestVersion != "" {
		currVer := a.currentVersion
		if currVer == "" {
			currVer = Head
		}

		args = append(args, fmt.Sprintf("%s..%s", a.latestVersion, currVer))
	}

	return args
}
