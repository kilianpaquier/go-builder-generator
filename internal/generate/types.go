package generate

import (
	"io/fs"
	"path/filepath"
	"strings"
)

const (
	// Rw represents a file permission of read/write for current user
	// and no access for user's group and other groups.
	Rw fs.FileMode = 0o600
	// RwRR represents a file permission of read/write for current user
	// and read-only access for user's group and other groups.
	RwRR fs.FileMode = 0o644
	// RwRwRw represents a file permission of read/write for current user
	// and read/write too for user's group and other groups.
	RwRwRw fs.FileMode = 0o666
	// RwxRxRxRx represents a file permission of read/write/execute for current user
	// and read/execute for user's group and other groups.
	RwxRxRxRx fs.FileMode = 0o755
)

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
	NoCMD        bool
	NoNotice     bool
	NoTool       bool
	PackageName  string
	Prefix       string
	ReturnCopy   bool
	Structs      []string
	ValidateFunc string
}

// FileRelPath returns the relative path of options File property and options Destdir.
func (c CLIOptions) FileRelPath() string {
	if strings.HasPrefix(c.File, modulePrefix) || strings.HasPrefix(c.File, stdPrefix) {
		return c.File
	}
	file, _ := filepath.Rel(c.Destdir, c.File) // relative path to destdir since RelPath is expected to be called in templates (in destination directory)
	return filepath.ToSlash(file)
}

// ToArgs serializes back the input options into its slice of string representation.
func (c CLIOptions) ToArgs(name string) []string {
	var args []string

	if c.Destdir != "" {
		args = append(args, "-d", ".") // . since ToString is expected to be called in templates (in destination directory)
	}

	if c.File != "" {
		args = append(args, "-f", c.FileRelPath())
	}

	if name != "" {
		args = append(args, "-s", name)
	} else if len(c.Structs) > 0 {
		args = append(args, "-s", strings.Join(c.Structs, ","))
	}

	if c.ValidateFunc != "" {
		args = append(args, "--validate-func", c.ValidateFunc)
	}

	if c.Prefix != "" {
		args = append(args, "-p", c.Prefix)
	}

	if c.PackageName != "" {
		args = append(args, "--package-name", c.PackageName)
	}

	if c.NoNotice {
		args = append(args, "--no-notice")
	}

	if c.NoCMD {
		args = append(args, "--no-cmd")
	}

	if c.NoTool {
		args = append(args, "--no-tool")
	}

	if c.ReturnCopy {
		args = append(args, "--return-copy")
	}

	return args
}

// ToString serializes back the input command into its string format.
func (c CLIOptions) ToString(name string) string {
	args := c.ToArgs(name)
	return strings.Join(args, " ")
}

// implData represents the struct for the _impl file to generate.
type implData struct {
	Builders []genData

	Opts     CLIOptions
	Packages packagesData
}

// genData represents the struct for a builder to generate.
type genData struct {
	DefaultFuncs []string
	Exported     bool
	TypeParams   []field
	Name         string
	Fields       []field

	Opts     CLIOptions
	Packages packagesData
}

type packagesData struct {
	Destdir       string
	DestName      string
	GeneratedFrom string
	HasGenerate   bool
	Imports       []string
	SameModule    bool
	SourceName    string
	ToolAvailable bool
}

// fieldOpts represents the available options to be put in `builder` tag at a field level.
type fieldOpts struct {
	Append      bool   `json:"append,omitempty"`
	DefaultFunc string `json:"default_func,omitempty"`
	Export      bool   `json:"export,omitempty"`
	FuncName    string `json:"func_name,omitempty"`
	Ignore      bool   `json:"ignore,omitempty"`
	Pointer     bool   `json:"pointer,omitempty"`
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
