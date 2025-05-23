package cobra

import (
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/kilianpaquier/go-builder-generator/internal/generate"
)

var (
	generateOpts = generate.CLIOptions{}

	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate builders for structs arguments present in file argument.",
		Args: func(cmd *cobra.Command, args []string) error {
			// now that we have the raw flags in args slice, reenable flag parsing for ParseFlags
			cmd.DisableFlagParsing = false

			// parse flags into generateOpts (yeah it's wobbly but cobra is missing this issue https://github.com/spf13/cobra/issues/1832)
			if err := cmd.ParseFlags(args); err != nil {
				return err
			}

			if help, _ := cmd.Flags().GetBool("help"); help {
				return flag.ErrHelp
			}
			return nil
		},
		Run: func(_ *cobra.Command, args []string) {
			if err := generate.Run(generateOpts, args); err != nil {
				logger.Fatal(err)
			}
		},
	}
)

func init() {
	generateCmd.DisableFlagParsing = true // disable flags parsing to get raw flags in args slice
	rootCmd.AddCommand(generateCmd)

	// dest flag
	generateCmd.Flags().StringVarP(&generateOpts.Destdir, "dest", "d", ".",
		"destination directory for the generated files")

	// files flag
	generateCmd.Flags().StringVarP(&generateOpts.File, "file", "f", "",
		"input file containing golang struct(s) to generate builders on")
	_ = generateCmd.MarkFlagRequired("file")

	// no cmd flag
	generateCmd.Flags().BoolVar(&generateOpts.NoCMD, "no-cmd", false,
		"removes top comment 'go:generate ...' from generated files")

	// no notice flag
	generateCmd.Flags().BoolVar(&generateOpts.NoNotice, "no-notice", false,
		"removes top notice 'Code generated by ...' from generated files")

	// no tool flag
	generateCmd.Flags().BoolVar(&generateOpts.NoTool, "no-tool", false,
		"avoid using go-build-generator as go tool when 'go:generate ...' is added in generated files")

	// specific package name flag
	generateCmd.Flags().StringVar(&generateOpts.PackageName, "package-name", "",
		"defines a specific package name instead of '--dest', '-d' directory name. Only available when generating files in another directory")

	// setter prefix flag
	generateCmd.Flags().StringVarP(&generateOpts.Prefix, "prefix", "p", "",
		"specific prefix to apply on setter functions")

	// copy before update and return copy
	generateCmd.Flags().BoolVar(&generateOpts.ReturnCopy, "return-copy", false,
		"returns a copy of the builder each time a setter function is called")

	// structs flag
	generateCmd.Flags().StringSliceVarP(&generateOpts.Structs, "structs", "s", []string{},
		"struct names to generate builders on (they must be contained in given input file)")
	_ = generateCmd.MarkFlagRequired("structs")

	// validate func flag
	generateCmd.Flags().StringVar(&generateOpts.ValidateFunc, "validate-func", "",
		"validate function name to be executed in Build, must have the signature 'func () error' and associated to built struct")
}
