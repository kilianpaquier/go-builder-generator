package cobra

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var (
	leveler   = new(slog.LevelVar)
	logger    = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: leveler}))
	logFormat = "text"
	logLevel  = "info"

	rootCmd = &cobra.Command{
		Use:               "go-builder-generator",
		SilenceErrors:     true, // don't print errors with cobra, let logger.Fatal handle them
		PersistentPreRunE: preRun,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", logLevel, "set logging level")
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", logFormat, `set logging format (either "text" or "json")`)

	_ = preRun(rootCmd, nil) // ensure logging is correctly configured with default values even when a bad input flag is given
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(1) //nolint:revive
	}
}

func preRun(*cobra.Command, []string) error {
	switch logFormat {
	case "text":
		// nothing specific to do since default logger is text
	case "json":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: leveler}))
	default:
		return errors.New(`invalid --log-format argument, must be either "json" or "text"`)
	}

	var level slog.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		level = slog.LevelInfo
	}
	leveler.Set(level)
	logger.Debug(fmt.Sprintf("running with level '%s'", level))
	return nil
}
