package prefixer

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"
)

// funcPrefixer implements Prefixer for FuncType.
type funcPrefixer struct {
	*ast.FuncType

	ParamsPrefixers  []Prefixer
	ResultsPrefixers []Prefixer
}

var _ Prefixer = &funcPrefixer{} // ensure interface is implemented

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (f *funcPrefixer) Valid() error {
	var errs []error

	editor := func(field *ast.Field) editor {
		return func(stringType string, exported bool) (string, bool) {
			var name string
			if len(field.Names) > 0 {
				name = field.Names[0].Name + " "
			}
			return name + stringType, exported
		}
	}

	// retrieve prefixers associated to func parameters
	if f.Params != nil {
		f.ParamsPrefixers = make([]Prefixer, 0, len(f.Params.List))
		for _, field := range f.Params.List {
			prefixer := NewPrefixerEditor(NewPrefixer(field.Type), editor(field))
			errs = append(errs, prefixer.Valid())
			f.ParamsPrefixers = append(f.ParamsPrefixers, prefixer)
		}
	}

	// retrieve prefixers associated to func outputs
	if f.Results != nil {
		f.ResultsPrefixers = make([]Prefixer, 0, len(f.Params.List))
		for _, field := range f.Results.List {
			prefixer := NewPrefixerEditor(NewPrefixer(field.Type), editor(field))
			errs = append(errs, prefixer.Valid())
			f.ResultsPrefixers = append(f.ResultsPrefixers, prefixer)
		}
	}

	return errors.Join(errs...)
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (f *funcPrefixer) ToString(sourcePackage string, typeParams []string, prefixes ...string) (_ string, _ bool) {
	exported := true

	// compute paramsTypes prefix part
	paramsTypes := make([]string, 0, len(f.ParamsPrefixers))
	for _, field := range f.ParamsPrefixers {
		stringType, paramExported := field.ToString(sourcePackage, typeParams)

		// don't affect directly because once exported is false, it should stays even if other fields are exported
		if !paramExported {
			exported = false
		}

		paramsTypes = append(paramsTypes, stringType)
	}

	// compute resultsTypes prefix part
	resultsTypes := make([]string, 0, len(f.ResultsPrefixers))
	for _, field := range f.ResultsPrefixers {
		stringType, resultExported := field.ToString(sourcePackage, typeParams)

		// don't affect directly because once exported is false, it should stays even if other fields are exported
		if !resultExported {
			exported = false
		}

		resultsTypes = append(resultsTypes, stringType)
	}

	return fmt.Sprintf(
		"%sfunc(%s) (%s)",
		strings.Join(prefixes, ""),
		strings.Join(paramsTypes, ", "),
		strings.Join(resultsTypes, ", "),
	), exported
}
