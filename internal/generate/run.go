package generate

import (
	"embed"
	"errors"
	"fmt"
	ast "go/parser"
	"go/token"
	"path/filepath"

	"github.com/huandu/xstrings"

	"github.com/kilianpaquier/go-builder-generator/internal/fs"
)

//go:embed all:templates
var tmpl embed.FS

// Run runs the builder generation with input options.
func Run(pwd string, options CLIOptions) error {
	// force first rune to lowercase in case of unexported types
	// it will be titled in gen template in case the type is exported
	options.Prefix = xstrings.FirstRuneToLower(options.Prefix)

	// retrieve destination full path
	destdir, err := filepath.Abs(options.Destdir)
	if err != nil {
		return fmt.Errorf("absolute '%s' path: %w", options.Destdir, err)
	}

	src, err := parseSrc(pwd, options.File)
	if err != nil {
		return err // error wrapping is handled in parseSrc function
	}
	srcdir := filepath.Dir(src)

	// parse source file as ast to retrieve golang code
	file, err := ast.ParseFile(token.NewFileSet(), src, nil, ast.SkipObjectResolution)
	if err != nil {
		return fmt.Errorf("parse file: %w", err)
	}

	// retrieve file imports to reuse them in template
	imports, err := parseImports(file, srcdir, destdir)
	if err != nil {
		return fmt.Errorf("parse imports: %w", err)
	}

	sourcePackage, destPackage := func() (string, string) {
		if destdir == srcdir {
			return "", file.Name.String()
		}
		return file.Name.String(), filepath.Base(destdir)
	}()

	var errs []error

	// generate all builders for input structs
	pkg := packageData{
		Destdir:       destdir,
		DestPackage:   destPackage,
		Imports:       imports,
		SourcePackage: sourcePackage,
	}
	builders, err := generateBuilders(file, pkg, options)
	if err != nil {
		errs = append(errs, err)
	}

	// generate implementation file$
	dest := filepath.Join(destdir, "builders_impl.go")
	if len(builders) > 0 && !fs.Exists(dest) {
		impl := &implData{
			Builders:      builders,
			DestPackage:   destPackage,
			Opts:          options,
			SourcePackage: sourcePackage,
		}
		if err := generateAny(ImplTemplate, dest, impl); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}
