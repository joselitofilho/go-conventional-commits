package transformers_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joselitofilho/go-conventional-commits/internal/changelogs"
	"github.com/joselitofilho/go-conventional-commits/internal/transformers"
)

func TestTransforms_ChangeLog_WithSimpleMessage(t *testing.T) {
	message := `feat: added a new feature`
	projectLink := "https://myproject.domain.com/board/"

	changeLog := transformers.TransformChangeLog(message, projectLink)

	expected := &changelogs.ChangeLog{
		Category: "Features",
		Refs:     "",
		Title:    "added a new feature",
		Link:     "",
	}
	require.Equal(t, expected, changeLog)
}

func TestTransforms_ChangeLog_WithFeatures(t *testing.T) {
	message := `feat: added a new feature

Refs #GCC-123
`
	projectLink := "https://myproject.domain.com/board/"

	changeLog := transformers.TransformChangeLog(message, projectLink)

	expected := &changelogs.ChangeLog{
		Category: "Features",
		Refs:     "GCC-123",
		Title:    "added a new feature",
		Link:     "[GCC-123](https://myproject.domain.com/board/GCC-123)",
	}
	require.Equal(t, expected, changeLog)
}

func TestTransforms_ChangeLog_WithFixes(t *testing.T) {
	message := `fix: fixed the problem

Refs #GCC-321
`
	projectLink := "https://myproject.domain.com/board/"

	changeLog := transformers.TransformChangeLog(message, projectLink)

	expected := &changelogs.ChangeLog{
		Category: "Fixes",
		Refs:     "GCC-321",
		Title:    "fixed the problem",
		Link:     "[GCC-321](https://myproject.domain.com/board/GCC-321)",
	}
	require.Equal(t, expected, changeLog)
}

func TestTransforms_ChangeLog_WithTitle(t *testing.T) {
	tests := []struct {
		name    string
		message string
	}{
		{
			name: "Title:",
			message: `feat: added a new feature

Description of the new feature
more details

Title: Amazing new feature
Refs #GCC-123
`,
		},
		{
			name: "Title #",
			message: `feat: added a new feature

Description of the new feature
more details

Title #Amazing new feature
Refs #GCC-123
`,
		},
	}

	projectLink := "https://myproject.domain.com/board/"
	expected := &changelogs.ChangeLog{
		Category: "Features",
		Refs:     "GCC-123",
		Title:    "Amazing new feature",
		Link:     "[GCC-123](https://myproject.domain.com/board/GCC-123)",
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			changeLog := transformers.TransformChangeLog(tc.message, projectLink)
			require.Equal(t, expected, changeLog)
		})
	}
}

func TestTransforms_ChangeLog_WithRefs(t *testing.T) {
	tests := []struct {
		name    string
		message string
	}{
		{
			name: "Refs:",
			message: `feat: added a new feature

Refs: GCC-123
`,
		},
		{
			name: "Refs #",
			message: `feat: added a new feature

Refs #GCC-123
`,
		},
	}

	projectLink := "https://myproject.domain.com/board/"
	expected := &changelogs.ChangeLog{
		Category: "Features",
		Refs:     "GCC-123",
		Title:    "added a new feature",
		Link:     "[GCC-123](https://myproject.domain.com/board/GCC-123)",
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			changeLog := transformers.TransformChangeLog(tc.message, projectLink)
			require.Equal(t, expected, changeLog)
		})
	}
}

func TestTransforms_ChangeLog_WithRefsLowerCase(t *testing.T) {
	message := `feat: added a new feature

refs #GCC-123
`
	projectLink := "https://myproject.domain.com/board/"

	changeLog := transformers.TransformChangeLog(message, projectLink)

	expected := &changelogs.ChangeLog{
		Category: "Features",
		Refs:     "GCC-123",
		Title:    "added a new feature",
		Link:     "[GCC-123](https://myproject.domain.com/board/GCC-123)",
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
	projectLink := "https://myproject.domain.com/board/"

	changeLogs := transformers.TransformChangeLogs(messages, projectLink)

	expected := changelogs.ChangeLogs{
		"GCC-123": &changelogs.ChangeLog{
			Category: "Features",
			Refs:     "GCC-123",
			Title:    "Amazing new feature",
			Link:     "[GCC-123](https://myproject.domain.com/board/GCC-123)",
		},
		"GCC-321": &changelogs.ChangeLog{
			Category: "Fixes",
			Refs:     "GCC-321",
			Title:    "Fix the problem",
			Link:     "[GCC-321](https://myproject.domain.com/board/GCC-321)",
		},
	}
	require.Equal(t, expected, changeLogs)
}

func TestTransforms_ChangeLogs_WithoutRefs(t *testing.T) {
	messages := []string{`feat: added a new feature #f4f7dec`, `fix: fixed the problem #f5f7dec`}
	projectLink := "https://myproject.domain.com/board/"

	changeLogs := transformers.TransformChangeLogs(messages, projectLink)

	expected := changelogs.ChangeLogs{
		"#f4f7dec": &changelogs.ChangeLog{
			Category: "Features",
			Refs:     "#f4f7dec",
			Title:    "added a new feature",
			Link:     "",
		},
		"#f5f7dec": &changelogs.ChangeLog{
			Category: "Fixes",
			Refs:     "#f5f7dec",
			Title:    "fixed the problem",
			Link:     "",
		},
	}
	require.Equal(t, expected, changeLogs)
}
