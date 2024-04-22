package prefixer

import (
	"go/ast"
)

// Prefixer represents an interface to transform a specific ast.Expr into its string representation.
type Prefixer interface {
	// Valid validates the prefixer and its subprefixers.
	//
	// An example would be a composition of a StarExpr with an ArrayType of an Ident.
	// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
	Valid() error

	// ToString transforms a Prefixer (ast.Expr) into its string representation.
	// It also returns a boolean indicating whether the type is exported.
	ToString(sourcePackage string, prefixes ...string) (stringType string, exported bool)
}

// NewPrefixer transforms the input expression into it's corresponding Prefixer interface.
//
// It returns an error if the input expression hasn't a Prefixer implementation.
func NewPrefixer(input ast.Expr) Prefixer {
	switch expr := input.(type) {
	// field type is *...
	case *ast.StarExpr:
		return &ptrPrefixer{StarExpr: expr}

	// field type is []... or [X]...
	case *ast.ArrayType:
		return &arrayPrefixer{ArrayType: expr}

	// field type is 'chan ...' or '<-chan ...' or 'chan<- ...'
	case *ast.ChanType:
		return &chanPrefixer{ChanType: expr}

	// field type is func(...) ...
	case *ast.FuncType:
		return &funcPrefixer{FuncType: expr}

	// field type is an anonymous interface{...}
	case *ast.InterfaceType:
		return &interfacePrefixer{InterfaceType: expr}

	// field type is map[...]...
	case *ast.MapType:
		return &mapPrefixer{MapType: expr}

	// field type is an anonymous struct{...}
	case *ast.StructType:
		return &structPrefixer{StructType: expr}

	// field type is a "simple" type (either primitive or current package) (string, int64, MyStruct, MyAlias, etc.)
	case *ast.Ident:
		return (*identPrefixer)(expr)

	// field type is a "simple" type coming from another package (package.MyStruct, package.MyAlias, etc.)
	case *ast.SelectorExpr:
		return &selectorPrefixer{SelectorExpr: expr}

	// field is just a string
	case *ast.BasicLit:
		return (*basicLitPrefixer)(expr)

	// field is a generic type
	case *ast.IndexExpr:
		return &genericsPrefixer{IndexExpr: expr}
	}

	// any other unanticipated types that could exist
	return &unimplementedPrefixer{Expr: input}
}
