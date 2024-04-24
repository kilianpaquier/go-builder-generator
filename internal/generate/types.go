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
	Prefix       string
	Structs      []string
	ValidateFunc string
	ReturnCopy   bool
}

// implData represents the struct for the _impl file to generate.
type implData struct {
	Builders    []genData
	DestPackage string
}

// genData represents the struct for a builder to generate.
type genData struct {
	DefaultFuncs []string
	Exported     bool
	TypeParams   []field
	Name         string
	Fields       []field

	Opts    CLIOptions
	Package packageData
}

type packageData struct {
	DestDir    string
	DestName   string
	Imports    []string
	SourceName string
}

// fieldOpts represents the available options to be put in `builder` tag at a field level.
type fieldOpts struct {
	Append      bool   `json:"append,omitempty"`
	DefaultFunc string `json:"default_func,omitempty"`
	Ignore      bool   `json:"ignore,omitempty"`
	Pointer     bool   `json:"pointer,omitempty"`
	FuncName    string `json:"func_name,omitempty"`
}

// field represents one parsed struct field.
type field struct {
	AlteredType string
	Exported    bool
	InitialType string
	Name        string
	ParamName   string

	Opts fieldOpts
}
