package success_struct

import "context"

//go:generate go tool go-builder-generator generate -f types.go -s Struct,StructNoFields -d builders

type Int64Alias int64

type Struct struct {
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
