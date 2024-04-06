package gitargs_test

import (
	"testing"

	"github.com/joselitofilho/go-conventional-commits/internal/gitargs"
	"github.com/stretchr/testify/require"
)

func TestArgs(t *testing.T) {
	tests := []struct {
		name           string
		latestVersion  string
		currentVersion string
		expected       []string
	}{
		{
			name:           "with latest version",
			latestVersion:  "v0.1.2",
			currentVersion: "",
			expected:       []string{"-s", "v0.1.2..HEAD"},
		},
		{
			name:           "with latest and current versions",
			latestVersion:  "v0.1.2",
			currentVersion: "v0.2.0",
			expected:       []string{"-s", "v0.1.2..v0.2.0"},
		},
		{
			name:           "without latest version",
			latestVersion:  "",
			currentVersion: "",
			expected:       []string{"-s"},
		},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			logArgs := gitargs.NewGitLogArgs(tc.latestVersion, tc.currentVersion)
			actual := logArgs.Args()
			require.Equal(t, tc.expected, actual)
		})
	}
}
