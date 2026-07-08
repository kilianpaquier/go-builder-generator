package prefixer

import "go/ast"

// parenPrefixer implements Prefixer for ParenExpr.
type parenPrefixer struct {
	*ast.ParenExpr

	XPrefixer Prefixer
}

var _ Prefixer = &parenPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (p *parenPrefixer) Valid() error {
	p.XPrefixer = NewPrefixer(p.X)
	return p.XPrefixer.Valid()
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
//
// Parentheses are dropped since they never alter the resulting type in the standalone
// slot they're rendered into (field type, array length, etc.).
func (p *parenPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	return p.XPrefixer.ToString(sourcePackage, typeParams, prefixes...)
}
