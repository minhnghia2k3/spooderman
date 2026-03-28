package main

import (
	"fmt"
	"slices"
	"testing"

	"github.com/minhnghia2k3/spooderman/pkg"
	"github.com/minhnghia2k3/spooderman/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNewSpooderManCommand(t *testing.T) {
	short := fmt.Sprintf("%s  - personal AI assistant v%s\n\n", pkg.Logo, config.Version)

	cmd := NewSpoodermanCommand()

	assert.Equal(t, cmd.Use, "spooderman")
	assert.Equal(t, cmd.Short, short)
	assert.Equal(t, cmd.Example, "spooderman version")

	allowCmds := []string{
		"version",
	}

	subCommands := cmd.Commands()
	assert.Equal(t, len(allowCmds), len(subCommands))

	for _, subCmd := range subCommands {
		found := slices.Contains(allowCmds, subCmd.Name())
		assert.True(t, found, "unknown command:", subCmd.Name())

		assert.False(t, subCmd.Hidden)
	}
}
