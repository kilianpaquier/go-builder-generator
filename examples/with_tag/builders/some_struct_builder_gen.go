// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

// Code generated from with_tag/types.go.

package builders

import "github.com/kilianpaquier/go-builder-generator/examples/with_tag"

// SomeStructBuilder represents SomeStruct's builder.
type SomeStructBuilder struct {
	build with_tag.SomeStruct
}

// NewSomeStructBuilder creates a new SomeStructBuilder.
func NewSomeStructBuilder() *SomeStructBuilder {
	return &SomeStructBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *SomeStructBuilder) Copy() *SomeStructBuilder {
	return &SomeStructBuilder{b.build}
}

// Build returns built SomeStruct.
func (b *SomeStructBuilder) Build() *with_tag.SomeStruct {
	b = b.SetTheChan()

	result := b.build
	return &result
}

// SomeSlice sets SomeStruct's SomeSlice.
func (b *SomeStructBuilder) SomeSlice(someSlice ...string) *SomeStructBuilder {
	b.build.SomeSlice = append(b.build.SomeSlice, someSlice...)
	return b
}
