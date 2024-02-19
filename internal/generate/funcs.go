package generate

import (
	"strings"
	"text/template"
)

// funcMap returns the template funcMap for go template generation.
func funcMap() template.FuncMap {
	return template.FuncMap{
		"hasPrefix": strings.HasPrefix,
	}
}
