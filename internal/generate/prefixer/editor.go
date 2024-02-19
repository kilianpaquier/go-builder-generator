package prefixer

// editor is the func taking as input a computed string type and
// returning the same string type or an altered one.
type editor func(stringType string, exported bool) (string, bool)

// editorPrefixer implements Prefixer to override after
// ToString execution the output string representation.
type editorPrefixer struct {
	Editor   editor
	Prefixer Prefixer
}

var _ Prefixer = &editorPrefixer{} // ensure interface is implemented

// NewPrefixerEditor creates a Prefixer wrapping another Prefixer.
//
// It can be used to alter the result of the wrapped Prefixer to add like a name,
// edit the exported bool result or even remove, replace, or change the resulted string type.
func NewPrefixerEditor(prefixer Prefixer, editor editor) Prefixer {
	return &editorPrefixer{
		Editor:   editor,
		Prefixer: prefixer,
	}
}

// Valid validates the prefixer and its subprefixers.
//
// An example would be a composition of a StarExpr with an ArrayType of an Ident.
// In that case, all three prefixers computed from those ast.Expr will be validated with Valid.
func (i *editorPrefixer) Valid() error {
	return i.Prefixer.Valid()
}

// ToString transforms a Prefixer (ast.Expr) into its string representation.
// It also returns a boolean indicating whether the type is exported.
func (i *editorPrefixer) ToString(sourcePackage string, _ ...string) (_ string, _ bool) {
	stringType, exported := i.Prefixer.ToString(sourcePackage)

	// apply editor on computed string type
	return i.Editor(stringType, exported)
}
