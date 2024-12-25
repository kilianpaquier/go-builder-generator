package cobra

import (
	"errors"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/kilianpaquier/cli-sdk/pkg/upgrade"
	"github.com/spf13/cobra"
)

var (
	dest        string
	major       string
	minor       string
	prereleases bool

	upgradeCmd = &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade or install go-builder-generator",
		Run: func(cmd *cobra.Command, _ []string) {
			options := []upgrade.RunOption{
				upgrade.WithDestination(dest),
				upgrade.WithHTTPClient(cleanhttp.DefaultClient()),
				upgrade.WithMajor(major),
				upgrade.WithMinor(minor),
				upgrade.WithPrereleases(prereleases),
			}

			version, err := upgrade.Run(cmd.Context(), "go-builder-generator", version, upgrade.GithubReleases("kilianpaquier", "go-builder-generator"), options...)
			if err != nil {
				switch {
				case errors.Is(err, upgrade.ErrNoNewVersion):
					logger.Info(err)
				case errors.Is(err, upgrade.ErrAlreadyInstalled):
					logger.Infof("version '%s' is already installed", version)
				default:
					logger.Fatal(err)
				}
				return
			}
			logger.Infof("successfully installed version '%s'", version)
		},
	}
)

func init() {
	rootCmd.AddCommand(upgradeCmd)

	upgradeCmd.Flags().StringVar(&dest, "dest", "", `destination directory where go-builder-generator will be upgraded / installed (by default "${HOME}/.local/bin")`)
	_ = upgradeCmd.MarkFlagDirname("dest")

	upgradeCmd.Flags().StringVar(&major, "major", "", `which major version to upgrade / install (must be of the form "v1", "v2", etc.) - mutually exclusive with --minor option`)
	upgradeCmd.Flags().StringVar(&minor, "minor", "", `which minor version to upgrade / install (must be of the form "v1.5", "v2.4", etc.) - mutually exclusive with --major option`)
	upgradeCmd.MarkFlagsMutuallyExclusive("major", "minor")

	upgradeCmd.Flags().BoolVar(&prereleases, "prereleases", false, "whether prereleases are accepted for installation or not")
}
