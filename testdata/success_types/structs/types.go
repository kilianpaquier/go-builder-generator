package testdata

import (
	"context"
	"time"
)

//go:generate ../../../go-builder-generator generate -f types.go -s Struct,StructNoFields -d builders

type Int64Alias int64

type Struct struct {
	TimeA, TimeB time.Time

	ExportedA, unexportedB struct { // ExportedA will be generated but not the unexportedB
		ExportedA, ExportedB Int64Alias
	}

	ExportedC, ExportedD struct {
		ExportedA, ExportedB Int64Alias
	}

	ExportedSubUnexportedA, ExportedSubUnexportedB struct { // not generated since there's one field unexported
		Exported   int
		unexported int64
		ExportedB  string
	}

	AStruct struct { // shouldn't be added since it contains an unexported field
		AField       Int64Alias
		privateField int64
	}

	AnotherStruct struct {
		Nested struct {
			Field string `json:"field,omitempty"`
			Ctx   context.Context
		}
		NotNested int64 `json:"not_nested"`
		Ctx       context.Context
		Alias     Int64Alias
	}
}

type StructNoFields struct {
	NoFields struct{}
}
