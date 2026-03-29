package status

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func NewStatusCommand() *cobra.Command {
// 	return &cobra.Command{
// 		Use:     "status",
// 		Aliases: []string{"s"},
// 		Short:   "Show spooderman status",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			statusCmd()
// 		},
// 	}
// }

func TestNewStatusCommand(t *testing.T) {
	cmd := NewStatusCommand()

	assert.Equal(t, "status", cmd.Use)
	assert.Equal(t, "s", cmd.Aliases[0])
	assert.Equal(t, "Show spooderman status", cmd.Short)
	assert.NotNil(t, cmd.Run)
}
