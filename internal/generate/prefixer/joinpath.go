package prefixer

import "go/ast"

// joinpathPrefixer implements Prefixer for JoinPath.
type joinpathPrefixer struct {
	*ast.Ellipsis

	EltPrefixer Prefixer
}

var _ Prefixer = &chanPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (c *joinpathPrefixer) Valid() error {
	c.EltPrefixer = NewPrefixer(c.Elt)
	return c.EltPrefixer.Valid()
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (c *joinpathPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	return c.EltPrefixer.ToString(sourcePackage, typeParams, append(prefixes, "...")...)
}
