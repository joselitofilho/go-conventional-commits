package changelogs

import (
	"fmt"
	"sort"

	"github.com/joselitofilho/go-conventional-commits/internal/common"
)

// ChangeLogs a slice of parsed change log messages
type ChangeLogs map[string]*ChangeLog

// ChangeLog a parsed change log message
type ChangeLog struct {
	Category string `json:"category"`
	Refs     string `json:"refs"`
	Title    string `json:"title"`
	Link     string `json:"link"`
}

func (clogs ChangeLogs) String() (message string) {
	fxs := make([]*ChangeLog, 0, len(clogs))
	fts := make([]*ChangeLog, 0, len(clogs))
	chgs := make([]*ChangeLog, 0, len(clogs))
	for i := range clogs {
		cl := clogs[i]

		switch cl.Category {
		case common.Fixes:
			fxs = append(fxs, cl)
		case common.Features:
			fts = append(fts, cl)
		case common.Changes:
			chgs = append(chgs, cl)
		}
	}

	if len(fxs) > 0 {
		message += fmt.Sprintf("\n\n### %s\n", common.Fixes)
		message += makeBlockString(fxs)
	}

	if len(fts) > 0 {
		message += fmt.Sprintf("\n\n### %s\n", common.Features)
		message += makeBlockString(fts)
	}

	if len(chgs) > 0 {
		message += fmt.Sprintf("\n\n### %s\n", common.Changes)
		message += makeBlockString(chgs)
	}

	return
}

func makeBlockString(block []*ChangeLog) (message string) {
	sort.Slice(block, func(i, j int) bool {
		return block[i].Refs < block[j].Refs
	})
	for i := range block {
		link := block[i].Link
		if link == "" {
			message += fmt.Sprintf("- %s %s\n", block[i].Title, block[i].Refs)
		} else {
			message += fmt.Sprintf("%s: %s\n", block[i].Link, block[i].Title)
		}
	}

	return
}
