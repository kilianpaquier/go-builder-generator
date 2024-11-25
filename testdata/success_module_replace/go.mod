module github.com/kilianpaquier/go-builder-generator/root

go 1.22.0

toolchain go1.22.2

require github.com/stretchr/testify v1.10.0

replace github.com/stretchr/testify v1.10.0 => github.com/stretchr/testify v1.10.0 // replace for testing purposes

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
