package transformers_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tsuyoshiwada/go-gitlog"

	"github.com/joselitofilho/go-conventional-commits/internal/transformers"
)

func TestTransforms_Messages(t *testing.T) {
	commits := []*gitlog.Commit{{
		Subject: "feat: added a new feature",
		Body:    "Refs #GCC-123",
	}}
	messages := transformers.TransformMessages(commits)
	require.Equal(t, []string{"feat: added a new feature\n\nRefs #GCC-123"}, messages)
}
