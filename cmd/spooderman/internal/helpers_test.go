package internal

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/minhnghia2k3/spooderman/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetConfigPath(t *testing.T) {
	t.Setenv("HOME", "/tmp/home")

	got := GetConfigPath()
	want := filepath.Join("/tmp/home", ".spooderman", "config.json")

	assert.Equal(t, want, got)
}

func TestGetConfigPath_WithSPOODERMAN_HOME(t *testing.T) {
	t.Setenv("SPOODERMAN_HOME", "/custom/spooderman")
	t.Setenv("HOME", "/tmp/home")

	got := GetConfigPath()
	want := filepath.Join("/custom/spooderman", "config.json")

	assert.Equal(t, want, got)
}

func TestGetConfigPath_WithSPOODERMAN_CONFIG(t *testing.T) {
	t.Setenv("SPOODERMAN_CONFIG", "/custom/config.json")
	t.Setenv(config.EnvHome, "/custom/picoclaw")
	t.Setenv("HOME", "/tmp/home")

	got := GetConfigPath()
	want := "/custom/config.json"

	assert.Equal(t, want, got)
}

func TestGetConfigPath_Windows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("windows-specific HOME behavior varies; run on windows")
	}

	testUserProfilePath := `C:\Users\Test`
	t.Setenv("USERPROFILE", testUserProfilePath)

	got := GetConfigPath()
	want := filepath.Join(testUserProfilePath, ".picoclaw", "config.json")

	require.True(t, strings.EqualFold(got, want), "GetConfigPath() = %q, want %q", got, want)
}
