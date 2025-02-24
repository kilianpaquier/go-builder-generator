<!-- This file is safe to edit. Once it exists it will not be overwritten. -->

# go-builder-generator <!-- omit in toc -->

<p align="center">
  <img alt="GitHub Release" src="https://img.shields.io/github/v/release/kilianpaquier/go-builder-generator?include_prereleases&sort=semver&style=for-the-badge">
  <img alt="GitHub Issues" src="https://img.shields.io/github/issues-raw/kilianpaquier/go-builder-generator?style=for-the-badge">
  <img alt="GitHub License" src="https://img.shields.io/github/license/kilianpaquier/go-builder-generator?style=for-the-badge">
  <img alt="Coverage" src="https://img.shields.io/codecov/c/github/kilianpaquier/go-builder-generator/main?style=for-the-badge">
  <img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/kilianpaquier/go-builder-generator/main?style=for-the-badge&label=Go+Version">
  <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/kilianpaquier/go-builder-generator?style=for-the-badge">
</p>

---

- [How to use ?](#how-to-use-)
  - [Go](#go)
  - [Linux](#linux)
- [Commands](#commands)
  - [Generate](#generate)
  - [Generation cases](#generation-cases)

## How to use ?

### Go

```sh
go install github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@latest
```

### Linux

```sh
OS="linux" # change it depending on your case
ARCH="amd64" # change it depending on your case
INSTALL_DIR="$HOME/.local/bin" # change it depending on your case

new_version=$(curl -fsSL "https://api.github.com/repos/kilianpaquier/go-builder-generator/releases/latest" | jq -r '.tag_name')
url="https://github.com/kilianpaquier/go-builder-generator/releases/download/$new_version/go-builder-generator_${OS}_${ARCH}.tar.gz"
curl -fsSL "$url" | (mkdir -p "/tmp/go-builder-generator/$new_version" && cd "/tmp/go-builder-generator/$new_version" && tar -xz)
cp "/tmp/go-builder-generator/$new_version/go-builder-generator" "$INSTALL_DIR/go-builder-generator"
```

## Commands

```
Usage:
  go-builder-generator [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Generate builders for structs arguments present in file argument.
  help        Help about any command
  upgrade     Upgrade or install go-builder-generator
  version     Show current go-builder-generator version

Flags:
  -h, --help                help for go-builder-generator
      --log-format string   set logging format (either "text" or "json") (default "text")
      --log-level string    set logging level (default "info")

Use "go-builder-generator [command] --help" for more information about a command.
```

### Generate

```
Generate builders for structs arguments present in file argument.

Usage:
  go-builder-generator generate [flags]

Flags:
  -d, --dest string            destination directory for the generated files (default ".")
  -f, --file string            input file containing golang struct(s) to generate builders on
  -h, --help                   help for generate
      --no-cmd                 removes top comment 'go:generate go run ...' from generated files
      --no-notice              removes top notice 'Code generated by ...' from generated files
      --package-name string    defines a specific package name instead of '--dest', '-d' directory name. Only available when generating files in another directory
  -p, --prefix string          specific prefix to apply on setter functions
      --return-copy            returns a copy of the builder each time a setter function is called
  -s, --structs strings        struct names to generate builders on (they must be contained in given input file)
      --validate-func string   validate function name to be executed in Build, must have the signature 'func () error' and associated to built struct

Global Flags:
      --log-format string   set logging format (either "text" or "json") (default "text")
      --log-level string    set logging level (default "info")
```

#### Tags

It's possible to tune builders generation with the struct tag `builder`. The available options are:

- `append`: when provided on a slice (only), the generated builder will append instead of reaffecting.
- `default_func`: when provided, an additional function will be generated in another file suffixed with `_impl.go` to allow manual affectation of field (or even other fields).
- `func_name`: when provided, the name of the function generated in the builder is set to the provided name, ignoring prefix or field name.
- `ignore`: when provided, the field will be ignored in builder generation.
- `pointer`: when provided on a pointer, the generated builder will keep it's pointer on the input parameter.
- `export`: when provided, the builder function associated to the given field (unexported field for that matter) will be exported.
  - It's only applicable when generation is done in the same package as the generated struct and the struct is exported.
  - In the case of an unexported field for an exported struct with the generation in another package, the field will always be ignored during generation.
  - In the case of an unexported field for an unexported struct with the generation in the same package, the builder function will always be exported (because the builder and the new function are unexported).

**Note:** `append` option and `pointer` option are exclusive with a priority for `append` if both provided. Also if `append` is provided on a field not being a slice, it will just be ignored.

Example:

```go
package pkg

//go:generate go run github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@latest generate -f types.go -s StructName -d builders

type StructName struct {
	Pointer               *int64   `builder:"pointer"` // generated builder will be 'SetPointer(pointer *int64)'
	NoPointer             *int64   // generated builder will be 'SetNoPointer(noPointer int64)'
	ASlice                []string `builder:"append"` // generated builder will be 'SetASlice(aSlice ...string)', additionally the affectation will be `b.ASlice = append(b.ASlice, aSlice...)`
	NoAppend              []string // generated builder will be 'SetASlice(noAppend []string)', additionally the affectation will be `b.NoAppend = noAppend`
	Ignore                int64    `builder:"ignore"`                            // no builder will be generated on this field
	DefaultFunc           int64    `builder:"default_func=SomeFuncName"`         // an additional function named 'SomeFuncName' will be generated in target package file '_impl.go' and associated to builder struct
	IgnoreWithDefaultFunc int64    `builder:"ignore,default_func=SomeOtherFunc"` // no builder will be generated and the additional function will be generated
	CustomFuncName        string   `builder:"func_name=FuncNameOverride"`        // generated builder will be 'FuncNameOverride(customFuncName string)'
}

//go:generate go run github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@latest generate -f types.go -s AnotherStructName

type AnotherStructName struct {
	unexportedField string `builder:"export"` // generated builder will be 'UnexportedField(unexportedField string)'
}
```

### Generation cases

#### Simple struct

```go
//go:generate go run github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@latest generate -f types.go -s StructName -d builders

type StructName struct {...}
```

#### Generic struct

```go
//go:generate go run github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@latest generate -f types.go -s GenericStruct,GenericStructUnary,GenericFieldStruct -d builders

type GenericStruct[T any, ...] struct {...}

type GenericStructUnary[T ~E, E any, ...] struct {...}

type GenericFieldStruct struct {
  Field GenericStruct[string, ...]
}
```

#### Imported module struct

In case it's needed to generate a builder on a struct not being one of the current module, it's possible to provide the `module::` prefix to tell `go-builder-generator` to generate the struct from an **imported** module.
The provided module must be imported in the current module `go.mod`.

This case works with both simple structs and generic structs. Under the hood, `go-builder-generator` will retrieve the appropriate version from the current module `go.mod` (it works with `replace` too)
and generate with those specific rules:
- If `replace` is provided with a custom path, then it will retrieve the file from that path
- If `GOPATH` environment variable exists, it will retrieve the file from `${GOPATH}/pkg/mod/module_name/...`
- If `GOPATH` environment variable doesn't exist, it will retrieve the file from `${HOME}/go/pkg/mod/module_name/...`

```go
//go:generate go run github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@latest generate -f module::github.com/kilianpaquier/go-builder-generator/path/to/file.go -s ExternalStructName -d builders
```

You may see the [examples](./examples/) folder for real generation cases.
