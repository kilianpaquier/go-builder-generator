package success_maps

import "context"

//go:generate go tool go-builder-generator generate -f types.go -s Map -d builders

type Int64Alias int64

type FuncAlias func()

type Map struct {
	MapField         map[int64]string
	MapFieldAlias    map[Int64Alias]FuncAlias
	MapFieldFunc     map[int64]func(in int64) error
	MapFieldPtrAlias *map[*Int64Alias]*FuncAlias
	MapFieldWithPkg  map[int64]context.Context
}
