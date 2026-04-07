package cmd

import (
	"fmt"
	"io"
	"slices"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a pattern from the exclude file",
	Long: `The del command removes a specified pattern from the .git/info/exclude file.
This is useful for un-excluding files or directories that were previously excluded.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, ok := ContextObjectFromContext(cmd.Context())
		if !ok {
			return fmt.Errorf("no exclude file found in context")
		}
		if len(args) == 0 {
			return fmt.Errorf("no pattern provided to delete")
		}
		pattern := args[0]
		return runDelCommand(ctx.Out, ctx.File, pattern)
	},
}

func init() {
	RootCmd.AddCommand(delCmd)
}

// runDelCommand deletes the specified pattern from the exclude file and writes the updated content back
func runDelCommand(out io.Writer, f readWriteTruncater, pattern string) error {
	existingPatterns, err := readExistingPatterns(f)
	if err != nil {
		return fmt.Errorf("error reading existing patterns: %v", err)
	}

	if !slices.Contains(existingPatterns, pattern) {
		fmt.Fprintf(out, "Pattern '%s' not found in the exclude file. Nothing to delete.\n", pattern)
		return nil
	}

	var updatedPatterns []string
	for _, p := range existingPatterns {
		if p != pattern {
			updatedPatterns = append(updatedPatterns, p)
		}
	}

	if err := f.Truncate(0); err != nil {
		return fmt.Errorf("error truncating exclude file: %w", err)
	}
	if _, err := f.Seek(0, 0); err != nil {
		return fmt.Errorf("error seeking to the beginning of exclude file: %w", err)
	}

	if err := writeDefaultExcludeContent(f); err != nil {
		return fmt.Errorf("error writing default exclude content: %w", err)
	}
	for _, p := range updatedPatterns {
		if _, err := fmt.Fprintln(f, p); err != nil {
			return fmt.Errorf("error writing updated patterns to exclude file: %v", err)
		}
	}
	fmt.Fprintf(out, "Deleted pattern '%s' from the exclude file.\n", pattern)
	return nil
}
