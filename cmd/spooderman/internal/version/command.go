package version

import (
	"fmt"

	"github.com/minhnghia2k3/spooderman/pkg"
	"github.com/minhnghia2k3/spooderman/pkg/config"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Show version information",
		Run: func(_ *cobra.Command, _ []string) {
			printVersion()
		},
	}

	return cmd
}

func printVersion() {
	fmt.Printf("\t\t\t%s  spooderman %s\n", pkg.Logo, config.FormatVersion())
	build, goVer := config.FormatBuildInfo()
	if build != "" {
		fmt.Printf("\t\t\t   Build: %s\n", build)
	}
	if goVer != "" {
		fmt.Printf("\t\t\t   Go version: %s\n", goVer)
	}
}
