package generate

import (
	"errors"
	"fmt"
	"go/ast"
	"os"
	"path/filepath"
	"slices"
	"strings"

	filesystem "github.com/kilianpaquier/filesystem/pkg"
	"golang.org/x/mod/modfile"
)

// parseSrc parses the input src and returns its absolute path.
func parseSrc(src string) (string, error) {
	file := src

	// handle home directory
	if strings.HasPrefix(file, "~") {
		home, _ := os.UserHomeDir()
		file = filepath.Join(home, file[2:])
	}

	// retrieve source file absolute path
	file, err := filepath.Abs(file)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve absolute '%s' path: %w", file, err)
	}

	return file, nil
}

// getImports returns the slice of imports associated to input ast file.
//
// If srcdir and destdir are different, it will search for the first go.mod in parent folders to retrieve the module name.
func getImports(file *ast.File, srcdir, destdir string) ([]string, error) {
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
			return nil, fmt.Errorf("failed to find %s module name: %w", srcdir, err)
		}
		fileImports = append(fileImports, fmt.Sprint(`"`, srcImport, `"`))
	}

	return fileImports, nil
}

// findSourceImport iterates over itself with input src package name to find the full package import path.
//
// Main purpose is to find the first parent go.mod and retrieve its module name to concatenate it with input src string.
func findSourceImport(srcdir string, packages ...string) (string, error) {
	imports := packages
	gomod := filepath.Join(srcdir, "go.mod")

	// go through parent directory to find go.mod in case it doesn't exist in current directory
	if !filesystem.Exists(gomod) {
		if slices.Contains([]string{".", "/"}, srcdir) {
			return "", errors.New("no go.mod found")
		}

		imports := append(imports, filepath.Base(srcdir))
		return findSourceImport(filepath.Dir(srcdir), imports...)
	}

	bytes, err := os.ReadFile(gomod)
	if err != nil {
		return "", fmt.Errorf("failed to read go.mod: %w", err)
	}
	file, err := modfile.Parse(gomod, bytes, nil)
	if err != nil {
		return "", fmt.Errorf("failed to parse go.mod: %w", err)
	}
	if file.Module == nil {
		return "", errors.New("invalid go.mod, module statement is missing")
	}

	if file.Module.Mod.Path != "std" { // specific exclusion for builtin
		imports = append(imports, file.Module.Mod.Path)
	}

	slices.Reverse(imports)
	return strings.Join(imports, "/"), nil
}
