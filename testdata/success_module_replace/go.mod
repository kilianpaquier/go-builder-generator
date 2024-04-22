module github.com/kilianpaquier/go-builder-generator/root

go 1.22.0

toolchain go1.22.2

require github.com/sirupsen/logrus v1.9.3

replace github.com/sirupsen/logrus v1.9.3 => github.com/sirupsen/logrus v1.9.3 // replace for testing purposes

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
