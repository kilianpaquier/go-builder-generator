package prefixer

import (
	"errors"
	"go/ast"
)

// selectorPrefixer implements Prefixer for SelectorExpr.
type selectorPrefixer struct {
	*ast.SelectorExpr

	XPrefixer   Prefixer
	SelPrefixer Prefixer
}

var _ Prefixer = &selectorPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (s *selectorPrefixer) Valid() error {
	var errs []error

	// retrieve package prefixer
	// with context.Context, X would be the "context" part
	s.XPrefixer = NewPrefixer(s.X)
	errs = append(errs, s.XPrefixer.Valid())

	// retrieve type name prefixer
	// with context.Context, Sel would be the "Context" part
	s.SelPrefixer = NewPrefixer(s.Sel)
	errs = append(errs, s.SelPrefixer.Valid())

	return errors.Join(errs...)
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (s *selectorPrefixer) ToString(_ string, prefixes ...string) (_ string, _ bool) {
	// parse package name
	packageName, _ := s.XPrefixer.ToString("")

	// parse type name
	return s.SelPrefixer.ToString(packageName, prefixes...)
}
