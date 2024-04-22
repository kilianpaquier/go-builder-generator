package success_generic_field

//go:generate ../../go-builder-generator generate -f types.go -s Struct -d builders

type GenericType[T any] struct {
	Value T
}

type GenericValue string

type Struct struct {
	GenericField  GenericType[int64]
	GenericFields []*GenericType[GenericValue]
}
