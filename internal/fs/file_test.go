package fs_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kilianpaquier/go-builder-generator/internal/fs"
)

func TestCopyFile(t *testing.T) {
	tmp := t.TempDir()
	src := filepath.Join(tmp, "file.txt")
	dest := filepath.Join(tmp, "copy.txt")

	err := os.WriteFile(src, []byte("hey file"), fs.RwRR)
	require.NoError(t, err)

	t.Run("error_src_not_exists", func(t *testing.T) {
		// Arrange
		src := filepath.Join(tmp, "invalid.txt")

		// Act
		err := fs.CopyFile(src, dest)

		// Assert
		assert.ErrorContains(t, err, "failed to read")
		assert.NoFileExists(t, dest)
	})

	t.Run("error_destdir_not_exists", func(t *testing.T) {
		// Arrange
		dest := filepath.Join(tmp, "invalid", "file.txt")

		// Act
		err := fs.CopyFile(src, dest)

		// Assert
		assert.ErrorContains(t, err, "failed to create")
		assert.NoFileExists(t, dest)
	})

	t.Run("success", func(t *testing.T) {
		// Act
		err := fs.CopyFile(src, dest)

		// Assert
		assert.NoError(t, err)
		assert.FileExists(t, dest)
	})
}

func TestExists(t *testing.T) {
	t.Run("false_not_exists", func(t *testing.T) {
		// Arrange
		invalid := filepath.Join(os.TempDir(), "invalid")

		// Act
		exists := fs.Exists(invalid)

		// Assert
		assert.False(t, exists)
	})

	t.Run("true_exists", func(t *testing.T) {
		// Arrange
		srcdir := t.TempDir()
		src := filepath.Join(srcdir, "file.txt")
		file, err := os.Create(src)
		require.NoError(t, err)
		require.NoError(t, file.Close())

		// Act
		exists := fs.Exists(src)

		// Assert
		assert.True(t, exists)
	})
}
