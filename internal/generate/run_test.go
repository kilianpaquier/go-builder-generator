package generate_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kilianpaquier/go-builder-generator/internal/fs"
	testfs "github.com/kilianpaquier/go-builder-generator/internal/fs/tests"
	"github.com/kilianpaquier/go-builder-generator/internal/generate"
)

func TestRun_Errors(t *testing.T) {
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	t.Run("error_no_src_gomod", func(t *testing.T) {
		// Arrange
		tmp := t.TempDir()

		src := filepath.Join(tmp, "no_gomod.go")
		err := os.WriteFile(src, []byte(
			`package no_gomod
			type NoGomod struct {
				Field string
			}
			`,
		), fs.RwRR)
		require.NoError(t, err)

		destdir := filepath.Join(testdata, "result")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(tmp, "no_gomod.go"),
			Structs: []string{},
		}

		// Act
		err = generate.Run(options, nil)

		// Assert
		assert.ErrorContains(t, err, "find src go.mod: no parent go.mod found")
		assert.NoDirExists(t, destdir)
	})

	t.Run("error_no_file", func(t *testing.T) {
		// Arrange
		tmp := t.TempDir()

		src := filepath.Join(tmp, "go.mod")
		err := os.WriteFile(src, []byte(``), fs.RwRR)
		require.NoError(t, err)

		destdir := filepath.Join(testdata, "result")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(tmp, "invalid.go"),
			Structs: []string{},
		}

		// Act
		err = generate.Run(options, nil)

		// Assert
		assert.ErrorContains(t, err, "parse file")
		assert.NoDirExists(t, destdir)
	})

	t.Run("error_no_dest_gomod", func(t *testing.T) {
		// Arrange
		destdir := filepath.Join(t.TempDir(), "builders")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    "",
			Structs: []string{"Invalid"},
		}

		// Act
		err := generate.Run(options, nil)

		// Assert
		assert.ErrorContains(t, err, "find dest go.mod: no parent go.mod found")
	})

	t.Run("error_not_required_module", func(t *testing.T) {
		// Arrange
		options := generate.CLIOptions{
			Destdir: filepath.Join(testdata, "result"),
			File:    "module::github.com/jarcoal/httpmock/match.go",
			Structs: []string{"Matcher"},
		}

		// Act
		err := generate.Run(options, nil)

		// Assert
		assert.ErrorContains(t, err, "missing module name 'github.com/jarcoal/httpmock/match.go")
	})

	t.Run("error_invalid_tags", func(t *testing.T) {
		// Arrange
		options := generate.CLIOptions{
			Destdir: filepath.Join(testdata, "result"),
			File:    filepath.Join(testdata, "errors", "invalid_tags.go"),
			Structs: []string{"Invalid"},
		}

		// Act
		err := generate.Run(options, nil)

		// Assert
		assert.ErrorContains(t, err, "field '[InvalidFlag]' options parsing")
		assert.ErrorContains(t, err, "tags parsing")
	})

	t.Run("error_unexported_type_generated_outside_package", func(t *testing.T) {
		// Arrange
		options := generate.CLIOptions{
			Destdir: filepath.Join(testdata, "result"),
			File:    filepath.Join(testdata, "errors", "unexported_type.go"),
			Structs: []string{"unexported"},
		}

		// Act
		err := generate.Run(options, nil)

		// Assert
		assert.ErrorContains(t, err, "is not exported (or one of its generic params is not) but generation destination is in an external package")
	})
}

func TestRun_DifferentPackage(t *testing.T) {
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	for _, tc := range []struct {
		generate.CLIOptions
		DirName string
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
				NoNotice:     true,
				Prefix:       "Set",
				ReturnCopy:   true,
				Structs:      []string{"Options", "Empty", "GenericOptions"},
				ValidateFunc: "Validate",
			},
		},
		{
			DirName: "success_generic",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"Struct", "SimpleGeneric", "AliasGeneric", "ComplexGeneric", "GenericAnonymousStruct", "ComplexSliceGeneric"},
			},
		},
	} {
		t.Run(tc.DirName, func(t *testing.T) {
			// Arrange
			tc.CLIOptions.File = filepath.Join(testdata, tc.DirName, "types.go")
			assertdir := filepath.Join(testdata, tc.DirName, "builders")
			tc.CLIOptions.Destdir = filepath.Join(testdata, tc.DirName, "result")
			tc.CLIOptions.PackageName = "builders"
			t.Cleanup(func() { require.NoError(t, os.RemoveAll(tc.CLIOptions.Destdir)) })

			// Act
			err := generate.Run(tc.CLIOptions, nil)

			// Assert
			assert.NoError(t, err)
			assert.NoError(t, testfs.EqualDirs(assertdir, tc.CLIOptions.Destdir))
		})
	}
}

func TestRun_ExternalModule(t *testing.T) {
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	for _, tc := range []struct {
		generate.CLIOptions
		DirName string
	}{
		{
			DirName: "success_module_replace",
			CLIOptions: generate.CLIOptions{
				File:    "module::github.com/stretchr/testify/mock/mock.go",
				Structs: []string{"Mock"},
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
				File:    "module::github.com/stretchr/testify/mock/mock.go",
				Structs: []string{"Mock"},
			},
		},
	} {
		t.Run(tc.DirName, func(t *testing.T) {
			// Arrange
			assertdir := filepath.Join(testdata, tc.DirName, "builders")
			tc.CLIOptions.Destdir = filepath.Join(testdata, tc.DirName, "result")
			tc.CLIOptions.PackageName = "builders"
			t.Cleanup(func() { require.NoError(t, os.RemoveAll(tc.CLIOptions.Destdir)) })

			// Act
			err := generate.Run(tc.CLIOptions, nil)

			// Assert
			assert.NoError(t, err)
			assert.NoError(t, testfs.EqualDirs(assertdir, tc.CLIOptions.Destdir))
		})
	}
}

func TestRun_SamePackage(t *testing.T) {
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	for _, tc := range []struct {
		generate.CLIOptions
		DirName string
	}{
		{
			DirName: "success_same_package",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"SamePackage", "unexportedType"},
			},
		},
		{
			DirName: "success_same_package_options",
			CLIOptions: generate.CLIOptions{
				Structs:     []string{"unexportedTypeOptions"},
				Prefix:      "Set",
				PackageName: "invalid", // shouldn't be used
			},
		},
	} {
		t.Run(tc.DirName, func(t *testing.T) {
			// Arrange
			assertdir := filepath.Join(testdata, tc.DirName)

			src := filepath.Join(testdata, tc.DirName, "types.go")
			tc.CLIOptions.Destdir = filepath.Join(testdata, tc.DirName, "result")
			tc.CLIOptions.File = filepath.Join(tc.CLIOptions.Destdir, "types.go")
			require.NoError(t, os.MkdirAll(tc.CLIOptions.Destdir, fs.RwxRxRxRx))
			require.NoError(t, fs.CopyFile(src, tc.CLIOptions.File))
			t.Cleanup(func() { require.NoError(t, os.RemoveAll(tc.CLIOptions.Destdir)) })

			// Act
			err := generate.Run(tc.CLIOptions, nil)

			// Assert
			assert.NoError(t, err)
			assert.NoError(t, testfs.EqualDirs(assertdir, tc.CLIOptions.Destdir))
		})
	}
}
