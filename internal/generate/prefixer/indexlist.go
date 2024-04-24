package prefixer

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"
)

// indexPrefixer implements Prefixer for IndexListExpr.
type indexListPrefixer struct {
	*ast.IndexListExpr

	XPrefixer        Prefixer
	IndicesPrefixers []Prefixer
}

var _ Prefixer = &indexListPrefixer{}

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (i *indexListPrefixer) Valid() error {
	errs := make([]error, 0, len(i.Indices)+1)

	i.XPrefixer = NewPrefixer(i.X)
	errs = append(errs, i.XPrefixer.Valid())

	i.IndicesPrefixers = make([]Prefixer, 0, len(i.Indices))
	for _, indice := range i.Indices {
		prefixer := NewPrefixer(indice)

		errs = append(errs, prefixer.Valid())
		i.IndicesPrefixers = append(i.IndicesPrefixers, prefixer)
	}

	return errors.Join(errs...)
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (i *indexListPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	x, xExported := i.XPrefixer.ToString(sourcePackage, typeParams, prefixes...)

	exported := true
	indices := make([]string, 0, len(i.IndicesPrefixers))
	for _, indice := range i.IndicesPrefixers {
		stringType, indiceExported := indice.ToString(sourcePackage, typeParams)

		// don't affect directly because once exported is false, it should stays even if other fields are exported
		if !indiceExported {
			exported = false
		}

		indices = append(indices, stringType)
	}

	return fmt.Sprintf("%s[%s]", x, strings.Join(indices, ", ")), xExported && exported
}
