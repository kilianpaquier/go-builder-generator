package generate

import (
	"cmp"
	"errors"
	"fmt"
	"go/ast"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/huandu/xstrings"
	filesystem "github.com/kilianpaquier/filesystem/pkg"
	"golang.org/x/tools/imports"
)

// generateStructs takes a go tree file as input and generates a builder for all the input structs slice.
//
// It returns a slice of ImplBuilder to generate aftermath an _impl file with custom functions.
func generateStructs(file *ast.File, structs []string, destdir string, opts genOpts) ([]*implBuilder, error) {
	// inspect whole file an retrieve all associated builders
	var errs []error
	builders := make([]*implBuilder, 0, len(structs))

	var validNodes int
	ast.Inspect(file, func(node ast.Node) bool {
		// stop recursive calls once all structs are retrieved
		if len(structs) == validNodes {
			return false
		}

		// go through next nodes if the current is not a `type`
		spec, ok := node.(*ast.TypeSpec)
		if !ok || !slices.Contains(structs, spec.Name.String()) {
			return true
		}

		// go through next nodes if current is not a struct
		s, ok := spec.Type.(*ast.StructType)
		if !ok || s.Fields == nil || len(s.Fields.List) == 0 {
			return true
		}
		validNodes++

		// initialize builder to avoid too many params in generateStruct
		builder := genBuilder{
			genOpts:     opts,
			implBuilder: implBuilder{Name: spec.Name.String()},
		}

		// generate struct builder
		impl, err := generateStruct(builder, s.Fields.List, destdir)
		errs = append(errs, err) // errors join handles nil errors

		// there may be some cases where impl is nil in case there're no default funcs to generate
		if impl != nil {
			builders = append(builders, impl)
		}
		return true
	})

	return builders, errors.Join(errs...)
}

// generateStruct parses all input fields at Properties and generate the resulting updated input builder.
//
// It returns an ImplBuilder only if there's no error and one of the Properties has at least one DefaultFunc.
func generateStruct(builder genBuilder, fields []*ast.Field, destdir string) (*implBuilder, error) {
	// iterate over struct properties and parse every one of them
	var errs []error
	for _, field := range fields {
		// compute property
		property, err := computeProperty(field, builder.SourcePackage)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		// move up validate option because it's common to all struct (Build function)
		// and not just the current field
		if builder.UseValidator && property.Validate {
			builder.HasValidate = true
			builder.Imports = append(builder.Imports, `"github.com/go-playground/validator/v10"`)
		}

		// move up default_func option one function could be use for multiple fields
		if property.DefaultFunc != "" && !slices.Contains(builder.DefaultFuncs, property.DefaultFunc) {
			builder.DefaultFuncs = append(builder.DefaultFuncs, property.DefaultFunc)
		}

		builder.Properties = append(builder.Properties, property)
	}

	// sort properties to have the same generation even if fields order change
	slices.SortStableFunc(builder.Properties, func(p1, p2 property) int {
		return cmp.Compare(strings.ToLower(p1.Name), strings.ToLower(p2.Name))
	})

	// not generating struct builder
	if len(errs) > 0 {
		err := errors.Join(errs...)
		return nil, fmt.Errorf("failed to parse builder for struct %s: %w", builder.Name, err)
	}

	// create destination directory
	if err := os.MkdirAll(destdir, filesystem.RwxRxRxRx); err != nil && !os.IsExist(err) {
		return nil, fmt.Errorf("failed to create %s: %w", destdir, err)
	}

	// generate struct builder
	dest := filepath.Join(destdir, xstrings.ToSnakeCase(builder.Name)+"_builder_gen.go")
	if err := generateAny(GenTemplate, dest, builder); err != nil {
		return nil, fmt.Errorf("failed to generate builder for struct %s: %w", builder.Name, err)
	}

	// return data for shared _impl.go file generation
	if len(builder.DefaultFuncs) > 0 {
		return &builder.implBuilder, nil
	}
	return nil, nil
}

// generateAny takes an input tmpl filename and generates a file at input destination
// with input data with go templating.
func generateAny(tmplName string, dest string, data any) error {
	// parse template file
	tpl, err := template.New(tmplName).
		Funcs(funcMap()).
		ParseFS(tmpl, path.Join("templates", tmplName))
	if err != nil {
		return fmt.Errorf("failed to parse template %s file: %w", tmplName, err)
	}

	// render file
	var content strings.Builder
	if err := tpl.Execute(&content, data); err != nil {
		return fmt.Errorf("failed to apply template %s: %w", tmplName, err)
	}

	// optimize file imports
	bytes := []byte(content.String())
	formatted, err := imports.Process(dest, bytes, nil)
	if err != nil {
		// also write file when imports optimization failed
		// better for debugging
		if err := os.WriteFile(dest, bytes, filesystem.RwRR); err != nil {
			return fmt.Errorf("failed to write file %s: %w", dest, err)
		}
		return fmt.Errorf("generated go file %s is incorrect: %w", dest, err)
	}

	// write file
	if err := os.WriteFile(dest, formatted, filesystem.RwRR); err != nil {
		return fmt.Errorf("failed to write file %s: %w", dest, err)
	}
	return nil
}
