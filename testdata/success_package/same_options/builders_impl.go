// This file is safe to edit. Once it exists it will not be overwritten.
// Deleting this file will force its generation again next time go-builder-generator is executed.

package testdata

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

// PrimitiveDef will be executed during Build function. It allows defining
// some fields of unexportedTypeOptions at the end of builder in case those would depend on other fields.
func (b *unexportedTypeOptionsBuilder) PrimitiveDef() *unexportedTypeOptionsBuilder {
	return b
}
