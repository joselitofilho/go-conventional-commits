package gitargs

import "fmt"

// GitLogArgs a struct to help us to build a git log command arguments
type GitLogArgs struct {
	LatestVersion string
}

// Args return a list of git log command arguments to get logs from the latest version
func (a *GitLogArgs) Args() []string {
	args := []string{"-s"}
	if a.LatestVersion != "" {
		args = append(args, fmt.Sprintf("%s..HEAD", a.LatestVersion))
	}
	return args
}
