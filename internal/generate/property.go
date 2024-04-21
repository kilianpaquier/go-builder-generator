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

// computeProperty computes the property struct depending on inputs.
//
// It's a special function because the target type will be altered depending on options.
func computeProperty(field *ast.Field, sourcePackage string) (property, error) {
	// parse property tags
	options, err := parseOptions(field.Tag)
	if err != nil {
		return property{}, fmt.Errorf("failed to parse field options: %w", err)
	}

	// retrieve typePrefixer for field type
	typePrefixer := prefixer.NewPrefixer(field.Type)
	if err := typePrefixer.Valid(); err != nil {
		return property{}, fmt.Errorf("property type prefixer is not implemented: %w", err)
	}

	// retrieve computed string type
	initialType, typeExported := typePrefixer.ToString(sourcePackage)
	alteredType := initialType

	switch {
	// checking if field is a slice with append option
	// append is exclusive with pointer
	case options.Append && (strings.HasPrefix(initialType, "*[]") || strings.HasPrefix(initialType, "[]")):
		options.Pointer = false
		// removing * from type because it's handle in template
		alteredType = strings.TrimPrefix(alteredType, "*")
		// removing [] from type because template is gonna use ...{{ .AlteredType }}
		// with append function for slices
		alteredType = strings.TrimPrefix(alteredType, "[]")

	// checking is field is a pointer with the pointer option
	// pointer is exclusive with append
	case options.Pointer && strings.HasPrefix(initialType, "*"):
		options.Append = false
		// removing * from type because it's handle in template
		alteredType = strings.TrimPrefix(alteredType, "*")

	default:
		options.Append = false
		options.Pointer = false
		// removing * from type because it's handle in template
		alteredType = strings.TrimPrefix(alteredType, "*")
	}

	propertyName := func() string {
		// returning name if it exists
		if len(field.Names) > 0 {
			return field.Names[0].Name
		}

		// first split type into package + real type
		split := strings.Split(alteredType, ".")
		// returning last element to cover two cases:
		// when altered type is a type from other package (t.SamePackage would be false)
		// when altered type is a type from the same package (t.SamePackage would be true)
		return split[len(split)-1]
	}()

	// check property export and ignore option in case generation is done in another package
	exported := typeExported && ast.IsExported(propertyName)
	if sourcePackage != "" && !exported {
		options.Ignore = true
	}

	paramName := func() string {
		// transform into camel case then put first letter in lowercase
		initial := xstrings.FirstRuneToLower(xstrings.ToCamelCase(propertyName))

		// handle builtin reserved keywords or functions
		if slices.Contains(models.Builtin(), initial) {
			return string(initial[0])
		}

		// for names full uppercase, change them to full lowercase
		if strings.ToUpper(propertyName) == propertyName {
			return strings.ToLower(propertyName)
		}

		// for all other names, keep initial value which is camelCase format
		return initial
	}()

	// returning property with computed types and options
	return property{
		AlteredType:  alteredType,
		Exported:     exported,
		InitialType:  initialType,
		Name:         propertyName,
		ParamName:    paramName,
		propertyOpts: options,
	}, nil
}

// parseOptions returns the property options for the input tags.
func parseOptions(tags *ast.BasicLit) (propertyOpts, error) {
	// check if there're tags
	if tags == nil {
		return propertyOpts{}, nil
	}
	value := strings.ReplaceAll(tags.Value, "`", "")

	// parse tags into something useable
	structtags, err := structtag.Parse(value)
	if err != nil {
		return propertyOpts{}, fmt.Errorf("failed to parse tags: %w", err)
	}

	// retrieve go-builder-generator specific tag
	builder, err := structtags.Get("builder")
	if err != nil && !strings.Contains(err.Error(), "tag does not exist") {
		return propertyOpts{}, fmt.Errorf("failed to retrieve 'builder' tag: %w", err)
	}

	var stringOptions []string
	if builder != nil {
		// adding "Name" because for builder tag it's also an option
		// and not a "Name" like it would be with 'json' or 'xml' tags
		stringOptions = append([]string{builder.Name}, builder.Options...)
	}

	// parse string options
	var errs []error
	var options propertyOpts
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

func getOptionValue(option string) (string, bool) {
	split := strings.Split(option, "=")
	if len(split) != 2 {
		return "", false
	}
	return split[1], true
}
