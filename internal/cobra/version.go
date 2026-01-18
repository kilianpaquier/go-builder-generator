package cobra

import (
	"github.com/spf13/cobra"

	"github.com/kilianpaquier/go-builder-generator/internal/build"
)

func version() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show current version",
		Run:   func(_ *cobra.Command, _ []string) { logger.Info(build.GetInfo().String()) },
	}
	return cmd
}
