package prefixer

import (
	"errors"
	"fmt"
	"go/ast"
)

// indexPrefixer implements Prefixer for IndexExpr.
type indexPrefixer struct {
	*ast.IndexExpr

	IndexPrefixer Prefixer
	XPrefixer     Prefixer
}

var _ Prefixer = &indexPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (g *indexPrefixer) Valid() error {
	g.XPrefixer = NewPrefixer(g.X)
	g.IndexPrefixer = NewPrefixer(g.Index)
	return errors.Join(g.XPrefixer.Valid(), g.IndexPrefixer.Valid())
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (g *indexPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (stringType string, exported bool) {
	x, xExported := g.XPrefixer.ToString(sourcePackage, typeParams, prefixes...)
	index, indexExported := g.IndexPrefixer.ToString(sourcePackage, typeParams)
	return fmt.Sprintf("%s[%s]", x, index), indexExported && xExported
}
