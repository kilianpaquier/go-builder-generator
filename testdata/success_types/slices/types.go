package testdata

import (
	"context"
	"unsafe"
)

//go:generate ../../../go-builder-generator generate -f types.go -s ArrayAndSlice -d builders

type Int64Alias int64

var SumInts = sumInts

var IntSlice = intSlice

func sumInts(first int, nums ...int) int { return len(nums) + 1 }

var intSlice = []int{1, 2, 3}

type ArrayAndSlice struct {
	ArrayField         [10]int64
	ArrayFieldAlias    [10]Int64Alias
	ArrayFieldPtrAlias *[10]*Int64Alias
	ArrayFieldWithPkg  [10]context.Context

	ArrayBinaryLen         [2 * 3]int64
	ArrayCallLen           [unsafe.Sizeof(uint64(0))]byte
	ArrayCallEllipsisLen   [unsafe.Sizeof(SumInts(55, IntSlice...))]byte
	ArrayCallUnexportedLen [unsafe.Sizeof(sumInts(63, intSlice...))]byte

	SliceField          []int64
	SliceFieldAlias     []Int64Alias
	SliceFieldAliasChan []chan<- Int64Alias
	SliceFieldPtrAlias  *[]*Int64Alias
	SliceFieldWithPkg   []context.Context
}
