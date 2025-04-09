// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

// Code generated from success_naming/types.go.

package builders

import (
	net_http "net/http"
	"net/url"

	"github.com/kilianpaquier/go-builder-generator/testdata/success_naming"
)

// NamingBuilder represents Naming's builder.
type NamingBuilder struct {
	build success_naming.Naming
}

// NewNamingBuilder creates a new NamingBuilder.
func NewNamingBuilder() *NamingBuilder {
	return &NamingBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *NamingBuilder) Copy() *NamingBuilder {
	return &NamingBuilder{b.build}
}

// Build returns built Naming.
func (b *NamingBuilder) Build() *success_naming.Naming {
	result := b.build
	return &result
}

// ACRONYMOUS sets Naming's ACRONYMOUS.
func (b *NamingBuilder) ACRONYMOUS(acronymous string) *NamingBuilder {
	b.build.ACRONYMOUS = acronymous
	return b
}

// AnotherACRONYMOUS sets Naming's AnotherACRONYMOUS.
func (b *NamingBuilder) AnotherACRONYMOUS(anotherACRONYMOUS string) *NamingBuilder {
	b.build.AnotherACRONYMOUS = anotherACRONYMOUS
	return b
}

// AnURL sets Naming's AnURL.
func (b *NamingBuilder) AnURL(anURL url.URL) *NamingBuilder {
	b.build.AnURL = &anURL
	return b
}

// ID sets Naming's ID.
func (b *NamingBuilder) ID(id int64) *NamingBuilder {
	b.build.ID = id
	return b
}

// SomeClientHTTP sets Naming's SomeClientHTTP.
func (b *NamingBuilder) SomeClientHTTP(someClientHTTP net_http.Client) *NamingBuilder {
	b.build.SomeClientHTTP = &someClientHTTP
	return b
}

// SomeID sets Naming's SomeID.
func (b *NamingBuilder) SomeID(someID int64) *NamingBuilder {
	b.build.SomeID = someID
	return b
}

// URL sets Naming's URL.
func (b *NamingBuilder) URL(url url.URL) *NamingBuilder {
	b.build.URL = &url
	return b
}
