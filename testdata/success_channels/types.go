package success_channels

import "context"

//go:generate ../../go-builder-generator generate -f types.go -s Chan -d builders

type Int64Alias int64

type FuncAlias func()

type Chan struct {
	ChanField           chan int64
	ChanFieldAlias      chan Int64Alias
	ChanFieldPtrAlias   *chan *Int64Alias
	ChanFieldSliceAlias chan []FuncAlias
	ChanFieldSliceFunc  chan []func(Int64Alias) (err error)
	ChanFieldWithPkg    chan context.Context

	RChanField         <-chan int64
	RChanFieldAlias    <-chan Int64Alias
	RChanFieldPtrAlias *<-chan *Int64Alias
	RChanFieldWithPkg  <-chan context.Context

	SChanField         chan<- int64
	SChanFieldAlias    chan<- Int64Alias
	SChanFieldPtrAlias *chan<- *Int64Alias
	SChanFieldWithPkg  chan<- context.Context
}
