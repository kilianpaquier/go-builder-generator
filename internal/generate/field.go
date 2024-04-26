package generate

import (
	"errors"
	"fmt"
	"go/ast"
	"slices"
	"strings"

	"github.com/fatih/structtag"
	"github.com/huandu/xstrings"

	"github.com/kilianpaquier/go-builder-generator/internal/generate/models"
	"github.com/kilianpaquier/go-builder-generator/internal/generate/prefixer"
)

// parseField parses and returns the struct field associated to input ast field.
func parseField(astField *ast.Field, sourcePackage string, typeParams []string) (field, error) {
	// parse field tags
	options, err := parseOptions(astField.Tag)
	if err != nil {
		return field{}, fmt.Errorf("field options parsing: %w", err)
	}

	// retrieve typePrefixer for field type
	typePrefixer := prefixer.NewPrefixer(astField.Type)
	if err := typePrefixer.Valid(); err != nil {
		return field{}, fmt.Errorf("field validation: %w", err)
	}

	// retrieve computed string type
	initialType, typeExported := typePrefixer.ToString(sourcePackage, typeParams)
	alteredType := initialType

	switch {
	// checking if field is a slice with append option
	// append is exclusive with pointer
	case options.Append && (strings.HasPrefix(initialType, prefixer.Star+prefixer.Slice) || strings.HasPrefix(initialType, prefixer.Slice)):
		options.Pointer = false
		// removing * from type because it's handled in template
		alteredType = strings.TrimPrefix(alteredType, prefixer.Star)
		// removing [] from type because template is gonna use ...{{ .AlteredType }}
		// with append function for slices
		alteredType = strings.TrimPrefix(alteredType, prefixer.Slice)

	// checking is field is a pointer with the pointer option
	// pointer is exclusive with append
	case options.Pointer && strings.HasPrefix(initialType, prefixer.Star):
		options.Append = false
		// removing * from type because it's handled in template
		alteredType = strings.TrimPrefix(alteredType, prefixer.Star)

	default:
		options.Append = false
		options.Pointer = false
		// removing * from type because it's handled in template
		alteredType = strings.TrimPrefix(alteredType, prefixer.Star)
	}

	fieldName := func() string {
		// returning name if it exists
		if len(astField.Names) > 0 {
			return astField.Names[0].Name
		}

		// handle composition fields (mainly those)
		// first split type into package + real type
		split := strings.Split(alteredType, ".")
		// returning last element to cover two cases:
		// when altered type is a type from other package (sourcePackage would be false)
		// when altered type is a type from the same package (sourcePackage would be true)
		return split[len(split)-1]
	}()

	// check field export and ignore option in case generation is done in another package
	exported := typeExported && ast.IsExported(fieldName)
	if sourcePackage != "" && !exported {
		options.Ignore = true
	}

	// returning field with computed types and options
	return field{
		AlteredType: alteredType,
		Exported:    exported,
		InitialType: initialType,
		Name:        fieldName,
		ParamName:   paramName(fieldName),

		Opts: options,
	}, nil
}

// paramName computes the parameter name for a function associated with the input fieldName.
//
// It takes care of acronyms, builtin reserved works and the simple cases.
//
// The resulted parameter name is in camelCase format.
func paramName(fieldName string) string {
	// for names full uppercase, change them to full lowercase
	// a fieldName being 'ID' would give 'id'
	// it's here to handle acronyms like ID, API, HTTP, etc.
	if strings.ToUpper(fieldName) == fieldName {
		return strings.ToLower(fieldName)
	}

	// transform into camel case then put first letter in lowercase
	initial := xstrings.FirstRuneToLower(xstrings.ToCamelCase(fieldName))

	// handle builtin reserved keywords or functions
	// a fieldName being 'Any' would give an initial 'any' and as such the paramName would be 'a'
	// it's not optimal but at least it works
	if slices.Contains(models.Builtin(), initial) {
		return string(initial[0])
	}

	// for all other names, keep initial value which is camelCase format
	// a fieldName being 'InputField' would give 'inputField'
	return initial
}

// parseOptions returns the field options for the input tags.
func parseOptions(astTags *ast.BasicLit) (fieldOpts, error) {
	// check if there're tags
	if astTags == nil {
		return fieldOpts{}, nil
	}
	value := strings.ReplaceAll(astTags.Value, "`", "")

	// parse tags into something useable
	structtags, err := structtag.Parse(value)
	if err != nil {
		return fieldOpts{}, fmt.Errorf("tags parsing: %w", err)
	}

	// retrieve go-builderTag-generator specific tag
	builderTag, err := structtags.Get("builder")
	if err != nil && !strings.Contains(err.Error(), "tag does not exist") {
		return fieldOpts{}, fmt.Errorf("tag 'builder' parsing: %w", err)
	}

	var stringOptions []string
	if builderTag != nil {
		// adding "Name" because for builder tag it's also an option
		// and not a "Name" like it would be with 'json' or 'xml' tags
		stringOptions = append([]string{builderTag.Name}, builderTag.Options...)
	}

	// parse string options
	var errs []error
	var options fieldOpts
	// used for option value parsing
	var ok bool
	for _, option := range stringOptions {
		switch {
		case option == "pointer":
			options.Pointer = true

		case option == "append":
			options.Append = true

		case option == "ignore":
			options.Ignore = true

		case option == "export":
			options.Export = true

		case strings.HasPrefix(option, "default_func"):
			if options.DefaultFunc, ok = getOptionValue(option); !ok {
				errs = append(errs, errors.New("found 'default_func' option but format is invalid, it should be of `default_func=func_name`"))
			}

		case strings.HasPrefix(option, "func_name"):
			if options.FuncName, ok = getOptionValue(option); !ok {
				errs = append(errs, errors.New("found 'func_name' option but format is invalid, it should be of `func_name=func_name`"))
			}
		}
	}
	return options, errors.Join(errs...)
}

// getOptionValue splits the input option to extract its value.
func getOptionValue(option string) (string, bool) {
	split := strings.Split(option, "=")
	if len(split) != 2 {
		return "", false
	}
	return split[1], true
}
