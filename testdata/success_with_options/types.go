package success_with_options

import "context"

//go:generate ../../go-builder-generator generate -f types.go -s Options,Empty,GenericOptions -d builders --validate-func Validate --return-copy

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

	ForceFuncName                string `builder:"func_name=FooBarForced"`
	ForceFuncNameWithDefaultFunc string `builder:"func_name=FooBarForceWithDefault,default_func=SetDefaultForceFuncName"`
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
