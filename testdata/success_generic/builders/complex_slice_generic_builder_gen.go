// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

//go:generate go run github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@latest generate -d . -f ../types.go -s ComplexSliceGeneric

package builders

import (
	"github.com/kilianpaquier/go-builder-generator/testdata/success_generic"
)

// ComplexSliceGenericBuilder represents ComplexSliceGeneric's builder.
type ComplexSliceGenericBuilder[S ~[]E, E success_generic.Struct] struct {
	build success_generic.ComplexSliceGeneric[S, E]
}

// NewComplexSliceGenericBuilder creates a new ComplexSliceGenericBuilder.
func NewComplexSliceGenericBuilder[S ~[]E, E success_generic.Struct]() *ComplexSliceGenericBuilder[S, E] {
	return &ComplexSliceGenericBuilder[S, E]{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *ComplexSliceGenericBuilder[S, E]) Copy() *ComplexSliceGenericBuilder[S, E] {
	return &ComplexSliceGenericBuilder[S, E]{b.build}
}

// Build returns built ComplexSliceGeneric.
func (b *ComplexSliceGenericBuilder[S, E]) Build() *success_generic.ComplexSliceGeneric[S, E] {
	result := b.build
	return &result
}

// ValueT sets ComplexSliceGeneric's ValueT.
func (b *ComplexSliceGenericBuilder[S, E]) ValueT(valueT func(S) E) *ComplexSliceGenericBuilder[S, E] {
	b.build.ValueT = valueT
	return b
}
