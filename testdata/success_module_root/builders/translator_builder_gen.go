// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import "github.com/huandu/xstrings"

// TranslatorBuilder represents Translator's builder.
type TranslatorBuilder struct {
	build xstrings.Translator
}

// NewTranslatorBuilder creates a new TranslatorBuilder.
func NewTranslatorBuilder() *TranslatorBuilder {
	return &TranslatorBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *TranslatorBuilder) Copy() *TranslatorBuilder {
	return &TranslatorBuilder{b.build}
}

// Build returns built Translator.
func (b *TranslatorBuilder) Build() *xstrings.Translator {
	result := b.build
	return &result
}
