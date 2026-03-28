package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVersionCommand(t *testing.T) {
	cmd := NewVersionCommand()

	require.NotNil(t, cmd)

	assert.Equal(t, cmd.Use, "version")
	assert.Equal(t, len(cmd.Aliases), 1)
	assert.Equal(t, cmd.Aliases[0], "v")
	assert.Equal(t, cmd.Short, cmd.Short)
	assert.NotNil(t, cmd.Run)
}
