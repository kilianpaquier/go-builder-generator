// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

//go:generate go run github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@latest generate -d . -f ../types.go -s InterfaceNoFields

package builders

import (
	"github.com/kilianpaquier/go-builder-generator/testdata/success_interface"
)

// InterfaceNoFieldsBuilder represents InterfaceNoFields's builder.
type InterfaceNoFieldsBuilder struct {
	build success_interface.InterfaceNoFields
}

// NewInterfaceNoFieldsBuilder creates a new InterfaceNoFieldsBuilder.
func NewInterfaceNoFieldsBuilder() *InterfaceNoFieldsBuilder {
	return &InterfaceNoFieldsBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *InterfaceNoFieldsBuilder) Copy() *InterfaceNoFieldsBuilder {
	return &InterfaceNoFieldsBuilder{b.build}
}

// Build returns built InterfaceNoFields.
func (b *InterfaceNoFieldsBuilder) Build() *success_interface.InterfaceNoFields {
	result := b.build
	return &result
}

// NoFields sets InterfaceNoFields's NoFields.
func (b *InterfaceNoFieldsBuilder) NoFields(noFields interface{}) *InterfaceNoFieldsBuilder {
	b.build.NoFields = noFields
	return b
}
