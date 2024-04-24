package generate

import (
	"errors"
	"fmt"
	"go/ast"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/huandu/xstrings"
	filesystem "github.com/kilianpaquier/filesystem/pkg"
	"github.com/samber/lo"
	"golang.org/x/tools/imports"
)

// generateBuilders takes a go tree file as input and generates a builder for all the input structs slice.
//
// It returns a slice to generate aftermath an _impl file with custom functions.
func generateBuilders(file *ast.File, pkg packageData, opts CLIOptions) ([]genData, error) {
	// inspect whole file an retrieve all associated builders
	var errs []error
	builders := make([]genData, 0, len(opts.Structs))

	var validNodes int
	ast.Inspect(file, func(node ast.Node) bool {
		// stop recursive calls once all structs are retrieved
		if len(opts.Structs) == validNodes {
			return false
		}

		// go through next nodes if the current is not a `type`
		typeSpec, ok := node.(*ast.TypeSpec)
		if !ok || !slices.Contains(opts.Structs, typeSpec.Name.String()) {
			return true
		}

		// go through next nodes if current is not a struct
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok || structType.Fields == nil || len(structType.Fields.List) == 0 {
			return true
		}
		validNodes++

		// compute builder data associated to struct name
		builder, err := parseStruct(typeSpec, structType, pkg, opts)
		if err != nil {
			errs = append(errs, err)
			return true
		}

		// create destination directory
		// only now because we don't want to create the directory unless at least one builder was successfully computed and ready for generation
		if err := os.MkdirAll(builder.Opts.Destdir, filesystem.RwxRxRxRx); err != nil && !os.IsExist(err) {
			errs = append(errs, fmt.Errorf("mkdir %s: %w", builder.Opts.Destdir, err))
			return false // since the destination directory couldn't be created, stop all
		}

		// generate struct builder
		dest := filepath.Join(builder.Opts.Destdir, xstrings.ToSnakeCase(builder.Name)+"_builder_gen.go")
		if err := generateAny(GenTemplate, dest, builder); err != nil {
			errs = append(errs, fmt.Errorf("generate builder %s: %w", builder.Name, err))
		}

		// there may be some cases where impl is nil in case there're no default funcs to generate
		if len(builder.DefaultFuncs) > 0 {
			builders = append(builders, builder)
		}
		return true
	})

	return builders, errors.Join(errs...)
}

// generateAny takes an input tmpl filename and generates a file at input destination
// with input data with go templating.
func generateAny(filename string, dest string, data any) error {
	// parse template file
	tpl, err := template.New(filename).
		Funcs(funcMap()).
		Funcs(sprig.FuncMap()).
		ParseFS(tmpl, path.Join("templates", filename))
	if err != nil {
		return fmt.Errorf("parse template %s file: %w", filename, err)
	}

	// render file
	var content strings.Builder
	if err := tpl.Execute(&content, data); err != nil {
		return fmt.Errorf("execute template %s: %w", filename, err)
	}

	writeFile := func(bytes []byte) error {
		if err := os.WriteFile(dest, bytes, filesystem.RwRR); err != nil {
			return fmt.Errorf("write file %s: %w", dest, err)
		}
		return nil
	}

	// optimize file imports
	bytes := []byte(content.String())
	formatted, err := imports.Process(dest, bytes, nil)
	if err != nil {
		// also write file when imports optimization failed
		// better for debugging
		_ = writeFile(bytes)
		return fmt.Errorf("generated builder '%s' is incorrect: %w", dest, err)
	}
	return writeFile(formatted)
}

// funcMap returns the functions to be used in go template generation to make it easier.
func funcMap() template.FuncMap {
	return template.FuncMap{
		"joinFields":      joinFields,
		"joinFieldsNames": joinFieldsNames,
	}
}

// joinFields joins all input fields with their names and associated types.
func joinFields(fields []field) string {
	if len(fields) == 0 {
		return ""
	}
	strParams := lo.Map(fields, func(param field, _ int) string { return fmt.Sprint(param.Name, " ", param.AlteredType) })
	return fmt.Sprintf("[%s]", strings.Join(strParams, ","))
}

// joinFieldsNames joins all input fields with only their names.
func joinFieldsNames(fields []field) string {
	if len(fields) == 0 {
		return ""
	}
	strParams := lo.Map(fields, func(param field, _ int) string { return param.Name })
	return fmt.Sprintf("[%s]", strings.Join(strParams, ","))
}
