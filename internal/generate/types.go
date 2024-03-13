package generate

const (
	// GenTemplate is the filename inside embedded fs for _gen files.
	GenTemplate = "template_gen.go.tmpl"
	// ImplTemplate is the filename inside embedded fs for _impl file.
	ImplTemplate = "template_impl.go.tmpl"
)

// CLIOptions represents the struct of available options to be given to go-builder-generator.
type CLIOptions struct {
	Destdir      string
	File         string
	NoNotice     bool
	SetterPrefix string
	Structs      []string
	ValidateFunc string
}

// implData represents the struct for the _impl file to generate.
type implData struct {
	Builders    []*implBuilder
	DestPackage string
}

// implBuilder represents a simplified struct of GenBuilder to be given to ImplData for _impl file generation.
type implBuilder struct {
	DefaultFuncs []string
	Name         string
}

// genBuilder represents the struct for a builder to generate.
type genBuilder struct {
	genOpts
	implBuilder

	Properties  []property
	HasValidate bool
}

// genOpts represents the global parsed options that are given to all builders when generating.
type genOpts struct {
	DestPackage   string
	Imports       []string
	NoNotice      bool
	SetterPrefix  string
	SourcePackage string
	ValidateFunc  string
}

// propertyOpts represents the available options to be put in `builder`
// tag at a property level.
type propertyOpts struct {
	Append      bool
	DefaultFunc string
	Ignore      bool
	Pointer     bool
}

// property represents one parsed struct field.
type property struct {
	propertyOpts

	AlteredType string
	InitialType string
	Name        string
	ParamName   string
}
