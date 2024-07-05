package cobra

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// version is substituted with -ldflags
	version = "v0.0.0"

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Shows current go-builder-generator version",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Print(version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
