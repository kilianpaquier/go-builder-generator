package cobra

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	flagLogFormat = "log-format"
	flagLogLevel  = "log-level"
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cmd := rootCmd()
	cmd.AddCommand(generateCmd())
	cmd.AddCommand(version())

	if err := cmd.Execute(); err != nil {
		subcmd, _, _ := cmd.Find(os.Args[1:])
		usage(cmd, subcmd, err)
		logger.Error(err.Error())
		os.Exit(1) //nolint:revive
	}
}

func rootCmd() *cobra.Command {
	logFormat, logLevel := "text", "info"

	cmd := &cobra.Command{
		Use:               "go-builder-generator",
		SilenceErrors:     true, // don't print errors with cobra, let logger.Error handle them
		SilenceUsage:      true, // don't print help on errors, let usage function handle printing depending on command error
		PersistentPreRunE: func(*cobra.Command, []string) error { return setupLogger(logFormat, logLevel) },
	}

	cmd.PersistentFlags().StringVar(&logFormat, flagLogFormat, coalesce(getenv(flagLogFormat), logFormat), `set logging format (either "text" or "json")`)
	cmd.PersistentFlags().StringVar(&logLevel, flagLogLevel, coalesce(getenv(flagLogLevel), logLevel), "set logging level")

	_ = setupLogger(logFormat, logLevel) // ensure logging is correctly configured with default values even when a bad input flag is given

	return cmd
}
