package config

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) {
	oldVersion, oldCommit := Version, GitCommit

	t.Cleanup(func() { Version, GitCommit = oldVersion, oldCommit })
}

func TestFormatVersion_NoGitCommit(t *testing.T) {
	setup(t)

	Version = "1.2.3"
	GitCommit = ""

	assert.Equal(t, "1.2.3", FormatVersion())
}

func TestFormatVersion(t *testing.T) {
	setup(t)

	Version = "1.2.3"
	GitCommit = "abc"

	expected := fmt.Sprintf("1.2.3 (git: %s)", GitCommit)

	assert.Equal(t, expected, FormatVersion())
}

func FormatBuildInfo_UsesBuildTimeAndGoVersion_WhenSet(t *testing.T) {
	setup(t)

	BuildTime = "2026-02-20T00:00:00Z"
	GoVersion = "go1.23.0"

	build, goVer := FormatBuildInfo()

	assert.Equal(t, build, BuildTime)
	assert.Equal(t, goVer, GoVersion)
}

func FormatBuildInfo_UsesBuildTimeAndWithoutGoVersion_WhenSet(t *testing.T) {
	setup(t)

	BuildTime = "2026-02-20T00:00:00Z"
	GoVersion = runtime.Version()

	build, goVer := FormatBuildInfo()

	assert.Equal(t, build, BuildTime)
	assert.Equal(t, goVer, GoVersion)
}
