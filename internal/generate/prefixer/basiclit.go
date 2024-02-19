package prefixer

import "go/ast"

// basicLitPrefixer implements Prefixer for BasicLit.
type basicLitPrefixer ast.BasicLit

var _ Prefixer = &basicLitPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (*basicLitPrefixer) Valid() error {
	return nil
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (b *basicLitPrefixer) ToString(_ string, _ ...string) (_ string, _ bool) {
	return b.Value, ast.IsExported(b.Value)
}
