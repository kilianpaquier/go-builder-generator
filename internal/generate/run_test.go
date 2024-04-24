package generate_test

import (
	"fmt"
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
		assert.ErrorContains(t, err, fmt.Sprintf("file %s parsing", options.File))
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
		assert.ErrorContains(t, err, "field options parsing")
		assert.ErrorContains(t, err, "tags parsing")
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
		assert.ErrorContains(t, err, "is not exported (or one of its generic params is not) but generation destination is in an external package")
	})

	for _, tc := range []struct {
		generate.CLIOptions
		DirName     string
		SamePackage bool
	}{
		{
			DirName: "success_channels",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"Chan"},
			},
		},
		{
			DirName: "success_export",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"Export"},
			},
		},
		{
			DirName: "success_funcs",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"Func"},
			},
		},
		{
			DirName: "success_interface",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"Interface", "InterfaceNoFields"},
			},
		},
		{
			DirName: "success_maps",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"Map"},
			},
		},
		{
			DirName: "success_module_replace",
			CLIOptions: generate.CLIOptions{
				File:    "module::github.com/sirupsen/logrus/hooks/test/test.go",
				Structs: []string{"Hook"},
			},
		},
		{
			DirName: "success_module_root",
			CLIOptions: generate.CLIOptions{
				File:    "module::github.com/go-playground/validator/v10/errors.go",
				Structs: []string{"InvalidValidationError"},
			},
		},
		{
			DirName: "success_module_subdir",
			CLIOptions: generate.CLIOptions{
				File:    "module::github.com/sirupsen/logrus/hooks/test/test.go",
				Structs: []string{"Hook"},
			},
		},
		{
			DirName: "success_naming",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"Naming"},
			},
		},
		{
			DirName: "success_root_gomod",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"RootType"},
			},
		},
		{
			DirName: "success_same_package",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"SamePackage", "unexportedType"},
			},
			SamePackage: true,
		},
		{
			DirName: "success_same_package_prefix",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"unexportedTypePrefix"},
				Prefix:  "Set",
			},
			SamePackage: true,
		},
		{
			DirName: "success_slices",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"ArrayAndSlice"},
			},
		},
		{
			DirName: "success_struct",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"Struct", "StructNoFields"},
			},
		},
		{
			DirName: "success_with_options",
			CLIOptions: generate.CLIOptions{
				ReturnCopy:   true,
				Structs:      []string{"Options", "Empty", "GenericOptions"},
				ValidateFunc: "Validate",
			},
		},
		{
			DirName: "success_generic",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"Struct", "SimpleGeneric", "AliasGeneric", "ComplexGeneric", "FuckedUpGeneric"},
			},
		},
	} {
		t.Run(tc.DirName, func(t *testing.T) {
			// Arrange
			if tc.CLIOptions.File == "" {
				tc.CLIOptions.File = filepath.Join(testdata, tc.DirName, "types.go")
			}
			var assertdir, destdir string
			if tc.SamePackage {
				src := tc.CLIOptions.File
				assertdir = filepath.Join(testdata, tc.DirName)
				destdir = t.TempDir()
				tc.CLIOptions.File = filepath.Join(destdir, "types.go")
				require.NoError(t, filesystem.CopyFile(src, tc.File))
			} else {
				assertdir = filepath.Join(testdata, tc.DirName, "builders")
				destdir = filepath.Join(t.TempDir(), "builders")
			}
			tc.CLIOptions.Destdir = destdir

			// Act
			err := generate.Run(pwd, tc.CLIOptions)

			// Assert
			assert.NoError(t, err)
			testfs.AssertEqualDir(t, assertdir, destdir)
		})
	}
}
