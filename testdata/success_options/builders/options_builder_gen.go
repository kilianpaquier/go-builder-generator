// Code generated from types.go.

package my_package

import (
	"fmt"

	"github.com/kilianpaquier/go-builder-generator/testdata"
)

//go:generate go tool go-builder-generator generate -d . -f ../types.go -s Options --validate-func Validate -p set --package-name my_package --no-notice --return-copy

// OptionsBuilder represents Options's builder.
type OptionsBuilder struct {
	build testdata.Options
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
func (b *OptionsBuilder) Build() (*testdata.Options, error) {
	b = b.GetDefaultString().GetDefaultContext().SharedFunc().SetDefaultForceFuncName()

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

// SetSharedFuncA sets Options's SharedFuncA.
func (b *OptionsBuilder) SetSharedFuncA(sharedFuncA string) *OptionsBuilder {
	b = b.Copy()
	b.build.SharedFuncA = sharedFuncA
	return b
}

// SetSharedFuncB sets Options's SharedFuncB.
func (b *OptionsBuilder) SetSharedFuncB(sharedFuncB string) *OptionsBuilder {
	b = b.Copy()
	b.build.SharedFuncB = sharedFuncB
	return b
}

// SetSimpleFieldAppend sets Options's SimpleFieldAppend.
func (b *OptionsBuilder) SetSimpleFieldAppend(simpleFieldAppend testdata.Int64Alias) *OptionsBuilder {
	b = b.Copy()
	b.build.SimpleFieldAppend = simpleFieldAppend
	return b
}

// SetSliceFieldAliasAppend sets Options's SliceFieldAliasAppend.
func (b *OptionsBuilder) SetSliceFieldAliasAppend(sliceFieldAliasAppend ...testdata.Int64Alias) *OptionsBuilder {
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
