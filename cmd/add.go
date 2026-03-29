package cmd

import (
	"fmt"
	"io"
	"slices"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add patterns to the exclude file",
	Long: `The add command allows you to add one or more patterns to the .git/info/exclude file.
This is useful for excluding files or directories from version control without modifying the .gitignore file.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, ok := ContextObjectFromContext(cmd.Context())
		if !ok {
			return fmt.Errorf("no exclude file found in context")
		}
		return runAddCommand(ctx.Out, ctx.File, args)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

// runAddCommand adds the specified patterns to the exclude file, ensuring no duplicates
// It handles both file and in-memory buffer cases for testing
func runAddCommand(out io.Writer, f io.ReadWriter, args []string) error {
	existingPatterns, err := readExistingPatterns(f)
	if err != nil {
		return fmt.Errorf("error reading existing patterns: %v", err)
	}

	for _, pattern := range args {
		if slices.Contains(existingPatterns, pattern) {
			fmt.Fprintf(out, "Pattern '%s' already exists in the exclude file. Skipping.\n", pattern)
			continue
		}

		_, err := fmt.Fprintln(f, pattern)
		if err != nil {
			return fmt.Errorf("error writing to exclude file: %v", err)
		}
		fmt.Fprintf(out, "Added pattern '%s' to the exclude file.\n", pattern)
	}
	return nil
}
