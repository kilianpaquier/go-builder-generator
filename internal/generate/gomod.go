package generate

import (
	"errors"
	"fmt"
	"go/ast"
	"go/version"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/samber/lo"
	"golang.org/x/mod/modfile"

	"github.com/kilianpaquier/go-builder-generator/internal/generate/files"
)

var (
	// ErrMissingGo indicates that 'go' statement is missing in a go.mod.
	ErrMissingGo = errors.New("missing go statement")

	// ErrMissingModule indicates that 'module' statement is missing in a go.mod.
	ErrMissingModule = errors.New("missing module statement")

	// ErrNilMod indicates that modfile.File is nil (should never happen).
	ErrNilMod = errors.New("nil go.mod")
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

	if modcache := os.Getenv("GOMODCACHE"); modcache != "" {
		return filepath.Join(modcache, module), nil
	}

	if gopath := os.Getenv("GOPATH"); gopath != "" {
		return filepath.Join(gopath, "pkg", "mod", module), nil
	}

	home, _ := os.UserHomeDir()
	return filepath.Join(home, "go", "pkg", "mod", module), nil
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

// findRequire finds the appropriate module name and version in input file for the input moduleName.
func findRequire(file *modfile.File, moduleName string) (string, error) {
	// find appropriate require in go.mod file to retrieve the version
	require, ok := lo.Find(file.Require, func(require *modfile.Require) bool {
		return strings.HasPrefix(moduleName, require.Mod.Path)
	})
	if !ok {
		return "", fmt.Errorf("missing module name '%s' in go.mod", moduleName)
	}
	subpath := strings.TrimPrefix(moduleName, require.Mod.Path)

	// find the appropriate replace version in go.mod (in such case it could exist)
	replace, ok := lo.Find(file.Replace, func(replace *modfile.Replace) bool {
		return strings.HasPrefix(moduleName, replace.Old.Path)
	})
	if !ok {
		// return required version if no replaced version found
		return filepath.Join(require.Mod.String(), subpath), nil
	}
	// return replaced version if provided
	return filepath.Join(replace.New.String(), subpath), nil
}

// findGomod finds the parent go.mod associated to input dir
// and returns the parsed modfile alongside the path between the go.mod and the input dir.
func findGomod(dir string, parts ...string) (*modfile.File, string, error) {
	mod := filepath.Join(dir, "go.mod")

	// go through parent directory to find go.mod in case it doesn't exist in current directory
	if !files.Exists(mod) {
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

	if err := validGomod(file); err != nil {
		return nil, "", fmt.Errorf("invalid '%s' go.mod: %w", mod, err)
	}

	slices.Reverse(parts)
	return file, strings.Join(parts, "/"), nil
}

// validGomod ensures that used properties in generation process aren't nil.
func validGomod(file *modfile.File) error {
	if file == nil {
		return ErrNilMod
	}
	if file.Module == nil {
		return ErrMissingModule
	}
	if file.Go == nil {
		return ErrMissingGo
	}
	return nil
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

// toolAvailable returns truthy when given modfile go version is above or equal to go1.24.
// Meaning it can handle go tool section.
func toolAvailable(file *modfile.File) bool {
	minGo := "go1.24"
	return version.Compare("go"+file.Go.Version, minGo) >= 0
}

// hasGenerate checks whether a 'go:generate' comment is present in input file for go-builder-generator.
func hasGenerate(file *ast.File, args []string) bool {
	options := regexp.QuoteMeta(strings.Join(args, " "))

	rexps := []*regexp.Regexp{
		regexp.MustCompile(fmt.Sprint(`^//go:generate ([^\s]+)?go-builder-generator generate `, options, "$")),
		regexp.MustCompile(fmt.Sprint(`^//go:generate go run github\.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@[^\s]+ generate `, options, "$")),
		regexp.MustCompile(fmt.Sprint(`^//go:generate go tool go-builder-generator generate `, options, "$")),
	}

	for _, group := range file.Comments {
		if group == nil {
			continue
		}

		for _, comment := range group.List {
			if comment == nil {
				continue
			}

			if _, ok := lo.Find(rexps, func(reg *regexp.Regexp) bool { return reg.MatchString(comment.Text) }); ok {
				return true
			}
		}
	}

	return false
}
