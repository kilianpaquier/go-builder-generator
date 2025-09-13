package testdata

//go:generate ../../../go-builder-generator generate -f types.go -s Export -d builders

type Int64Alias int64

type unexportedAlias int64

type FuncAlias func()

type Export struct {
	Int64Alias // should be added to builder

	Exported            unexportedAlias // should not be added to builder
	unexported          FuncAlias       `builder:"export"` // should not be added to builder
	unexportedPrimitive int64           // should not be added to builder

	unexportWithFuncName string `builder:"export,func_name=ShouldNotBeAddedToBuilder"` // should not be added to builder
}
