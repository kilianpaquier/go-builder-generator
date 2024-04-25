package success_funcs

//go:generate ../../go-builder-generator generate -f types.go -s Func -d builders

type Int64Alias int64

type FuncAlias func()

type Func struct {
	FuncField         func(int64) string
	FuncFieldCResult  func(int64, string) (func(), error)
	FuncFieldMultiple func(int64, string) (string, error)
	FuncFieldNamed    func(in int64) (out string)

	FuncFieldAlias         func(Int64Alias) string
	FuncFieldAliasMultiple func(Int64Alias, FuncAlias) (string, error)
	FuncFieldNoNames       func(map[string]int, func(Int64Alias)) error
	FuncFieldAliasNamed    func(in Int64Alias) (out FuncAlias)
	FuncFieldChan          func(c chan<- Int64Alias) error

	FuncFieldPtrAlias *func(in *Int64Alias) (out *FuncAlias, err error)
}
