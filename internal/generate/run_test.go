package generate_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	compare "github.com/kilianpaquier/compare/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kilianpaquier/go-builder-generator/internal/generate"
	"github.com/kilianpaquier/go-builder-generator/internal/generate/files"
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
		), files.RwRR)
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

	t.Run("errors_module", func(t *testing.T) {
		// Arrange
		tmp := t.TempDir()

		destdir := filepath.Join(testdata, "result")
		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(tmp, "invalid.go"),
			Structs: []string{},
		}

		testcases := []struct {
			Name        string
			GoMod       string
			ErrContains string
		}{
			{
				Name:        "error_module_statement",
				ErrContains: generate.ErrMissingModule.Error(),
			},
			{
				Name:        "error_go_statement",
				GoMod:       "module test",
				ErrContains: generate.ErrMissingGo.Error(),
			},
			{
				Name: "error_no_file",
				GoMod: `module test
				go 1.22`,
				ErrContains: "parse file",
			},
		}
		for _, tc := range testcases {
			t.Run(tc.Name, func(t *testing.T) {
				// Arrange
				src := filepath.Join(tmp, "go.mod")
				err := os.WriteFile(src, []byte(tc.GoMod), files.RwRR)
				require.NoError(t, err)

				// Act
				err = generate.Run(options, nil)

				// Assert
				assert.ErrorContains(t, err, tc.ErrContains)
				assert.NoDirExists(t, destdir)
			})
		}
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
		CLIOptions generate.CLIOptions
		DirName    string
	}{
		{
			DirName: "success_channels",
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"Chan"},
			},
		},
		{
			DirName: "success_export",
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"Export"},
			},
		},
		{
			DirName: "success_funcs",
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"Func"},
			},
		},
		{
			DirName: "success_generic",
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"Struct", "SimpleGeneric", "AliasGeneric", "ComplexGeneric", "GenericAnonymousStruct", "ComplexSliceGeneric"},
			},
		},
		{
			DirName: "success_interface",
			CLIOptions: generate.CLIOptions{
				NoCMD:   false, // enforce testing at least with one case that the cmd can be printed (and is right) in generated files
				NoTool:  true,  // force go run ...
				Structs: []string{"Interface", "InterfaceNoFields"},
			},
		},
		{
			DirName: "success_maps",
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"Map"},
			},
		},
		{
			DirName: "success_naming",
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"Naming"},
			},
		},
		{
			DirName: "success_root_gomod",
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"RootType"},
			},
		},
		{
			DirName: "success_slices",
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"ArrayAndSlice"},
			},
		},
		{
			DirName: "success_struct",
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"Struct", "StructNoFields"},
			},
		},
		{
			DirName: "success_with_options",
			CLIOptions: generate.CLIOptions{
				NoCMD:        true,
				NoNotice:     true,
				Prefix:       "Set",
				ReturnCopy:   true,
				Structs:      []string{"Options", "Empty", "GenericOptions"},
				ValidateFunc: "Validate",
			},
		},
	} {
		t.Run(tc.DirName, func(t *testing.T) {
			// Arrange
			assertdir := filepath.Join(testdata, tc.DirName, "builders")

			tc.CLIOptions.PackageName = "builders"
			tc.CLIOptions.File = filepath.Join(testdata, tc.DirName, "types.go")
			tc.CLIOptions.Destdir = filepath.Join(testdata, tc.DirName, "result")
			t.Cleanup(func() { require.NoError(t, os.RemoveAll(tc.CLIOptions.Destdir)) })

			// Act
			err := generate.Run(tc.CLIOptions, nil)

			// Assert
			assert.NoError(t, err)
			assert.NoError(t, compare.Dirs(assertdir, tc.CLIOptions.Destdir))
		})
	}
}

func TestRun_ExternalModule(t *testing.T) {
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	for _, tc := range []struct {
		CLIOptions generate.CLIOptions
		DirName    string
	}{
		{
			DirName: "success_module_replace",
			CLIOptions: generate.CLIOptions{
				File:    "module::github.com/stretchr/testify/mock/mock.go",
				NoCMD:   true,
				Structs: []string{"Mock"},
			},
		},
		{
			DirName: "success_module_root",
			CLIOptions: generate.CLIOptions{
				File:    "module::github.com/huandu/xstrings/translate.go",
				NoCMD:   true,
				Structs: []string{"Translator"},
			},
		},
		{
			DirName: "success_module_subdir",
			CLIOptions: generate.CLIOptions{
				File:    "module::github.com/stretchr/testify/mock/mock.go",
				NoCMD:   true,
				Structs: []string{"Mock"},
			},
		},
	} {
		t.Run(tc.DirName, func(t *testing.T) {
			// Arrange
			assertdir := filepath.Join(testdata, tc.DirName, "builders")

			tc.CLIOptions.PackageName = "builders"
			tc.CLIOptions.Destdir = filepath.Join(testdata, tc.DirName, "result")
			t.Cleanup(func() { require.NoError(t, os.RemoveAll(tc.CLIOptions.Destdir)) })

			// Act
			err := generate.Run(tc.CLIOptions, nil)

			// Assert
			assert.NoError(t, err)
			assert.NoError(t, compare.Dirs(assertdir, tc.CLIOptions.Destdir))
		})
	}
}

func TestRun_SamePackage(t *testing.T) {
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	for _, tc := range []struct {
		Args       string
		CLIOptions generate.CLIOptions
		DirName    string
	}{
		{
			DirName: "success_same_package",
			Args:    "-f types.go -s SamePackage,unexportedType",
			CLIOptions: generate.CLIOptions{
				Structs: []string{"SamePackage", "unexportedType"},
			},
		},
		{
			DirName: "success_same_package_options",
			Args:    "-f types.go -s unexportedTypeOptions -p set --package-name unused",
			CLIOptions: generate.CLIOptions{
				Structs:     []string{"unexportedTypeOptions"},
				Prefix:      "Set",
				PackageName: "invalid", // shouldn't be used
			},
		},
	} {
		t.Run(tc.DirName, func(t *testing.T) {
			// Arrange
			assertdir := t.TempDir()
			require.NoError(t, os.CopyFS(assertdir, os.DirFS(filepath.Join(testdata, tc.DirName))))

			tc.CLIOptions.Destdir = filepath.Join(testdata, tc.DirName, "result")
			require.NoError(t, os.MkdirAll(tc.CLIOptions.Destdir, files.RwxRxRxRx)) // the only reason we need to create the directory is because types.go is copied before generation
			t.Cleanup(func() { require.NoError(t, os.RemoveAll(tc.CLIOptions.Destdir)) })

			tc.CLIOptions.File = filepath.Join(tc.CLIOptions.Destdir, "types.go")
			require.NoError(t, files.Copy(filepath.Join(testdata, tc.DirName, "types.go"), tc.CLIOptions.File))

			// Act
			err := generate.Run(tc.CLIOptions, strings.Split(tc.Args, " "))

			// Assert
			assert.NoError(t, err)
			assert.NoError(t, compare.Dirs(assertdir, tc.CLIOptions.Destdir))
		})
	}
}
