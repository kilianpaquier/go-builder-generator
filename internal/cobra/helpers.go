package cobra

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

var (
	leveler = new(slog.LevelVar)
	logger  = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: leveler}))
)

func setupLogger(logFormat, logLevel string) error {
	switch logFormat {
	case "text":
		// nothing specific to do since default logger is text
	case "json":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: leveler}))
	default:
		return fmt.Errorf(`invalid argument %q for "--%s" flag`, logFormat, flagLogFormat)
	}

	var level slog.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		return fmt.Errorf(`invalid argument %q for "--%s" flag`, logLevel, flagLogLevel)
	}
	leveler.Set(level)

	return nil
}

func coalesce(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func getenv(flag string) string {
	key := strings.ToUpper(strings.ReplaceAll(flag, "-", "_"))
	return os.Getenv(key)
}
