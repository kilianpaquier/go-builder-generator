package generate_test

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	compare "github.com/kilianpaquier/compare/pkg"
	"github.com/samber/lo"

	"github.com/kilianpaquier/go-builder-generator/internal/generate"
	"github.com/kilianpaquier/go-builder-generator/internal/generate/files"
)

func TestRun_Errors(t *testing.T) {
	ctx := t.Context()
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	t.Run("error_no_src_gomod", func(t *testing.T) {
		// Arrange
		destdir := t.TempDir()
		NoError(t, os.WriteFile(filepath.Join(destdir, "go.mod"), []byte("module test\ngo 1.22"), files.RwRR))

		srcdir := t.TempDir()
		NoError(t, os.CopyFS(srcdir, os.DirFS(filepath.Join(testdata, "errors", "no_gomod"))))

		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(srcdir, "types.go"),
			Structs: []string{"NoGomod"},
		}

		// Act
		err := generate.Run(ctx, options, nil)

		// Assert
		ErrorContains(t, err, "find src go.mod: no parent go.mod found")
	})

	t.Run("error_no_dest_gomod", func(t *testing.T) {
		// Act
		err := generate.Run(ctx, generate.CLIOptions{Destdir: t.TempDir()}, nil)

		// Assert
		ErrorContains(t, err, "find dest go.mod: no parent go.mod found")
	})

	t.Run("error_no_module_statement", func(t *testing.T) {
		// Arrange
		destdir := t.TempDir()
		file, err := os.Create(filepath.Join(destdir, "go.mod"))
		NoError(t, err)
		NoError(t, file.Close())

		// Act
		err = generate.Run(ctx, generate.CLIOptions{Destdir: destdir}, nil)

		// Assert
		ErrorIs(t, err, generate.ErrMissingModule)
	})

	t.Run("error_no_go_statement", func(t *testing.T) {
		// Arrange
		destdir := t.TempDir()
		NoError(t, os.WriteFile(filepath.Join(destdir, "go.mod"), []byte("module test"), files.RwRR))

		// Act
		err := generate.Run(ctx, generate.CLIOptions{Destdir: destdir}, nil)

		// Assert
		ErrorIs(t, err, generate.ErrMissingGo)
	})

	t.Run("error_missing_struct_file", func(t *testing.T) {
		// Arrange
		destdir := t.TempDir()
		NoError(t, os.WriteFile(filepath.Join(destdir, "go.mod"), []byte("module test\ngo 1.22"), files.RwRR))

		options := generate.CLIOptions{
			Destdir: destdir,
			File:    filepath.Join(destdir, "missing.go"),
			Structs: []string{},
		}

		// Act
		err := generate.Run(ctx, options, nil)

		// Assert
		ErrorContains(t, err, "parse file")
	})

	t.Run("error_not_required_module", func(t *testing.T) {
		// Arrange
		destdir := t.TempDir()
		NoError(t, os.WriteFile(filepath.Join(destdir, "go.mod"), []byte("module test\ngo 1.22"), files.RwRR))

		options := generate.CLIOptions{
			Destdir: destdir,
			File:    "module::github.com/jarcoal/httpmock/match.go",
			Structs: []string{"Matcher"},
		}

		// Act
		err := generate.Run(ctx, options, nil)

		// Assert
		ErrorContains(t, err, "missing module name 'github.com/jarcoal/httpmock/match.go")
	})

	t.Run("error_invalid", func(t *testing.T) {
		// Arrange
		destdir := t.TempDir()
		srcdir := filepath.Join(testdata, "errors", "invalids")
		NoError(t, os.CopyFS(destdir, os.DirFS(srcdir)))

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
				err := generate.Run(ctx, options, nil)

				// Assert
				for _, contain := range tc.ErrContains {
					ErrorContains(t, err, contain)
				}
			})
		}
	})
}

func TestRun_GeneratedAnotherPackage(t *testing.T) {
	ctx := t.Context()
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	for _, tc := range []struct {
		CLIOptions   generate.CLIOptions
		Dependencies []string
		DirName      string
	}{
		{
			DirName: path.Join("success_types", "builtin"),
			CLIOptions: generate.CLIOptions{
				File:    filepath.Join("pkg", "types.go"),
				NoCMD:   true,
				Structs: []string{"Builtin"},
			},
		},
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
			DirName: path.Join("success_imports", "std"),
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"STD"},
			},
		},
		{
			DirName:      path.Join("success_imports", "local_dir"),
			Dependencies: []string{"dir"},
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Structs: []string{"LocalImport"},
			},
		},
		{
			DirName:      path.Join("success_imports", "local_same_dir"),
			Dependencies: []string{"dir"},
			CLIOptions: generate.CLIOptions{
				NoCMD:   true,
				Destdir: "local_same_dir",
				Structs: []string{"LocalImport"},
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
			types := tc.CLIOptions.File
			if types == "" {
				types = "types.go"
			}
			srcdir := filepath.Join(testdata, tc.DirName)
			assertdir := filepath.Join(srcdir, filepath.Dir(types), lo.CoalesceOrEmpty(tc.CLIOptions.Destdir, "builders"))

			destdir := t.TempDir()
			for _, file := range []string{"go.mod", types} {
				NoError(t, os.MkdirAll(filepath.Join(destdir, filepath.Dir(file)), files.RwxRxRxRx))
				err := CopyFile(filepath.Join(srcdir, file), filepath.Join(destdir, file))
				NoError(t, err)
			}
			for _, dependency := range tc.Dependencies {
				NoError(t, os.CopyFS(filepath.Join(destdir, dependency), os.DirFS(filepath.Join(srcdir, dependency))))
			}
			tc.CLIOptions.Destdir = filepath.Join(destdir, filepath.Dir(types), lo.CoalesceOrEmpty(tc.CLIOptions.Destdir, "builders"))
			tc.CLIOptions.File = filepath.Join(destdir, types)

			// Act
			err := generate.Run(ctx, tc.CLIOptions, nil)

			// Assert
			NoError(t, err)
			NoError(t, compare.Dirs(assertdir, tc.CLIOptions.Destdir))
		})
	}
}

func TestRun_GeneratedFromAnotherModule(t *testing.T) {
	ctx := t.Context()
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	for _, tc := range []struct {
		CLIOptions generate.CLIOptions
		DirName    string
	}{
		{
			DirName: path.Join("success_module", "replace"),
			CLIOptions: generate.CLIOptions{
				File:    "module::github.com/spf13/cobra/command.go",
				NoCMD:   true,
				Structs: []string{"Command"},
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
			DirName: path.Join("success_module", "std"),
			CLIOptions: generate.CLIOptions{
				File:    "std::go/build/build.go",
				Structs: []string{"Context"},
			},
		},
		{
			DirName: path.Join("success_module", "subdirectory"),
			CLIOptions: generate.CLIOptions{
				File:    "module::golang.org/x/mod/modfile/rule.go",
				Structs: []string{"File"},
			},
		},
	} {
		t.Run(tc.DirName, func(t *testing.T) {
			// Arrange
			srcdir := filepath.Join(testdata, tc.DirName)
			assertdir := filepath.Join(srcdir, "builders")

			destdir := t.TempDir()
			for _, file := range []string{"go.mod", "go.sum"} {
				err := CopyFile(filepath.Join(srcdir, file), filepath.Join(destdir, file))
				NoError(t, err)
			}
			tc.CLIOptions.Destdir = filepath.Join(destdir, "builders")

			// Act
			err := generate.Run(ctx, tc.CLIOptions, nil)

			// Assert
			NoError(t, err)
			NoError(t, compare.Dirs(assertdir, tc.CLIOptions.Destdir))
		})
	}
}

func TestRun_GeneratedSamePackage(t *testing.T) {
	ctx := t.Context()
	pwd, _ := os.Getwd()
	testdata := filepath.Join(pwd, "..", "..", "testdata")

	for _, tc := range []struct {
		CLIOptions   generate.CLIOptions
		Dependencies []string
		DirName      string
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
		{
			DirName:      path.Join("success_imports", "local"),
			Dependencies: []string{"dir"},
			CLIOptions: generate.CLIOptions{
				Structs: []string{"LocalImport"},
			},
		},
	} {
		t.Run(tc.DirName, func(t *testing.T) {
			// Arrange
			assertdir := filepath.Join(testdata, tc.DirName)

			destdir := t.TempDir()
			tc.CLIOptions.Destdir = destdir
			for _, file := range []string{"go.mod", "types.go"} {
				err := CopyFile(filepath.Join(assertdir, file), filepath.Join(destdir, file))
				NoError(t, err)
			}
			for _, dependency := range tc.Dependencies {
				NoError(t, os.CopyFS(filepath.Join(destdir, dependency), os.DirFS(filepath.Join(assertdir, dependency))))
			}
			tc.CLIOptions.File = filepath.Join(destdir, "types.go")

			// Act
			err := generate.Run(ctx, tc.CLIOptions, tc.CLIOptions.ToArgs("")[2:])

			// Assert
			NoError(t, err)
			NoError(t, compare.Dirs(assertdir, tc.CLIOptions.Destdir))
		})
	}
}

// CopyFile copies a provided file from src to dest with a default permission of 0o644. It fails if it's a directory.
func CopyFile(src, dst string) error {
	sfile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer sfile.Close()

	dfile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer dfile.Close()

	// copy buffer from src to dest
	if _, err := io.Copy(dfile, sfile); err != nil {
		return fmt.Errorf("copy: %w", err)
	}

	// update dest permissions
	if err := dfile.Chmod(files.RwRR); err != nil {
		return fmt.Errorf("chmod: %w", err)
	}
	return nil
}

// NoError fails t in case input err is not nil.
func NoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

// ErrorContains fails t in case input err is nil or doesn't contain the expected contains.
func ErrorContains(t testing.TB, err error, contains string) {
	t.Helper()

	if err == nil {
		t.Error("error is nil")
		t.FailNow()
	}
	if !strings.Contains(err.Error(), contains) {
		t.Errorf("error '%s' doesn't contain '%s'", err.Error(), contains)
		t.FailNow()
	}
}

// ErrorIs fails t in case input err doesn't validate errors.Is for expected is error.
func ErrorIs(t testing.TB, err, is error) {
	t.Helper()

	if !errors.Is(err, is) {
		t.Errorf("error '%s' doesn't validate 'errors.Is' for '%s'", err.Error(), is.Error())
		t.FailNow()
	}
}
