package prefixer

import (
	"fmt"
	"go/ast"
)

// unimplementedPrefixer represents a prefixer for an expression that isn't implemented yet.
//
// This Prefixer returns an error on Valid call
// and panics on ToString call (because it obviously should never be called).
type unimplementedPrefixer struct{ ast.Expr }

var _ Prefixer = &unimplementedPrefixer{} // ensure interface is implemented

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (*unimplementedPrefixer) ToString(_ string, _ []string, _ ...string) (_ string, _ bool) {
	panic("should not be called")
}

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (i *unimplementedPrefixer) Valid() error {
	return fmt.Errorf("expression '%T' not implemented", i.Expr)
}

// nooPrefixer implements Prefixer for nil expressions.
type noopPrefixer struct{}

var _ Prefixer = &noopPrefixer{}

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (*noopPrefixer) Valid() error {
	return nil
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (*noopPrefixer) ToString(_ string, _ []string, _ ...string) (_ string, _ bool) {
	return "", true
}
