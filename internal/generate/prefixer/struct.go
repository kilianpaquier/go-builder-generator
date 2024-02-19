package prefixer

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"
)

// structPrefixer implements Prefixer for StructType.
type structPrefixer struct {
	*ast.StructType

	FieldFields []Prefixer
}

var _ Prefixer = &structPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (s *structPrefixer) Valid() error {
	var errs []error

	// retrieve prefixers associated to struct fields
	if s.Fields != nil {
		editor := func(field *ast.Field) editor {
			return func(stringType string, exported bool) (string, bool) {
				if len(field.Names) > 0 {
					// for an anonymous struct, exported means the field name
					// starts with an uppercase and the string type is exported too
					exported = exported && ast.IsExported(field.Names[0].Name)
					return fmt.Sprint(field.Names[0].Name, " ", stringType), exported
				}

				// this case shouldn't happen since we're in a struct{}
				return "", false
			}
		}

		s.FieldFields = make([]Prefixer, 0, len(s.Fields.List))
		for _, field := range s.Fields.List {
			// create a prefixer editor to ensure struct export and fields' name are added / updated with name uppercasing (or not)
			prefixer := NewPrefixerEditor(NewPrefixer(field.Type), editor(field))

			errs = append(errs, prefixer.Valid())
			s.FieldFields = append(s.FieldFields, prefixer)
		}
	}

	return errors.Join(errs...)
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (s *structPrefixer) ToString(sourcePackage string, prefixes ...string) (_ string, _ bool) {
	exported := true
	types := make([]string, 0, len(s.FieldFields))

	// compute fields prefix
	for _, field := range s.FieldFields {
		stringType, fieldExported := field.ToString(sourcePackage)

		// don't affect directly because once exported is false, it should stays even if other fields are exported
		if !fieldExported {
			exported = false
		}

		types = append(types, stringType)
	}

	// specific case to avoid unnecessary newlines
	if len(types) == 0 {
		return strings.Join(prefixes, "") + "struct{}", exported
	}
	return fmt.Sprintf("%sstruct{\n%s\n}", strings.Join(prefixes, ""), strings.Join(types, "\n")), exported
}
