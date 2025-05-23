package success_generic

//go:generate go tool go-builder-generator generate -f types.go -s Struct,SimpleGeneric,AliasGeneric,GenericAnonymousStruct,ComplexGeneric,ComplexSliceGeneric -d builders

type SimpleGeneric[T any] struct {
	Value T
}

type GenericValue string

type Struct struct {
	GenericField  SimpleGeneric[int64]
	GenericFields []*SimpleGeneric[GenericValue]
}

type AliasGeneric[t any, GV GenericValue] struct {
	ValueT t
	ValueX GV
}

type ComplexGeneric[T Struct, Y comparable] struct {
	ValueT             map[Y]AliasGeneric[T, GenericValue]
	ValueY             [10]*[]SimpleGeneric[Y]
	FuncT              func(T, SimpleGeneric[T]) (T, error)
	AnonymousInterface interface {
		HeyFunc(T) error
		HeyFuncMulti(AliasGeneric[T, GenericValue]) AliasGeneric[T, GenericValue]
	}
}

type ComplexSliceGeneric[S ~[]E, E Struct] struct {
	ValueT func(S) E
}

type GenericAnonymousStruct[T struct {
	Property string `builder:"pointer"` // tag won't be taken into account
}] struct {
	Prop T
}
