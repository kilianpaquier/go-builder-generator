package local_and_dir

import "github.com/kilianpaquier/go-builder-generator/testdata/local_same_dir/dir/local_same_dir"

//go:generate ../../../go-builder-generator generate -f types.go -s LocalImport -d local_same_dir

type LocalImport struct {
	Field2 local_same_dir.AnotherNaming
}
