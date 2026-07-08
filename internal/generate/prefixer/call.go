package prefixer

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"
)

// callPrefixer implements Prefixer for CallExpr.
//
// It's mainly used for constant array length calls (e.g. [unsafe.Sizeof(uint64(0))]byte).
type callPrefixer struct {
	*ast.CallExpr

	ArgsPrefixers []Prefixer
	FunPrefixer   Prefixer
}

var _ Prefixer = &callPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (c *callPrefixer) Valid() error {
	errs := make([]error, 0, len(c.Args)+1)

	c.FunPrefixer = NewPrefixer(c.Fun)
	errs = append(errs, c.FunPrefixer.Valid())

	c.ArgsPrefixers = make([]Prefixer, 0, len(c.Args))
	for _, arg := range c.Args {
		prefixer := NewPrefixer(arg)
		errs = append(errs, prefixer.Valid())
		c.ArgsPrefixers = append(c.ArgsPrefixers, prefixer)
	}

	return errors.Join(errs...)
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (c *callPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	fun, exported := c.FunPrefixer.ToString(sourcePackage, typeParams)

	args := make([]string, 0, len(c.ArgsPrefixers))
	for _, arg := range c.ArgsPrefixers {
		stringType, argExported := arg.ToString(sourcePackage, typeParams)

		// don't affect directly because once exported is false, it should stays even if other args are exported
		if !argExported {
			exported = false
		}

		args = append(args, stringType)
	}

	if c.Ellipsis.IsValid() {
		args[len(args)-1] += "..."
	}

	return fmt.Sprintf("%s%s(%s)", strings.Join(prefixes, ""), fun, strings.Join(args, ", ")), exported
}
