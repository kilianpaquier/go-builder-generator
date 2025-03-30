package with_options

//go:generate go tool go-builder-generator generate -f types.go -s Options -d builders --validate-func Validate --prefix set --no-cmd

type Options struct {
	SomeField int64
	SomeSlice []string
}

func (o *Options) Validate() error {
	return nil
}
