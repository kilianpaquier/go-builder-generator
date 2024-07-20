// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

//go:generate go run github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@latest generate -d . -f ../types.go -s Chan

package builders

import (
	"context"

	"github.com/kilianpaquier/go-builder-generator/testdata/success_channels"
)

// ChanBuilder represents Chan's builder.
type ChanBuilder struct {
	build success_channels.Chan
}

// NewChanBuilder creates a new ChanBuilder.
func NewChanBuilder() *ChanBuilder {
	return &ChanBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *ChanBuilder) Copy() *ChanBuilder {
	return &ChanBuilder{b.build}
}

// Build returns built Chan.
func (b *ChanBuilder) Build() *success_channels.Chan {
	result := b.build
	return &result
}

// ChanField sets Chan's ChanField.
func (b *ChanBuilder) ChanField(chanField chan int64) *ChanBuilder {
	b.build.ChanField = chanField
	return b
}

// ChanFieldAlias sets Chan's ChanFieldAlias.
func (b *ChanBuilder) ChanFieldAlias(chanFieldAlias chan success_channels.Int64Alias) *ChanBuilder {
	b.build.ChanFieldAlias = chanFieldAlias
	return b
}

// ChanFieldPtrAlias sets Chan's ChanFieldPtrAlias.
func (b *ChanBuilder) ChanFieldPtrAlias(chanFieldPtrAlias chan *success_channels.Int64Alias) *ChanBuilder {
	b.build.ChanFieldPtrAlias = &chanFieldPtrAlias
	return b
}

// ChanFieldSliceAlias sets Chan's ChanFieldSliceAlias.
func (b *ChanBuilder) ChanFieldSliceAlias(chanFieldSliceAlias chan []success_channels.FuncAlias) *ChanBuilder {
	b.build.ChanFieldSliceAlias = chanFieldSliceAlias
	return b
}

// ChanFieldSliceFunc sets Chan's ChanFieldSliceFunc.
func (b *ChanBuilder) ChanFieldSliceFunc(chanFieldSliceFunc chan []func(success_channels.Int64Alias) (err error)) *ChanBuilder {
	b.build.ChanFieldSliceFunc = chanFieldSliceFunc
	return b
}

// ChanFieldWithPkg sets Chan's ChanFieldWithPkg.
func (b *ChanBuilder) ChanFieldWithPkg(chanFieldWithPkg chan context.Context) *ChanBuilder {
	b.build.ChanFieldWithPkg = chanFieldWithPkg
	return b
}

// RChanField sets Chan's RChanField.
func (b *ChanBuilder) RChanField(rchanField <-chan int64) *ChanBuilder {
	b.build.RChanField = rchanField
	return b
}

// RChanFieldAlias sets Chan's RChanFieldAlias.
func (b *ChanBuilder) RChanFieldAlias(rchanFieldAlias <-chan success_channels.Int64Alias) *ChanBuilder {
	b.build.RChanFieldAlias = rchanFieldAlias
	return b
}

// RChanFieldPtrAlias sets Chan's RChanFieldPtrAlias.
func (b *ChanBuilder) RChanFieldPtrAlias(rchanFieldPtrAlias <-chan *success_channels.Int64Alias) *ChanBuilder {
	b.build.RChanFieldPtrAlias = &rchanFieldPtrAlias
	return b
}

// RChanFieldWithPkg sets Chan's RChanFieldWithPkg.
func (b *ChanBuilder) RChanFieldWithPkg(rchanFieldWithPkg <-chan context.Context) *ChanBuilder {
	b.build.RChanFieldWithPkg = rchanFieldWithPkg
	return b
}

// SChanField sets Chan's SChanField.
func (b *ChanBuilder) SChanField(schanField chan<- int64) *ChanBuilder {
	b.build.SChanField = schanField
	return b
}

// SChanFieldAlias sets Chan's SChanFieldAlias.
func (b *ChanBuilder) SChanFieldAlias(schanFieldAlias chan<- success_channels.Int64Alias) *ChanBuilder {
	b.build.SChanFieldAlias = schanFieldAlias
	return b
}

// SChanFieldPtrAlias sets Chan's SChanFieldPtrAlias.
func (b *ChanBuilder) SChanFieldPtrAlias(schanFieldPtrAlias chan<- *success_channels.Int64Alias) *ChanBuilder {
	b.build.SChanFieldPtrAlias = &schanFieldPtrAlias
	return b
}

// SChanFieldWithPkg sets Chan's SChanFieldWithPkg.
func (b *ChanBuilder) SChanFieldWithPkg(schanFieldWithPkg chan<- context.Context) *ChanBuilder {
	b.build.SChanFieldWithPkg = schanFieldWithPkg
	return b
}
