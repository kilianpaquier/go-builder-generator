package cobra //nolint:testpackage

import (
	"testing"

	"github.com/kilianpaquier/go-builder-generator/internal/testutils"
)

func TestRootFlags(t *testing.T) {
	t.Run("from_env", func(t *testing.T) {
		// Arrange
		t.Setenv("LOG_LEVEL", "debug")
		t.Setenv("LOG_FORMAT", "json")

		cmd := rootCmd()

		// Act
		err := cmd.Execute()

		// Assert
		testutils.NoError(testutils.Require(t), err)

		format, err := cmd.PersistentFlags().GetString(flagLogFormat)
		testutils.NoError(testutils.Require(t), err)
		testutils.Equal(t, "json", format)

		level, err := cmd.PersistentFlags().GetString(flagLogLevel)
		testutils.NoError(testutils.Require(t), err)
		testutils.Equal(t, "debug", level)
	})

	t.Run("from_flags", func(t *testing.T) {
		// Arrange
		cmd := rootCmd()
		cmd.SetArgs([]string{"--" + flagLogFormat, "json", "--" + flagLogLevel, "debug"})

		// Act
		err := cmd.Execute()

		// Assert
		testutils.NoError(testutils.Require(t), err)

		format, err := cmd.PersistentFlags().GetString(flagLogFormat)
		testutils.NoError(testutils.Require(t), err)
		testutils.Equal(t, "json", format)

		level, err := cmd.PersistentFlags().GetString(flagLogLevel)
		testutils.NoError(testutils.Require(t), err)
		testutils.Equal(t, "debug", level)
	})
}
