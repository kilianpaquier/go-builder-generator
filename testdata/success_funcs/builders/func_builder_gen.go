// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

// Code generated from testdata/success_funcs/types.go.

package builders

import "github.com/kilianpaquier/go-builder-generator/testdata/success_funcs"

// FuncBuilder represents Func's builder.
type FuncBuilder struct {
	build success_funcs.Func
}

// NewFuncBuilder creates a new FuncBuilder.
func NewFuncBuilder() *FuncBuilder {
	return &FuncBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *FuncBuilder) Copy() *FuncBuilder {
	return &FuncBuilder{b.build}
}

// Build returns built Func.
func (b *FuncBuilder) Build() *success_funcs.Func {
	result := b.build
	return &result
}

// FuncField sets Func's FuncField.
func (b *FuncBuilder) FuncField(funcField func(int64) string) *FuncBuilder {
	b.build.FuncField = funcField
	return b
}

// FuncFieldAlias sets Func's FuncFieldAlias.
func (b *FuncBuilder) FuncFieldAlias(funcFieldAlias func(success_funcs.Int64Alias) string) *FuncBuilder {
	b.build.FuncFieldAlias = funcFieldAlias
	return b
}

// FuncFieldAliasMultiple sets Func's FuncFieldAliasMultiple.
func (b *FuncBuilder) FuncFieldAliasMultiple(funcFieldAliasMultiple func(success_funcs.Int64Alias, success_funcs.FuncAlias) (string, error)) *FuncBuilder {
	b.build.FuncFieldAliasMultiple = funcFieldAliasMultiple
	return b
}

// FuncFieldAliasNamed sets Func's FuncFieldAliasNamed.
func (b *FuncBuilder) FuncFieldAliasNamed(funcFieldAliasNamed func(in success_funcs.Int64Alias) (out success_funcs.FuncAlias)) *FuncBuilder {
	b.build.FuncFieldAliasNamed = funcFieldAliasNamed
	return b
}

// FuncFieldChan sets Func's FuncFieldChan.
func (b *FuncBuilder) FuncFieldChan(funcFieldChan func(c chan<- success_funcs.Int64Alias) error) *FuncBuilder {
	b.build.FuncFieldChan = funcFieldChan
	return b
}

// FuncFieldCResult sets Func's FuncFieldCResult.
func (b *FuncBuilder) FuncFieldCResult(funcFieldCResult func(int64, string) (func(), error)) *FuncBuilder {
	b.build.FuncFieldCResult = funcFieldCResult
	return b
}

// FuncFieldMultiple sets Func's FuncFieldMultiple.
func (b *FuncBuilder) FuncFieldMultiple(funcFieldMultiple func(int64, string) (string, error)) *FuncBuilder {
	b.build.FuncFieldMultiple = funcFieldMultiple
	return b
}

// FuncFieldNamed sets Func's FuncFieldNamed.
func (b *FuncBuilder) FuncFieldNamed(funcFieldNamed func(in int64) (out string)) *FuncBuilder {
	b.build.FuncFieldNamed = funcFieldNamed
	return b
}

// FuncFieldNoNames sets Func's FuncFieldNoNames.
func (b *FuncBuilder) FuncFieldNoNames(funcFieldNoNames func(map[string]int, func(success_funcs.Int64Alias)) error) *FuncBuilder {
	b.build.FuncFieldNoNames = funcFieldNoNames
	return b
}

// FuncFieldPtrAlias sets Func's FuncFieldPtrAlias.
func (b *FuncBuilder) FuncFieldPtrAlias(funcFieldPtrAlias func(in *success_funcs.Int64Alias) (out *success_funcs.FuncAlias, err error)) *FuncBuilder {
	b.build.FuncFieldPtrAlias = &funcFieldPtrAlias
	return b
}
