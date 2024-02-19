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

	InputFields  []Prefixer
	OutputFields []Prefixer
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
		f.InputFields = make([]Prefixer, 0, len(f.Params.List))
		for _, field := range f.Params.List {
			prefixer := NewPrefixerEditor(NewPrefixer(field.Type), editor(field))
			errs = append(errs, prefixer.Valid())
			f.InputFields = append(f.InputFields, prefixer)
		}
	}

	// retrieve prefixers associated to func outputs
	if f.Results != nil {
		f.OutputFields = make([]Prefixer, 0, len(f.Params.List))
		for _, field := range f.Results.List {
			prefixer := NewPrefixerEditor(NewPrefixer(field.Type), editor(field))
			errs = append(errs, prefixer.Valid())
			f.OutputFields = append(f.OutputFields, prefixer)
		}
	}

	return errors.Join(errs...)
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (f *funcPrefixer) ToString(sourcePackage string, prefixes ...string) (_ string, _ bool) {
	exported := true

	// compute inputs prefix part
	inputs := make([]string, 0, len(f.InputFields))
	for _, field := range f.InputFields {
		stringType, fieldExported := field.ToString(sourcePackage)

		// don't affect directly because once exported is false, it should stays even if other fields are exported
		if !fieldExported {
			exported = false
		}

		inputs = append(inputs, stringType)
	}

	// compute outputs prefix part
	outputs := make([]string, 0, len(f.OutputFields))
	for _, field := range f.OutputFields {
		stringType, fieldExported := field.ToString(sourcePackage)

		// don't affect directly because once exported is false, it should stays even if other fields are exported
		if !fieldExported {
			exported = false
		}

		outputs = append(outputs, stringType)
	}

	return fmt.Sprintf(
		"%sfunc(%s) (%s)",
		strings.Join(prefixes, ""),
		strings.Join(inputs, ", "),
		strings.Join(outputs, ", "),
	), exported
}
