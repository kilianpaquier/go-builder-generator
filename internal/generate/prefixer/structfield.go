package prefixer

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"
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
	stringType, exportedType := f.TypePrefixer.ToString(sourcePackage, typeParams, prefixes...)
	tag, _ := f.TagPrefixer.ToString("", nil)

	exported := exportedType // if the type is not exported, then the structfield is not exported even if names are exported
	names := make([]string, 0, len(f.Names))
	for _, name := range f.Names {
		// unexport the whole structfield if at least one of its fields is not
		// i.e. anonymous struct{ Start, end time.Time } cannot be used outside of its own package
		if !ast.IsExported(name.Name) {
			exported = false
		}
		names = append(names, fmt.Sprint(name, " ", stringType, tag))
	}
	return strings.Join(names, "\n"), exported
}
