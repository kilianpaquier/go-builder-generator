package testdata

//go:generate ../../go-builder-generator generate -f types.go -s Options,Empty,GenericOptions -d builders --package-name my_package --validate-func Validate --return-copy -p Set --no-notice
