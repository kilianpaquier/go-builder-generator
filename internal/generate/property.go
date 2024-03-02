package generate

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"

	"github.com/fatih/structtag"
	"github.com/huandu/xstrings"
	"golang.org/x/exp/slices"

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

	// retrieve tprefixer for field type
	tprefixer := prefixer.NewPrefixer(field.Type)
	if err := tprefixer.Valid(); err != nil {
		return property{}, fmt.Errorf("property type prefixer is not implemented: %w", err)
	}

	// retrieve computed string type
	initialType, exported := tprefixer.ToString(sourcePackage)
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

	name := func() string {
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

	// check property export and ignore option
	if !exported || !ast.IsExported(name) {
		options.Ignore = true
	}

	paramName := func() string {
		// transform into camel case then put first letter in lowercase
		initial := xstrings.FirstRuneToLower(xstrings.ToCamelCase(name))

		// handle builtin reserved keywords or functions
		if slices.Contains(models.Builtin(), initial) {
			return string(initial[0])
		}

		// for names full uppercase, change them to full lowercase
		if strings.ToUpper(name) == name {
			return strings.ToLower(name)
		}

		// for all other names, keep initial value which is camelCase format
		return initial
	}()

	// returning property with computed types and options
	return property{
		AlteredType:  alteredType,
		InitialType:  initialType,
		Name:         name,
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
	for _, option := range stringOptions {
		switch {
		case option == "pointer":
			options.Pointer = true

		case option == "append":
			options.Append = true

		case option == "ignore":
			options.Ignore = true

		case strings.HasPrefix(option, "default_func"):
			split := strings.Split(option, "=")
			if len(split) != 2 {
				errs = append(errs, errors.New("found 'default_func' option but format is invalid, it should be of `default_func=func_name`"))
			} else {
				options.DefaultFunc = split[1]
			}
		}
	}

	// check if validator needs to be added to options
	if _, err := structtags.Get("validate"); err == nil {
		options.Validate = true
	}

	return options, errors.Join(errs...)
}
