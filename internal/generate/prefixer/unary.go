package prefixer

import "go/ast"

// unaryPrefixer implements Prefixer for UnaryExpr.
type unaryPrefixer struct {
	*ast.UnaryExpr

	XPrefixer Prefixer
}

var _ Prefixer = &unaryPrefixer{}

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (u *unaryPrefixer) Valid() error {
	u.XPrefixer = NewPrefixer(u.X)
	return u.XPrefixer.Valid() //nolint:wrapcheck
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (u *unaryPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	return u.XPrefixer.ToString(sourcePackage, typeParams, append(prefixes, u.Op.String())...)
}
