package generate

import (
	"context"
	"errors"
	"fmt"
	"go/ast"
	"go/version"
	"os"
	"os/exec"
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

const (
	modulePrefix = "module::"
	stdPrefix    = "std::"
)

// modFile is an enriched modfile.File with its local directory path.
type modFile struct {
	*modfile.File
	Dir string
}

// modulePath finds the appropriate required in modfile for the input module name
// and returns its path in the current filesystem.
func modulePath(ctx context.Context, file modFile, modulepathfile string) (string, error) {
	// find required in go/src since it's prefixed by "std::"
	if require, ok := strings.CutPrefix(modulepathfile, stdPrefix); ok {
		if root := os.Getenv("GOROOT"); root != "" {
			return filepath.Join(root, "src", require), nil
		}

		// runtime.GOROOT() being deprecated, the only way to get GOROOT is by retrieving it from 'go env' command
		output, err := exec.CommandContext(ctx, "go", "env", "GOROOT").CombinedOutput()
		if err == nil {
			return filepath.Join(strings.TrimSpace(string(output)), "src", require), nil
		}
		if len(output) > 0 {
			return "", fmt.Errorf("get 'GOROOT' through 'go env GOROOT': %s: %w", string(output), err)
		}
		return "", fmt.Errorf("get 'GOROOT' through 'go env GOROOT': %w", err)
	}

	// find modulepathfile in go.mod file since it's prefixed by "module::"
	if modulepathfile, ok := strings.CutPrefix(modulepathfile, modulePrefix); ok {
		module, err := findRequire(file, modulepathfile)
		if err != nil {
			return "", err
		}

		// handle local replace
		//  - could be provided as absolute path (but it would be weird): replace github.com/module/name => /path/to/module/name
		//	- could be provided as relative path: replace github.com/module/name => ../../path/to/module/name
		if filepath.IsAbs(module) {
			return module, nil
		}
		if abs, err := filepath.Abs(filepath.Join(file.Dir, module)); err == nil && files.Exists(abs) {
			return abs, nil
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

	// no require to find since it should be a local module file
	return modulepathfile, nil
}

// getImports returns the slice of imports
// and the slice of imports' names associated to input ast file (it takes care of alias'ed imports).
func getImports(file *ast.File) (names []string, imports []string, _ error) {
	names, imports = make([]string, 0, len(file.Imports)), make([]string, 0, len(file.Imports))
	for _, item := range file.Imports {
		if item.Name != nil {
			names = append(names, item.Name.Name)
			imports = append(imports, fmt.Sprint(item.Name.Name, " ", item.Path.Value))
			continue
		}
		names = append(names, strings.Trim(path.Base(item.Path.Value), `"`))
		imports = append(imports, item.Path.Value)
	}
	return names, imports, nil
}

// fileImport returns the import and its alias (in case the import name was already taken) for the input pkg with the associated go.mod modfile.
func fileImport(file modFile, pkg string, names []string) (alias string, imp string) {
	m := lo.FromPtr(file.Module)

	// when working with std package, go doesn't add the go.mod module name to the import
	if m.Mod.Path == "std" {
		return "", fmt.Sprint(`"`, pkg, `"`)
	}

	// ensure import name is not already used by another import
	var changed bool
	alias = lo.CoalesceOrEmpty(pkg, path.Base(m.Mod.Path))
	for _, name := range names {
		if alias == name {
			alias = "builded" // Note: might need a generated name in case 'builded' is also taken ... (but should do the trick for now? Right? ...)
			changed = true
		}
	}
	if !changed {
		alias = "" // remove unnecessary alias since it's not used by other imports names
	}

	// when working with a standard go module, the package name is prefixed with the module name from where it comes
	return alias, alias + " " + fmt.Sprint(`"`, path.Join(m.Mod.Path, pkg), `"`)
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
func findGomod(dir string, parts ...string) (modFile, string, error) {
	mod := filepath.Join(dir, "go.mod")

	// go through parent directory to find go.mod in case it doesn't exist in current directory
	if !files.Exists(mod) {
		// handle root directory -> VolumeName (e.g "C:") + os.PathSeparator
		if dir == filepath.VolumeName(dir)+string(os.PathSeparator) {
			return modFile{}, "", errors.New("no parent go.mod found")
		}
		return findGomod(filepath.Dir(dir), append(parts, filepath.Base(dir))...)
	}
	bytes, _ := os.ReadFile(mod) // don't need to handle error since file exists

	file, err := modfile.Parse(mod, bytes, nil)
	if err != nil {
		return modFile{}, "", fmt.Errorf("go.mod '%s' parsing: %w", mod, err)
	}

	if err := validGomod(file); err != nil {
		return modFile{}, "", fmt.Errorf("invalid '%s' go.mod: %w", mod, err)
	}

	slices.Reverse(parts)
	return modFile{File: file, Dir: dir}, strings.Join(parts, "/"), nil
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
		regexp.MustCompile(fmt.Sprint(`^//go:generate (?:\S+)?go-builder-generator generate `, options, "$")),
		regexp.MustCompile(fmt.Sprint(`^//go:generate go run github\.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@\S+ generate `, options, "$")),
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
