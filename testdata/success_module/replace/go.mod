module github.com/kilianpaquier/go-builder-generator/testdata

go 1.25.1

require github.com/spf13/cobra v1.10.0

replace github.com/spf13/cobra => github.com/spf13/cobra v1.10.1 // test replace feature

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
)
