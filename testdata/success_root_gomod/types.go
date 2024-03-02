package root

//go:generate ../../go-builder-generator generate -f types.go -s RootType -d builders

type RootType struct {
	Field int64
}
