package conventionalcommits

import (
	"encoding/json"
	"fmt"
)

var (
	Marshal = json.Marshal
)

// ConventionalCommits a slice of parsed conventional commit messages
type ConventionalCommits []*ConventionalCommit

// ConventionalCommit a parsed conventional commit message
type ConventionalCommit struct {
	Category    string   `json:"category"`
	Scope       string   `json:"scope"`
	Description string   `json:"description"`
	Body        string   `json:"body"`
	Footer      []string `json:"footer"`
	Major       bool     `json:"major"`
	Minor       bool     `json:"minor"`
	Patch       bool     `json:"patch"`
}

func (cc *ConventionalCommit) String() string {
	data, err := Marshal(cc)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return string(data)
}
