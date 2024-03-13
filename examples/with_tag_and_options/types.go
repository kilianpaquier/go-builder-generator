package with_tag_and_options

//go:generate ../../go-builder-generator generate -f types.go -s SomeStruct -d builders --validate-func Validate

type SomeAlias int64

type SomeStruct struct {
	SomeChan  chan<- SomeAlias `builder:"ignore,default_func=SetTheChan"`
	SomeSlice []string         `builder:"append"`
}

func (s *SomeStruct) Validate() error {
	return nil
}
