package testdata

import (
	net_http "net/http"
	"net/url"
)

//go:generate ../../../go-builder-generator generate -f types.go -s STD -d builders

type STD struct {
	ACRONYMOUS        string
	AnotherACRONYMOUS string
	AnURL             *url.URL
	ID                int64
	SomeClientHTTP    *net_http.Client
	SomeID            int64
	URL               *url.URL
}
