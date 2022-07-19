package transformers_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joselitofilho/go-conventional-commits/internal/transformers"
)

func TestTransforms_ConventionalCommit(t *testing.T) {
	message := "feat: added a new feature"
	convetionalCommit := transformers.TransformConventionalCommit(message)
	require.Equal(t, "feat", convetionalCommit.Category)
}

func TestTransforms_ConventionalCommit_WithPatchChange(t *testing.T) {
	message := "fix: fixed the problem"
	convetionalCommit := transformers.TransformConventionalCommit(message)
	require.True(t, convetionalCommit.Patch)
}

func TestTransforms_ConventionalCommit_WithMinorChange(t *testing.T) {
	message := "feat: added a new feature"
	convetionalCommit := transformers.TransformConventionalCommit(message)
	require.True(t, convetionalCommit.Minor)
}

func TestTransforms_ConventionalCommit_WithMajorChange(t *testing.T) {
	message := "feat!: added a new feature"
	convetionalCommit := transformers.TransformConventionalCommit(message)
	require.True(t, convetionalCommit.Major)
}

func TestTransforms_ConventionalCommit_WithFooter(t *testing.T) {
	message := `feat: added a new feature

Refs #GCC-123
`
	convetionalCommit := transformers.TransformConventionalCommit(message)
	require.Equal(t, []string{"Refs #GCC-123"}, convetionalCommit.Footer)
}
