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
		Subject:  "<put the task title here>",
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
		Subject:  "<put the task title here>",
		Link:     "[GCC-321](https://myproject.application.com/board/GCC-321)",
	}
	require.Equal(t, expected, changeLog)
}

func TestTransforms_ChangeLog_WithSubject(t *testing.T) {
	message := `feat: added a new feature

Description of the new feature
more details

Refs #GCC-123
`
	projectLink := "https://myproject.application.com/board/"

	changeLog := transformers.TransformChangeLog(message, projectLink)

	expected := &changelogs.ChangeLog{
		Category: "Features",
		Refs:     "GCC-123",
		Subject:  "Description of the new feature\nmore details",
		Link:     "[GCC-123](https://myproject.application.com/board/GCC-123)",
	}
	require.Equal(t, expected, changeLog)
}
