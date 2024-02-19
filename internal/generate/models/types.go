package models

// PrimaryTypes represents the slice of string representations of all golang primary types.
//
// All types are considered primary type from the moment they doesn't belong to a specific go package.
func PrimaryTypes() []string {
	return []string{
		"uintptr", "uint", "uint8", "uint16", "uint32", "uint64",
		"int", "int8", "int16", "int32", "int64",
		"float32", "float64",
		"complex64", "complex128",
		"byte", "rune", "string",
		"bool",
		"error", "any",
	}
}
