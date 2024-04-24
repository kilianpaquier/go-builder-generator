package prefixer

import (
	"errors"
	"fmt"
	"go/ast"
)

// structFieldPrefixer implements Prefixer for a struct field.
type structFieldPrefixer struct {
	*ast.Field

	TagPrefixer  Prefixer
	TypePrefixer Prefixer
}

var _ Prefixer = &structFieldPrefixer{} // ensure interface is implemented

// newStructFieldPrefixer creates a new prefixer for a field.
//
// It has it's specific function because ast.Field doesn't implement ast.Expr, as such a field prefixer can't be created with NewPrefixer.
func newStructFieldPrefixer(field *ast.Field) Prefixer {
	return &structFieldPrefixer{Field: field}
}

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (f *structFieldPrefixer) Valid() error {
	f.TagPrefixer = func() Prefixer {
		if f.Tag != nil {
			return NewPrefixer(f.Tag)
		}
		return &noopPrefixer{}
	}()
	f.TypePrefixer = NewPrefixer(f.Type)
	return errors.Join(f.TagPrefixer.Valid(), f.TypePrefixer.Valid())
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (f *structFieldPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	stringType, exported := f.TypePrefixer.ToString(sourcePackage, typeParams, prefixes...)
	tag, _ := f.TagPrefixer.ToString("", nil)

	if len(f.Names) > 0 {
		// for an anonymous struct, exported means the field name
		// starts with an uppercase and the string type is exported too
		exported = exported && ast.IsExported(f.Names[0].Name)
		return fmt.Sprint(f.Names[0].Name, " ", stringType, tag), exported
	}

	// this case shouldn't happen since we're in a struct{}
	return "", false
}
