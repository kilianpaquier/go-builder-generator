// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kilianpaquier/go-builder-generator/testdata/success_with_options"
)

// OptionsBuilder is an alias of Options to build Options with builder-pattern.
type OptionsBuilder success_with_options.Options

// NewOptionsBuilder creates a new OptionsBuilder.
func NewOptionsBuilder() *OptionsBuilder {
	return &OptionsBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *OptionsBuilder) Copy() *OptionsBuilder {
	c := *b
	return &c
}

// Build returns built Options.
func (b *OptionsBuilder) Build() (*success_with_options.Options, error) {
	b = b.GetDefaultString().GetDefaultContext()
	c := (success_with_options.Options)(*b)
	if err := validator.New().Struct(c); err != nil {
		return nil, fmt.Errorf("failed to validate 'Options' struct: %w", err)
	}
	return &c, nil
}

// SetDefaultField sets Options's DefaultField.
func (b *OptionsBuilder) SetDefaultField(defaultField int64) *OptionsBuilder {
	b.DefaultField = defaultField
	return b
}

// SetDefaultFieldFunc sets Options's DefaultFieldFunc.
func (b *OptionsBuilder) SetDefaultFieldFunc(defaultFieldFunc string) *OptionsBuilder {
	b.DefaultFieldFunc = defaultFieldFunc
	return b
}

// SetPtrField sets Options's PtrField.
func (b *OptionsBuilder) SetPtrField(ptrField *string) *OptionsBuilder {
	b.PtrField = ptrField
	return b
}

// SetPtrFieldValidatedToo sets Options's PtrFieldValidatedToo.
func (b *OptionsBuilder) SetPtrFieldValidatedToo(ptrFieldValidatedToo string) *OptionsBuilder {
	b.PtrFieldValidatedToo = &ptrFieldValidatedToo
	return b
}

// SetSimpleFieldAppend sets Options's SimpleFieldAppend.
func (b *OptionsBuilder) SetSimpleFieldAppend(simpleFieldAppend success_with_options.Int64Alias) *OptionsBuilder {
	b.SimpleFieldAppend = simpleFieldAppend
	return b
}

// SetSliceFieldAliasAppend sets Options's SliceFieldAliasAppend.
func (b *OptionsBuilder) SetSliceFieldAliasAppend(sliceFieldAliasAppend ...success_with_options.Int64Alias) *OptionsBuilder {
	b.SliceFieldAliasAppend = append(b.SliceFieldAliasAppend, sliceFieldAliasAppend...)
	return b
}

// SetSliceFieldAppend sets Options's SliceFieldAppend.
func (b *OptionsBuilder) SetSliceFieldAppend(sliceFieldAppend ...int64) *OptionsBuilder {
	b.SliceFieldAppend = append(b.SliceFieldAppend, sliceFieldAppend...)
	return b
}

// SetSliceFieldAppendPtr sets Options's SliceFieldAppendPtr.
func (b *OptionsBuilder) SetSliceFieldAppendPtr(sliceFieldAppendPtr ...*int64) *OptionsBuilder {
	b.SliceFieldAppendPtr = append(b.SliceFieldAppendPtr, sliceFieldAppendPtr...)
	return b
}

// SetSliceFieldNoPtrAppend sets Options's SliceFieldNoPtrAppend.
func (b *OptionsBuilder) SetSliceFieldNoPtrAppend(sliceFieldNoPtrAppend ...int64) *OptionsBuilder {
	*b.SliceFieldNoPtrAppend = append(*b.SliceFieldNoPtrAppend, sliceFieldNoPtrAppend...)
	return b
}

// SetSliceFieldPtrAppend sets Options's SliceFieldPtrAppend.
func (b *OptionsBuilder) SetSliceFieldPtrAppend(sliceFieldPtrAppend ...int64) *OptionsBuilder {
	*b.SliceFieldPtrAppend = append(*b.SliceFieldPtrAppend, sliceFieldPtrAppend...)
	return b
}
