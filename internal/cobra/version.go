package cobra

import (
	"github.com/spf13/cobra"

	"github.com/kilianpaquier/go-builder-generator/internal/build"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current go-builder-generator version",
	Run:   func(_ *cobra.Command, _ []string) { logger.Info(build.GetInfo()) },
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
