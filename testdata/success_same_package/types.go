package success_same_package

import "context"

//go:generate ../../go-builder-generator generate -f types.go -s SamePackage

type Int64Alias int64

type SamePackage struct {
	Int64Alias

	Ctx       context.Context
	Primitive string

	nonExported string // should not be added to builder
}
