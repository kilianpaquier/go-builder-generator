package success_slices

import "context"

//go:generate go tool go-builder-generator generate -f types.go -s ArrayAndSlice -d builders

type Int64Alias int64

type ArrayAndSlice struct {
	ArrayField         [10]int64
	ArrayFieldAlias    [10]Int64Alias
	ArrayFieldPtrAlias *[10]*Int64Alias
	ArrayFieldWithPkg  [10]context.Context

	SliceField          []int64
	SliceFieldAlias     []Int64Alias
	SliceFieldAliasChan []chan<- Int64Alias
	SliceFieldPtrAlias  *[]*Int64Alias
	SliceFieldWithPkg   []context.Context
}
