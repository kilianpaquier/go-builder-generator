package prefixer

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"
)

// mapPrefixer implements Prefixer for MapType.
type mapPrefixer struct {
	*ast.MapType

	KeyPrefixer   Prefixer
	ValuePrefixer Prefixer
}

var _ Prefixer = &mapPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (m *mapPrefixer) Valid() error {
	var errs []error //nolint:prealloc

	// retrieve map key prefixer
	m.KeyPrefixer = NewPrefixer(m.Key)
	errs = append(errs, m.KeyPrefixer.Valid())

	// retrieve map value prefixer
	m.ValuePrefixer = NewPrefixer(m.Value)
	errs = append(errs, m.ValuePrefixer.Valid())

	return errors.Join(errs...)
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (m *mapPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	// compute map key type
	key, keyExported := m.KeyPrefixer.ToString(sourcePackage, typeParams)

	// compute map value type
	value, valueExported := m.ValuePrefixer.ToString(sourcePackage, typeParams)

	// to be exported, both key and value of map must be exported
	exported := keyExported && valueExported

	stringType := fmt.Sprintf("%smap[%s]%s", strings.Join(prefixes, ""), key, value)
	return stringType, exported
}
