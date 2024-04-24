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

	FieldsPrefixers []Prefixer
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
		s.FieldsPrefixers = make([]Prefixer, 0, len(s.Fields.List))
		for _, field := range s.Fields.List {
			prefixer := newStructFieldPrefixer(field)

			errs = append(errs, prefixer.Valid())
			s.FieldsPrefixers = append(s.FieldsPrefixers, prefixer)
		}
	}

	return errors.Join(errs...)
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (s *structPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	// compute fields prefix
	exported := true
	fieldsTypes := make([]string, 0, len(s.FieldsPrefixers))
	for _, field := range s.FieldsPrefixers {
		stringType, fieldExported := field.ToString(sourcePackage, typeParams)

		// don't affect directly because once exported is false, it should stays even if other fields are exported
		if !fieldExported {
			exported = false
		}

		fieldsTypes = append(fieldsTypes, stringType)
	}

	// specific case to avoid unnecessary newlines
	if len(fieldsTypes) == 0 {
		return strings.Join(prefixes, "") + "struct{}", exported
	}
	return fmt.Sprintf("%sstruct{\n%s\n}", strings.Join(prefixes, ""), strings.Join(fieldsTypes, "\n")), exported
}
