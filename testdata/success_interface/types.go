package success_interface

import "context"

// execute this command from your terminal
// ./go-builder-generator generate -f testdata/success_interface/types.go -s Interface,InterfaceNoFields -d testdata/success_interface/builders --package-name builders

type Int64Alias int64

type Interface struct {
	AnInterface interface { // shouldn't be added since there's a private function
		AFunction() string
		AnotherFunc() context.Context
		privateFunc() Int64Alias
	}

	AnotherInterface interface {
		SomeFunc() chan<- Int64Alias
		SomeOtherFunc(ctx context.Context, c <-chan int64) error
		AFunc(i Int64Alias) context.Context
	}
}

type InterfaceNoFields struct {
	NoFields interface{}
}
