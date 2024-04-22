package root

import "github.com/sirupsen/logrus/hooks/test"

//go:generate ../../go-builder-generator generate -f module::github.com/sirupsen/logrus/hooks/test/test.go -s Hook -d builders

type Hook test.Hook
