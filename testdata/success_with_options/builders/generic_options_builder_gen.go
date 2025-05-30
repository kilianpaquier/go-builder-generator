// Code generated from success_with_options/types.go.

package builders

import (
	"fmt"

	"github.com/kilianpaquier/go-builder-generator/testdata/success_with_options"
)

// GenericOptionsBuilder represents GenericOptions's builder.
type GenericOptionsBuilder[T any] struct {
	build success_with_options.GenericOptions[T]
}

// NewGenericOptionsBuilder creates a new GenericOptionsBuilder.
func NewGenericOptionsBuilder[T any]() *GenericOptionsBuilder[T] {
	return &GenericOptionsBuilder[T]{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *GenericOptionsBuilder[T]) Copy() *GenericOptionsBuilder[T] {
	return &GenericOptionsBuilder[T]{b.build}
}

// Build returns built GenericOptions.
func (b *GenericOptionsBuilder[T]) Build() (*success_with_options.GenericOptions[T], error) {
	b = b.GetDefaultString()

	result := b.build
	if err := result.Validate(); err != nil {
		return nil, fmt.Errorf("validation of 'GenericOptions''s struct: %w", err)
	}
	return &result, nil
}

// SetDefaultFieldFunc sets GenericOptions's DefaultFieldFunc.
func (b *GenericOptionsBuilder[T]) SetDefaultFieldFunc(defaultFieldFunc T) *GenericOptionsBuilder[T] {
	b = b.Copy()
	b.build.DefaultFieldFunc = defaultFieldFunc
	return b
}
