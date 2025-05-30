// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

// Code generated from success_maps/types.go.

package builders

import (
	"context"

	"github.com/kilianpaquier/go-builder-generator/testdata/success_maps"
)

// MapBuilder represents Map's builder.
type MapBuilder struct {
	build success_maps.Map
}

// NewMapBuilder creates a new MapBuilder.
func NewMapBuilder() *MapBuilder {
	return &MapBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *MapBuilder) Copy() *MapBuilder {
	return &MapBuilder{b.build}
}

// Build returns built Map.
func (b *MapBuilder) Build() *success_maps.Map {
	result := b.build
	return &result
}

// MapField sets Map's MapField.
func (b *MapBuilder) MapField(mapField map[int64]string) *MapBuilder {
	b.build.MapField = mapField
	return b
}

// MapFieldAlias sets Map's MapFieldAlias.
func (b *MapBuilder) MapFieldAlias(mapFieldAlias map[success_maps.Int64Alias]success_maps.FuncAlias) *MapBuilder {
	b.build.MapFieldAlias = mapFieldAlias
	return b
}

// MapFieldFunc sets Map's MapFieldFunc.
func (b *MapBuilder) MapFieldFunc(mapFieldFunc map[int64]func(in int64) error) *MapBuilder {
	b.build.MapFieldFunc = mapFieldFunc
	return b
}

// MapFieldPtrAlias sets Map's MapFieldPtrAlias.
func (b *MapBuilder) MapFieldPtrAlias(mapFieldPtrAlias map[*success_maps.Int64Alias]*success_maps.FuncAlias) *MapBuilder {
	b.build.MapFieldPtrAlias = &mapFieldPtrAlias
	return b
}

// MapFieldWithPkg sets Map's MapFieldWithPkg.
func (b *MapBuilder) MapFieldWithPkg(mapFieldWithPkg map[int64]context.Context) *MapBuilder {
	b.build.MapFieldWithPkg = mapFieldWithPkg
	return b
}
