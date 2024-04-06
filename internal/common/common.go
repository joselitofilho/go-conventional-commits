package common

const (
	Features = "Features"
	Fixes    = "Fixes"
	Changes  = "Other Changes"
)

var (
	// PatchCategories is the list of categories used to determine if this commit is a patch bump or not
	PatchCategories = []string{
		"fix",
		"refactor",
		"perf",
		"docs",
		"style",
		"bug",
		"test",
	}

	// MinorCategories is the list of categories used to determine if this commit is a minor bump or not
	MinorCategories = []string{
		"feat",
		"feature",
		"story",
	}

	// MajorCategories is the list of categories used to determine if this commit is a major bump or not
	MajorCategories = []string{
		"breaking",
	}
)
