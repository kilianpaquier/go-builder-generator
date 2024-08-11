package generate

import (
	"embed"
	"errors"
	"fmt"
	ast "go/parser"
	"go/token"
	"path"
	"path/filepath"
	"strings"

	"github.com/huandu/xstrings"

	"github.com/kilianpaquier/go-builder-generator/internal/fs"
)

//go:embed all:templates
var tmpl embed.FS

// Run runs the builder generation with input options.
func Run(options CLIOptions, rawArgs []string) error {
	// force first rune to lowercase in case of unexported types
	// it will be titled in gen template in case the type is exported
	options.Prefix = xstrings.FirstRuneToLower(options.Prefix)

	// retrieve destination full path
	destdir, err := filepath.Abs(options.Destdir)
	if err != nil {
		return fmt.Errorf("absolute path: %w", err)
	}

	// retrieve destination modfile and path parts to it
	destfile, _, err := findGomod(destdir)
	if err != nil {
		return fmt.Errorf("find dest go.mod: %w", err)
	}

	// retrieve source full path (specific since src can use "module::" specific key)
	src, err := func() (string, error) {
		if !strings.HasPrefix(options.File, modulePrefix) {
			return filepath.Abs(options.File)
		}

		p, err := modulePath(destfile, options.File)
		if err != nil {
			return "", err
		}
		return filepath.Abs(p)
	}()
	if err != nil {
		return err
	}
	srcdir := filepath.Dir(src)

	// retrieve source modfile and path parts to it
	srcfile, srcpkg, err := findGomod(srcdir)
	if err != nil {
		return fmt.Errorf("find src go.mod: %w", err)
	}

	// parse source file as ast to retrieve golang code
	file, err := ast.ParseFile(token.NewFileSet(), src, nil, ast.ParseComments)
	if err != nil {
		return fmt.Errorf("parse file: %w", err)
	}

	// retrieve source file imports to reuse them in template
	imports, err := getImports(file)
	if err != nil {
		return fmt.Errorf("get imports: %w", err)
	}
	imports = append(imports, fileImport(srcfile, srcpkg))

	sourcePackage, destPackage := func() (string, string) {
		if destdir == srcdir {
			return "", file.Name.String()
		}
		return file.Name.String(), filepath.Base(destdir)
	}()

	var errs []error

	// generate all builders for input structs
	data := packagesData{
		Destdir:       destdir,
		DestName:      destPackage,
		GeneratedFrom: path.Join(srcpkg, filepath.Base(options.File)),
		HasGenerate:   hasGenerate(file, rawArgs),
		Imports:       imports,
		SameModule:    destfile.Module.Mod.String() == srcfile.Module.Mod.String(),
		SourceName:    sourcePackage,
	}
	builders, err := generateBuilders(file, data, options)
	if err != nil {
		errs = append(errs, err)
	}

	// generate implementation file
	dest := filepath.Join(destdir, "builders_impl.go")
	if len(builders) > 0 && !fs.Exists(dest) {
		impl := &implData{
			Builders: builders,
			Opts:     options,
			Packages: data,
		}
		if err := generateAny(ImplTemplate, dest, impl); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}
