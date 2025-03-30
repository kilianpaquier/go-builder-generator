// Code generated from testdata/success_with_options/types.go.

package builders

import (
	"fmt"

	"github.com/kilianpaquier/go-builder-generator/testdata/success_with_options"
)

// OptionsBuilder represents Options's builder.
type OptionsBuilder struct {
	build success_with_options.Options
}

// NewOptionsBuilder creates a new OptionsBuilder.
func NewOptionsBuilder() *OptionsBuilder {
	return &OptionsBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *OptionsBuilder) Copy() *OptionsBuilder {
	return &OptionsBuilder{b.build}
}

// Build returns built Options.
func (b *OptionsBuilder) Build() (*success_with_options.Options, error) {
	b = b.GetDefaultString().GetDefaultContext().SetDefaultForceFuncName()

	result := b.build
	if err := result.Validate(); err != nil {
		return nil, fmt.Errorf("validation of 'Options''s struct: %w", err)
	}
	return &result, nil
}

// SetDefaultFieldFunc sets Options's DefaultFieldFunc.
func (b *OptionsBuilder) SetDefaultFieldFunc(defaultFieldFunc string) *OptionsBuilder {
	b = b.Copy()
	b.build.DefaultFieldFunc = defaultFieldFunc
	return b
}

// FooBarForced sets Options's ForceFuncName.
func (b *OptionsBuilder) FooBarForced(forceFuncName string) *OptionsBuilder {
	b = b.Copy()
	b.build.ForceFuncName = forceFuncName
	return b
}

// FooBarForceWithDefault sets Options's ForceFuncNameWithDefaultFunc.
func (b *OptionsBuilder) FooBarForceWithDefault(forceFuncNameWithDefaultFunc string) *OptionsBuilder {
	b = b.Copy()
	b.build.ForceFuncNameWithDefaultFunc = forceFuncNameWithDefaultFunc
	return b
}

// SetPtrField sets Options's PtrField.
func (b *OptionsBuilder) SetPtrField(ptrField *string) *OptionsBuilder {
	b = b.Copy()
	b.build.PtrField = ptrField
	return b
}

// SetPtrFieldValidatedToo sets Options's PtrFieldValidatedToo.
func (b *OptionsBuilder) SetPtrFieldValidatedToo(ptrFieldValidatedToo string) *OptionsBuilder {
	b = b.Copy()
	b.build.PtrFieldValidatedToo = &ptrFieldValidatedToo
	return b
}

// SetSimpleFieldAppend sets Options's SimpleFieldAppend.
func (b *OptionsBuilder) SetSimpleFieldAppend(simpleFieldAppend success_with_options.Int64Alias) *OptionsBuilder {
	b = b.Copy()
	b.build.SimpleFieldAppend = simpleFieldAppend
	return b
}

// SetSliceFieldAliasAppend sets Options's SliceFieldAliasAppend.
func (b *OptionsBuilder) SetSliceFieldAliasAppend(sliceFieldAliasAppend ...success_with_options.Int64Alias) *OptionsBuilder {
	b = b.Copy()
	b.build.SliceFieldAliasAppend = append(b.build.SliceFieldAliasAppend, sliceFieldAliasAppend...)
	return b
}

// SetSliceFieldAppend sets Options's SliceFieldAppend.
func (b *OptionsBuilder) SetSliceFieldAppend(sliceFieldAppend ...int64) *OptionsBuilder {
	b = b.Copy()
	b.build.SliceFieldAppend = append(b.build.SliceFieldAppend, sliceFieldAppend...)
	return b
}

// SetSliceFieldAppendPtr sets Options's SliceFieldAppendPtr.
func (b *OptionsBuilder) SetSliceFieldAppendPtr(sliceFieldAppendPtr ...*int64) *OptionsBuilder {
	b = b.Copy()
	b.build.SliceFieldAppendPtr = append(b.build.SliceFieldAppendPtr, sliceFieldAppendPtr...)
	return b
}

// SetSliceFieldNoPtrAppend sets Options's SliceFieldNoPtrAppend.
func (b *OptionsBuilder) SetSliceFieldNoPtrAppend(sliceFieldNoPtrAppend ...int64) *OptionsBuilder {
	b = b.Copy()
	*b.build.SliceFieldNoPtrAppend = append(*b.build.SliceFieldNoPtrAppend, sliceFieldNoPtrAppend...)
	return b
}

// SetSliceFieldPtrAppend sets Options's SliceFieldPtrAppend.
func (b *OptionsBuilder) SetSliceFieldPtrAppend(sliceFieldPtrAppend ...int64) *OptionsBuilder {
	b = b.Copy()
	*b.build.SliceFieldPtrAppend = append(*b.build.SliceFieldPtrAppend, sliceFieldPtrAppend...)
	return b
}
