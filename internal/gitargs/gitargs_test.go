package gitargs_test

import (
	"testing"

	"github.com/joselitofilho/go-conventional-commits/internal/gitargs"
	"github.com/stretchr/testify/require"
)

func TestArgs(t *testing.T) {
	tests := []struct {
		name          string
		latestVersion string
		expected      []string
	}{
		{
			name:          "with latest version",
			latestVersion: "v0.1.2",
			expected:      []string{"-s", "v0.1.2..HEAD"},
		},
		{
			name:          "without latest version",
			latestVersion: "",
			expected:      []string{"-s"},
		},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			logArgs := gitargs.GitLogArgs{LatestVersion: tc.latestVersion}
			actual := logArgs.Args()
			require.Equal(t, tc.expected, actual)
		})
	}
}
