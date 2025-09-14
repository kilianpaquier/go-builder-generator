// This file is safe to edit. Once it exists it will not be overwritten.
// Deleting this file will force its generation again next time go-builder-generator is executed.

package my_package

/*
Additional functions present here are called during a specific builder's
Build function in the order of struct fields.

For instance, with the below struct, `some_function` will be called first and
then `some_other_function` will be called.

type MyStruct struct {
    FirstField  string `builder:"default_func=some_function"`
    SecondField string `builder:"default_func=some_other_function"`
}
*/

// GetDefaultString will be executed during Build function. It allows defining
// some fields of Options at the end of builder in case those would depend on other fields.
func (b *OptionsBuilder) GetDefaultString() *OptionsBuilder {
	return b
}

// GetDefaultContext will be executed during Build function. It allows defining
// some fields of Options at the end of builder in case those would depend on other fields.
func (b *OptionsBuilder) GetDefaultContext() *OptionsBuilder {
	return b
}

// SharedFunc will be executed during Build function. It allows defining
// some fields of Options at the end of builder in case those would depend on other fields.
func (b *OptionsBuilder) SharedFunc() *OptionsBuilder {
	return b
}

// SetDefaultForceFuncName will be executed during Build function. It allows defining
// some fields of Options at the end of builder in case those would depend on other fields.
func (b *OptionsBuilder) SetDefaultForceFuncName() *OptionsBuilder {
	return b
}

// GetDefaultString will be executed during Build function. It allows defining
// some fields of GenericOptions at the end of builder in case those would depend on other fields.
func (b *GenericOptionsBuilder[T]) GetDefaultString() *GenericOptionsBuilder[T] {
	return b
}
