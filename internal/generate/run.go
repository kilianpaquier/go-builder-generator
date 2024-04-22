package generate

import (
	"embed"
	"errors"
	"fmt"
	"go/parser"
	"go/token"
	"path/filepath"

	"github.com/huandu/xstrings"
	filesystem "github.com/kilianpaquier/filesystem/pkg"
)

//go:embed all:templates
var tmpl embed.FS

// Run runs the builder generation with input options.
func Run(pwd string, options CLIOptions) error {
	// retrieve destination full path
	destdir, err := filepath.Abs(options.Destdir)
	if err != nil {
		return fmt.Errorf("failed to retrieve absolute '%s' path: %w", options.Destdir, err)
	}

	src, err := parseSrc(pwd, options.File)
	if err != nil {
		return err // error wrapping is handled in parseSrc function
	}
	srcdir := filepath.Dir(src)

	// parse source file as ast to retrieve golang code
	file, err := parser.ParseFile(token.NewFileSet(), src, nil, parser.SkipObjectResolution)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %w", src, err)
	}

	// retrieve file imports to reuse them in template
	fileImports, err := fileImports(file, srcdir, destdir)
	if err != nil {
		return fmt.Errorf("failed to find %s module name: %w", srcdir, err)
	}

	sourcePackage, destPackage := func() (string, string) {
		if destdir == srcdir {
			return "", file.Name.String()
		}
		return file.Name.String(), filepath.Base(destdir)
	}()

	// generate all builders for input structs
	opts := genOpts{
		DestPackage:   destPackage,
		Imports:       fileImports,
		NoNotice:      options.NoNotice,
		SourcePackage: sourcePackage,
		ValidateFunc:  options.ValidateFunc,
		ReturnCopy:    options.ReturnCopy,

		// force first rune to lowercase in case of unexported types
		// it will be titled in gen template in case the type is exported
		SetterPrefix: xstrings.FirstRuneToLower(options.SetterPrefix),
	}
	var errs []error
	builders, err := generateStructs(file, options.Structs, destdir, opts)
	if err != nil {
		errs = append(errs, err)
	}

	// generate implementation file$
	dest := filepath.Join(destdir, "builders_impl.go")
	if len(builders) > 0 && !filesystem.Exists(dest) {
		impl := &implData{
			Builders:    builders,
			DestPackage: filepath.Base(destdir),
		}
		if err := generateAny(ImplTemplate, dest, impl); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}
