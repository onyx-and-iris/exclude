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
// It handles both file and in-memory buffer cases for testing
func runDelCommand(out io.Writer, f any, pattern string) error {
	r, ok := f.(io.Reader)
	if !ok {
		return fmt.Errorf("provided file does not support Reader")
	}
	existingPatterns, err := readExistingPatterns(r)
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

	var w io.Writer
	if t, ok := f.(truncater); ok {
		if err := t.Truncate(0); err != nil {
			return fmt.Errorf("error truncating exclude file: %w", err)
		}
		if s, ok := f.(io.Seeker); ok {
			if _, err := s.Seek(0, 0); err != nil {
				return fmt.Errorf("error seeking to the beginning of exclude file: %w", err)
			}
		}
		w, _ = f.(io.Writer)
	} else if buf, ok := f.(interface{ Reset() }); ok {
		buf.Reset()
		w, _ = f.(io.Writer)
	} else {
		return fmt.Errorf("provided file does not support writing")
	}

	if err := writeDefaultExcludeContent(w); err != nil {
		return fmt.Errorf("error writing default exclude content: %w", err)
	}
	for _, p := range updatedPatterns {
		if _, err := fmt.Fprintln(w, p); err != nil {
			return fmt.Errorf("error writing updated patterns to exclude file: %v", err)
		}
	}
	fmt.Fprintf(out, "Deleted pattern '%s' from the exclude file.\n", pattern)
	return nil
}
