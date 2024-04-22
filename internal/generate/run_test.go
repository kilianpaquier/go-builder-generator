package generate_test

import (
	"os"
	"path/filepath"
	"testing"

	filesystem "github.com/kilianpaquier/filesystem/pkg"
	testfs "github.com/kilianpaquier/filesystem/pkg/tests"
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
		err := generate.Run(pwd, options)

		// Assert
		assert.ErrorContains(t, err, "failed to parse")
		assert.NoDirExists(t, destdir)
	})

	t.Run("error_no_gomod", func(t *testing.T) {
		// Arrange
		src := filepath.Join(t.TempDir(), "no_gomod.go")
		err := os.WriteFile(src, []byte(
			`package no_gomod
			type NoGomod struct {
				Field string
			}
			`,
		), filesystem.RwRR)
		require.NoError(t, err)

		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    src,
			Structs: []string{"Invalid"},
		}

		// Act
		err = generate.Run(pwd, options)

		// Assert
		assert.ErrorContains(t, err, "no go.mod found")
	})

	t.Run("error_not_required_module", func(t *testing.T) {
		// Arrange
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    "module::github.com/jarcoal/httpmock/match.go",
			Structs: []string{"Matcher"},
		}

		// Act
		err := generate.Run(pwd, options)

		// Assert
		assert.ErrorContains(t, err, "failed to find appropriate require")
		assert.ErrorContains(t, err, "make sure you are importing this module")
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
		err := generate.Run(pwd, options)

		// Assert
		assert.ErrorContains(t, err, "failed to parse tags")
		assert.ErrorContains(t, err, "failed to parse builder for struct Invalid")
	})

	t.Run("error_unexported_type_generated_outside_package", func(t *testing.T) {
		// Arrange
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(testdata, "errors", "unexported_type.go"),
			Structs: []string{"unexported"},
		}

		// Act
		err := generate.Run(pwd, options)

		// Assert
		assert.ErrorContains(t, err, "is not exported but generation destination is in an external package")
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
		err := generate.Run(pwd, options)

		// Assert
		assert.NoError(t, err)
		testfs.AssertEqualDir(t, assertdir, destdir)
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
		err := generate.Run(pwd, options)

		// Assert
		assert.NoError(t, err)
		testfs.AssertEqualDir(t, assertdir, destdir)
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
		err := generate.Run(pwd, options)

		// Assert
		assert.NoError(t, err)
		testfs.AssertEqualDir(t, assertdir, destdir)
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
		err := generate.Run(pwd, options)

		// Assert
		assert.NoError(t, err)
		testfs.AssertEqualDir(t, assertdir, destdir)
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
		err := generate.Run(pwd, options)

		// Assert
		assert.NoError(t, err)
		testfs.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_module_replace", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_module_replace", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    "module::github.com/sirupsen/logrus/hooks/test/test.go",
			Structs: []string{"Hook"},
		}

		// Act
		err := generate.Run(assertdir, options)

		// Assert
		assert.NoError(t, err)
		testfs.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_module_root", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_module_root", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    "module::github.com/go-playground/validator/v10/errors.go",
			Structs: []string{"InvalidValidationError"},
		}

		// Act
		err := generate.Run(pwd, options)

		// Assert
		require.NoError(t, err)

		// Assert
		testfs.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_module_subdir", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_module_subdir", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    "module::github.com/sirupsen/logrus/hooks/test/test.go",
			Structs: []string{"Hook"},
		}

		// Act
		err := generate.Run(pwd, options)

		// Assert
		require.NoError(t, err)

		// Assert
		testfs.AssertEqualDir(t, assertdir, destdir)
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
		err := generate.Run(pwd, options)

		// Assert
		assert.NoError(t, err)
		testfs.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_root_gomod", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_root_gomod", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(testdata, "success_root_gomod", "types.go"),
			Structs: []string{"RootType"},
		}

		// Act
		err := generate.Run(pwd, options)

		// Assert
		assert.NoError(t, err)
		testfs.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_same_package", func(t *testing.T) {
		// Arrange
		destdir := t.TempDir()
		assertdir := filepath.Join(testdata, "success_same_package")

		src := filepath.Join(testdata, "success_same_package", "types.go")
		dest := filepath.Join(destdir, "types.go")
		require.NoError(t, filesystem.CopyFile(src, dest))

		options := generate.CLIOptions{
			Destdir: destdir,
			File:    dest,
			Structs: []string{"SamePackage", "unexportedType"},
		}

		// Act
		err := generate.Run(pwd, options)
		require.NoError(t, err)

		// Assert
		testfs.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_same_package_prefix", func(t *testing.T) {
		// Arrange
		destdir := t.TempDir()
		assertdir := filepath.Join(testdata, "success_same_package_prefix")

		src := filepath.Join(testdata, "success_same_package_prefix", "types.go")
		dest := filepath.Join(destdir, "types.go")
		require.NoError(t, filesystem.CopyFile(src, dest))

		options := generate.CLIOptions{
			Destdir:      destdir,
			File:         dest,
			SetterPrefix: "Set",
			Structs:      []string{"unexportedTypePrefix"},
		}

		// Act
		err := generate.Run(pwd, options)
		require.NoError(t, err)

		// Assert
		testfs.AssertEqualDir(t, assertdir, destdir)
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
		err := generate.Run(pwd, options)

		// Assert
		assert.NoError(t, err)
		testfs.AssertEqualDir(t, assertdir, destdir)
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
		err := generate.Run(pwd, options)

		// Assert
		assert.NoError(t, err)
		testfs.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_with_options", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_with_options", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir:      destdir,
			File:         filepath.Join(testdata, "success_with_options", "types.go"),
			Structs:      []string{"Options", "Empty"},
			ValidateFunc: "Validate",
		}

		// Act
		err := generate.Run(pwd, options)

		// Assert
		assert.NoError(t, err)
		testfs.AssertEqualDir(t, assertdir, destdir)
	})

	t.Run("success_return_copy", func(t *testing.T) {
		// Arrange
		assertdir := filepath.Join(testdata, "success_return_copy", "builders")
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir:    destdir,
			File:       filepath.Join(testdata, "success_return_copy", "types.go"),
			Structs:    []string{"ReturnCopy"},
			ReturnCopy: true,
		}

		// Act
		err := generate.Run(pwd, options)

		// Assert
		assert.NoError(t, err)
		testfs.AssertEqualDir(t, assertdir, destdir)
	})
}
