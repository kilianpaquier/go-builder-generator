package generate

import (
	"cmp"
	"errors"
	"fmt"
	"go/ast"
	"slices"
	"strings"

	"github.com/samber/lo"
)

// parseStruct parses and returns a gen data (to generate _gen file). Computing is done from input ast type spec and struct type.
func parseStruct(typeSpec *ast.TypeSpec, structType *ast.StructType, pkg packageData, opts CLIOptions) (genData, error) {
	var errs []error

	// initialize builder to avoid too many params in generateStruct
	builder := genData{
		Name:    typeSpec.Name.String(),
		Opts:    opts,
		Package: pkg,
	}

	// compute generic params for builder structure
	if typeSpec.TypeParams != nil {
		builder.TypeParams = make([]field, 0, len(typeSpec.TypeParams.List))
		for _, typeParam := range typeSpec.TypeParams.List {
			field, err := parseField(typeParam, pkg.SourceName, nil)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			builder.TypeParams = append(builder.TypeParams, field)
		}
	}
	typeParamsNames := lo.Map(builder.TypeParams, func(typeParam field, _ int) string { return typeParam.Name })
	typeParamsExported := lo.CountBy(builder.TypeParams, func(param field) bool { return param.Exported }) == len(builder.TypeParams)

	// add an error if destination package is not the same as the source one
	// and the struct to generate is not exported
	builder.Exported = typeParamsExported && ast.IsExported(builder.Name)
	if pkg.SourceName != "" && !builder.Exported {
		errs = append(errs, fmt.Errorf("%s is not exported (or one of its generic params is not) but generation destination is in an external package", builder.Name))
	}

	// compute all fields associated to builder
	builder.Fields = make([]field, 0, len(structType.Fields.List))
	for _, astField := range structType.Fields.List {
		field, err := parseField(astField, pkg.SourceName, typeParamsNames)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		// move up default_func option one function could be use for multiple fields
		if field.Opts.DefaultFunc != "" && !slices.Contains(builder.DefaultFuncs, field.Opts.DefaultFunc) {
			builder.DefaultFuncs = append(builder.DefaultFuncs, field.Opts.DefaultFunc)
		}

		builder.Fields = append(builder.Fields, field)
	}

	// sort fields to have the same generation even if fields order change
	slices.SortStableFunc(builder.Fields, func(f1, f2 field) int {
		return cmp.Compare(strings.ToLower(f1.Name), strings.ToLower(f2.Name))
	})
	return builder, errors.Join(errs...)
}
