package generate_test

import (
	"os"
	"path"
	"path/filepath"
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
		destdir := t.TempDir()
		require.NoError(t, os.WriteFile(filepath.Join(destdir, "go.mod"), []byte("module test\ngo 1.22"), files.RwRR))

		srcdir := t.TempDir()
		require.NoError(t, os.CopyFS(srcdir, os.DirFS(filepath.Join(testdata, "errors", "no_gomod"))))

		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(srcdir, "types.go"),
			Structs: []string{"NoGomod"},
		}

		// Act
		err := generate.Run(options, nil)

		// Assert
		assert.ErrorContains(t, err, "find src go.mod: no parent go.mod found")
	})

	t.Run("error_no_dest_gomod", func(t *testing.T) {
		// Act
		err := generate.Run(generate.CLIOptions{Destdir: t.TempDir()}, nil)

		// Assert
		assert.ErrorContains(t, err, "find dest go.mod: no parent go.mod found")
	})

	t.Run("error_no_module_statement", func(t *testing.T) {
		// Arrange
		destdir := t.TempDir()
		file, err := os.Create(filepath.Join(destdir, "go.mod"))
		require.NoError(t, err)
		require.NoError(t, file.Close())

		// Act
		err = generate.Run(generate.CLIOptions{Destdir: destdir}, nil)

		// Assert
		assert.ErrorIs(t, err, generate.ErrMissingModule)
	})

	t.Run("error_no_go_statement", func(t *testing.T) {
		// Arrange
		destdir := t.TempDir()
		require.NoError(t, os.WriteFile(filepath.Join(destdir, "go.mod"), []byte("module test"), files.RwRR))

		// Act
		err := generate.Run(generate.CLIOptions{Destdir: destdir}, nil)

		// Assert
		assert.ErrorIs(t, err, generate.ErrMissingGo)
	})

	t.Run("error_missing_struct_file", func(t *testing.T) {
		// Arrange
		destdir := t.TempDir()
		require.NoError(t, os.WriteFile(filepath.Join(destdir, "go.mod"), []byte("module test\ngo 1.22"), files.RwRR))

		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(destdir, "missing.go"),
			Structs: []string{},
		}

		// Act
		err := generate.Run(options, nil)

		// Assert
		assert.ErrorContains(t, err, "parse file")
	})

	t.Run("error_not_required_module", func(t *testing.T) {
		// Arrange
		destdir := t.TempDir()
		require.NoError(t, os.WriteFile(filepath.Join(destdir, "go.mod"), []byte("module test\ngo 1.22"), files.RwRR))

		options := generate.CLIOptions{
			Destdir: destdir,
			File:    "module::github.com/jarcoal/httpmock/match.go",
			Structs: []string{"Matcher"},
		}

		// Act
		err := generate.Run(options, nil)

		// Assert
		assert.ErrorContains(t, err, "missing module name 'github.com/jarcoal/httpmock/match.go")
	})

	t.Run("error_invalid", func(t *testing.T) {
		// Arrange
		destdir := t.TempDir()
		srcdir := filepath.Join(testdata, "errors", "invalids")
		require.NoError(t, os.CopyFS(destdir, os.DirFS(srcdir)))

		type testcase struct {
			ErrContains []string
			Name        string
		}
		cases := []testcase{
			{
				Name:        "InvalidTag",
				ErrContains: []string{"field '[Flag]' options parsing: tags parsing: bad syntax for struct tag value"},
			},
			{
				Name: "InvalidOption",
				ErrContains: []string{
					"field '[SimpleOption]' options parsing: unknown option 'invalid_option'",
					"field '[EqualOption]' options parsing: unknown option 'equal_option=prop'",
				},
			},
			{
				Name:        "unexported",
				ErrContains: []string{"is not exported (or one of its generic params is not) but generation destination is in an external package"},
			},
		}
		for _, tc := range cases {
			t.Run(tc.Name, func(t *testing.T) {
				// Arrange
				options := generate.CLIOptions{
					Destdir: filepath.Join(destdir, "builders"),
					File:    filepath.Join(destdir, "types.go"),
					Structs: []string{tc.Name},
				}

				// Act
				err := generate.Run(options, nil)

				// Assert
				for _, contain := range tc.ErrContains {
					assert.ErrorContains(t, err, contain)
				}
			})
		}
	})
}

func TestRun_Types(t *testing.T) {
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	for _, tc := range []struct {
		CLIOptions generate.CLIOptions
		DirName    string
	}{
		{
			DirName: path.Join("success_types", "channels"),
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"Chan"},
			},
		},
		{
			DirName: path.Join("success_types", "exports"),
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"Export"},
			},
		},
		{
			DirName: path.Join("success_types", "funcs"),
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"Func"},
			},
		},
		{
			DirName: path.Join("success_types", "generics"),
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"Struct", "SimpleGeneric", "AliasGeneric", "ComplexGeneric", "GenericAnonymousStruct", "ComplexSliceGeneric"},
			},
		},
		{
			DirName: path.Join("success_types", "interfaces"),
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"Interface", "InterfaceNoFields"},
			},
		},
		{
			DirName: path.Join("success_types", "maps"),
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
			DirName: path.Join("success_tool", "tool"),
			CLIOptions: generate.CLIOptions{
				Structs: []string{"SimpleTypeTool"},
			},
		},
		{
			DirName: path.Join("success_tool", "no_tool"),
			CLIOptions: generate.CLIOptions{
				Structs: []string{"SimpleTypeNoTool"},
			},
		},
		{
			DirName: path.Join("success_types", "slices"),
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"ArrayAndSlice"},
			},
		},
		{
			DirName: path.Join("success_types", "structs"),
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"Struct", "StructNoFields"},
			},
		},
		{
			DirName: "success_options",
			CLIOptions: generate.CLIOptions{
				NoNotice:     true,
				PackageName:  "my_package",
				Prefix:       "Set",
				ReturnCopy:   true,
				Structs:      []string{"Options", "Empty", "GenericOptions"},
				ValidateFunc: "Validate",
			},
		},
	} {
		t.Run(tc.DirName, func(t *testing.T) {
			// Arrange
			srcdir := filepath.Join(testdata, tc.DirName)
			assertdir := filepath.Join(srcdir, "builders")

			destdir := t.TempDir()
			for _, file := range []string{"go.mod", "types.go"} {
				err := files.Copy(filepath.Join(srcdir, file), filepath.Join(destdir, file))
				require.NoError(t, err)
			}
			tc.CLIOptions.Destdir = filepath.Join(destdir, "builders")
			tc.CLIOptions.File = filepath.Join(destdir, "types.go")

			// Act
			err := generate.Run(tc.CLIOptions, nil)

			// Assert
			assert.NoError(t, err)
			assert.NoError(t, compare.Dirs(assertdir, tc.CLIOptions.Destdir))
		})
	}
}

func TestRun_Module(t *testing.T) {
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	for _, tc := range []struct {
		CLIOptions generate.CLIOptions
		DirName    string
	}{
		{
			DirName: path.Join("success_module", "replace"),
			CLIOptions: generate.CLIOptions{
				File:    "module::github.com/stretchr/testify/mock/mock.go",
				NoCMD:   true,
				Structs: []string{"Mock"},
			},
		},
		{
			DirName: path.Join("success_module", "basedirectory"),
			CLIOptions: generate.CLIOptions{
				File:    "module::github.com/huandu/xstrings/translate.go",
				NoCMD:   true,
				Structs: []string{"Translator"},
			},
		},
		{
			DirName: path.Join("success_module", "subdirectory"),
			CLIOptions: generate.CLIOptions{
				File:    "module::github.com/stretchr/testify/mock/mock.go",
				NoCMD:   true,
				Structs: []string{"Mock"},
			},
		},
	} {
		t.Run(tc.DirName, func(t *testing.T) {
			// Arrange
			srcdir := filepath.Join(testdata, tc.DirName)
			assertdir := filepath.Join(srcdir, "builders")

			destdir := t.TempDir()
			for _, file := range []string{"go.mod", "go.sum", "types.go"} {
				err := files.Copy(filepath.Join(srcdir, file), filepath.Join(destdir, file))
				require.NoError(t, err)
			}
			tc.CLIOptions.Destdir = filepath.Join(destdir, "builders")

			// Act
			err := generate.Run(tc.CLIOptions, nil)

			// Assert
			assert.NoError(t, err)
			assert.NoError(t, compare.Dirs(assertdir, tc.CLIOptions.Destdir))
		})
	}
}

func TestRun_Package(t *testing.T) {
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	for _, tc := range []struct {
		CLIOptions generate.CLIOptions
		DirName    string
	}{
		{
			DirName: path.Join("success_package", "same"),
			CLIOptions: generate.CLIOptions{
				Structs: []string{"SamePackage", "unexportedType"},
			},
		},
		{
			DirName: path.Join("success_package", "same_options"),
			CLIOptions: generate.CLIOptions{
				Structs:     []string{"unexportedTypeOptions"},
				Prefix:      "set",
				PackageName: "unused",
			},
		},
	} {
		t.Run(tc.DirName, func(t *testing.T) {
			// Arrange
			assertdir := filepath.Join(testdata, tc.DirName)

			destdir := t.TempDir()
			tc.CLIOptions.Destdir = destdir
			for _, file := range []string{"go.mod", "types.go"} {
				err := files.Copy(filepath.Join(assertdir, file), filepath.Join(destdir, file))
				require.NoError(t, err)
			}
			tc.CLIOptions.File = filepath.Join(destdir, "types.go")

			// Act
			err := generate.Run(tc.CLIOptions, tc.CLIOptions.ToArgs("")[2:])

			// Assert
			assert.NoError(t, err)
			assert.NoError(t, compare.Dirs(assertdir, tc.CLIOptions.Destdir))
		})
	}
}
