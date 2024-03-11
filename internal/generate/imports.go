package generate

import (
	"errors"
	"fmt"
	"go/ast"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/hashicorp/go-getter"
	"github.com/samber/lo"
	"golang.org/x/mod/modfile"
)

const git = "git::"

// parseSrc parses the input src and returns its absolute path.
func parseSrc(src string) (string, error) {
	file := src

	// handle home relative paths
	if strings.HasPrefix(file, "~") {
		home, _ := os.UserHomeDir()
		file = filepath.Join(home, file[1:])
	}

	// handle git repositories files
	if strings.HasPrefix(file, git) {
		u := strings.TrimPrefix(file, git)

		// parse file as URL since it's a remote file
		initial, err := url.Parse(u)
		if err != nil {
			return "", fmt.Errorf("failed to parse url '%s': %w", file, err)
		}
		remote, _ := url.Parse(initial.String()) // ignoring error since it's sure initial is an URL
		remote.Path = path.Dir(remote.Path)      // remove filename from URL

		// download remote git repository into temporary dir update file to read
		destdir := filepath.Join(os.TempDir(), remote.Path)
		if err := getter.Get(destdir, git+remote.String()); err != nil {
			return "", fmt.Errorf("failed to download git repository '%s': %w", remote.String(), err)
		}
		file = filepath.Join(destdir, filepath.Base(initial.Path))
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
	if _, err := os.Stat(gomod); err != nil {
		if !os.IsNotExist(err) {
			return "", fmt.Errorf("go.mod seems to exist but is not readable: %w", err)
		}

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
		return "", fmt.Errorf("failed to parse go.mod '%s': %w", gomod, err)
	}
	module := lo.FromPtr(file.Module)

	// specific exclusion for builtin
	if module.Mod.Path != "std" {
		imports = append(imports, module.Mod.Path)
	}
	slices.Reverse(imports)
	return strings.Join(imports, "/"), nil
}
