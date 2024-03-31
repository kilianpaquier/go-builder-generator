package success_same_package

import "context"

//go:generate ../../go-builder-generator generate -f types.go -s SamePackage,unexportedType

//go:generate ../../go-builder-generator generate -f types.go -s unexportedTypePrefix -p Set

type Int64Alias int64

type SamePackage struct {
	Int64Alias

	Ctx       context.Context
	Primitive string

	nonExported string // should be added to builder but not exported
}

type unexportedType struct {
	Int64Alias

	Ctx       context.Context
	Primitive string

	nonExported string // should be added to builder since builder won't be exported anyway
}

type unexportedTypePrefix struct {
	Int64Alias

	Ctx       context.Context
	Primitive string

	nonExported string // should be added to builder since builder won't be exported anyway
}
