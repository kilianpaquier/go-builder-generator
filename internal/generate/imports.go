package generate

import (
	"errors"
	"fmt"
	"go/ast"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/samber/lo"
	"golang.org/x/mod/modfile"

	"github.com/kilianpaquier/go-builder-generator/internal/fs"
)

const modulePrefix = "module::"

// parseSrc parses the input dest and returns its absolute path.
func parseSrc(pwd, dest string) (string, error) {
	file := dest

	if strings.HasPrefix(file, modulePrefix) {
		noPrefix := strings.TrimPrefix(file, modulePrefix)

		module, err := findModule(pwd, noPrefix)
		if err != nil {
			return "", err // err is wrapped in subfunction
		}

		gopath := func() string {
			gopath := os.Getenv("GOPATH")
			if gopath != "" {
				return gopath
			}
			home, _ := os.UserHomeDir()
			return filepath.Join(home, "go")
		}()
		gopkg := filepath.Join(gopath, "pkg", "mod")

		file = filepath.Join(gopkg, module.ModulePath, module.SubPath)
	}

	// retrieve source file absolute path
	file, err := filepath.Abs(file)
	if err != nil {
		return "", fmt.Errorf("absolute '%s' path: %w", file, err)
	}
	return file, nil
}

// parseImports returns the slice of imports associated to input ast file.
//
// If srcdir and destdir are different, it will search for the first go.mod in parent folders to retrieve the module name.
func parseImports(file *ast.File, srcdir, destdir string) ([]string, error) {
	// get file imports as string
	fileImports := make([]string, 0, len(file.Imports))
	for _, item := range file.Imports {
		if item.Name != nil {
			fileImports = append(fileImports, fmt.Sprint(item.Name.Name, " ", item.Path.Value))
			continue
		}
		fileImports = append(fileImports, item.Path.Value)
	}

	// check if destination is the same as src
	if srcdir != destdir {
		// find source package path to add it as import in builder package
		srcImport, err := findSourceImport(srcdir)
		if err != nil {
			return nil, fmt.Errorf("find imports: %w", err)
		}
		fileImports = append(fileImports, fmt.Sprint(`"`, srcImport, `"`))
	}

	return fileImports, nil
}

type module struct {
	ModulePath string
	SubPath    string
}

// findModule returns the associated module in srcdir go.mod (or parent(s) directory) for dest.
//
// dest can be the direct module name or any subdirectory (or subfiles) of it (github.com/kilianpaquier/go-builder-generator/internal/generate/types.go).
//
// An error is returned in such case where no go.mod is found or where dest is not in the go.mod require imports.
func findModule(srcdir, dest string) (*module, error) {
	gomod := filepath.Join(srcdir, "go.mod")
	if !fs.Exists(gomod) {
		// handle root directory -> VolumeName (e.g "C:") + os.PathSeparator
		if srcdir == filepath.VolumeName(srcdir)+string(os.PathSeparator) {
			return nil, errors.New("current module go.mod not found")
		}
		return findModule(filepath.Dir(srcdir), dest)
	}
	bytes, _ := os.ReadFile(gomod) // don't need to handle error since file exists

	file, err := modfile.Parse(gomod, bytes, nil)
	if err != nil {
		return nil, fmt.Errorf("go.mod '%s' parsing: %w", gomod, err)
	}

	// find appropriate require in go.mod file to retrieve the version
	require, ok := lo.Find(file.Require, func(require *modfile.Require) bool {
		return strings.HasPrefix(dest, require.Mod.Path)
	})
	if !ok {
		return nil, fmt.Errorf("failed to find appropriate require in '%s', make sure you are importing base module of '%s'", gomod, dest)
	}
	subpath := strings.TrimPrefix(dest, require.Mod.Path)

	// find the appropriate replace version in go.mod (in such case it could exist)
	replace, ok := lo.Find(file.Replace, func(replace *modfile.Replace) bool {
		return strings.HasPrefix(dest, replace.Old.Path)
	})
	if !ok {
		// return required version if no replaced version found
		return &module{
			ModulePath: require.Mod.String(),
			SubPath:    subpath,
		}, nil
	}

	// return replaced version if provided
	return &module{
		ModulePath: replace.New.String(),
		SubPath:    subpath,
	}, nil
}

// findSourceImport iterates over itself with input src package name to find the full package import path.
//
// Main purpose is to find the first parent go.mod and retrieve its module name to concatenate it with input src string.
func findSourceImport(srcdir string, packages ...string) (string, error) {
	imports := packages
	gomod := filepath.Join(srcdir, "go.mod")

	// go through parent directory to find go.mod in case it doesn't exist in current directory
	if !fs.Exists(gomod) {
		// handle root directory -> VolumeName (e.g "C:") + os.PathSeparator
		if srcdir == filepath.VolumeName(srcdir)+string(os.PathSeparator) {
			return "", errors.New("no go.mod found")
		}
		imports := append(imports, filepath.Base(srcdir))
		return findSourceImport(filepath.Dir(srcdir), imports...)
	}
	bytes, _ := os.ReadFile(gomod) // don't need to handle error since file exists

	file, err := modfile.Parse(gomod, bytes, nil)
	if err != nil {
		return "", fmt.Errorf("go.mod '%s' parsing: %w", gomod, err)
	}
	module := lo.FromPtr(file.Module)

	// specific exclusion for builtin
	if module.Mod.Path != "std" {
		imports = append(imports, module.Mod.Path)
	}
	slices.Reverse(imports)
	return strings.Join(imports, "/"), nil
}
