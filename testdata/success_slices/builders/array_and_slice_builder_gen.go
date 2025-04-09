// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

// Code generated from success_slices/types.go.

package builders

import (
	"context"

	"github.com/kilianpaquier/go-builder-generator/testdata/success_slices"
)

// ArrayAndSliceBuilder represents ArrayAndSlice's builder.
type ArrayAndSliceBuilder struct {
	build success_slices.ArrayAndSlice
}

// NewArrayAndSliceBuilder creates a new ArrayAndSliceBuilder.
func NewArrayAndSliceBuilder() *ArrayAndSliceBuilder {
	return &ArrayAndSliceBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *ArrayAndSliceBuilder) Copy() *ArrayAndSliceBuilder {
	return &ArrayAndSliceBuilder{b.build}
}

// Build returns built ArrayAndSlice.
func (b *ArrayAndSliceBuilder) Build() *success_slices.ArrayAndSlice {
	result := b.build
	return &result
}

// ArrayField sets ArrayAndSlice's ArrayField.
func (b *ArrayAndSliceBuilder) ArrayField(arrayField [10]int64) *ArrayAndSliceBuilder {
	b.build.ArrayField = arrayField
	return b
}

// ArrayFieldAlias sets ArrayAndSlice's ArrayFieldAlias.
func (b *ArrayAndSliceBuilder) ArrayFieldAlias(arrayFieldAlias [10]success_slices.Int64Alias) *ArrayAndSliceBuilder {
	b.build.ArrayFieldAlias = arrayFieldAlias
	return b
}

// ArrayFieldPtrAlias sets ArrayAndSlice's ArrayFieldPtrAlias.
func (b *ArrayAndSliceBuilder) ArrayFieldPtrAlias(arrayFieldPtrAlias [10]*success_slices.Int64Alias) *ArrayAndSliceBuilder {
	b.build.ArrayFieldPtrAlias = &arrayFieldPtrAlias
	return b
}

// ArrayFieldWithPkg sets ArrayAndSlice's ArrayFieldWithPkg.
func (b *ArrayAndSliceBuilder) ArrayFieldWithPkg(arrayFieldWithPkg [10]context.Context) *ArrayAndSliceBuilder {
	b.build.ArrayFieldWithPkg = arrayFieldWithPkg
	return b
}

// SliceField sets ArrayAndSlice's SliceField.
func (b *ArrayAndSliceBuilder) SliceField(sliceField []int64) *ArrayAndSliceBuilder {
	b.build.SliceField = sliceField
	return b
}

// SliceFieldAlias sets ArrayAndSlice's SliceFieldAlias.
func (b *ArrayAndSliceBuilder) SliceFieldAlias(sliceFieldAlias []success_slices.Int64Alias) *ArrayAndSliceBuilder {
	b.build.SliceFieldAlias = sliceFieldAlias
	return b
}

// SliceFieldAliasChan sets ArrayAndSlice's SliceFieldAliasChan.
func (b *ArrayAndSliceBuilder) SliceFieldAliasChan(sliceFieldAliasChan []chan<- success_slices.Int64Alias) *ArrayAndSliceBuilder {
	b.build.SliceFieldAliasChan = sliceFieldAliasChan
	return b
}

// SliceFieldPtrAlias sets ArrayAndSlice's SliceFieldPtrAlias.
func (b *ArrayAndSliceBuilder) SliceFieldPtrAlias(sliceFieldPtrAlias []*success_slices.Int64Alias) *ArrayAndSliceBuilder {
	b.build.SliceFieldPtrAlias = &sliceFieldPtrAlias
	return b
}

// SliceFieldWithPkg sets ArrayAndSlice's SliceFieldWithPkg.
func (b *ArrayAndSliceBuilder) SliceFieldWithPkg(sliceFieldWithPkg []context.Context) *ArrayAndSliceBuilder {
	b.build.SliceFieldWithPkg = sliceFieldWithPkg
	return b
}
