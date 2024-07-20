package success_spo

import "context"

//go:generate ../../go-builder-generator generate -f types.go -s unexportedTypeOptions -p set --package-name invalid

type Int64Alias int64

type unexportedTypeOptions struct {
	Int64Alias

	Ctx       context.Context
	Primitive string `builder:"default_func=PrimitiveDef"`

	nonExported          string // should be added to builder since builder won't be exported anyway
	nonExportedExportOpt string `builder:"export"` // should be added to builder since builder won't be exported anyway
}
