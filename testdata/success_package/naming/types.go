package naming

import (
	"github.com/kilianpaquier/go-builder-generator/testdata/success_package/naming/another_folder/naming"
)

//go:generate ../../../go-builder-generator generate --no-cmd -f types.go -s Naming --return-copy
type Naming struct {
	field2 naming.AnotherNaming `builder:"func_name=Field2,pointer,export"`
}
