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

	MethodsPrefixers []Prefixer
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

		i.MethodsPrefixers = make([]Prefixer, 0, len(i.Methods.List))
		for _, method := range i.Methods.List {
			// create a prefixer to remove func( prefix and add name prefix
			prefixer := NewPrefixerEditor(NewPrefixer(method.Type), editor(method))

			errs = append(errs, prefixer.Valid())
			i.MethodsPrefixers = append(i.MethodsPrefixers, prefixer)
		}
	}

	return errors.Join(errs...)
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (i *interfacePrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	// compute fields prefix
	exported := true
	methodsTypes := make([]string, 0, len(i.MethodsPrefixers))
	for _, field := range i.MethodsPrefixers {
		stringType, fieldExported := field.ToString(sourcePackage, typeParams)

		// don't affect directly because once exported is false, it should stays even if other fields are exported
		if !fieldExported {
			exported = false
		}

		methodsTypes = append(methodsTypes, stringType)
	}

	// specific case to avoid unnecessary newlines
	if len(methodsTypes) == 0 {
		return strings.Join(prefixes, "") + "interface{}", exported
	}
	return fmt.Sprintf("%sinterface{\n%s\n}", strings.Join(prefixes, ""), strings.Join(methodsTypes, "\n")), exported
}
