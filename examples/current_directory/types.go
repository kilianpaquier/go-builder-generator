package examples

//go:generate ../../go-builder-generator generate -f types.go -s SomeStruct

type SomeAlias int64

type SomeStruct struct {
	SomeSlice []string
	SomeChan  chan<- SomeAlias
}
