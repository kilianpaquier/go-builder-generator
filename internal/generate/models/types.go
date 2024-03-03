package models

// PrimaryTypes represents the slice of string representations of all golang primary types.
//
// All types are considered primary type from the moment they doesn't belong to a specific go package.
func PrimaryTypes() []string {
	return []string{
		"any", "bool", "byte",
		"comparable", "complex64", "complex128",
		"error",
		"float32", "float64",
		"int", "int8", "int16", "int32", "int64",
		"rune", "string",
		"uintptr", "uint", "uint8", "uint16", "uint32", "uint64",
	}
}

// Builtin returns all golang reserved keywords or standard functions (accessible without any import).
func Builtin() []string {
	return []string{
		// builtin functions
		"append", "cap", "clear", "close", "complex", "copy",
		"delete", "imag", "len", "make", "max", "min", "new",
		"panic", "print", "println", "real", "recover",

		// keywords
		"break", "default", "func", "interface", "select",
		"case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var",
	}
}
