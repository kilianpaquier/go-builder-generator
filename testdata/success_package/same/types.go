package testdata

import "context"

//go:generate ../../../go-builder-generator generate -f types.go -s SamePackage,unexportedType

type Int64Alias int64

type SamePackage struct {
	Int64Alias

	Ctx       context.Context
	Primitive string

	nonExported          string // should be added to builder but not exported
	nonExportedExportOpt string `builder:"export"` // should be added to builder and exported since the option says so
}

type unexportedType struct {
	Int64Alias

	Ctx       context.Context
	Primitive string

	nonExported          string // should be added to builder since builder won't be exported anyway
	nonExportedExportOpt string `builder:"export"` // should be added to builder since builder won't be exported anyway
}
