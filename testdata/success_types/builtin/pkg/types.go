package pkg

import "unsafe"

//go:generate ../../../../go-builder-generator generate -f types.go -s Builtin -d builders

type Builtin struct {
	Any           any
	Bool          bool
	Byte          byte
	Bytes         []byte
	Complex128    complex128
	Complex64     complex64
	Error         error
	Float32       float32
	Float64       float64
	Int           int
	Int32         int32
	Int64         int64
	Int8          int8
	Rune          rune
	Runes         []rune
	String        string
	Uint          uint
	Uint16        uint16
	Uint32        uint32
	Uint64        uint64
	Uint8         uint8
	Uintptr       uintptr
	UnsafePointer unsafe.Pointer
}
