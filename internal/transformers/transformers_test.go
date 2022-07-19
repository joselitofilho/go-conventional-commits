package transformers_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joselitofilho/go-conventional-commits/internal/changelogs"
	"github.com/joselitofilho/go-conventional-commits/internal/transformers"
)

func TestTransformsConventionalCommit(t *testing.T) {
	message := "feat: added a new feature"
	convetionalCommit := transformers.TransformConventionalCommit(message)
	require.Equal(t, "feat", convetionalCommit.Category)
}

func TestTransformsConventionalCommitWithPatchChange(t *testing.T) {
	message := "fix: fixed the problem"
	convetionalCommit := transformers.TransformConventionalCommit(message)
	require.True(t, convetionalCommit.Patch)
}

func TestTransformsConventionalCommitWithMinorChange(t *testing.T) {
	message := "feat: added a new feature"
	convetionalCommit := transformers.TransformConventionalCommit(message)
	require.True(t, convetionalCommit.Minor)
}

func TestTransformsConventionalCommitWithMajorChange(t *testing.T) {
	message := "feat!: added a new feature"
	convetionalCommit := transformers.TransformConventionalCommit(message)
	require.True(t, convetionalCommit.Major)
}

func TestTransformsConventionalCommitWithFooter(t *testing.T) {
	message := `feat: added a new feature

Refs #GCC-123
`
	convetionalCommit := transformers.TransformConventionalCommit(message)
	require.Equal(t, []string{"Refs #GCC-123"}, convetionalCommit.Footer)
}

func TestTransformsChangeLogWithSimpleMessage(t *testing.T) {
	message := `feat: added a new feature`
	projectLink := "https://myproject.application.com/board/"

	changeLog := transformers.TransformChangeLog(message, projectLink)

	var expected *changelogs.ChangeLog
	require.Equal(t, expected, changeLog)
}

func TestTransformsChangeLogWithFeatures(t *testing.T) {
	message := `feat: added a new feature

Refs #GCC-123
`
	projectLink := "https://myproject.application.com/board/"

	changeLog := transformers.TransformChangeLog(message, projectLink)

	expected := &changelogs.ChangeLog{
		Category: "Features",
		Refs:     "GCC-123",
		Link:     "[GCC-123](https://myproject.application.com/board/GCC-123)",
	}
	require.Equal(t, expected, changeLog)
}

func TestTransformsChangeLogWithFixes(t *testing.T) {
	message := `fix: fixed the problem

Refs #GCC-321
`
	projectLink := "https://myproject.application.com/board/"

	changeLog := transformers.TransformChangeLog(message, projectLink)

	expected := &changelogs.ChangeLog{
		Category: "Fixes",
		Refs:     "GCC-321",
		Link:     "[GCC-321](https://myproject.application.com/board/GCC-321)",
	}
	require.Equal(t, expected, changeLog)
}
