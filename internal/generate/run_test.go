package generate_test

import (
	"os"
	"path/filepath"
	"testing"

	filesystem "github.com/kilianpaquier/filesystem/pkg"
	filesystem_tests "github.com/kilianpaquier/filesystem/pkg/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kilianpaquier/go-builder-generator/internal/generate"
)

func TestRun(t *testing.T) {
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	t.Run("error_no_file", func(t *testing.T) {
		// Arrange
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(t.TempDir(), "invalid.go"),
			Structs: []string{},
		}

		// Act
		err := generate.Run(options)

		// Assert
		assert.ErrorContains(t, err, "failed to parse")
		assert.ErrorContains(t, err, "no such file or directory")
		assert.NoDirExists(t, destdir)
	})

	t.Run("error_invalid_tags", func(t *testing.T) {
		// Arrange
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(testdata, "errors", "invalid_tags.go"),
			Structs: []string{"Invalid"},
		}

		// Act
		err := generate.Run(options)

		// Assert
		assert.ErrorContains(t, err, "failed to parse tags")
		assert.ErrorContains(t, err, "failed to parse builder for struct Invalid")
	})

	t.Run("success_channels", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_channels", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(testdata, "success_channels", "types.go"),
			Structs: []string{"Chan"},
		}

		// Act
		err := generate.Run(options)

		// Assert
		assert.NoError(t, err)
		filesystem_tests.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_export", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_export", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(testdata, "success_export", "types.go"),
			Structs: []string{"Export"},
		}

		// Act
		err := generate.Run(options)

		// Assert
		assert.NoError(t, err)
		filesystem_tests.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_funcs", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_funcs", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(testdata, "success_funcs", "types.go"),
			Structs: []string{"Func"},
		}

		// Act
		err := generate.Run(options)

		// Assert
		assert.NoError(t, err)
		filesystem_tests.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_interface", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_interface", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(testdata, "success_interface", "types.go"),
			Structs: []string{"Interface", "InterfaceNoFields"},
		}

		// Act
		err := generate.Run(options)

		// Assert
		assert.NoError(t, err)
		filesystem_tests.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_maps", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_maps", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(testdata, "success_maps", "types.go"),
			Structs: []string{"Map"},
		}

		// Act
		err := generate.Run(options)

		// Assert
		assert.NoError(t, err)
		filesystem_tests.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_naming", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_naming", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(testdata, "success_naming", "types.go"),
			Structs: []string{"Naming"},
		}

		// Act
		err := generate.Run(options)

		// Assert
		assert.NoError(t, err)
		filesystem_tests.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_same_package", func(t *testing.T) {
		// Arrange
		tmp := t.TempDir()
		assertdir := filepath.Join(testdata, "success_same_package")

		src := filepath.Join(testdata, "success_same_package", "types.go")
		dest := filepath.Join(tmp, "types.go")

		err := filesystem.CopyFile(src, dest)
		require.NoError(t, err)

		options := generate.CLIOptions{
			Destdir: tmp,
			File:    dest,
			Structs: []string{"SamePackage"},
		}

		// Act
		err = generate.Run(options)

		// Assert
		assert.NoError(t, err)
		filesystem_tests.AssertEqualDir(t, assertdir, tmp)
	})

	t.Run("success_slices", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_slices", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(testdata, "success_slices", "types.go"),
			Structs: []string{"ArrayAndSlice"},
		}

		// Act
		err := generate.Run(options)

		// Assert
		assert.NoError(t, err)
		filesystem_tests.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_struct", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_struct", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(testdata, "success_struct", "types.go"),
			Structs: []string{"Struct", "StructNoFields"},
		}

		// Act
		err := generate.Run(options)

		// Assert
		assert.NoError(t, err)
		filesystem_tests.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_with_options", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_with_options", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir:       destdir,
			File:          filepath.Join(testdata, "success_with_options", "types.go"),
			Structs:       []string{"Options", "Empty"},
			UserValidator: true,
		}

		// Act
		err := generate.Run(options)

		// Assert
		assert.NoError(t, err)
		filesystem_tests.AssertEqualDir(t, assertdir, destdir)
	})
}
