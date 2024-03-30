<!-- This file is safe to edit. Once it exists it will not be overwritten. -->

# go-builder-generator <!-- omit in toc -->

<p align="center">
  <img alt="GitHub Actions" src="https://img.shields.io/github/actions/workflow/status/kilianpaquier/go-builder-generator/integration.yml?branch=main&style=for-the-badge">
  <img alt="GitHub Release" src="https://img.shields.io/github/v/release/kilianpaquier/go-builder-generator?include_prereleases&sort=semver&style=for-the-badge">
  <img alt="GitHub Issues" src="https://img.shields.io/github/issues-raw/kilianpaquier/go-builder-generator?style=for-the-badge">
  <img alt="GitHub License" src="https://img.shields.io/github/license/kilianpaquier/go-builder-generator?style=for-the-badge">
  <img alt="Coverage" src="https://img.shields.io/codecov/c/github/kilianpaquier/go-builder-generator/main?style=for-the-badge">
  <img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/kilianpaquier/go-builder-generator/main?style=for-the-badge&label=Go+Version">
  <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/kilianpaquier/go-builder-generator?style=for-the-badge">
</p>

---

- [How to use ?](#how-to-use-)
- [Commands](#commands)
  - [Generate](#generate)
    - [Tags](#tags)

## How to use ?

```sh
go install github.com/kilianpaquier/go-builder-generator/cmd/go-builder-generator@latest
```

## Commands

```
go-builder-generator stands here to easily generate builders for your golang struct types.

Usage:
  go-builder-generator [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Generate builders for structs arguments present in file argument.
  help        Help about any command
  version     Shows current go-builder-generator version

Flags:
  -h, --help               help for go-builder-generator
  -l, --log-level string   set logging level

Use "go-builder-generator [command] --help" for more information about a command.
```

### Generate

```
Generate builders for structs arguments present in file argument.

Usage:
  go-builder-generator generate [flags]

Flags:
  -d, --dest string            destination directory in which generate builders (default ".")
  -f, --file string            input files containing struct types on which generate builders
  -h, --help                   help for generate
      --no-notice              remove top notice ('generated by ...') in generated files
  -p, --prefix string          specific prefix to apply on setter functions
  -s, --structs strings        golang struct names on which generate builders (they must be contained in given input files)
      --validate-func string   validate function name to be executed in Build, must have the signature 'func () error' and associated to built struct

Global Flags:
  -l, --log-level string   set logging level
```

#### Tags

It's possible to tune builders generation with the struct tag `builder`. The available options are:

- `ignore`: when provided, the field will be ignored in builder generation.
- `append`: when provided on a slice (only), the generated builder will append instead of reaffecting.
- `pointer`: when provided on a pointer, the generated builder will keep it's pointer on the input parameter.
- `default_func`: when provided, an additional function will be generated in another file suffixed with `_impl.go` to allow manual affectation of field (or even other fields).

Example:

```go
package pkg

//go:generate go-builder-generator -f types.go -s StructName -d builders

type StructName struct {
	Pointer               *int64   `builder:"pointer"` // generated builder will be 'SetPointer(pointer *int64)'
	NoPointer             *int64   // generated builder will be 'SetNoPointer(noPointer int64)'
	ASlice                []string `builder:"append"` // generated builder will be 'SetASlice(aSlice ...string)', additionally the affectation will be `b.ASlice = append(b.ASlice, aSlice...)`
	NoAppend              []string // generated builder will be 'SetASlice(noAppend []string)', additionally the affectation will be `b.NoAppend = noAppend`
	Ignore                int64    `builder:"ignore"`                            // no builder will be generated on this field
	DefaultFunc           int64    `builder:"default_func=SomeFuncName"`         // an additional function named 'SomeFuncName' will be generated in target package file '_impl.go' and associated to builder struct
	IgnoreWithDefaultFunc int64    `builder:"ignore,default_func=SomeOtherFunc"` // no builder will be generated and the additional function will be generated
}
```

**Note:** `append` option and `pointer` option are exclusive with a priority for `append` if both provided. Also if `append` is provided on a field not being a slice, it will just be ignored.

For more examples, you can check in `examples` package at project root !