package prefixer

import (
	"fmt"
	"go/ast"
	"slices"
	"strings"

	"github.com/kilianpaquier/go-builder-generator/internal/generate/models"
)

// identPrefixer implements Prefixer for Ident.
type identPrefixer ast.Ident

var _ Prefixer = &identPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (*identPrefixer) Valid() error {
	return nil
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (i *identPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	// field type is a primary type, no package prefix is needed
	// field type is a reserved name (dyanmic type param - generic), no package prefix is needed
	if slices.Contains(models.PrimaryTypes(), i.Name) || slices.Contains(typeParams, i.Name) {
		return strings.Join(prefixes, "") + i.Name, true
	}

	// compute final type with package prefix if necessary
	finalType := func() string {
		if sourcePackage == "" {
			return strings.Join(prefixes, "") + i.Name
		}
		return fmt.Sprint(strings.Join(prefixes, ""), sourcePackage, ".", i.Name)
	}()
	return finalType, ast.IsExported(i.Name)
}
