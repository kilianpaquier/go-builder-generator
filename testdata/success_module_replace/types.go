package root

import (
	"github.com/stretchr/testify/mock"
)

//go:generate ../../go-builder-generator generate -f module::github.com/stretchr/testify/mock/mock.go -s Mock -d builders --no-cmd

// Mock is just an alias of testify Mock.
type Mock mock.Mock
