// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"github.com/kilianpaquier/go-builder-generator/testdata/success_generic"
)

// ComplexGenericBuilder represents ComplexGeneric's builder.
type ComplexGenericBuilder[T success_generic.Struct, Y comparable] struct {
	build success_generic.ComplexGeneric[T, Y]
}

// NewComplexGenericBuilder creates a new ComplexGenericBuilder.
func NewComplexGenericBuilder[T success_generic.Struct, Y comparable]() *ComplexGenericBuilder[T, Y] {
	return &ComplexGenericBuilder[T, Y]{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *ComplexGenericBuilder[T, Y]) Copy() *ComplexGenericBuilder[T, Y] {
	return &ComplexGenericBuilder[T, Y]{b.build}
}

// Build returns built ComplexGeneric.
func (b *ComplexGenericBuilder[T, Y]) Build() *success_generic.ComplexGeneric[T, Y] {
	result := b.build
	return &result
}

// AnonymousInterface sets ComplexGeneric's AnonymousInterface.
func (b *ComplexGenericBuilder[T, Y]) AnonymousInterface(anonymousInterface interface {
	HeyFunc(T) error
	HeyFuncMulti(success_generic.AliasGeneric[T, success_generic.GenericValue]) success_generic.AliasGeneric[T, success_generic.GenericValue]
}) *ComplexGenericBuilder[T, Y] {
	b.build.AnonymousInterface = anonymousInterface
	return b
}

// FuncT sets ComplexGeneric's FuncT.
func (b *ComplexGenericBuilder[T, Y]) FuncT(funcT func(T, success_generic.SimpleGeneric[T]) (T, error)) *ComplexGenericBuilder[T, Y] {
	b.build.FuncT = funcT
	return b
}

// ValueT sets ComplexGeneric's ValueT.
func (b *ComplexGenericBuilder[T, Y]) ValueT(valueT map[Y]success_generic.AliasGeneric[T, success_generic.GenericValue]) *ComplexGenericBuilder[T, Y] {
	b.build.ValueT = valueT
	return b
}

// ValueY sets ComplexGeneric's ValueY.
func (b *ComplexGenericBuilder[T, Y]) ValueY(valueY [10]*[]success_generic.SimpleGeneric[Y]) *ComplexGenericBuilder[T, Y] {
	b.build.ValueY = valueY
	return b
}