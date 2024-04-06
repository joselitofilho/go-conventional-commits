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
	for i := range clogs {
		cl := clogs[i]

		switch cl.Category {
		case common.Fixes:
			fxs = append(fxs, cl)
		case common.Features:
			fts = append(fts, cl)
		}
	}

	if len(fxs) > 0 {
		message += fmt.Sprintf("\n\n### %s\n", common.Fixes)
		sort.Slice(fxs, func(i, j int) bool {
			return fxs[i].Refs < fxs[j].Refs
		})
		for i := range fxs {
			link := fxs[i].Link
			if link == "" {
				message += fmt.Sprintf("- %s %s\n", fxs[i].Title, fxs[i].Refs)
			} else {
				message += fmt.Sprintf("%s: %s\n", fxs[i].Link, fxs[i].Title)
			}
		}
	}

	if len(fts) > 0 {
		message += fmt.Sprintf("\n\n### %s\n", common.Features)
		sort.Slice(fts, func(i, j int) bool {
			return fts[i].Refs < fts[j].Refs
		})
		for i := range fts {
			link := fts[i].Link
			if link == "" {
				message += fmt.Sprintf("- %s %s\n", fts[i].Title, fts[i].Refs)
			} else {
				message += fmt.Sprintf("%s: %s\n", fts[i].Link, fts[i].Title)
			}
		}
	}

	return
}
