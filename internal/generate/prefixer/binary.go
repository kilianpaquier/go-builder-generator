package prefixer

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"
)

// binaryPrefixer implements Prefixer for BinaryExpr.
//
// It's mainly used for array length arithmetic (e.g. [N+1]int) and inline generic
// type parameter constraints type-set unions (e.g. type Number interface{ ~int | ~float64 }).
type binaryPrefixer struct {
	*ast.BinaryExpr

	XPrefixer Prefixer
	YPrefixer Prefixer
}

var _ Prefixer = &binaryPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (b *binaryPrefixer) Valid() error {
	b.XPrefixer = NewPrefixer(b.X)
	b.YPrefixer = NewPrefixer(b.Y)
	return errors.Join(b.XPrefixer.Valid(), b.YPrefixer.Valid())
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (b *binaryPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	x, xExported := b.XPrefixer.ToString(sourcePackage, typeParams)
	y, yExported := b.YPrefixer.ToString(sourcePackage, typeParams)

	// to be exported, both sides of the binary expression must be exported
	return fmt.Sprintf("%s%s %s %s", strings.Join(prefixes, ""), x, b.Op.String(), y), xExported && yExported
}
