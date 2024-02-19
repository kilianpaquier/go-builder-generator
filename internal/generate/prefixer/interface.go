package prefixer

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"
)

// interfacePrefixer implements Prefixer for InterfaceType.
type interfacePrefixer struct {
	*ast.InterfaceType

	MethodFields []Prefixer
}

var _ Prefixer = &interfacePrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (i *interfacePrefixer) Valid() error {
	var errs []error

	// retrieve prefixers associated to interface methods
	if i.Methods != nil {
		editor := func(method *ast.Field) editor {
			return func(stringType string, _ bool) (string, bool) {
				if len(method.Names) > 0 {
					// for an anonymous interface, exported means the function name starts with an uppercase
					exported := ast.IsExported(method.Names[0].Name)

					// remove first func prefix reference
					// because for interfaces, the func is named, like SomeFunc() error
					return method.Names[0].Name + strings.Replace(stringType, "func(", "(", 1), exported
				}

				// this case shouldn't happen since we're in an interface type
				return "", false
			}
		}

		i.MethodFields = make([]Prefixer, 0, len(i.Methods.List))
		for _, method := range i.Methods.List {
			// create a prefixer prefixer to remove func( prefix and add name prefix
			prefixer := NewPrefixerEditor(NewPrefixer(method.Type), editor(method))

			errs = append(errs, prefixer.Valid())
			i.MethodFields = append(i.MethodFields, prefixer)
		}
	}

	return errors.Join(errs...)
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (i *interfacePrefixer) ToString(sourcePackage string, prefixes ...string) (_ string, _ bool) {
	exported := true
	types := make([]string, 0, len(i.MethodFields))

	// compute fields prefix
	for _, field := range i.MethodFields {
		stringType, fieldExported := field.ToString(sourcePackage)

		// don't affect directly because once exported is false, it should stays even if other fields are exported
		if !fieldExported {
			exported = false
		}

		types = append(types, stringType)
	}

	// specific case to avoid unnecessary newlines
	if len(types) == 0 {
		return strings.Join(prefixes, "") + "interface{}", exported
	}
	return fmt.Sprintf("%sinterface{\n%s\n}", strings.Join(prefixes, ""), strings.Join(types, "\n")), exported
}
