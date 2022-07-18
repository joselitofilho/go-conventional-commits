package changelogs

import (
	"fmt"
	"sort"
	"strings"

	"github.com/joselitofilho/go-conventional-commits/internal/common"
)

// ChangeLogs a slice of parsed change log messages
type ChangeLogs map[string]*ChangeLog

// ChangeLog a parsed change log message
type ChangeLog struct {
	Category string `json:"category"`
	Refs     string `json:"refs"`
	Link     string `json:"link"`
}

func (clogs ChangeLogs) String() (message string) {
	fxs := make([]string, 0, len(clogs))
	fts := make([]string, 0, len(clogs))
	for i := range clogs {
		cl := clogs[i]

		switch cl.Category {
		case common.Fixes:
			fxs = append(fxs, cl.Link)
		case common.Features:
			fts = append(fts, cl.Link)
		}
	}

	if len(fxs) > 0 {
		sort.Strings(fxs)
		message += fmt.Sprintf("\n\n### Fixes\n%s: <put the task title here>", strings.Join(fxs, ": <put the task title here>\n"))
	}

	if len(fts) > 0 {
		sort.Strings(fts)
		message += fmt.Sprintf("\n\n### Features\n%s: <put the task title here>", strings.Join(fts, ": <put the task title here>\n"))
	}

	return
}
