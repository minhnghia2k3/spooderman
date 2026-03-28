package config

import (
	"fmt"
	"runtime"
)

// Build-time variables injected via ldflags during build process.
// These are set by the Makefile or .goreleaser.yml using the -X flag:
//
// -X github.com/minhnghia2k3/spooderman/pkg/config.Version=<version>
var (
	Version   = "dev"
	GitCommit string // git commit SHA (short)
	BuildTime string // Build time in RFC3399 format (e.g. 2023-08-13T16:07:54Z)
	GoVersion string // Go version used for building
)

// FormatVersion prints out version string with optional short git sha commit
func FormatVersion() string {
	v := Version
	if GitCommit != "" {
		v += fmt.Sprintf(" (git: %s)", GitCommit)
	}

	return v
}

// FormatBuildInfo prints out build time and go version info
func FormatBuildInfo() (string, string) {
	build := BuildTime
	goVer := GoVersion
	if goVer == "" {
		goVer = runtime.Version()
	}

	return build, goVer
}
