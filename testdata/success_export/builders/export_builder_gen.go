// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"github.com/kilianpaquier/go-builder-generator/testdata/success_export"
)

// ExportBuilder is an alias of Export to build Export with builder-pattern.
type ExportBuilder success_export.Export

// NewExportBuilder creates a new ExportBuilder.
func NewExportBuilder() *ExportBuilder {
	return &ExportBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *ExportBuilder) Copy() *ExportBuilder {
	c := *b
	return &c
}

// Build returns built Export.
func (b *ExportBuilder) Build() *success_export.Export {
	c := (success_export.Export)(*b)
	return &c
}

// SetInt64Alias sets Export's Int64Alias.
func (b *ExportBuilder) SetInt64Alias(int64Alias success_export.Int64Alias) *ExportBuilder {
	b.Int64Alias = int64Alias
	return b
}
