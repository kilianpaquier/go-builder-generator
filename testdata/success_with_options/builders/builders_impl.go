// This file is safe to edit. Once it exists it will not be overwritten.
// Deleting this file will force its generation again next time go-builder-generator is executed.

package builders

/*
Additional functions present here are called during a specific builder's
Build function in the order of struct properties.

For instance, with the below struct, `some_function` will be called first and
then `some_other_function` will be called.

type MyStruct struct {
    FirstField  string `builder:"default_func=some_function"`
    SecondField string `builder:"default_func=some_other_function"`
}
*/

// GetDefaultString will be executed during Build function. It allows defining
// some properties of Options at the end of builder in case those would depend on other properties.
func (b *OptionsBuilder) GetDefaultString() *OptionsBuilder {
	return b
}

// GetDefaultContext will be executed during Build function. It allows defining
// some properties of Options at the end of builder in case those would depend on other properties.
func (b *OptionsBuilder) GetDefaultContext() *OptionsBuilder {
	return b
}
