package success_with_options

import "context"

//go:generate ../../go-builder-generator generate -f types.go -s Options -d builders --use-validator

type Int64Alias int64

type Options struct {
	DefaultField              int64           `builder:"default=10"`
	DefaultFieldFunc          string          `builder:"default_func=GetDefaultString"`
	DefaultIgnoreAndFieldFunc context.Context `builder:"ignore,default_func=GetDefaultContext"`
	IgnoreField               string          `builder:"ignore"`

	PtrField             *string `builder:"pointer" validate:"required"`
	PtrFieldValidatedToo *string `validate:"required"`

	SimpleFieldAppend Int64Alias `builder:"append"` // option shouldn't be used

	SliceFieldAliasAppend []Int64Alias `builder:"append"`
	SliceFieldAppend      []int64      `builder:"append"`
	SliceFieldAppendPtr   []*int64     `builder:"append"`
	SliceFieldPtrAppend   *[]int64     `builder:"append,pointer"` // should be the same as below
	SliceFieldNoPtrAppend *[]int64     `builder:"append"`         // should be the same as above
}
