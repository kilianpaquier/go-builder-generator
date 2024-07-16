// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"github.com/kilianpaquier/go-builder-generator/testdata/success_generic"
)

// AliasGenericBuilder represents AliasGeneric's builder.
type AliasGenericBuilder[t any, GV success_generic.GenericValue] struct {
	build success_generic.AliasGeneric[t, GV]
}

// NewAliasGenericBuilder creates a new AliasGenericBuilder.
func NewAliasGenericBuilder[t any, GV success_generic.GenericValue]() *AliasGenericBuilder[t, GV] {
	return &AliasGenericBuilder[t, GV]{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *AliasGenericBuilder[t, GV]) Copy() *AliasGenericBuilder[t, GV] {
	return &AliasGenericBuilder[t, GV]{b.build}
}

// Build returns built AliasGeneric.
func (b *AliasGenericBuilder[t, GV]) Build() *success_generic.AliasGeneric[t, GV] {
	result := b.build
	return &result
}

// ValueT sets AliasGeneric's ValueT.
func (b *AliasGenericBuilder[t, GV]) ValueT(valueT t) *AliasGenericBuilder[t, GV] {
	b.build.ValueT = valueT
	return b
}

// ValueX sets AliasGeneric's ValueX.
func (b *AliasGenericBuilder[t, GV]) ValueX(valueX GV) *AliasGenericBuilder[t, GV] {
	b.build.ValueX = valueX
	return b
}