package transformers_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joselitofilho/go-conventional-commits/internal/changelogs"
	"github.com/joselitofilho/go-conventional-commits/internal/transformers"
)

func TestTransforms_ChangeLog_WithSimpleMessage(t *testing.T) {
	message := `feat: added a new feature`
	projectLink := "https://myproject.application.com/board/"

	changeLog := transformers.TransformChangeLog(message, projectLink)

	var expected *changelogs.ChangeLog
	require.Equal(t, expected, changeLog)
}

func TestTransforms_ChangeLog_WithFeatures(t *testing.T) {
	message := `feat: added a new feature

Refs #GCC-123
`
	projectLink := "https://myproject.application.com/board/"

	changeLog := transformers.TransformChangeLog(message, projectLink)

	expected := &changelogs.ChangeLog{
		Category: "Features",
		Refs:     "GCC-123",
		Title:    "<put the task title here>",
		Link:     "[GCC-123](https://myproject.application.com/board/GCC-123)",
	}
	require.Equal(t, expected, changeLog)
}

func TestTransforms_ChangeLog_WithFixes(t *testing.T) {
	message := `fix: fixed the problem

Refs #GCC-321
`
	projectLink := "https://myproject.application.com/board/"

	changeLog := transformers.TransformChangeLog(message, projectLink)

	expected := &changelogs.ChangeLog{
		Category: "Fixes",
		Refs:     "GCC-321",
		Title:    "<put the task title here>",
		Link:     "[GCC-321](https://myproject.application.com/board/GCC-321)",
	}
	require.Equal(t, expected, changeLog)
}

func TestTransforms_ChangeLog_WithSubject(t *testing.T) {
	message := `feat: added a new feature

Description of the new feature
more details

Title: Amazing new feature
Refs #GCC-123
`
	projectLink := "https://myproject.application.com/board/"

	changeLog := transformers.TransformChangeLog(message, projectLink)

	expected := &changelogs.ChangeLog{
		Category: "Features",
		Refs:     "GCC-123",
		Title:    "Amazing new feature",
		Link:     "[GCC-123](https://myproject.application.com/board/GCC-123)",
	}
	require.Equal(t, expected, changeLog)
}

func TestTransforms_ChangeLogs(t *testing.T) {
	messages := []string{`feat: added a new feature

Description of the new feature
more details

Title: Amazing new feature
Refs #GCC-123
`,
		`fix: fixed the problem

Description of the fix
more details

Title: Fix the problem
Refs #GCC-321`,
	}
	projectLink := "https://myproject.application.com/board/"

	changeLogs := transformers.TransformChangeLogs(messages, projectLink)

	expected := changelogs.ChangeLogs{
		"GCC-123": &changelogs.ChangeLog{
			Category: "Features",
			Refs:     "GCC-123",
			Title:    "Amazing new feature",
			Link:     "[GCC-123](https://myproject.application.com/board/GCC-123)",
		},
		"GCC-321": &changelogs.ChangeLog{
			Category: "Fixes",
			Refs:     "GCC-321",
			Title:    "Fix the problem",
			Link:     "[GCC-321](https://myproject.application.com/board/GCC-321)",
		},
	}
	require.Equal(t, expected, changeLogs)
}

func TestTransforms_ChangeLogs_WithoutRefs(t *testing.T) {
	messages := []string{`feat: added a new feature`, `fix: fixed the problem`}
	projectLink := "https://myproject.application.com/board/"

	changeLogs := transformers.TransformChangeLogs(messages, projectLink)

	require.Empty(t, changeLogs)
}
