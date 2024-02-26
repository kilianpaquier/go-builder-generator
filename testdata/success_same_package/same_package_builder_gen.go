// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package success_same_package

import (
	"context"
)

// SamePackageBuilder is an alias of SamePackage to build SamePackage with builder-pattern.
type SamePackageBuilder SamePackage

// NewSamePackageBuilder creates a new SamePackageBuilder.
func NewSamePackageBuilder() *SamePackageBuilder {
	return &SamePackageBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *SamePackageBuilder) Copy() *SamePackageBuilder {
	c := *b
	return &c
}

// Build returns built SamePackage.
func (b *SamePackageBuilder) Build() *SamePackage {
	c := (SamePackage)(*b)
	return &c
}

// SetCtx sets SamePackage's Ctx.
func (b *SamePackageBuilder) SetCtx(ctx context.Context) *SamePackageBuilder {
	b.Ctx = ctx
	return b
}

// SetInt64Alias sets SamePackage's Int64Alias.
func (b *SamePackageBuilder) SetInt64Alias(int64Alias Int64Alias) *SamePackageBuilder {
	b.Int64Alias = int64Alias
	return b
}

// SetPrimitive sets SamePackage's Primitive.
func (b *SamePackageBuilder) SetPrimitive(primitive string) *SamePackageBuilder {
	b.Primitive = primitive
	return b
}
