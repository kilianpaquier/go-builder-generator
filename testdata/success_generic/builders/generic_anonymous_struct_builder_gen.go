// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

// Code generated from success_generic/types.go.

package builders

import "github.com/kilianpaquier/go-builder-generator/testdata/success_generic"

// GenericAnonymousStructBuilder represents GenericAnonymousStruct's builder.
type GenericAnonymousStructBuilder[T struct {
	Property string `builder:"pointer"`
}] struct {
	build success_generic.GenericAnonymousStruct[T]
}

// NewGenericAnonymousStructBuilder creates a new GenericAnonymousStructBuilder.
func NewGenericAnonymousStructBuilder[T struct {
	Property string `builder:"pointer"`
}]() *GenericAnonymousStructBuilder[T] {
	return &GenericAnonymousStructBuilder[T]{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *GenericAnonymousStructBuilder[T]) Copy() *GenericAnonymousStructBuilder[T] {
	return &GenericAnonymousStructBuilder[T]{b.build}
}

// Build returns built GenericAnonymousStruct.
func (b *GenericAnonymousStructBuilder[T]) Build() *success_generic.GenericAnonymousStruct[T] {
	result := b.build
	return &result
}

// Prop sets GenericAnonymousStruct's Prop.
func (b *GenericAnonymousStructBuilder[T]) Prop(prop T) *GenericAnonymousStructBuilder[T] {
	b.build.Prop = prop
	return b
}
