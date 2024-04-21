package success_return_copy

//go:generate ../../go-builder-generator generate -f types.go -s ReturnCopy -d builders

type ReturnCopy struct {
	Field string
	Slice []string `builder:"append"`
}
