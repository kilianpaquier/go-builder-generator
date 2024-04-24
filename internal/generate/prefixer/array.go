package prefixer

import (
	"errors"
	"fmt"
	"go/ast"
)

// arrayPrefixer implements Prefixer for ArrayType.
type arrayPrefixer struct {
	*ast.ArrayType

	LenPrefixer Prefixer
	EltPrefixer Prefixer
}

var _ Prefixer = &arrayPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (a *arrayPrefixer) Valid() error {
	var errs []error
	if a.Len != nil {
		// retrieve prefixer associated to ellipsis [X]
		a.LenPrefixer = NewPrefixer(a.Len)
		errs = append(errs, a.LenPrefixer.Valid())
	}

	// retrieve prefixer associated to slice/array element
	a.EltPrefixer = NewPrefixer(a.Elt)
	errs = append(errs, a.EltPrefixer.Valid())

	return errors.Join(errs...)
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (a *arrayPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	var prefix string
	if a.LenPrefixer == nil {
		prefix = "[]"
	} else {
		ellipsis, _ := a.LenPrefixer.ToString("", nil)
		prefix = fmt.Sprintf("[%s]", ellipsis)
	}

	// retrieve prefixer associated to slice/array element
	return a.EltPrefixer.ToString(sourcePackage, typeParams, append(prefixes, prefix)...)
}
