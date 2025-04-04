package with_tag

//go:generate go tool go-builder-generator generate -f types.go -s SomeStruct -d builders

type SomeAlias int64

type SomeStruct struct {
	SomeChan  chan<- SomeAlias `builder:"ignore,default_func=SetTheChan"`
	SomeSlice []string         `builder:"append"`
}
