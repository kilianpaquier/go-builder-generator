package generate

import (
	"errors"
	"fmt"
	"go/ast"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/samber/lo"
	"golang.org/x/mod/modfile"

	"github.com/kilianpaquier/go-builder-generator/internal/fs"
)

const modulePrefix = "module::"

// modulePath finds the appropriate required in modfile for the input module name
// and returns its path in the current filesystem.
func modulePath(file *modfile.File, moduleName string) (string, error) {
	require := strings.TrimPrefix(moduleName, modulePrefix)

	module, err := findRequire(file, require)
	if err != nil {
		return "", err
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		home, _ := os.UserHomeDir()
		gopath = filepath.Join(home, "go")
	}
	gopkg := filepath.Join(gopath, "pkg", "mod")

	return filepath.Join(gopkg, module.ModulePath, module.SubPath), nil
}

// getImports returns the slice of imports associated to input ast file.
//
// If srcdir and destdir are different, it will search for the first go.mod in parent folders to retrieve the module name.
func getImports(file *ast.File) ([]string, error) {
	// get file imports as string
	imports := make([]string, 0, len(file.Imports))
	for _, item := range file.Imports {
		if item.Name != nil {
			imports = append(imports, fmt.Sprint(item.Name.Name, " ", item.Path.Value))
			continue
		}
		imports = append(imports, item.Path.Value)
	}
	return imports, nil
}

type module struct {
	ModulePath string
	SubPath    string
}

// findRequire finds the appropriate module name and version in input file for the input moduleName.
func findRequire(file *modfile.File, moduleName string) (module, error) {
	// find appropriate require in go.mod file to retrieve the version
	require, ok := lo.Find(file.Require, func(require *modfile.Require) bool {
		return strings.HasPrefix(moduleName, require.Mod.Path)
	})
	if !ok {
		return module{}, fmt.Errorf("missing module name '%s' in go.mod", moduleName)
	}
	subpath := strings.TrimPrefix(moduleName, require.Mod.Path)

	// find the appropriate replace version in go.mod (in such case it could exist)
	replace, ok := lo.Find(file.Replace, func(replace *modfile.Replace) bool {
		return strings.HasPrefix(moduleName, replace.Old.Path)
	})
	if !ok {
		// return required version if no replaced version found
		return module{
			ModulePath: require.Mod.String(),
			SubPath:    subpath,
		}, nil
	}

	// return replaced version if provided
	return module{
		ModulePath: replace.New.String(),
		SubPath:    subpath,
	}, nil
}

// findGomod finds the parent go.mod associated to input dir
// and returns the parsed modfile alongside the path between the go.mod and the input dir.
func findGomod(dir string, parts ...string) (*modfile.File, string, error) {
	mod := filepath.Join(dir, "go.mod")

	// go through parent directory to find go.mod in case it doesn't exist in current directory
	if !fs.Exists(mod) {
		// handle root directory -> VolumeName (e.g "C:") + os.PathSeparator
		if dir == filepath.VolumeName(dir)+string(os.PathSeparator) {
			return nil, "", errors.New("no parent go.mod found")
		}
		return findGomod(filepath.Dir(dir), append(parts, filepath.Base(dir))...)
	}
	bytes, _ := os.ReadFile(mod) // don't need to handle error since file exists

	file, err := modfile.Parse(mod, bytes, nil)
	if err != nil {
		return nil, "", fmt.Errorf("go.mod '%s' parsing: %w", mod, err)
	}

	slices.Reverse(parts)
	return file, strings.Join(parts, "/"), nil
}

// fileImport returns the import for the input pkg with the associated go.mod modfile.
func fileImport(file *modfile.File, pkg string) string {
	m := lo.FromPtr(file.Module)

	// when working with std package, go doesn't add the go.mod module name to the import
	if m.Mod.Path == "std" {
		return pkg
	}

	// when working with a standard go module, the package name is prefixed with the module name from where it comes
	i := path.Join(m.Mod.Path, pkg)
	return fmt.Sprint(`"`, i, `"`)
}

// hasGenerate checks whether a 'go:generate' comment is present in input file for go-builder-generator.
func hasGenerate(file *ast.File, rawArgs []string) bool {
	rexps := []*regexp.Regexp{
		regexp.MustCompile(`//go:generate go run github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@[^\s]+ generate ` + strings.Join(rawArgs, " ")),
		regexp.MustCompile(`//go:generate ([^\s]+)?go-builder-generator generate ` + strings.Join(rawArgs, " ")),
	}

	for _, group := range file.Comments {
		if group == nil {
			continue
		}

		for _, comment := range group.List {
			if comment == nil {
				continue
			}

			_, ok := lo.Find(rexps, func(reg *regexp.Regexp) bool { return reg.MatchString(comment.Text) })
			if ok {
				return true
			}
		}
	}

	return false
}
