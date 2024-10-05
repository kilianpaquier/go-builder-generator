package prefixer

import (
	"go/ast"
)

// chanPrefixer implements Prefixer for ChanType.
type chanPrefixer struct {
	*ast.ChanType

	ValuePrefixer Prefixer
}

var _ Prefixer = &chanPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (c *chanPrefixer) Valid() error {
	c.ValuePrefixer = NewPrefixer(c.Value)
	return c.ValuePrefixer.Valid() //nolint:wrapcheck
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (c *chanPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	prefix := func() string {
		if c.Dir == ast.SEND {
			return "chan<- "
		}
		if c.Dir == ast.RECV {
			return "<-chan "
		}
		return "chan "
	}()
	return c.ValuePrefixer.ToString(sourcePackage, typeParams, append(prefixes, prefix)...)
}
