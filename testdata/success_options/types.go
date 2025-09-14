package testdata

import "context"

type Int64Alias int64

type Empty struct{}

type Options struct {
	DefaultFieldFunc          string          `builder:"default_func=GetDefaultString"`
	DefaultIgnoreAndFieldFunc context.Context `builder:"ignore,default_func=GetDefaultContext"`
	IgnoreField               string          `builder:"ignore"`

	PtrField             *string `builder:"pointer" validate:"required"`
	PtrFieldValidatedToo *string

	SimpleFieldAppend Int64Alias `builder:"append"` // option shouldn't be used

	SliceFieldAliasAppend []Int64Alias `builder:"append"`
	SliceFieldAppend      []int64      `builder:"append"`
	SliceFieldAppendPtr   []*int64     `builder:"append"`
	SliceFieldPtrAppend   *[]int64     `builder:"append,pointer"` // should be the same as below
	SliceFieldNoPtrAppend *[]int64     `builder:"append"`         // should be the same as above

	SharedFuncA, SharedFuncB string `builder:"default_func=SharedFunc"`

	ForceFuncName                string `builder:"func_name=FooBarForced"`
	ForceFuncNameWithDefaultFunc string `builder:"func_name=FooBarForceWithDefault,default_func=SetDefaultForceFuncName"`

	privateField           string // should not be generated at all
	privateFieldWithExport string `builder:"export"` // should not be generated at all even with export option
}

func (o *Options) Validate() error {
	return nil
}

type GenericOptions[T any] struct {
	DefaultFieldFunc T `builder:"default_func=GetDefaultString"`
}

func (o *GenericOptions[T]) Validate() error {
	return nil
}
