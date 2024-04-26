package success_same_package

import "context"

//go:generate ../../go-builder-generator generate -f types.go -s unexportedTypePrefix -p set

type Int64Alias int64

type unexportedTypePrefix struct {
	Int64Alias

	Ctx       context.Context
	Primitive string

	nonExported          string // should be added to builder since builder won't be exported anyway
	nonExportedExportOpt string `builder:"export"` // should be added to builder since builder won't be exported anyway
}
