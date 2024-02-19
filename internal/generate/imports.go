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

// getImports returns the slice of imports associated to input ast file.
//
// If srcdir and destdir are different, it will search for the first go.mod in parent folders to retrieve the module name.
func getImports(file *ast.File, srcdir, destdir string) ([]string, error) {
	// get file imports as string
	fileImports := make([]string, 0, len(file.Imports))
	for _, item := range file.Imports {
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
func findSourceImport(src string, packages ...string) (string, error) {
	dir := filepath.Dir(src)
	gomod := filepath.Join(dir, "go.mod")

	packages = append(packages, filepath.Base(src)) // nolint:revive
	if !filesystem.Exists(gomod) {
		return findSourceImport(dir, packages...)
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

	packages = append(packages, file.Module.Mod.Path) // nolint:revive
	slices.Reverse(packages)
	return strings.Join(packages, "/"), nil
}
