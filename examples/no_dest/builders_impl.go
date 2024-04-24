// This file is safe to edit. Once it exists it will not be overwritten.
// Deleting this file will force its generation again next time go-builder-generator is executed.

package no_dest

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
