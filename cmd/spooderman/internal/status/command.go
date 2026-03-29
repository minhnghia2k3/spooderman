package status

import (
	"fmt"
	"os"

	"github.com/minhnghia2k3/spooderman/cmd/spooderman/internal"
	"github.com/minhnghia2k3/spooderman/cmd/spooderman/internal/version"
	"github.com/spf13/cobra"
)

func NewStatusCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "status",
		Aliases: []string{"s"},
		Short:   "Show spooderman status",
		Run: func(cmd *cobra.Command, args []string) {
			statusCmd()
		},
	}
}

// statusCmd prints out spooderman version, config, workspace, authentication
func statusCmd() {
	cfg, err := internal.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	configPath := internal.GetConfigPath()

	version.PrintVersion()

	if _, err := os.Stat(configPath); err != nil {
		fmt.Println("Config:", configPath, "✗")
	} else {
		fmt.Println("Config:", configPath, "✓")
	}

	workspace := cfg.WorkspacePath()
	if _, err := os.Stat(workspace); err != nil {
		fmt.Println("Workspace:", workspace, "✗")
	} else {
		fmt.Println("Workspace:", workspace, "✓")
	}

	if _, err := os.Stat(configPath); err == nil {
		fmt.Printf("Model: %s ✓\n", cfg.Agents.Defaults.GetModelName())
		// TODO: load auth store
	}
}
