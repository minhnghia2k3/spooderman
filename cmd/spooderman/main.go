package main

import (
	"fmt"
	"os"

	"github.com/minhnghia2k3/spooderman/cmd/spooderman/internal/version"
	"github.com/minhnghia2k3/spooderman/pkg"
	"github.com/minhnghia2k3/spooderman/pkg/config"
	"github.com/spf13/cobra"
)

func NewSpoodermanCommand() *cobra.Command {
	short := fmt.Sprintf("%s  - personal AI assistant v%s\n\n", pkg.Logo, config.Version)

	cmd := &cobra.Command{
		Use:     "spooderman",
		Short:   short,
		Example: "spooderman version",
	}

	cmd.AddCommand(
		version.NewVersionCommand(),
	)

	return cmd
}

func main() {
	cmd := NewSpoodermanCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
