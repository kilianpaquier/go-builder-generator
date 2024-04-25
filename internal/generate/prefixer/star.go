package prefixer

import (
	"go/ast"
)

const Star = "*"

// starPrefixer implements Prefixer for StarExpr.
type starPrefixer struct {
	*ast.StarExpr

	XPrefixer Prefixer
}

var _ Prefixer = &starPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (p *starPrefixer) Valid() error {
	p.XPrefixer = NewPrefixer(p.X)
	return p.XPrefixer.Valid()
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (p *starPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (stringType string, exported bool) {
	return p.XPrefixer.ToString(sourcePackage, typeParams, append(prefixes, Star)...)
}
