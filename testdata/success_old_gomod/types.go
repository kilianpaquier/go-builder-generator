package root

// execute this command from your terminal
// this case ensures that go run ... is added to generated files and not go tool (since go.mod is older than go1.24)

// ../../go-builder-generator generate -f types.go -s RootType -d builders

type RootType struct {
	Field int64
}
