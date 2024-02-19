package cobra

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// version is substituted with -ldflags
	version = "v0.0.0"

	versionCmd = &cobra.Command{
		Use:    "version",
		Short:  "Shows current go-builder-generator version",
		PreRun: SetLogLevel,
		Run: func(_ *cobra.Command, _ []string) {
			logrus.Info(version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
