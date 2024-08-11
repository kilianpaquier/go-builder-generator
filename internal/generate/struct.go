package generate

import (
	"cmp"
	"errors"
	"fmt"
	"go/ast"
	"slices"
	"strings"

	"github.com/samber/lo"

	"github.com/kilianpaquier/go-builder-generator/internal/generate/prefixer"
)

// parseStruct parses and returns a gen data (to generate _gen file). Computing is done from input ast type spec and struct type.
func parseStruct(typeSpec *ast.TypeSpec, structType *ast.StructType, pkg packagesData, opts CLIOptions) (genData, error) {
	var errs []error

	// initialize builder to avoid too many params in generateStruct
	builder := genData{
		Name:     typeSpec.Name.String(),
		Opts:     opts,
		Packages: pkg,
	}

	// compute generic params for builder structure
	var genericNames []string
	if typeSpec.TypeParams != nil {
		// first a first time over all type params (generic types) to retrieve only the names
		// in case some type params depends on the others, we must build this slice before computing any field
		genericNames = make([]string, len(typeSpec.TypeParams.List))
		for i, typeParam := range typeSpec.TypeParams.List {
			if len(typeParam.Names) == 0 {
				errs = append(errs, fmt.Errorf("struct '%s' has unnamed type parameter (generic parameter)", builder.Name))
			}
			genericNames[i] = typeParam.Names[0].Name
		}

		// build final type params with type prefixing if it applies
		builder.TypeParams = make([]field, 0, len(typeSpec.TypeParams.List))
		for _, typeParam := range lo.Zip2(genericNames, typeSpec.TypeParams.List) {
			// retrieve typePrefixer for field type
			typePrefixer := prefixer.NewPrefixer(typeParam.B.Type)
			if err := typePrefixer.Valid(); err != nil {
				errs = append(errs, fmt.Errorf("field validation: %w", err))
				continue
			}

			// retrieve computed string type
			finalType, exported := typePrefixer.ToString(pkg.SourceName, genericNames)
			field := field{
				AlteredType: finalType,
				Exported:    exported,
				InitialType: finalType,
				Name:        typeParam.A,
				ParamName:   paramName(typeParam.A),
			}
			builder.TypeParams = append(builder.TypeParams, field)
		}
	}
	genericExported := lo.CountBy(builder.TypeParams, func(param field) bool { return param.Exported }) == len(builder.TypeParams)

	// add an error if destination package is not the same as the source one
	// and the struct to generate is not exported
	builder.Exported = genericExported && ast.IsExported(builder.Name)
	if pkg.SourceName != "" && !builder.Exported {
		errs = append(errs, fmt.Errorf("%s is not exported (or one of its generic params is not) but generation destination is in an external package", builder.Name))
	}

	// compute all fields associated to builder
	builder.Fields = make([]field, 0, len(structType.Fields.List))
	for _, astField := range structType.Fields.List {
		field, err := parseField(astField, pkg.SourceName, genericNames)
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
