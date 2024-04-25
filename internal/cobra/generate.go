package cobra

import (
	"os"

	"github.com/huandu/xstrings"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/kilianpaquier/go-builder-generator/internal/generate"
)

var (
	generateOpts = generate.CLIOptions{}

	generateCmd = &cobra.Command{
		Use:    "generate",
		Short:  "Generate builders for structs arguments present in file argument.",
		PreRun: SetLogLevel,
		Run: func(cmd *cobra.Command, _ []string) {
			// force first rune to lowercase in case of unexported types
			// it will be titled in gen template in case the type is exported
			generateOpts.Prefix = xstrings.FirstRuneToLower(generateOpts.Prefix)

			pwd, _ := os.Getwd()
			if err := generate.Run(pwd, generateOpts); err != nil {
				logrus.WithContext(cmd.Context()).
					WithError(err).
					Fatal("failed to generate builders")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)

	// files flag
	generateCmd.Flags().StringVarP(&generateOpts.File, "file", "f", "",
		"input file containing the structures to generate builders on")
	_ = generateCmd.MarkFlagRequired("file")

	// structs flag
	generateCmd.Flags().StringSliceVarP(&generateOpts.Structs, "structs", "s", []string{},
		"structures names to generate builders on (they must be contained in given input files)")
	_ = generateCmd.MarkFlagRequired("structs")

	// dest flag
	generateCmd.Flags().StringVarP(&generateOpts.Destdir, "dest", "d", ".",
		"destination directory for the generated files")

	// no notice flag
	generateCmd.Flags().BoolVar(&generateOpts.NoNotice, "no-notice", false,
		"remove top notice 'Code generated by ...' in generated files")

	// validate func flag
	generateCmd.Flags().StringVar(&generateOpts.ValidateFunc, "validate-func", "",
		"validate function name to be executed in Build, must have the signature 'func () error' and associated to built struct")

	// setter prefix flag
	generateCmd.Flags().StringVarP(&generateOpts.Prefix, "prefix", "p", "",
		"specific prefix to apply on setter functions")

	// copy before update and return copy
	generateCmd.Flags().BoolVar(&generateOpts.ReturnCopy, "return-copy", false,
		"returns a copy of the builder each time a setter function is called")
}
