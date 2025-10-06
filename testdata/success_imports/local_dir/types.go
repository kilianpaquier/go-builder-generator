package local_dir

import "github.com/kilianpaquier/go-builder-generator/testdata/local_dir/dir/local_dir"

//go:generate ../../../go-builder-generator generate -f types.go -s LocalImport -d builders

type LocalImport struct {
	Field2 local_dir.AnotherNaming
}
