package success_pkg

//go:generate ../../go-builder-generator generate -f ~/.local/go/src/os/error.go -s SyscallError -d builders

//go:generate ../../go-builder-generator generate -f ~/.local/go/src/database/sql/sql.go -s NamedArg -d builders
