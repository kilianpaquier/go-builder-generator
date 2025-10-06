package local

import "github.com/kilianpaquier/go-builder-generator/testdata/local/dir/local"

//go:generate ../../../go-builder-generator generate -f types.go -s LocalImport

type LocalImport struct {
	unexported       local.AnotherNaming
	unexportedTagged local.AnotherNaming `builder:"export"`

	Exported local.AnotherNaming
}
