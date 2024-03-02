package success_pkg

//go:generate ../../go-builder-generator generate -f ~/.local/go/src/os/error.go -s SyscallError -d builders

//go:generate ../../go-builder-generator generate -f ~/.local/go/src/database/sql/sql.go -s NamedArg -d builders

//go:generate ../../go-builder-generator generate -f ~/go/pkg/mod/github.com/go-playground/validator/v10@v10.19.0/errors.go -s InvalidValidationError -d builders
