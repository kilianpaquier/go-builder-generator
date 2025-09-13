module github.com/kilianpaquier/go-builder-generator/testdata

go 1.25.1

require github.com/stretchr/testify v1.10.0

replace github.com/stretchr/testify => github.com/stretchr/testify v1.11.1 // test replace feature

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
